package main

import (
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
)

func get_shaders() (vertex uint32, fragment uint32) {
	vertex_code := `
	attribute vec2 position;

	void main() {
		gl_Position = vec4(position, 0.0, 1.0);
	}
	`

	fragment_code := `
	void main() {
		gl_FragColor = vec4(0.0, 0.0, 0.0, 1.0);
	}
	`

	c_vertex_code, c_vertex_free := gl.Strs(vertex_code)
	c_fragment_code, c_fragment_free := gl.Strs(fragment_code)

	defer c_vertex_free()
	defer c_fragment_free()

	new_vertex := gl.CreateShader(gl.VERTEX_SHADER)
	new_fragment := gl.CreateShader(gl.FRAGMENT_SHADER)

	gl.ShaderSource(new_vertex, 1, c_vertex_code, nil)
	gl.ShaderSource(new_fragment, 1, c_fragment_code, nil)

	return new_vertex, new_fragment
}

func compile_shader(shader_obj *uint32) {
	gl.CompileShader(*shader_obj)

	var status int32
	gl.GetShaderiv(*shader_obj, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var log_length int32
		gl.GetShaderiv(*shader_obj, gl.INFO_LOG_LENGTH, &log_length)

		log := strings.Repeat("\x00", int(log_length))

		gl.GetShaderInfoLog(*shader_obj, log_length, nil, gl.Str(log))

		print("gl.CompileShader(*shader_obj): ", log)
		panic("shader_obj compile error.")
	}
}

func link_program(program_obj *uint32) {
	gl.LinkProgram(*program_obj)

	var status int32
	gl.GetProgramiv(*program_obj, gl.LINK_STATUS, &status)

	if status == gl.FALSE {
		var log_length int32
		gl.GetShaderiv(*program_obj, gl.INFO_LOG_LENGTH, &log_length)

		log := strings.Repeat("\x00", int(log_length))

		gl.GetShaderInfoLog(*program_obj, log_length, nil, gl.Str(log))

		print("gl.LinkProgram(*program_obj): ", log)
		panic("Program link error.")
	}
}
