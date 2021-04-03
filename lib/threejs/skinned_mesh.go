package threejs

import (
	"errors"
	"syscall/js"
)

// SkinnedMesh is a mesh that has a Skeleton with bones that can then be used to animate the vertices of the geometry.
// The material must support skinning and have skinning enabled - see MeshStandardMaterial.skinning.
type SkinnedMesh interface {
	Mesh

	// BindMode gets Either "attached" or "detached". "attached" uses the SkinnedMesh.matrixWorld property
	// for the base transform matrix of the bones.
	// "detached" uses the SkinnedMesh.bindMatrix. Default is "attached".
	BindMode() string

	// Skeleton gets skeleton representing the bone hierarchy of the skinned mesh.
	Skeleton() (Skeleton, error)

	// Pose sets the skinned mesh in the rest pose (resets the pose).
	Pose()
}

type skinnedMeshImp struct {
	Mesh
}

// NewSkinnedMesh creates SkinnedMesh.
func NewSkinnedMesh(geometry BufferGeometry, material Material) SkinnedMesh {
	return &skinnedMeshImp{
		NewMeshFromJSValue(Threejs("SkinnedMesh").New(geometry.JSValue(), material.JSValue())),
	}
}

// NewSkinnedMeshWithMultiMaterial creates SkinnedMesh with multi materials.
func NewSkinnedMeshWithMultiMaterial(geometry BufferGeometry, materialSlice []Material) SkinnedMesh {

	var a []interface{} = make([]interface{}, len(materialSlice))
	for i, v := range materialSlice {
		a[i] = v.JSValue()
	}

	return &skinnedMeshImp{
		NewMeshFromJSValue(Threejs("SkinnedMesh").New(geometry.JSValue(), a)),
	}
}

// NewSkinnedMeshFromJSValue creates SkinnedMesh with js.Value.
func NewSkinnedMeshFromJSValue(value js.Value) SkinnedMesh {
	return &skinnedMeshImp{
		NewMeshFromJSValue(value),
	}
}

// BindMode gets Either "attached" or "detached". "attached" uses the SkinnedMesh.matrixWorld property
// for the base transform matrix of the bones.
// "detached" uses the SkinnedMesh.bindMatrix. Default is "attached".
func (c *skinnedMeshImp) BindMode() string {
	return c.JSValue().Get("bindMode").String()
}

// Skeleton gets skeleton representing the bone hierarchy of the skinned mesh.
func (c *skinnedMeshImp) Skeleton() (Skeleton, error) {
	s := c.JSValue().Get("skeleton")
	if s.IsUndefined() || s.IsNull() {
		return nil, errors.New("skeleton is null or undefined")
	}
	return NewSkeletonFromJSValue(s), nil

}

// Pose sets the skinned mesh in the rest pose (resets the pose).
func (c *skinnedMeshImp) Pose() {
	c.JSValue().Call("pose")
}

func (c *skinnedMeshImp) DisposeAll() {
	c.Mesh.DisposeAll()

	if s, err := c.Skeleton(); err == nil {
		if tx, err := s.BoneTexture(); err == nil {
			tx.Dispose()
		}
		s.Dispose()
	}
}
