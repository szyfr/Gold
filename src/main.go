package main


import (
	_"fmt"
	_"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/veandco/go-sdl2/sdl"
)


func main() {
	initialize_program();

//	rect := sdl.Rect{0,0,56,56};
//	program.testImage.Blit(&rect, program.sdlSurface, &rect);

	program.sdlWindow.UpdateSurface();

	for program.running {
		event_polling();
	}

	program.sdlWindow.Destroy();
	sdl.Quit();
}


func initialize_program() {
	program              = ProgramData{};
	program.windowHeight = 720;
	program.windowWidth  = 1280;
	program.sdlWindow    = init_sdl2();
	program.sdlSurface   = get_surface(program.sdlWindow);
	program.running      = true;

//	program.testImage  = load_media("data/TheBoi.bmp");
	
//	if err  := gl.Init(); err != nil { panic(err); }
//	version := gl.GoStr(gl.GetString(gl.VERSION));
//	fmt.Printf("OpenGL version: %s", version);

}