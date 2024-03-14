package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func key_event(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
}

func move(window *glfw.Window, x *float32, y *float32) {
	if window.GetKey(glfw.KeyUp) == glfw.Press {
		*y += 0.01
	}
	if window.GetKey(glfw.KeyDown) == glfw.Press {
		*y -= 0.01
	}
	if window.GetKey(glfw.KeyRight) == glfw.Press {
		*x += 0.01
	}
	if window.GetKey(glfw.KeyLeft) == glfw.Press {
		*x -= 0.01
	}
}

func mouse_event(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	print("mouse")
}
