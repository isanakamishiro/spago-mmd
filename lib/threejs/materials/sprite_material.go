package materials

import (
	"app/lib/threejs"
	"syscall/js"
)

// SpriteMaterialParameters is ...
type SpriteMaterialParameters interface {
}

// SpriteMaterial is a material for a use with a Sprite.
type SpriteMaterial interface {
	threejs.Material

	// Color of the material, by default set to white (0xffffff). The .map is mutiplied by the color.
	Color() threejs.Color

	// SetRotation sets the rotation of the sprite in radians. Default is 0.
	SetRotation(v float64)
	// SetRotation gets the rotation of the sprite in radians. Default is 0.
	Rotation() float64

	// Whether the size of the sprite is attenuated by the camera depth. (Perspective camera only.) Default is true.
	SetSizeAttenuation(b bool)
	// Whether the size of the sprite is attenuated by the camera depth. (Perspective camera only.) Default is true.
	SizeAttenuation() bool
}

// spriteMaterialImp is a implementation of SpriteMaterial.
type spriteMaterialImp struct {
	threejs.Material
}

// NewSpriteMaterial is constructor with parameter.
// The exception is the property color, which can be passed in as a hexadecimal string
// and is 0xffffff (white) by default. Color.set( color ) is called internally.
// SpriteMaterials are not clipped by using Material.clippingPlanes.
//
// parameters - (optional) an object with one or more properties defining the material's appearance.
// Any property of the material (including any property inherited from Material) can be passed in here.
func NewSpriteMaterial(parameters SpriteMaterialParameters) SpriteMaterial {
	return &spriteMaterialImp{
		threejs.NewDefaultMaterialFromJSValue(threejs.GetJsObject("SpriteMaterial").New(parameters)),
	}
}

// NewSpriteMaterialFromJSValue creates SpriteMaterial from js.Value object.
func NewSpriteMaterialFromJSValue(v js.Value) SpriteMaterial {
	return &spriteMaterialImp{
		threejs.NewDefaultMaterialFromJSValue(v),
	}

}

// func (c *spriteMaterialImp) Map() threejs.Texture {
// 	return threejs.NewDefaultTextureFromJSValue(
// 		c.JSValue().Get("map"),
// 	)
// }

func (c *spriteMaterialImp) Color() threejs.Color {
	return threejs.NewColorFromJSValue(
		c.JSValue().Get("color"),
	)
}

func (c *spriteMaterialImp) SetRotation(v float64) {
	c.JSValue().Set("rotation", v)
}

func (c *spriteMaterialImp) Rotation() float64 {
	return c.JSValue().Get("rotation").Float()
}

func (c *spriteMaterialImp) SetSizeAttenuation(b bool) {
	c.JSValue().Set("sizeAttenuation", b)
}

func (c *spriteMaterialImp) SizeAttenuation() bool {
	return c.JSValue().Get("sizeAttenuation").Bool()
}
