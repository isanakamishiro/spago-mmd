package stats

import (
	"app/lib/threejs"
	"log"
	"syscall/js"
)

const statsModulePath = "./assets/threejs/ex/jsm/libs/stats.module.js"

var statsModule js.Value

func init() {

	m := threejs.LoadModule([]string{"Stats"}, statsModulePath)
	if len(m) == 0 {
		log.Fatal("Stats module could not be loaded.")
	}
	statsModule = m[0]
}

// Stats is JavaScript Performance Monitor
// This class provides a simple info box that will help you monitor your code performance.
//
// FPS Frames rendered in the last second. The higher the number the better.
// MS Milliseconds needed to render a frame. The lower the number the better.
// MB MBytes of allocated memory. (Run Chrome with --enable-precise-memory-info)
// CUSTOM User-defined panel support.
type Stats interface {
	JSValue() js.Value

	ShowPanel(panel int)
	DomElement() js.Value

	Begin()
	End()
	Update()
}

// statsImp is default implementation of GUI
type statsImp struct {
	js.Value
}

// NewStats is ...
func NewStats() Stats {
	return &statsImp{
		Value: statsModule.New(),
	}
}

// ShowPanel is ...
func (c *statsImp) ShowPanel(panel int) {
	c.Call("showPanel", panel)
}

// DomElement is ...
func (c *statsImp) DomElement() js.Value {
	return c.Get("dom")
}

// Begin is ...
func (c *statsImp) Begin() {
	c.Call("begin")
}

// End is ...
func (c *statsImp) End() {
	c.Call("end")
}

// Update is ...
func (c *statsImp) Update() {
	c.Call("update")
}
