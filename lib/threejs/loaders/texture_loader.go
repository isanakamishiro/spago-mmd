package loaders

import (
	"app/lib/threejs"
	"syscall/js"
)

// TextureLoader is class for loading a texture.
// This uses the ImageLoader internally for loading files.
type TextureLoader interface {
	threejs.Loader

	Load(url string, onLoad js.Func, onProgress js.Func, onError js.Func) threejs.Texture
	LoadSimply(url string) threejs.Texture
}

type textureLoaderImp struct {
	threejs.Loader
}

// NewTextureLoader creates a new TextureLoader.
func NewTextureLoader() TextureLoader {
	return &textureLoaderImp{
		Loader: threejs.NewDefaultLoaderFromJSValue(threejs.GetJsObject("TextureLoader").New()),
	}
}

// NewTextureLoaderWithManager creates a new TextureLoader with LoadingManager.
func NewTextureLoaderWithManager(m threejs.LoadingManager) TextureLoader {
	return &textureLoaderImp{
		Loader: threejs.NewDefaultLoaderFromJSValue(threejs.GetJsObject("TextureLoader").New(m.JSValue())),
	}
}

// Load loads texture.
// url — the path or URL to the file. This can also be a Data URI.
// onLoad — Will be called when load completes. The argument will be the loaded texture.
// onProgress — Will be called while load progresses. The argument will be the XMLHttpRequest instance, which contains .total and .loaded bytes.
// onError — Will be called when load errors.
//
// Begin loading from the given URL and pass the fully loaded texture to onLoad.
// The method also returns a new texture object which can directly be used for material creation.
// If you do it this way, the texture may pop up in your scene once the respective loading process is finished.
func (c *textureLoaderImp) Load(url string, onLoad js.Func, onProgress js.Func, onError js.Func) threejs.Texture {
	return threejs.NewDefaultTextureFromJSValue(
		c.JSValue().Call("load", url, onLoad, onProgress, onError),
	)
}

// LoadSimply loads texture.
// url — the path or URL to the file. This can also be a Data URI.
//
// Begin loading from the given URL and pass the fully loaded texture to onLoad.
// The method also returns a new texture object which can directly be used for material creation.
// If you do it this way, the texture may pop up in your scene once the respective loading process is finished.
func (c *textureLoaderImp) LoadSimply(url string) threejs.Texture {
	return threejs.NewDefaultTextureFromJSValue(
		c.JSValue().Call("load", url),
	)
}
