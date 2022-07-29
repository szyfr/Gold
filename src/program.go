package main


// Imports
import (
	_"runtime"
	_"fmt"
	_"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)


// Global
var program *ProgramData;


// Structures
type ProgramData struct {
	// General
	windowHeight int;
	windowWidth  int;

	// GLFW
	glfwWindow    *glfw.Window;

	// OpenGL
	programID  uint32;
	vao        uint32;
	vertShader uint32;
	fragShader uint32;

}


// Procedures
func initialize_gold() {
	program = new(ProgramData);

	program.windowHeight =  720;
	program.windowWidth  = 1280;

	program.init_glfw();

	program.programID    = init_opengl();
	program.vao          = make_vao(triangle);
}