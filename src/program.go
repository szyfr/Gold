package main


import (
	_"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)


var program ProgramData;


type ProgramData struct {
	// General
	windowHeight int32
	windowWidth  int32
	running      bool;

	// SDL2
	sdlWindow    *sdl.Window;
	sdlSurface   *sdl.Surface;

	// OpenGL
	glContext sdl.GLContext;
	programID uint32;

//	testImage  *sdl.Surface;
}