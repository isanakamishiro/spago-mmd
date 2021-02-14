package materials

import "app/lib/threejs"

// ShaderMaterialParameters is ...
type ShaderMaterialParameters interface{}

// ShaderMaterial is a material rendered with custom shaders. A shader is a small program written in GLSL that runs on the GPU. You may want to use a custom shader if you need to:
//
// - implement an effect not included with any of the built-in materials
// - combine many objects into a single Geometry or BufferGeometry in order to improve performance
type ShaderMaterial interface {
	threejs.Material
}

type shaderMaterialImp struct {
	threejs.Material
}

// NewShaderMaterial creates ShaderMaterial.
// Any property of the material (including any property inherited from Material) can be passed in here.
// parameters - (optional) an object with one or more properties defining the material's appearance.
func NewShaderMaterial(parameters ShaderMaterialParameters) ShaderMaterial {
	return &shaderMaterialImp{
		threejs.NewDefaultMaterialFromJSValue(threejs.GetJsObject("ShaderMaterial").New(parameters)),
	}
}
