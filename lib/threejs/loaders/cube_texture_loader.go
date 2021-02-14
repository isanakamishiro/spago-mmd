package loaders

import (
	"app/lib/threejs"
	"app/lib/threejs/textures"
	"syscall/js"
)

// CubeTextureLoader is class for loading a cube texture.
// This uses the ImageLoader internally for loading files.
type CubeTextureLoader interface {
	threejs.Loader

	Load(urls []string, onLoad js.Func, onProgress js.Func, onError js.Func) textures.CubeTexture
	LoadSimply(urls []string) textures.CubeTexture
}

type cubeTextureLoaderImp struct {
	threejs.Loader
}

// NewCubeTextureLoader creates a new CubeTextureLoader.
func NewCubeTextureLoader() CubeTextureLoader {

	return &cubeTextureLoaderImp{
		Loader: threejs.NewDefaultLoaderFromJSValue(threejs.GetJsObject("CubeTextureLoader").New()),
	}
}

func (c *cubeTextureLoaderImp) Load(urls []string, onLoad js.Func, onProgress js.Func, onError js.Func) textures.CubeTexture {

	var jsImages []interface{} = make([]interface{}, len(urls))
	for i, url := range urls {
		if i > 5 {
			break // index を最大 0 - 5 の範囲でデータ入れる
		}
		jsImages[i] = url
	}

	return textures.NewCubeTextureFromJSValue(
		c.JSValue().Call("load", jsImages, onLoad, onProgress, onError),
	)
}

func (c *cubeTextureLoaderImp) LoadSimply(urls []string) textures.CubeTexture {

	var jsImages []interface{} = make([]interface{}, len(urls))
	for i, url := range urls {
		if i > 5 {
			break // index を最大 0 - 5 の範囲でデータ入れる
		}
		jsImages[i] = url
	}

	return textures.NewCubeTextureFromJSValue(
		c.JSValue().Call("load", jsImages),
	)
}
