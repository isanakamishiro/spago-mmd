package controls

import (
	"app/lib/threejs"
	"log"
	"syscall/js"
)

const orbitControlModulePath = "./assets/threejs/ex/jsm/controls/OrbitControls.js"

var orbitControlModule js.Value

func init() {

	m := threejs.LoadModule([]string{"OrbitControls"}, orbitControlModulePath)
	if len(m) == 0 {
		log.Fatal("OrbitControls module could not be loaded.")
	}
	orbitControlModule = m[0]
}

// OrbitControls allow the camera to orbit around a target.
// To use this, as with all files in the /examples directory,
// you will have to include the file seperately in your HTML.
type OrbitControls interface {
	JSValue() js.Value

	// Target is the focus point of the controls,
	// the .object orbits around this.
	// It can be updated manually at any point to change the focus of the controls.
	Target() *threejs.Vector3

	// Update updates the controls.
	// Must be called after any manual changes to the camera's transform,
	// or in the update loop if .autoRotate or .enableDamping are set.
	Update() bool

	// Reset the controls to their state from either the last time the .saveState was called, or the initial state.
	Reset()

	// Save the current state of the controls. This can later be recovered with .reset.
	SaveState()

	// EnableDamping gets to true to enable damping (inertia), which can be used to give a sense of weight to the controls. Default is false.
	// Note that if this is enabled, you must call .update () in your animation loop.
	EnableDamping() bool

	// SetEnableDamping sets to true to enable damping (inertia), which can be used to give a sense of weight to the controls. Default is false.
	// Note that if this is enabled, you must call .update () in your animation loop.
	SetEnableDamping(e bool)

	AddEventListener(event string, fn js.Func)
	RemoveEventListener(event string, fn js.Func)

	Dispose()
}

// orbitControlsImp is implementation of OrbitControls
type orbitControlsImp struct {
	js.Value
	// orbitControlModule js.Value
}

// NewOrbitControls creates OrbitControls.
// camera: (required) The camera to be controlled.
// The camera must not be a child of another object, unless that object is the scene itself.
//
// domElement: The HTML element used for event listeners.
func NewOrbitControls(camera threejs.Camera, domElement js.Value) OrbitControls {

	// m := spago.LoadModule([]string{"OrbitControls"}, orbitControlModulePath)
	// if len(m) == 0 {
	// 	log.Fatal("OrbitControls module could not be loaded.")
	// }

	return &orbitControlsImp{
		Value: orbitControlModule.New(camera.JSValue(), domElement),
	}
}

// JSValue is ...
func (c *orbitControlsImp) JSValue() js.Value {
	return c.Value
}

// Target is the focus point of the controls,
// the .object orbits around this.
// It can be updated manually at any point to change the focus of the controls.
func (c *orbitControlsImp) Target() *threejs.Vector3 {
	return threejs.NewVector3FromJSValue(
		c.Get("target"),
	)
}

// Update updates the controls.
// Must be called after any manual changes to the camera's transform,
// or in the update loop if .autoRotate or .enableDamping are set.
func (c *orbitControlsImp) Update() bool {
	return c.Call("update").Bool()
}

// Reset the controls to their state from either the last time the .saveState was called, or the initial state.
func (c *orbitControlsImp) Reset() {
	c.Call("reset")
}

// Save the current state of the controls. This can later be recovered with .reset.
func (c *orbitControlsImp) SaveState() {
	c.Call("saveState")
}

// EnableDamping set to true to enable damping (inertia), which can be used to give a sense of weight to the controls. Default is false.
// Note that if this is enabled, you must call .update () in your animation loop.
func (c *orbitControlsImp) EnableDamping() bool {
	return c.Get("enableDamping").Bool()
}

// SetEnableDamping is ...
func (c *orbitControlsImp) SetEnableDamping(d bool) {
	c.Set("enableDamping", d)
}

// AddEventListener is ...
func (c *orbitControlsImp) AddEventListener(event string, fn js.Func) {
	c.Call("addEventListener", event, fn)
}

// RemoveEventListener is ...
func (c *orbitControlsImp) RemoveEventListener(event string, fn js.Func) {
	c.Call("removeEventListener", event, fn)
}

// Dispose remove all the event listeners.
func (c *orbitControlsImp) Dispose() {
	c.Call("dispose")
}
