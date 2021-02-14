package materials

import (
	"app/lib/threejs"
)

// MeshBasicMaterialParameters is ...
type MeshBasicMaterialParameters interface {
}

// MeshBasicMaterial is a material for drawing geometries in a simple shaded (flat or wireframe) way.
//
// This material is not affected by lights.
type MeshBasicMaterial interface {
	threejs.Material

	Color() threejs.Color
}

// meshPhongMaterialImp is a implementation of MeshPhongMaterial.
type meshBasicMaterialImp struct {
	threejs.Material
}

// NewMeshBasicMaterial is constructor.
// parameters - (optional) an object with one or more properties defining the material's appearance.
// Any property of the material (including any property inherited from Material) can be passed in here.
//
// The exception is the property color,
// which can be passed in as a hexadecimal string and is 0xffffff (white) by default.
// Color.set( color ) is called internally.
func NewMeshBasicMaterial(parameters MeshBasicMaterialParameters) MeshBasicMaterial {
	return &meshBasicMaterialImp{
		threejs.NewDefaultMaterialFromJSValue(threejs.GetJsObject("MeshBasicMaterial").New(parameters)),
	}

}

// Color of the material, by default set to white (0xffffff).
func (c *meshBasicMaterialImp) Color() threejs.Color {
	return threejs.NewColorFromJSValue(
		c.JSValue().Get("color"),
	)
}
