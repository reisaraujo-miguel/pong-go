package main

import (
	"log"

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

	window, err := glfw.CreateWindow(912, 560, "Pong-Go!", nil, nil)

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
		// left square
		{-0.89, -0.18},
		{-0.89, 0.18},
		{-0.95, -0.18},
		{-0.95, 0.18},
		// right square
		{0.89, -0.18},
		{0.89, 0.18},
		{0.95, -0.18},
		{0.95, 0.18},
	}

	send_to_gpu(&vertices, &program)

	glfw.GetCurrentContext().Show()

	const (
		BG_RED   float32 = 0.03
		BG_BLUE  float32 = 0.03
		BG_GREEN float32 = 0.03
		BG_ALPHA float32 = 1.0
	)

	var (
		t_y float32 = 0.0
		t_x float32 = 0.0
	)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		gl.ClearColor(BG_RED, BG_BLUE, BG_GREEN, BG_ALPHA)

		move(window, &t_x, &t_y)

		translation := mgl32.Mat4{
			1.0, 0.0, 0.0, t_x,
			0.0, 1.0, 0.0, t_y,
			0.0, 0.0, 1.0, 0.0,
			0.0, 0.0, 0.0, 1.0,
		}

		var loc uint8
		gl.GetUniformLocation(program, &loc)

		gl.UniformMatrix4fv(int32(loc), 1, true, &translation[0])

		gl.DrawArrays(gl.TRIANGLE_STRIP, 0, 4)
		gl.DrawArrays(gl.TRIANGLE_STRIP, 4, 4)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
