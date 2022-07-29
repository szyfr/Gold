package main


import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
)


func init_opengl() uint32 {
	if err  := gl.Init(); err != nil { panic(err); }
	version := gl.GoStr(gl.GetString(gl.VERSION));
	fmt.Printf("OpenGL version %s\n", version);

	prog := gl.CreateProgram();
	gl.LinkProgram(prog);
	return prog;
}

func make_vao(points []float32) uint32 {
	var vbo uint32;
	gl.GenBuffers(1, &vbo);
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo);
	gl.BufferData(gl.ARRAY_BUFFER, 4 * len(points), gl.Ptr(points), gl.STATIC_DRAW);

	var vao uint32;
	gl.GenVertexArrays(1, &vao);
	gl.BindVertexArray(vao);
	gl.EnableVertexAttribArray(0);
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo);
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil);

	return vao;
}