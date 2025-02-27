import glfw
import OpenGL.GL as gl
import numpy as np
from xvfbwrapper import Xvfb

class GLFWRenderer:
    def __init__(self, width=800, height=600):
        self.width = width
        self.height = height
        
        # 初始化 GLFW
        if not glfw.init():
            raise Exception("GLFW initialization failed")
            
        # 配置 GLFW
        glfw.window_hint(glfw.CONTEXT_VERSION_MAJOR, 3)
        glfw.window_hint(glfw.CONTEXT_VERSION_MINOR, 3)
        glfw.window_hint(glfw.OPENGL_PROFILE, glfw.OPENGL_CORE_PROFILE)
        glfw.window_hint(glfw.VISIBLE, glfw.FALSE)  # 离屏渲染
        
        # 创建窗口
        self.window = glfw.create_window(width, height, "GLFW Window", None, None)
        if not self.window:
            glfw.terminate()
            raise Exception("Failed to create GLFW window")
            
        # 设置上下文
        glfw.make_context_current(self.window)
        
        # 编译着色器
        self.shader_program = self.create_shader_program()
        
        # 创建 VAO 和 VBO
        self.setup_buffers()

    def create_shader_program(self):
        # 顶点着色器
        vertex_shader = """
        #version 330 core
        layout (location = 0) in vec3 position;
        layout (location = 1) in vec3 color;
        out vec3 vertex_color;
        
        void main() {
            gl_Position = vec4(position, 1.0);
            vertex_color = color;
        }
        """
        
        # 片段着色器
        fragment_shader = """
        #version 330 core
        in vec3 vertex_color;
        out vec4 fragment_color;
        
        void main() {
            fragment_color = vec4(vertex_color, 1.0);
        }
        """
        
        # 编译顶点着色器
        vertex_shader_id = gl.glCreateShader(gl.GL_VERTEX_SHADER)
        gl.glShaderSource(vertex_shader_id, vertex_shader)
        gl.glCompileShader(vertex_shader_id)
        self.check_shader_compilation(vertex_shader_id, "vertex")
        
        # 编译片段着色器
        fragment_shader_id = gl.glCreateShader(gl.GL_FRAGMENT_SHADER)
        gl.glShaderSource(fragment_shader_id, fragment_shader)
        gl.glCompileShader(fragment_shader_id)
        self.check_shader_compilation(fragment_shader_id, "fragment")
        
        # 链接着色器程序
        program_id = gl.glCreateProgram()
        gl.glAttachShader(program_id, vertex_shader_id)
        gl.glAttachShader(program_id, fragment_shader_id)
        gl.glLinkProgram(program_id)
        
        # 清理着色器
        gl.glDeleteShader(vertex_shader_id)
        gl.glDeleteShader(fragment_shader_id)
        
        return program_id

    def check_shader_compilation(self, shader_id, shader_type):
        success = gl.glGetShaderiv(shader_id, gl.GL_COMPILE_STATUS)
        if not success:
            info_log = gl.glGetShaderInfoLog(shader_id)
            print(f"{shader_type} shader compilation failed: {info_log}")
            raise Exception(f"{shader_type} shader compilation failed")

    def setup_buffers(self):
        # 顶点数据：位置和颜色
        vertices = np.array([
            # 位置             # 颜色
            -0.5, -0.5, 0.0,  1.0, 0.0, 0.0,  # 底部左边 - 红色
             0.5, -0.5, 0.0,  0.0, 1.0, 0.0,  # 底部右边 - 绿色
             0.0,  0.5, 0.0,  0.0, 0.0, 1.0   # 顶部     - 蓝色
        ], dtype=np.float32)
        
        # 创建并绑定 VAO
        self.vao = gl.glGenVertexArrays(1)
        gl.glBindVertexArray(self.vao)
        
        # 创建并绑定 VBO
        self.vbo = gl.glGenBuffers(1)
        gl.glBindBuffer(gl.GL_ARRAY_BUFFER, self.vbo)
        gl.glBufferData(gl.GL_ARRAY_BUFFER, vertices.nbytes, vertices, gl.GL_STATIC_DRAW)
        
        # 设置顶点属性指针
        # 位置属性
        gl.glVertexAttribPointer(0, 3, gl.GL_FLOAT, gl.GL_FALSE, 24, None)
        gl.glEnableVertexAttribArray(0)
        # 颜色属性
        gl.glVertexAttribPointer(1, 3, gl.GL_FLOAT, gl.GL_FALSE, 24, gl.ctypes.c_void_p(12))
        gl.glEnableVertexAttribArray(1)

    def render_frame(self):
        # 清除颜色缓冲
        gl.glClearColor(0.2, 0.3, 0.3, 1.0)
        gl.glClear(gl.GL_COLOR_BUFFER_BIT)
        
        # 使用着色器程序
        gl.glUseProgram(self.shader_program)
        
        # 绘制三角形
        gl.glBindVertexArray(self.vao)
        gl.glDrawArrays(gl.GL_TRIANGLES, 0, 3)
        
        # 交换缓冲
        glfw.swap_buffers(self.window)
        glfw.poll_events()

    def cleanup(self):
        # 清理资源
        gl.glDeleteVertexArrays(1, [self.vao])
        gl.glDeleteBuffers(1, [self.vbo])
        gl.glDeleteProgram(self.shader_program)
        glfw.terminate()

def main():
    # 启动虚拟显示
    vdisplay = Xvfb()
    vdisplay.start()
    
    try:
        # 创建渲染器
        renderer = GLFWRenderer()
        
        # 打印 OpenGL 信息
        print(f"OpenGL Version: {gl.glGetString(gl.GL_VERSION).decode()}")
        print(f"OpenGL Vendor: {gl.glGetString(gl.GL_VENDOR).decode()}")
        print(f"OpenGL Renderer: {gl.glGetString(gl.GL_RENDERER).decode()}")
        
        # 渲染循环
        for frame in range(100):  # 渲染100帧
            renderer.render_frame()
            if frame % 10 == 0:  # 每10帧打印一次
                print(f"Rendered frame {frame}")
            
            # 检查是否应该关闭窗口
            if glfw.window_should_close(renderer.window):
                break
                
    except Exception as e:
        print(f"Error occurred: {e}")
        
    finally:
        # 清理资源
        renderer.cleanup()
        vdisplay.stop()

if __name__ == "__main__":
    main()