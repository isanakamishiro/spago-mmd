package threejs

import (
	"syscall/js"
)

// MorphColor is ...
type MorphColor interface {
}

// MorphNormals is ...
type MorphNormals interface {
}

// MorphTarget is ...
type MorphTarget interface {
}

// Geometry is a user-friendly alternative to BufferGeometry.
// Geometries store attributes (vertex positions, faces, colors, etc.) using objects
// like Vector3 or Color that are easier to read and edit,
// but less efficient than typed arrays.
//
// Prefer BufferGeometry for large or serious projects.
type Geometry interface {
	JSValue() js.Value

	// SetUVsNeedUpdate set to true if the uvs array has been updated.
	SetUVsNeedUpdate(b bool)
}

// geometryImpl extend: [EventDispatcher]
type geometryImpl struct {
	js.Value
}

// NewGeometry creates a new Geometry.
func NewGeometry() Geometry {
	return NewDefaultGeometryFromJSValue(
		GetJsObject("Geometry").New(),
	)
}

// NewDefaultGeometryFromJSValue creates a new Geometry.
func NewDefaultGeometryFromJSValue(value js.Value) Geometry {
	return &geometryImpl{
		Value: value,
	}
}

// JSValue is ...
func (gg *geometryImpl) JSValue() js.Value {
	return gg.Value
}

func (gg *geometryImpl) SetUVsNeedUpdate(b bool) {
	gg.Set("uvsNeedUpdate", b)
}
