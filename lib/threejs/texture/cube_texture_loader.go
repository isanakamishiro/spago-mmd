package texture

import (
	"app/lib/threejs"
	"syscall/js"
)

// CubeLoader is class for loading a cube texture.
// This uses the ImageLoader internally for loading files.
type CubeLoader interface {
	threejs.Loader

	Load(urls []string, onLoad js.Func, onProgress js.Func, onError js.Func) CubeTexture
	LoadSimply(urls []string) CubeTexture
}

type cubeTextureLoaderImp struct {
	threejs.Loader
}

// NewCubeLoader creates a new CubeTextureLoader.
func NewCubeLoader() CubeLoader {

	return &cubeTextureLoaderImp{
		Loader: threejs.NewDefaultLoaderFromJSValue(threejs.Threejs("CubeTextureLoader").New()),
	}
}

func (c *cubeTextureLoaderImp) Load(urls []string, onLoad js.Func, onProgress js.Func, onError js.Func) CubeTexture {

	var jsImages []interface{} = make([]interface{}, len(urls))
	for i, url := range urls {
		if i > 5 {
			break // index を最大 0 - 5 の範囲でデータ入れる
		}
		jsImages[i] = url
	}

	return NewCubeTextureFromJSValue(
		c.JSValue().Call("load", jsImages, onLoad, onProgress, onError),
	)
}

func (c *cubeTextureLoaderImp) LoadSimply(urls []string) CubeTexture {

	var jsImages []interface{} = make([]interface{}, len(urls))
	for i, url := range urls {
		if i > 5 {
			break // index を最大 0 - 5 の範囲でデータ入れる
		}
		jsImages[i] = url
	}

	return NewCubeTextureFromJSValue(
		c.JSValue().Call("load", jsImages),
	)
}
