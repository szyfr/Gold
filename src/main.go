package main


import (
	"runtime"
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)


func main() {
	runtime.LockOSThread();

	window  := initGlfw();
	defer glfw.Terminate();
	program := initOpenGL();

	for !window.ShouldClose() {
		draw(window, program);
	}
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil { panic(err); }

	glfw.WindowHint(glfw.Resizable, glfw.False);
	glfw.WindowHint(glfw.ContextVersionMajor, 4);
	glfw.WindowHint(glfw.ContextVersionMinor, 1);
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile);
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True);

	window, err := glfw.CreateWindow(1280, 720, "GLFW and OpenGL test", nil, nil);
	if err != nil { panic(err); }

	window.MakeContextCurrent();

	return window;
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil { panic(err); }
	version := gl.GoStr(gl.GetString(gl.VERSION));
	fmt.Printf("OpenGL version %s\n", version);

	prog := gl.CreateProgram();
	gl.LinkProgram(prog);
	return prog;
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT);
	gl.UseProgram(program);

	glfw.PollEvents();
	window.SwapBuffers();
}