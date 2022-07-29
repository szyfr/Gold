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

	initialize_gold();

	for !program.glfwWindow.ShouldClose() {
		draw();
	}
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
	gl.UseProgram(program.programID);

	gl.BindVertexArray(program.vao);
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle) / 3))

	glfw.PollEvents();
	program.glfwWindow.SwapBuffers();
}