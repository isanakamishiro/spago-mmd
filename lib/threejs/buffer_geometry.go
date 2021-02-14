package threejs

import (
	"syscall/js"
)

// BufferGeometry is an efficient representation of mesh, line, or point geometry.
// Includes vertex positions, face indices, normals, colors, UVs, and custom attributes
// within buffers, reducing the cost of passing all this data to the GPU.
//
// To read and edit data in BufferGeometry attributes, see BufferAttribute documentation.
//
// For a less efficient but easier-to-use representation of geometry, see Geometry.
type BufferGeometry interface {
	Geometry

	// SetAttribute sets an attribute to this geometry. Use this rather than the attributes property, because an internal hashmap of .attributes is maintained to speed up iterating over attributes.
	SetAttribute(name string, attribute *BufferAttribute)
}

// geometryImpl extend: [EventDispatcher]
type bufferGeometryImpl struct {
	Geometry
}

// NewBufferGeometry creates a new BufferGeometry.
func NewBufferGeometry() BufferGeometry {
	return NewDefaultBufferGeometryFromJSValue(
		GetJsObject("BufferGeometry").New(),
	)
}

// NewDefaultBufferGeometryFromJSValue creates a new BufferGeometry.
func NewDefaultBufferGeometryFromJSValue(value js.Value) BufferGeometry {
	return &bufferGeometryImpl{
		NewDefaultGeometryFromJSValue(value),
	}
}

func (c *bufferGeometryImpl) SetAttribute(name string, attribute *BufferAttribute) {

	c.JSValue().Call("setAttribute", name, attribute.JSValue())

}
