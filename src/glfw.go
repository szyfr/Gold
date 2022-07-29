package main


import (
	_"fmt"
	"github.com/go-gl/glfw/v3.3/glfw"
)


func (prg *ProgramData) init_glfw() {
	if err := glfw.Init(); err != nil { panic(err); }

	glfw.WindowHint(glfw.Resizable, glfw.False);
	glfw.WindowHint(glfw.ContextVersionMajor, 4);
	glfw.WindowHint(glfw.ContextVersionMinor, 1);
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile);
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True);

	window, err := glfw.CreateWindow(
		program.windowWidth, program.windowHeight,
		"GLFW and OpenGL Testing", nil, nil,
	);
	if err != nil { panic(err); }

	window.MakeContextCurrent();

	prg.glfwWindow = window;
}