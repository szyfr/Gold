package main


import (
	"github.com/veandco/go-sdl2/sdl"
)


var program ProgramData;


type ProgramData struct {
	sdlWindow  *sdl.Window;
	sdlSurface *sdl.Surface;
	running     bool;

//	testImage  *sdl.Surface;
}