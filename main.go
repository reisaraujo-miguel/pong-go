package main

import (
	"log"
	"unsafe"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("Failed to initialize glfw: ", err)
	}

	defer glfw.Terminate()

	glfw.WindowHint(glfw.Visible, glfw.False)

	window, err := glfw.CreateWindow(720, 600, "Pong-Go!", nil, nil)

	if err != nil {
		log.Fatalln("Could not create window: ", err)
	}

	window.MakeContextCurrent()

	glfw.GetCurrentContext().SetKeyCallback(key_event)
	glfw.GetCurrentContext().SetMouseButtonCallback(mouse_event)

	if err := gl.Init(); err != nil {
		log.Fatalln("Could not init gl: ", err)
	}

	program := gl.CreateProgram()
	vertex, fragment := get_shaders()

	compile_shader(&vertex)
	compile_shader(&fragment)

	gl.AttachShader(program, vertex)
	gl.AttachShader(program, fragment)

	link_program(&program)

	gl.UseProgram(program)

	vertices := []mgl32.Vec2{
		{0.0, -1.0},
		{0.7, 1.0},
		{-0.95, -0.32},
		{0.95, -0.32},
		{-0.7, 1.0},
	}

	var buffer uint32
	gl.GenBuffers(1, &buffer)

	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)

	vertex_data := gl.Ptr(vertices)
	vertex_size := int(unsafe.Sizeof(vertices[0])) * len(vertices)

	gl.BufferData(gl.ARRAY_BUFFER, vertex_size, vertex_data, gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer)

	loc := gl.GetAttribLocation(program, gl.Str("position\x00"))

	gl.EnableVertexAttribArray(uint32(loc))

	gl.VertexAttribPointer(uint32(loc), 2, gl.FLOAT, false, 0, nil)

	glfw.GetCurrentContext().Show()

	const (
		BG_RED   float32 = 1.0
		BG_BLUE  float32 = 1.0
		BG_GREEN float32 = 1.0
		BG_ALPHA float32 = 1.0
	)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.ClearColor(BG_RED, BG_BLUE, BG_GREEN, BG_ALPHA)

		gl.DrawArrays(gl.LINE_LOOP, 0, 5)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
