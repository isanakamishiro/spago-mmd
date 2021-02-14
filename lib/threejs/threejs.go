package threejs

import (
	"fmt"
	"strings"
	"syscall/js"
)

// RenderID is ...
type RenderID js.Value

const modulePath = "./assets/threejs/build/three.module.js"

var module js.Value

func init() {

	// module = spago.LoadModuleAs("THREE", modulePath)
	module = LoadModuleAs("THREE", modulePath)
}

// LoadModule equivalent `import {'name1', 'name2', ...} from 'url'`
func LoadModule(names []string, url string) []js.Value {

	document := js.Global().Get("document")

	ch := make(chan js.Value, len(names))
	var sendFunc js.Func
	sendFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- args[0]
		return nil
	})
	var closeFunc js.Func
	closeFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer sendFunc.Release()
		defer closeFunc.Release()
		close(ch)
		return nil
	})
	js.Global().Set("__threejs_send__", sendFunc)
	js.Global().Set("__threejs_close__", closeFunc)

	lines := []string{}
	lines = append(lines, fmt.Sprintf("import { %s } from %q;", strings.Join(names, ", "), url))
	for _, name := range names {
		lines = append(lines, fmt.Sprintf("__threejs_send__(%s);", name))
	}
	lines = append(lines, "__threejs_close__();")

	script := document.Call("createElement", "script")
	script.Set("type", "module")
	code := document.Call("createTextNode", strings.Join(lines, "\n"))
	script.Call("appendChild", code)

	document.Get("head").Call("appendChild", script)

	res := make([]js.Value, 0, len(names))
	for v := range ch {
		res = append(res, v)
	}
	return res
}

// LoadModuleAs equivalent `import * as 'name' from 'url'`
func LoadModuleAs(name string, url string) js.Value {

	document := js.Global().Get("document")

	ch := make(chan js.Value, 1)

	var sendFunc js.Func
	sendFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ch <- args[0]

		return nil
	})
	var closeFunc js.Func
	closeFunc = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer sendFunc.Release()
		defer closeFunc.Release()
		close(ch)
		return nil
	})
	js.Global().Set("__threejs_send__", sendFunc)
	js.Global().Set("__threejs_close__", closeFunc)
	lines := []string{}
	lines = append(lines, fmt.Sprintf("import * as %s from %q;", name, url))
	lines = append(lines, fmt.Sprintf("__threejs_send__(%s);", name))
	lines = append(lines, "__threejs_close__();")

	script := document.Call("createElement", "script")
	script.Set("type", "module")
	code := document.Call("createTextNode", strings.Join(lines, "\n"))
	script.Call("appendChild", code)

	document.Get("head").Call("appendChild", script)

	// chanのループ処理で、closeが呼ばれるまで待つ
	res := make([]js.Value, 0, 1)
	for v := range ch {
		res = append(res, v)
	}
	return res[0] // 最初の結果のみ
}

// GetJsObject is getter for JavaScript object of ThreeJs.
func GetJsObject(key string) js.Value {
	return module.Get(key)
}

// setJsObject is setter for JavaScript object of ThreeJs.
func setJsObject(key string, v interface{}) {
	module.Set(key, v)
}

// RequestAnimationFrame is ...
func RequestAnimationFrame(render js.Func) RenderID {
	return RenderID(js.Global().Call("requestAnimationFrame", render))
}

// CancelAnimationFrame is ...
func CancelAnimationFrame(renderID RenderID) {
	js.Global().Call("cancelAnimationFrame", js.Value(renderID))
}
