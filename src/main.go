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
	vertexShaderSource = `
		#version 410
		in vec3 vp;
		void main() {
		    gl_Position = vec4(vp, 1.0);
		}
	` + "\x00"

	fragmentShaderSource = `
	    #version 410
	    out vec4 frag_colour;
	    void main() {
	        frag_colour = vec4(1, 1, 1, 1);
	    }
	` + "\x00"
)


func main() {
	runtime.LockOSThread();

	initialize_gold();

	for !program.glfwWindow.ShouldClose() {
		draw();
	}

	glfw.Terminate();
}

func draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
	gl.UseProgram(program.programID);

	gl.BindVertexArray(program.vao);
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle) / 3))

	glfw.PollEvents();
	program.glfwWindow.SwapBuffers();
}