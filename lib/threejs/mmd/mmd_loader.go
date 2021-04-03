package mmd

import (
	"app/lib/threejs"
	"app/lib/threejs/animation"
	"context"
	"errors"
	"log"
	"sync"
	"syscall/js"
)

const (
	modulePath = "./assets/threejs/ex/jsm/loaders/MMDLoader.js"
)

var (
	module js.Value
)

func init() {

	m := threejs.LoadModule([]string{"MMDLoader"}, modulePath)
	if len(m) == 0 {
		log.Fatal("MMDLoader module could not be loaded.")
	}
	module = m[0]

}

// Loader creates Three.js Objects from MMD resources as PMD, PMX, VMD, and VPD files. See MMDAnimationHelper for MMD animation handling as IK, Grant, and Physics.
//
// If you want raw content of MMD resources, use .loadPMD/PMX/VMD/VPD methods.
type Loader interface {
	threejs.Loader

	// LoadModel begin loading PMD/PMX model file from url and fire the callback function with the parsed SkinnedMesh containing BufferGeometry and an array of MeshToonMaterial.
	//
	// ctx - Context
	// url — A string containing the path/URL of the .pmd or .pmx file.
	LoadModel(ctx context.Context, url string) <-chan FutureMesh

	// LoadWithAnimation begin loading PMD/PMX model file and VMD motion file(s) from urls and fire the callback function with an Object containing parsed SkinnedMesh and AnimationClip fitting to the SkinnedMesh.
	//
	// modelURL — A string containing the path/URL of the .pmd or .pmx file.
	// vmdURLs — an array of string containing the path/URL of the .vmd file(s).
	// onLoad — A function to be called after the loading is successfully completed.
	// onProgress — (optional) A function to be called while the loading is in progress. The argument will be the XMLHttpRequest instance, that contains .total and .loaded bytes.
	// onError — (optional) A function to be called if an error occurs during loading. The function receives error as an argument.
	LoadWithAnimation(modelURL string, vmdURLs []string, onLoad func(mesh threejs.SkinnedMesh, clip animation.Clip), onProgress func(loadedBytes int, totalBytes int), onError func(err error))

	// LoadCameraAnimation begin loading VMD motion file(s) from url(s) and fire the callback function with the parsed AnimationClip.
	//
	// urls — A string or an array of string containing the path/URL of the .vmd file(s).If two or more files are specified, they'll be merged.
	// camera — Clip and its tracks will be fitting to this object.
	LoadCameraAnimation(ctx context.Context, urls []string, camera threejs.Camera) <-chan FutureClip

	// LoadMotionAnimation begin loading VMD motion file(s) from url(s) and fire the callback function with the parsed AnimationClip.
	//
	// urls — A string or an array of string containing the path/URL of the .vmd file(s).If two or more files are specified, they'll be merged.
	// model — Clip and its tracks will be fitting to this object(SkinnedMesh).
	LoadMotionAnimation(ctx context.Context, urls []string, model threejs.Mesh) <-chan FutureClip

	// LoadVPD load vpd files.
	LoadVPDs(ctx context.Context, urls []string, isUnicode bool) <-chan FutureVpd
}

type mmdLoaderImp struct {
	threejs.Loader
}

// NewLoader creates a new MMDLoader.
func NewLoader() Loader {
	return &mmdLoaderImp{
		Loader: threejs.NewDefaultLoaderFromJSValue(module.New()),
	}
}

// NewLoaderWithManager creates a new MMDLoader with MMDLoadingManager.
func NewLoaderWithManager(manager threejs.LoadingManager) Loader {
	return &mmdLoaderImp{
		Loader: threejs.NewDefaultLoaderFromJSValue(module.New(manager.JSValue())),
	}
}

/*

	Methods

*/

func (c *mmdLoaderImp) LoadModel(ctx context.Context, url string) <-chan FutureMesh {

	ch := c.loadChannelGenerator(ctx, []string{url})
	pipeline := c.loadMMDModel(ctx, ch)

	return pipeline

}

// LoadWithAnimation begin loading PMD/PMX model file and VMD motion file(s) from urls and fire the callback function with an Object containing parsed SkinnedMesh and AnimationClip fitting to the SkinnedMesh.
//
// modelURL — A string containing the path/URL of the .pmd or .pmx file.
// vmdURLs — an array of string containing the path/URL of the .vmd file(s).
// onLoad — A function to be called after the loading is successfully completed.
// onProgress — (optional) A function to be called while the loading is in progress. The argument will be the XMLHttpRequest instance, that contains .total and .loaded bytes.
// onError — (optional) A function to be called if an error occurs during loading. The function receives error as an argument.
func (c *mmdLoaderImp) LoadWithAnimation(modelURL string, vmdURLs []string, onLoad func(mesh threejs.SkinnedMesh, clip animation.Clip), onProgress func(loadedBytes int, totalBytes int), onError func(err error)) {

	// すべてのコールバックを設定する
	var jsfnOnProgress js.Func = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

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
		mesh := threejs.NewSkinnedMeshFromJSValue(mmd.Get("mesh"))
		// animation := newAnimationFromJSValue(mmd.Get("animation"))
		clip, err := animation.NewClipFromJSValue(mmd.Get("animation"))
		if err != nil {
			if onError != nil {
				onError(errors.New("clip could not be loaded"))
			}

			return nil
		}

		if onLoad != nil {
			onLoad(mesh, clip)
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

// LoadCameraAnimation
// urls — A string or an array of string containing the path/URL of the .vmd file(s).If two or more files are specified, they'll be merged.
// camera — Clip and its tracks will be fitting to this object.
func (c *mmdLoaderImp) LoadCameraAnimation(ctx context.Context, urls []string, camera threejs.Camera) <-chan FutureClip {

	ch := c.loadChannelGenerator(ctx, urls)
	return c.loadAnimation(ctx, ch, camera.JSValue())

}

// LoadMotionAnimation
// urls — A string or an array of string containing the path/URL of the .vmd file(s).If two or more files are specified, they'll be merged.
// mesh — Clip and its tracks will be fitting to this object(SkinnedMesh).
func (c *mmdLoaderImp) LoadMotionAnimation(ctx context.Context, urls []string, model threejs.Mesh) <-chan FutureClip {

	ch := c.loadChannelGenerator(ctx, urls)
	return c.loadAnimation(ctx, ch, model.JSValue())

}

func (c *mmdLoaderImp) LoadVPDs(ctx context.Context, urls []string, isUnicode bool) <-chan FutureVpd {

	ch := c.loadChannelGenerator(ctx, urls)
	pipeline := c.loadVPD(ctx, ch, isUnicode)

	return pipeline

}

func (c *mmdLoaderImp) loadChannelGenerator(ctx context.Context, urls []string) <-chan string {

	ch := make(chan string)

	go func() {
		defer close(ch)

		for _, url := range urls {
			select {
			case <-ctx.Done():
				return
			case ch <- url:
			}
		}
	}()

	return ch
}

func (c *mmdLoaderImp) loadVPD(ctx context.Context, urlCh <-chan string, isUnicode bool) <-chan FutureVpd {

	result := make(chan FutureVpd)

	go func() {
		var wg sync.WaitGroup

		jsfnOnLoad := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			defer wg.Done()

			result <- NewFutureVpd(args[0], 0, 0, nil)
			return nil
		})

		jsfnOnProgress := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			xhr := args[0]
			loadedBytes := xhr.Get("loaded").Int()
			totalBytes := xhr.Get("total").Int()

			result <- NewFutureVpd(
				nil,
				uint(loadedBytes),
				uint(totalBytes),
				nil,
			)
			return nil
		})

		jsfnOnError := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			defer wg.Done()

			result <- NewFutureVpd(nil, 0, 0, errors.New(args[0].Get("message").String()))
			return nil
		})

		for url := range urlCh {

			select {
			case <-ctx.Done():
				return
			default:
				wg.Add(1)
				c.JSValue().Call("loadVPD", url, isUnicode, jsfnOnLoad, jsfnOnProgress, jsfnOnError)
			}
		}

		// 終了処理
		go func() {
			wg.Wait()

			log.Println("Release funcs in loadVPD.")

			jsfnOnLoad.Release()
			jsfnOnProgress.Release()
			jsfnOnError.Release()
			close(result)
		}()

	}()

	return result

}

func (c *mmdLoaderImp) loadMMDModel(ctx context.Context, urlCh <-chan string) <-chan FutureMesh {

	result := make(chan FutureMesh)

	go func() {
		var wg sync.WaitGroup

		jsfnOnLoad := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			defer wg.Done()

			result <- NewFutureMesh(
				threejs.NewSkinnedMeshFromJSValue(args[0]),
				0,
				0,
				nil)
			return nil
		})

		jsfnOnProgress := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			xhr := args[0]
			loadedBytes := xhr.Get("loaded").Int()
			totalBytes := xhr.Get("total").Int()

			result <- NewFutureMesh(
				nil,
				uint(loadedBytes),
				uint(totalBytes),
				nil,
			)
			return nil
		})

		jsfnOnError := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			defer wg.Done()

			result <- NewFutureMesh(nil, 0, 0, errors.New(args[0].Get("message").String()))
			return nil
		})

		for url := range urlCh {

			select {
			case <-ctx.Done():
				return
			default:
				wg.Add(1)
				c.JSValue().Call("load", url, jsfnOnLoad, jsfnOnProgress, jsfnOnError)
			}
		}

		// 終了処理
		go func() {
			wg.Wait()

			log.Println("Release funcs in loadMMDModel")

			jsfnOnLoad.Release()
			jsfnOnProgress.Release()
			jsfnOnError.Release()
			close(result)
		}()

	}()

	return result
}

func (c *mmdLoaderImp) loadAnimation(ctx context.Context, urlCh <-chan string, obj js.Value) <-chan FutureClip {

	result := make(chan FutureClip)

	go func() {
		var wg sync.WaitGroup

		jsfnOnLoad := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			defer wg.Done()

			clip, err := animation.NewClipFromJSValue(args[0])
			if err != nil {
				result <- NewFutureClip(
					nil, 0, 0,
					errors.New("file is loaded but clip creation is failed"),
				)

				return nil
			}

			result <- NewFutureClip(clip, 0, 0, nil)
			return nil
		})

		jsfnOnProgress := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			xhr := args[0]
			loadedBytes := xhr.Get("loaded").Int()
			totalBytes := xhr.Get("total").Int()

			result <- NewFutureClip(
				nil,
				uint(loadedBytes), uint(totalBytes),
				nil,
			)
			return nil
		})

		jsfnOnError := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			defer wg.Done()

			result <- NewFutureClip(
				nil, 0, 0,
				errors.New(args[0].Get("message").String()),
			)
			return nil
		})

		for url := range urlCh {

			select {
			case <-ctx.Done():
				return
			default:
				wg.Add(1)

				c.JSValue().Call("loadAnimation", url, obj, jsfnOnLoad, jsfnOnProgress, jsfnOnError)
			}
		}

		// 終了処理
		go func() {
			wg.Wait()

			log.Println("Release funcs in loadAnimation")

			jsfnOnLoad.Release()
			jsfnOnProgress.Release()
			jsfnOnError.Release()
			close(result)
		}()

	}()

	return result

}
