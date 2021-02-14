package materials

import "app/lib/threejs"

// MeshStandardMaterialParameters is ...
type MeshStandardMaterialParameters interface {
}

// MeshStandardMaterial is a standard physically based material, using Metallic-Roughness workflow.
// Physically based rendering (PBR) has recently become the standard in many 3D applications, such as Unity, Unreal and 3D Studio Max.
// This approach differs from older approaches in that instead of using approximations for the way in which light interacts with a surface,
// a physically correct model is used. The idea is that, instead of tweaking materials to look good under specific lighting,
// a material can be created that will react 'correctly' under all lighting scenarios.
//
// In practice this gives a more accurate and realistic looking result than the MeshLambertMaterial or MeshPhongMaterial,
// at the cost of being somewhat more computationally expensive.
//
// Shading is calculated in the same way as for the MeshPhongMaterial, using a Phong shading model.
// This calculates shading per pixel (i.e. in the fragment shader, AKA pixel shader) which gives more accurate results
// than the Gouraud model used by MeshLambertMaterial, at the cost of some performance.
//
// Note that for best results you should always specify an environment map when using this material.
// For a non-technical introduction to the concept of PBR and how to set up a PBR material, check out these articles by the people at marmoset:
// - Basic Theory of Physically Based Rendering
// - Physically Based Rendering and You Can Too
// Technical details of the approach used in three.js (and most other PBR systems) can be found is this paper from Disney (pdf), by Brent Burley.
type MeshStandardMaterial interface {
	threejs.Material

	Color() threejs.Color
}

type meshStandardMaterialImp struct {
	threejs.Material
}

// NewMeshStandardMaterial is constructor.
// parameters - (optional) an object with one or more properties defining the material's appearance.
// Any property of the material (including any property inherited from Material) can be passed in here.
// The exception is the property color, which can be passed in as a hexadecimal string and is 0xffffff (white) by default. Color.set( color ) is called internally.
func NewMeshStandardMaterial(parameters MeshStandardMaterialParameters) MeshStandardMaterial {
	return &meshStandardMaterialImp{
		threejs.NewDefaultMaterialFromJSValue(threejs.GetJsObject("MeshStandardMaterial").New(parameters)),
	}
}

// Color of the material, by default set to white (0xffffff).
func (c *meshStandardMaterialImp) Color() threejs.Color {
	return threejs.NewColorFromJSValue(
		c.JSValue().Get("color"),
	)
}
