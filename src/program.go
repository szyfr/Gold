package main


import (
	_"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)


const (
	vertShaderSource = "#version 330 core\n layout (location = 0) in vec3 Position; void main() { gl_Position = vec4(Position, 1.0); }";
	fragShaderSource = "#version 330 core\n out vec4 Color; void main() { Color = vec4(1.0f, 0.5f, 0.2f, 1.0f); }"
)


var program ProgramData;


type ProgramData struct {
	// General
	windowHeight int32;
	windowWidth  int32;
	running      bool;

	// SDL2
	sdlWindow    *sdl.Window;
	sdlSurface   *sdl.Surface;

	// OpenGL
	glContext  sdl.GLContext;
	programID  uint32;
	vertShader uint32;
	fragShader uint32;

//	testImage  *sdl.Surface;
}