from fastapi import FastAPI, Request
from fastapi.responses import HTMLResponse, JSONResponse
from fastapi.templating import Jinja2Templates
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

# Configure CORS
app.add_middleware(
    CORSMiddleware,
    # allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Configure Jinja2 templating
templates = Jinja2Templates(directory="templates")


@app.get("/", response_class=HTMLResponse)
def read_root(request: Request):
    return templates.TemplateResponse(
        "index.html",
        {
            "request": request,
            "title": "Python Backend",
        },
    )


@app.post("/api/addr", response_class=JSONResponse)
def remote_addr(request: Request):
    c = request.client
    if c:
        remote_addr = f"{c.host}:{c.port}"
    return JSONResponse(content={"addr": f"{remote_addr}"})


import argparse
import uvicorn

if __name__ == "__main__":
    parse = argparse.ArgumentParser()
    parse.add_argument("--subpath", type=str, help="run in subpath for reverse proxy")
    args = parse.parse_args()
    if args.subpath != "" :
        app0 = FastAPI()
        app0.mount(args.subpath, app)
        uvicorn.run(app0, host="0.0.0.0", port=8080)
    else:
        uvicorn.run(app, host="0.0.0.0", port=8080)
