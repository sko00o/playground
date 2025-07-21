import asyncio
from ctypes import cast
from typing import Dict, Optional, Iterable, Any, cast
import json
from contextlib import AsyncExitStack

from mcp import ClientSession, StdioServerParameters
from mcp.client.stdio import stdio_client
from mcp.types import CallToolResult

from anthropic import Anthropic
from anthropic.types import MessageParam, TextBlock, ToolParam
from dotenv import load_dotenv

load_dotenv()


class MCPClient:
    def __init__(self) -> None:
        self.session: Optional[ClientSession] = None
        self.exit_stack = AsyncExitStack()
        self.anthropic = Anthropic()
        self.model = "claude-3-5-sonnet-20241022"

    async def connet_to_server(self, server_params: StdioServerParameters):
        stdio_transport = await self.exit_stack.enter_async_context(
            stdio_client(server_params)
        )
        self.stdio, self.write = stdio_transport
        self.session = await self.exit_stack.enter_async_context(
            ClientSession(self.stdio, self.write)
        )
        await self.session.initialize()

        # List available tools
        response = await self.session.list_tools()
        tools = response.tools
        print("\nConnected to server with tools:", [tool.name for tool in tools])

    async def process_query(self, query: str) -> str:
        """Process a query using Claude and available tools"""
        if self.session is None:
            raise RuntimeError("Not connected to server")

        messages: list[MessageParam] = [
            MessageParam(
                role="user",
                content=[TextBlock(text=query, type="text")],
            )
        ]

        response = await self.session.list_tools()
        available_tools: list[ToolParam] = [
            ToolParam(
                input_schema=tool.inputSchema,
                name=tool.name,
                description=tool.description or "",
            )
            for tool in response.tools
        ]

        # Initial Claude API call
        response = self.anthropic.messages.create(
            model=self.model,
            max_tokens=1000,
            messages=messages,
            tools=available_tools,
        )

        # Process response and handle tool calls
        final_text = []

        for content in response.content:
            if content.type == "text":
                final_text.append(content.text)
            elif content.type == "tool_use":
                tool_name = content.name
                tool_args = cast(Dict[str, Any], content.input)

                # Execute tool call
                result: CallToolResult = await self.session.call_tool(
                    tool_name, dict(tool_args)
                )
                final_text.append(f"[Calling tool {tool_name} with args {tool_args}]")

                messages.append(
                    MessageParam(
                        role="user",
                        content=cast(Iterable[TextBlock], result.content),
                    )
                )

                # Get next response from Claude
                response = self.anthropic.messages.create(
                    model=self.model,
                    max_tokens=1000,
                    messages=messages,
                )

                if isinstance(response.content[0], TextBlock):
                    final_text.append(response.content[0].text)

        return "\n".join(final_text)

    async def chat_loop(self):
        """Run an interactive chat loop"""
        print("\nMCP Client Started!")
        print("Type your queries or 'quit' to exit.")

        while True:
            try:
                query = input("\nQuery: ").strip()

                if query.lower() == "quit":
                    break

                response = await self.process_query(query)
                print("\n" + response)

            except Exception as e:
                print(f"\nError: {str(e)}")

    async def cleanup(self):
        """Clean up resources"""
        await self.exit_stack.aclose()


async def main():
    try:
        # Read mcp.json configuration
        with open("mcp.json", "r") as f:
            config = json.load(f)

        # Get the first server configuration
        servers = config.get("mcpServers", {})
        if not servers:
            print("No MCP servers found in mcp.json")
            return

        server_name = list(servers.keys())[0]
        server_config = servers[server_name]

        print(f"Connecting to MCP server: {server_name}")

        # Configure server parameters from mcp.json
        command = server_config["command"]
        args = server_config.get("args", [])
        env = server_config.get("env")

        server_params = StdioServerParameters(command=command, args=args, env=env)

        client = MCPClient()

        try:
            await client.connet_to_server(server_params)
            await client.chat_loop()
        finally:
            await client.cleanup()

    except FileNotFoundError:
        print("mcp.json not found in current directory")
    except json.JSONDecodeError:
        print("Invalid JSON in mcp.json")
    except KeyError as e:
        print(f"Missing required configuration in mcp.json: {e}")


if __name__ == "__main__":
    asyncio.run(main())
