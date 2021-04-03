package threejs

import (
	"errors"
	"syscall/js"
)

// Mesh is class representing triangular polygon mesh based objects.
// Also serves as a base for other classes such as SkinnedMesh.
type Mesh interface {
	Object3D

	// Geometry gets an instance of Geometry or BufferGeometry (or derived classes),
	// defining the object's structure.
	// It's recommended to always use a BufferGeometry if possible for best performance.
	Geometry() (BufferGeometry, error)

	// Materials gets an instance of material derived from the Materials base class
	// or an array of materials, defining the object's appearance.
	// Default is a MeshBasicMaterial.
	Materials() ([]Material, error)

	// DisposeAll dispose all geometry, material, textures for this mesh.
	DisposeAll()
}

// MeshImpl extend: [Object3D]
type meshImpl struct {
	Object3D
}

// NewMesh is factory method for MeshImpl.
func NewMesh(geometry BufferGeometry, material Material) Mesh {
	return &meshImpl{
		NewObject3DFromJSValue(Threejs("Mesh").New(geometry.JSValue(), material.JSValue())),
	}
}

// NewMeshWithMultiMaterial is factory method for MeshImpl.
func NewMeshWithMultiMaterial(geometry BufferGeometry, materialSlice []Material) Mesh {

	var a []interface{} = make([]interface{}, len(materialSlice))
	for i, v := range materialSlice {
		a[i] = v.JSValue()
	}

	return &meshImpl{
		NewObject3DFromJSValue(Threejs("Mesh").New(geometry.JSValue(), a)),
	}
}

// NewMeshFromJSValue is factory method for MeshImpl.
func NewMeshFromJSValue(value js.Value) Mesh {
	return &meshImpl{
		NewObject3DFromJSValue(value),
	}
}

func (c *meshImpl) Geometry() (BufferGeometry, error) {

	geom := c.JSValue().Get("geometry")
	if geom.IsUndefined() || geom.IsNull() {
		return nil, errors.New("geometry is null or undefined")
	}
	return NewBufferGeometryFromJSValue(geom), nil
}

func (c *meshImpl) Materials() ([]Material, error) {

	mat := c.JSValue().Get("material")
	if mat.IsUndefined() || mat.IsNull() {
		return nil, errors.New("material is null or undefined")
	}

	// Not array
	length := mat.Length()
	if length == 0 {
		return []Material{NewDefaultMaterialFromJSValue(mat)}, nil
	}

	var mats []Material = make([]Material, length)
	for i := 0; i < length; i++ {
		mats[i] = NewDefaultMaterialFromJSValue(mat.Index(i))
	}
	return mats, nil

}

func (c *meshImpl) DisposeAll() {

	// Dispose Geometry
	if g, err := c.Geometry(); err == nil {
		g.Dispose()
	}
	// Dispose Material/Texture
	if m, err := c.Materials(); err == nil {
		for _, v := range m {

			// Dispose textures
			if tx, err := v.Map(); err == nil {
				tx.Dispose()
			}
			if tx, err := v.EnvMap(); err == nil {
				tx.Dispose()
			}
			if tx, err := v.GradientMap(); err == nil {
				tx.Dispose()
			}
			if tx, err := v.AlphaMap(); err == nil {
				tx.Dispose()
			}
			if tx, err := v.AOMap(); err == nil {
				tx.Dispose()
			}
			if tx, err := v.BumpMap(); err == nil {
				tx.Dispose()
			}
			if tx, err := v.EmissiveMap(); err == nil {
				tx.Dispose()
			}
			if tx, err := v.LightMap(); err == nil {
				tx.Dispose()
			}
			if tx, err := v.NormalMap(); err == nil {
				tx.Dispose()
			}

			// Dispose material
			v.Dispose()
		}
	}

}
