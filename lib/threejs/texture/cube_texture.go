package texture

import (
	"app/lib/threejs"
	"syscall/js"
)

// CubeTexture is a cube texture made up of six images.
type CubeTexture interface {
	threejs.Texture
}

type cubeTextureImp struct {
	threejs.Texture
}

// NewCubeTextureFromJSValue creates new CubeTexture from js.Value.
func NewCubeTextureFromJSValue(v js.Value) CubeTexture {
	return &cubeTextureImp{
		Texture: threejs.NewDefaultTextureFromJSValue(v),
	}
}
