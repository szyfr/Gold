package main


import (
	_"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)


func init_opengl() {
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4);
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 1);
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE);

	glContext, err := program.sdlWindow.GLCreateContext();
	if err != nil { panic(err); }
	program.glContext = glContext;
}