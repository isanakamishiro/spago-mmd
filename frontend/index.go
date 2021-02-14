package main

import (
	"app/frontend/actions"
	"app/frontend/views"
	"syscall/js"

	"github.com/nobonobo/spago"
	"github.com/nobonobo/spago/dispatcher"
	"github.com/nobonobo/spago/router"
)

var (
	topView *views.Top = views.NewTop()

	currentView spago.Component = topView
)

func init() {
	dispatcher.Register(actions.Refresh, func(args ...interface{}) {
		spago.Rerender(currentView)
	})

	loadScript("./assets/threejs/ex/js/libs/ammo.wasm.js")

}

func main() {

	// spago.VerboseMode = true

	r := router.New()
	r.Handle("/", func(key string) {
		spago.SetTitle("Top")
		currentView = topView
		spago.RenderBody(topView)
	})

	if err := r.Start(); err != nil {
		println(err)
		spago.RenderBody(router.NotFoundPage())
	}

	select {}
}

// loadScript synchronous javascript loader
func loadScript(url string) {

	document := js.Global().Get("document")

	ch := make(chan bool)
	script := document.Call("createElement", "script")
	script.Set("src", url)
	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer fn.Release()
		close(ch)
		return nil
	})
	script.Call("addEventListener", "load", fn)
	document.Get("head").Call("appendChild", script)
	<-ch
}
