package threejs

import (
	"errors"
	"syscall/js"
)

// Skeleton uses an array of bones to create a skeleton that can be used by a SkinnedMesh.
type Skeleton interface {
	JSValue() js.Value

	// BoneTexture gets the DataTexture holding the bone data when using a vertex texture.
	BoneTexture() (Texture, error)

	// BoneTextureSize gets The size of the .boneTexture.
	BoneTextureSize() int

	// Pose returns the skeleton to the base pose.
	Pose()

	// Update updates the boneMatrices and boneTexture after changing the bones.
	// This is called automatically by the WebGLRenderer if the skeleton is used with a SkinnedMesh.
	Update()

	// Dispose can be used if an instance of Skeleton becomes obsolete in an application.
	// The method will free internal resources.
	Dispose()
}

type skeletonImp struct {
	js.Value
}

// NewSkeletonFromJSValue creates default implementation of Skeleton from js.Value.
func NewSkeletonFromJSValue(v js.Value) Skeleton {
	return &skeletonImp{
		Value: v,
	}
}

// func (c *skeletonImp) JSValue() js.Value {
// 	panic("not implemented") // TODO: Implement
// }

// BoneTexture gets the DataTexture holding the bone data when using a vertex texture.
func (c *skeletonImp) BoneTexture() (Texture, error) {

	tx := c.Get("boneTexture")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("bone texture is null or undefined")
	}

	return NewDefaultTextureFromJSValue(tx), nil

}

// BoneTextureSize gets The size of the .boneTexture.
func (c *skeletonImp) BoneTextureSize() int {
	return c.Get("boneTextureSize").Int()
}

// Pose returns the skeleton to the base pose.
func (c *skeletonImp) Pose() {
	c.Call("pose")
}

// Update updates the boneMatrices and boneTexture after changing the bones.
// This is called automatically by the WebGLRenderer if the skeleton is used with a SkinnedMesh.
func (c *skeletonImp) Update() {
	c.Call("update")
}

// Dispose can be used if an instance of Skeleton becomes obsolete in an application.
// The method will free internal resources.
func (c *skeletonImp) Dispose() {
	c.Call("dispose")
}
