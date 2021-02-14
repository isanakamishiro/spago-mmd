package mmdloaders

import (
	"app/lib/threejs"
	"errors"
	"log"
	"syscall/js"
)

const (
	mmdLoaderModulePath = "./assets/threejs/ex/jsm/loaders/MMDLoader.js"
)

var (
	mmdLoaderModule js.Value
)

func init() {

	m := threejs.LoadModule([]string{"MMDLoader"}, mmdLoaderModulePath)
	if len(m) == 0 {
		log.Fatal("MMDLoader module could not be loaded.")
	}
	mmdLoaderModule = m[0]
}

// MMDLoader creates Three.js Objects from MMD resources as PMD, PMX, VMD, and VPD files. See MMDAnimationHelper for MMD animation handling as IK, Grant, and Physics.
//
// If you want raw content of MMD resources, use .loadPMD/PMX/VMD/VPD methods.
type MMDLoader interface {
	threejs.Loader

	// Load begin loading PMD/PMX model file from url and fire the callback function with the parsed SkinnedMesh containing BufferGeometry and an array of MeshToonMaterial.
	//
	// url — A string containing the path/URL of the .pmd or .pmx file.
	// onLoad — A function to be called after the loading is successfully completed.
	// onProgress — (optional) A function to be called while the loading is in progress. The argument will be the XMLHttpRequest instance, that contains .total and .loaded bytes.
	// onError — (optional) A function to be called if an error occurs during loading. The function receives error as an argument.
	Load(url string, onLoad func(mesh MMDMesh), onProgress func(loadedBytes int, totalBytes int), onError func(err error))

	// LoadWithAnimation begin loading PMD/PMX model file and VMD motion file(s) from urls and fire the callback function with an Object containing parsed SkinnedMesh and AnimationClip fitting to the SkinnedMesh.
	//
	// modelURL — A string containing the path/URL of the .pmd or .pmx file.
	// vmdURLs — an array of string containing the path/URL of the .vmd file(s).
	// onLoad — A function to be called after the loading is successfully completed.
	// onProgress — (optional) A function to be called while the loading is in progress. The argument will be the XMLHttpRequest instance, that contains .total and .loaded bytes.
	// onError — (optional) A function to be called if an error occurs during loading. The function receives error as an argument.
	LoadWithAnimation(modelURL string, vmdURLs []string, onLoad func(mesh MMDMesh, animation MMDAnimation), onProgress func(loadedBytes int, totalBytes int), onError func(err error))

	// LoadCameraAnimation begin loading VMD motion file(s) from url(s) and fire the callback function with the parsed AnimationClip.
	//
	// urls — A string or an array of string containing the path/URL of the .vmd file(s).If two or more files are specified, they'll be merged.
	// camera — Clip and its tracks will be fitting to this object.
	// onLoad — A function to be called after the loading is successfully completed.
	// onProgress — (optional) A function to be called while the loading is in progress. The argument will be the XMLHttpRequest instance, that contains .total and .loaded bytes.
	// onError — (optional) A function to be called if an error occurs during loading. The function receives error as an argument.
	LoadCameraAnimation(urls []string, camera threejs.Camera, onLoad func(cameraAnimation MMDAnimation), onProgress func(loadedBytes int, totalBytes int), onError func(err error))
}

type mmdLoaderImp struct {
	threejs.Loader
}

// NewMMDLoader creates a new MMDLoader.
func NewMMDLoader() MMDLoader {
	return &mmdLoaderImp{
		Loader: threejs.NewDefaultLoaderFromJSValue(mmdLoaderModule.New()),
	}
}

// NewMMDLoaderWithManager creates a new MMDLoader with MMDLoadingManager.
func NewMMDLoaderWithManager(manager MMDLoadingManager) MMDLoader {
	return &mmdLoaderImp{
		Loader: threejs.NewDefaultLoaderFromJSValue(mmdLoaderModule.New(manager.JSValue())),
	}
}

/*

	Methods

*/

// Load begin loading PMD/PMX model file from url and fire the callback function with the parsed SkinnedMesh containing BufferGeometry and an array of MeshToonMaterial.
//
// url — A string containing the path/URL of the .pmd or .pmx file.
// onLoad — A function to be called after the loading is successfully completed.
// onProgress — (optional) A function to be called while the loading is in progress. The argument will be the XMLHttpRequest instance, that contains .total and .loaded bytes.
// onError — (optional) A function to be called if an error occurs during loading. The function receives error as an argument.
func (c *mmdLoaderImp) Load(url string, onLoad func(mesh MMDMesh), onProgress func(loadedBytes int, totalBytes int), onError func(err error)) {

	// すべてのコールバックを設定する
	var jsfnOnProgress js.Func
	jsfnOnProgress = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		xhr := args[0]
		loadedBytes := xhr.Get("loaded").Int()
		totalBytes := xhr.Get("total").Int()

		if onProgress != nil {
			onProgress(loadedBytes, totalBytes)
		}

		return nil
	})

	var jsfnOnLoad js.Func
	jsfnOnLoad = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		defer jsfnOnLoad.Release()
		defer jsfnOnProgress.Release()

		mesh := newMMDMeshFromJSValue(args[0])

		if onLoad != nil {
			onLoad(mesh)
		}

		return nil
	})

	var jsfnOnError js.Func
	jsfnOnError = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		defer jsfnOnError.Release()
		defer jsfnOnLoad.Release()
		defer jsfnOnProgress.Release()

		errorMessage := args[0].Get("message").String()

		if onError != nil {
			onError(errors.New(errorMessage))
		}

		return nil
	})

	c.JSValue().Call("load", url, jsfnOnLoad, jsfnOnProgress, jsfnOnError)
}

// LoadWithAnimation begin loading PMD/PMX model file and VMD motion file(s) from urls and fire the callback function with an Object containing parsed SkinnedMesh and AnimationClip fitting to the SkinnedMesh.
//
// modelURL — A string containing the path/URL of the .pmd or .pmx file.
// vmdURLs — an array of string containing the path/URL of the .vmd file(s).
// onLoad — A function to be called after the loading is successfully completed.
// onProgress — (optional) A function to be called while the loading is in progress. The argument will be the XMLHttpRequest instance, that contains .total and .loaded bytes.
// onError — (optional) A function to be called if an error occurs during loading. The function receives error as an argument.
func (c *mmdLoaderImp) LoadWithAnimation(modelURL string, vmdURLs []string, onLoad func(mesh MMDMesh, animation MMDAnimation), onProgress func(loadedBytes int, totalBytes int), onError func(err error)) {

	// すべてのコールバックを設定する
	var jsfnOnProgress js.Func
	jsfnOnProgress = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		xhr := args[0]
		loadedBytes := xhr.Get("loaded").Int()
		totalBytes := xhr.Get("total").Int()

		if onProgress != nil {
			onProgress(loadedBytes, totalBytes)
		}

		return nil
	})

	var jsfnOnLoad js.Func
	jsfnOnLoad = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		defer jsfnOnLoad.Release()
		defer jsfnOnProgress.Release()

		mmd := args[0]
		mesh := threejs.NewMeshFromJSValue(mmd.Get("mesh"))
		animation := newMMDAnimationFromJSValue(mmd.Get("animation"))

		if onLoad != nil {
			onLoad(mesh, animation)
		}

		return nil
	})

	var jsfnOnError js.Func
	jsfnOnError = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		defer jsfnOnError.Release()
		defer jsfnOnLoad.Release()
		defer jsfnOnProgress.Release()

		errorMessage := args[0].Get("message").String()

		if onError != nil {
			onError(errors.New(errorMessage))
		}

		return nil
	})

	// type conversion
	var iVmdURLs []interface{} = make([]interface{}, len(vmdURLs))
	for i, d := range vmdURLs {
		iVmdURLs[i] = d
	}

	c.JSValue().Call("loadWithAnimation", modelURL, iVmdURLs, jsfnOnLoad, jsfnOnProgress, jsfnOnError)
}

// urls — A string or an array of string containing the path/URL of the .vmd file(s).If two or more files are specified, they'll be merged.
// camera — Clip and its tracks will be fitting to this object.
// onLoad — A function to be called after the loading is successfully completed.
// onProgress — (optional) A function to be called while the loading is in progress. The argument will be the XMLHttpRequest instance, that contains .total and .loaded bytes.
// onError — (optional) A function to be called if an error occurs during loading. The function receives error as an argument.
func (c *mmdLoaderImp) LoadCameraAnimation(urls []string, camera threejs.Camera, onLoad func(cameraAnimation MMDAnimation), onProgress func(loadedBytes int, totalBytes int), onError func(err error)) {

	// すべてのコールバックを設定する
	var jsfnOnProgress js.Func
	jsfnOnProgress = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		xhr := args[0]
		loadedBytes := xhr.Get("loaded").Int()
		totalBytes := xhr.Get("total").Int()

		if onProgress != nil {
			onProgress(loadedBytes, totalBytes)
		}

		return nil
	})

	var jsfnOnLoad js.Func
	jsfnOnLoad = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		defer jsfnOnLoad.Release()
		defer jsfnOnProgress.Release()

		cameraAnimation := newMMDAnimationFromJSValue(args[0])

		if onLoad != nil {
			onLoad(cameraAnimation)
		}

		return nil
	})

	var jsfnOnError js.Func
	jsfnOnError = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		defer jsfnOnError.Release()
		defer jsfnOnLoad.Release()
		defer jsfnOnProgress.Release()

		errorMessage := args[0].Get("message").String()

		if onError != nil {
			onError(errors.New(errorMessage))
		}

		return nil
	})

	// type conversion
	var iURLs []interface{} = make([]interface{}, len(urls))
	for i, d := range urls {
		iURLs[i] = d
	}

	c.JSValue().Call("loadAnimation", iURLs, camera.JSValue(), jsfnOnLoad, jsfnOnProgress, jsfnOnError)
}
