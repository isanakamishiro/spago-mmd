package threejs

import (
	"syscall/js"
)

// Loader is base class for implementing loaders.
type Loader interface {
	JSValue() js.Value

	SetPath(path string)
}

// LoadingManager handles and keeps track of loaded and pending data.
// A default global instance of this class is created and used by loaders
// if not supplied manually - see DefaultLoadingManager.
//
// In general that should be sufficient,
// however there are times when it can be useful to have seperate loaders
// - for example if you want to show seperate loading bars for objects and textures.
type LoadingManager interface {
	JSValue() js.Value

	SetOnStart(fn func(url string, itemsLoaded int, itemsTotal int))
	SetOnProgress(fn func(url string, itemsLoaded int, itemsTotal int))
	SetOnLoad(fn func())
	SetOnError(fn func(url string))
}

type defaultLoaderImp struct {
	js.Value
}

type defaultLoadingManagerImp struct {
	js.Value

	onStart    func(url string, itemsLoaded int, itemsTotal int)
	onProgress func(url string, itemsLoaded int, itemsTotal int)
	onLoad     func()
	onError    func(url string)
}

// NewDefaultLoaderFromJSValue is constructor.
func NewDefaultLoaderFromJSValue(v js.Value) Loader {
	return &defaultLoaderImp{
		Value: v,
	}
}

// NewLoadingManager is constructor.
func NewLoadingManager() LoadingManager {

	c := &defaultLoadingManagerImp{
		Value: GetJsObject("LoadingManager").New(),
	}

	// 初期化時にすべてのコールバックを設定する

	var jsfnOnStart js.Func
	jsfnOnStart = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		defer jsfnOnStart.Release()

		url := args[0].String()
		itemsLoaded := args[1].Int()
		itemsTotal := args[2].Int()

		if c.onStart != nil {
			c.onStart(url, itemsLoaded, itemsTotal)
		}

		return nil
	})
	var jsfnOnProgress js.Func
	jsfnOnProgress = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		url := args[0].String()
		itemsLoaded := args[1].Int()
		itemsTotal := args[2].Int()

		if c.onProgress != nil {
			c.onProgress(url, itemsLoaded, itemsTotal)
		}

		return nil
	})
	var jsfnOnLoad js.Func
	jsfnOnLoad = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		// defer jsfnOnStart.Release()
		defer jsfnOnLoad.Release()
		defer jsfnOnProgress.Release()

		if c.onLoad != nil {
			c.onLoad()
		}

		return nil
	})
	var jsfnOnError js.Func
	jsfnOnError = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		// defer jsfnOnStart.Release()
		defer jsfnOnError.Release()
		defer jsfnOnLoad.Release()
		defer jsfnOnProgress.Release()

		url := args[0].String()
		if c.onError != nil {
			c.onError(url)
		}

		return nil
	})

	c.Set("onStart", jsfnOnStart)
	c.Set("onProgress", jsfnOnProgress)
	c.Set("onLoad", jsfnOnLoad)
	c.Set("onError", jsfnOnError)

	return c
}

func (c *defaultLoaderImp) SetPath(path string) {
	c.JSValue().Call("setPath", path)
}

// SetOnStart is ...
func (c *defaultLoadingManagerImp) SetOnStart(fn func(url string, itemsLoaded int, itemsTotal int)) {
	c.onStart = fn
}

// SetOnProgress is ...
func (c *defaultLoadingManagerImp) SetOnProgress(fn func(url string, itemsLoaded int, itemsTotal int)) {
	c.onProgress = fn
}

// SetOnLoad is ...
func (c *defaultLoadingManagerImp) SetOnLoad(fn func()) {
	c.onLoad = fn
}

// SetOnError is ...
func (c *defaultLoadingManagerImp) SetOnError(fn func(url string)) {
	c.onError = fn
}
