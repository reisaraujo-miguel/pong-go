package main

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

func key_event(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	print("key")
}

func mouse_event(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {
	print("mouse")
}
