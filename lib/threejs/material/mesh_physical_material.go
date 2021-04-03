package material

import "app/lib/threejs"

// MeshPhysicalMaterialParameters is ...
type MeshPhysicalMaterialParameters interface {
}

// MeshPhysicaldMaterial is an extension of the MeshStandardMaterial,
// providing more advanced physically-based rendering properties:
//
// Clearcoat: Some materials — like car paints, carbon fiber,
// and wet surfaces — require a clear, reflective layer on top of another layer
// that may be irregular or rough. Clearcoat approximates this effect,
// without the need for a separate transparent surface.
//
// Physically-based transparency: One limitation of .opacity is that highly transparent materials
// are less reflective. Physically-based .transmission provides a more realistic option for thin,
// transparent surfaces like glass.
//
// Advanced reflectivity: More flexible reflectivity for non-metallic materials.
//
// As a result of these complex shading features,
// MeshPhysicalMaterial has a higher performance cost, per pixel,
// than other three.js materials. Most effects are disabled by default,
// and add cost as they are enabled. For best results, always specify an environment map
// when using this material.
type MeshPhysicaldMaterial interface {
	threejs.Material

	Color() threejs.Color
}

type meshPhysicalMaterialImp struct {
	threejs.Material
}

// NewMeshPhysicalMaterial is constructor.
// parameters - (optional) an object with one or more properties defining the material's appearance. Any property of the material (including any property inherited from Material and MeshStandardMaterial) can be passed in here.
//
// The exception is the property color, which can be passed in as a hexadecimal string and is 0xffffff (white) by default. Color.set( color ) is called internally.
func NewMeshPhysicalMaterial(parameters MeshPhysicalMaterialParameters) MeshPhysicaldMaterial {
	return &meshPhysicalMaterialImp{
		threejs.NewDefaultMaterialFromJSValue(threejs.Threejs("MeshPhysicalMaterial").New(parameters)),
	}
}

// Color of the material, by default set to white (0xffffff).
func (c *meshPhysicalMaterialImp) Color() threejs.Color {
	return threejs.NewColorFromJSValue(
		c.JSValue().Get("color"),
	)
}
