package main


import (
	"runtime"
	_"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)


var (
	triangle = []float32{
		 0.0,  0.5, 0.0,
		-0.5, -0.5, 0.0,
		 0.5, -0.5, 0.0,
	}
)


func main() {
	runtime.LockOSThread();

//	initialize_program();
	program = ProgramData{};
	program.windowHeight =  720;
	program.windowWidth  = 1280;

	program.glfwWindow   = init_glfw();
	defer glfw.Terminate();

	program.programID    = init_opengl();
	program.vao          = make_vao(triangle);

	for !program.glfwWindow.ShouldClose() {
		draw();
	}
}

func initialize_program() {
	program.windowHeight =  720;
	program.windowWidth  = 1280;

	program.glfwWindow   = init_glfw();
	defer glfw.Terminate();

	program.programID    = init_opengl();
	program.vao          = make_vao(triangle);
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
	gl.UseProgram(program.programID);

	gl.BindVertexArray(program.vao);
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle) / 3))

	glfw.PollEvents();
	program.glfwWindow.SwapBuffers();
}