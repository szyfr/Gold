package main

import (
	"fmt"
	_ "github.com/go-gl/gl/v4.1-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SCREEN_HEIGHT = 480
	SCREEN_WIDTH  = 640
)

var sdlWindow  *sdl.Window;
var sdlSurface *sdl.Surface;

func main() {
	init_sdl2();

	sdlSurface, _ := sdlWindow.GetSurface();
	fmt.Printf("%p\n",sdlSurface);

	rect := sdl.Rect{50,50,200,200};
	sdlSurface.FillRect(&rect, 0xffff00ff);
	sdlWindow.UpdateSurface();

	running := true;
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("quit");
				running = false;
				break;
			}
		}
	}

	fmt.Printf("%p\n",sdlSurface);

	sdlWindow.Destroy();
	sdl.Quit();
}

func init_sdl2() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil { panic(err); }

	tempWindow, err := sdl.CreateWindow(
		"SDL Testing",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH, SCREEN_HEIGHT,
		sdl.WINDOW_SHOWN,
	);
	if err != nil { panic(err); }

	tempSurface, _ := sdlWindow.GetSurface();
	fmt.Printf("%p\n",tempSurface);
//	if err != nil { panic(err); }

	sdlWindow  = tempWindow;
	sdlSurface = tempSurface;
}