package threejs

import (
	"syscall/js"
)

// Mesh is class representing triangular polygon mesh based objects.
// Also serves as a base for other classes such as SkinnedMesh.
type Mesh interface {
	Object3D

	// Geometry gets an instance of Geometry or BufferGeometry (or derived classes),
	// defining the object's structure.
	// It's recommended to always use a BufferGeometry if possible for best performance.
	Geometry() Geometry

	// Material gets an instance of material derived from the Material base class
	// or an array of materials, defining the object's appearance.
	// Default is a MeshBasicMaterial.
	Material() Material
}

// MeshImpl extend: [Object3D]
type meshImpl struct {
	Object3D
}

// NewMesh is factory method for MeshImpl.
func NewMesh(geometry Geometry, material Material) Mesh {
	return &meshImpl{
		NewObject3DFromJSValue(GetJsObject("Mesh").New(geometry.JSValue(), material.JSValue())),
	}
}

// NewMeshWithMultiMaterial is factory method for MeshImpl.
func NewMeshWithMultiMaterial(geometry Geometry, materialSlice []Material) Mesh {

	var a []interface{} = make([]interface{}, len(materialSlice))
	for i, v := range materialSlice {
		a[i] = v.JSValue()
	}

	return &meshImpl{
		NewObject3DFromJSValue(GetJsObject("Mesh").New(geometry.JSValue(), a)),
	}
}

// NewMeshFromJSValue is factory method for MeshImpl.
func NewMeshFromJSValue(value js.Value) Mesh {
	return &meshImpl{
		NewObject3DFromJSValue(value),
	}
}

func (c *meshImpl) Geometry() Geometry {
	return NewDefaultGeometryFromJSValue(
		c.JSValue().Get("geometry"),
	)
}

func (c *meshImpl) Material() Material {
	return NewDefaultMaterialFromJSValue(
		c.JSValue().Get("material"),
	)
}
