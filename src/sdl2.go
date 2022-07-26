package main


import (
	"github.com/veandco/go-sdl2/sdl"
)


func init_sdl2() *sdl.Window {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil { panic(err); }

//	sdl.GLSetAttribute(sdl., 5);

	tempWindow, err := sdl.CreateWindow(
		WINDOW_NAME,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		program.windowWidth, program.windowHeight,
		sdl.WINDOW_SHOWN,
	);
	if err != nil { panic(err); }

	return tempWindow;
}

func get_surface(window *sdl.Window) *sdl.Surface {
	tempSurface, err := window.GetSurface();
	if err != nil { panic(err); }

	return tempSurface;
}

func event_polling() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("quit");
			program.running = false;
			break;
		}
	}
}

func load_media(path string) *sdl.Surface {
	surface, err := sdl.LoadBMP(path);
	if err != nil { panic(err); }

	return surface;
}
func close_media(image *sdl.Surface) {
	image.Free();
}