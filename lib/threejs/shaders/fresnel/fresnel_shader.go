package fresnel

import (
	"app/lib/threejs"
	"log"
	"syscall/js"
)

const modulePath = "./assets/threejs/ex/jsm/shaders/FresnelShader.js"

var module js.Value

func init() {

	m := threejs.LoadModule([]string{"FresnelShader"}, modulePath)
	if len(m) == 0 {
		log.Fatal("FresnelShader module could not be loaded.")
	}
	module = m[0]
}

// Shader is ...
type Shader struct {
	js.Value
}

// NewFresnelShader creates fresnel shader.
func NewFresnelShader() *Shader {
	return &Shader{
		Value: module,
	}
}

// Uniforms gets ...
func (c *Shader) Uniforms() js.Value {
	return c.Get("uniforms")
}

// VertexShader gets ...
func (c *Shader) VertexShader() js.Value {
	return c.Get("vertexShader")
}

// FragmentShader gets ...
func (c *Shader) FragmentShader() js.Value {
	return c.Get("fragmentShader")
}
