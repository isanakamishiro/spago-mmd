package sprite

import (
	"app/lib/threejs"
	"app/lib/threejs/material"
	"syscall/js"
)

// Sprite is is a plane that always faces towards the camera, generally with a partially transparent texture applied.
//
// Sprites do not cast shadows, setting 'castShadow = true' will have no effect.
type Sprite interface {
	threejs.Object3D

	// Material gets An instance of SpriteMaterial, defining the object's appearance. Default is a white SpriteMaterial.
	Material() threejs.Material

	// Center gets the sprite's anchor point, and the point around which the sprite rotates. A value of (0.5, 0.5) corresponds to the midpoint of the sprite. A value of (0, 0) corresponds to the lower left corner of the sprite. The default is (0.5, 0.5).
	Center() *threejs.Vector2

	// Geometry gets sprite buffered geometry
	Geometry() threejs.BufferGeometry

	CloneGeometryToOwn()
}

// SpriteImpl extend: [Object3D]
type spriteImpl struct {
	threejs.Object3D
}

// NewSprite is factory method for SpriteImpl.
func NewSprite(material material.SpriteMaterial) Sprite {
	return &spriteImpl{
		threejs.NewObject3DFromJSValue(
			threejs.Threejs("Sprite").New(material.JSValue()),
		),
	}
}

// NewSpriteFromJSValue is factory method for SpriteImpl.
func NewSpriteFromJSValue(value js.Value) Sprite {
	return &spriteImpl{
		threejs.NewObject3DFromJSValue(value),
	}
}

func (c *spriteImpl) Material() threejs.Material {
	return material.NewSpriteMaterialFromJSValue(
		c.JSValue().Get("material"),
	)
}

func (c *spriteImpl) Center() *threejs.Vector2 {
	return &threejs.Vector2{
		Value: c.JSValue().Get("center"),
	}
}

func (c *spriteImpl) Geometry() threejs.BufferGeometry {
	return threejs.NewBufferGeometryFromJSValue(
		c.JSValue().Get("geometry"),
	)
}

func (c *spriteImpl) CloneGeometryToOwn() {

	g := c.JSValue().Get("geometry").Call("clone")
	c.JSValue().Set("geometry", g)
}
