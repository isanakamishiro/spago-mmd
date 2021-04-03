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
	JSValue() js.Value

	// SetAttribute sets an attribute to this geometry. Use this rather than the attributes property, because an internal hashmap of .attributes is maintained to speed up iterating over attributes.
	SetAttribute(name string, attribute *BufferAttribute)

	// Disposes the object from memory.
	// You need to call this when you want the BufferGeometry removed while the application is running.
	Dispose()
}

// geometryImpl extend: [EventDispatcher]
type bufferGeometryImpl struct {
	js.Value
}

// NewBufferGeometry creates a new BufferGeometry.
func NewBufferGeometry() BufferGeometry {
	return NewBufferGeometryFromJSValue(
		Threejs("BufferGeometry").New(),
	)
}

// NewBufferGeometryFromJSValue creates a new BufferGeometry.
func NewBufferGeometryFromJSValue(value js.Value) BufferGeometry {
	return &bufferGeometryImpl{
		Value: value,
	}
}

func (c *bufferGeometryImpl) SetAttribute(name string, attribute *BufferAttribute) {

	c.Call("setAttribute", name, attribute.JSValue())

}

func (c *bufferGeometryImpl) Dispose() {
	c.Call("dispose")
}
