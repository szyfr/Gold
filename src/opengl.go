package main


import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)


func init_opengl() {
	// Set Attributes
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4);
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 1);
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE);

	// Create Context
	glContext, err := program.sdlWindow.GLCreateContext();
	if err != nil { panic(err); }
	program.glContext = glContext;

	// Set VSync
//	if sdl.GLSetSwapInterval(1) != nil { fmt.Printf("Unable to set VSync!\n"); }

	// Init OpenGL
	if err = gl.Init(); err != nil { panic(err); }

	// Create Program
	program.programID  = gl.CreateProgram();

	// Create and compile Vertex Shader
	program.vertShader = gl.CreateShader(gl.VERTEX_SHADER);
	csources, free := gl.Strs(vertShaderSource);
	gl.ShaderSource(program.vertShader, 1, csources, nil);
	free();
	gl.CompileShader(program.vertShader);

	var shaderCompiled int32 = gl.FALSE;
	gl.GetShaderiv(program.vertShader, gl.COMPILE_STATUS, &shaderCompiled);
	if shaderCompiled != gl.TRUE { fmt.Printf("Vertex Shader failed to compile %d!\n", program.vertShader); }

	gl.AttachShader(program.programID, program.vertShader);

	// Create and compile Fragment Shader
	program.fragShader = gl.CreateShader(gl.FRAGMENT_SHADER);
	csources, free = gl.Strs(fragShaderSource);
	gl.ShaderSource(program.fragShader, 1, csources, nil);
	free();
	gl.CompileShader(program.fragShader);

	shaderCompiled = gl.FALSE;
	gl.GetShaderiv(program.fragShader, gl.COMPILE_STATUS, &shaderCompiled);
	if shaderCompiled != gl.TRUE { fmt.Printf("Fragment Shader failed to compile %d!\n", program.fragShader); }

	gl.AttachShader(program.programID, program.fragShader);

	// Link program
	gl.LinkProgram(program.programID);

	// Error Checking
	var programSuccess int32 = gl.TRUE;
	gl.GetProgramiv(program.programID, gl.LINK_STATUS, &programSuccess);
	if programSuccess != gl.TRUE { fmt.Printf("Failed to Link Program!\n"); }
}