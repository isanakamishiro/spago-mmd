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

	// AutoRotate gets to true to automatically rotate around the target.
	// Note that if this is enabled, you must call .update () in your animation loop.
	AutoRotate() bool

	// SetAutoRotate sets to true to automatically rotate around the target.
	// Note that if this is enabled, you must call .update () in your animation loop.
	SetAutoRotate(b bool)

	// AutoRotateSpeed gets how fast to rotate around the target if .autoRotate is true. Default is 2.0, which equates to 30 seconds per orbit at 60fps.
	// Note that if .autoRotate is enabled, you must call .update () in your animation loop.
	AutoRotateSpeed() float64

	// SetAutoRotateSpeed sets how fast to rotate around the target if .autoRotate is true. Default is 2.0, which equates to 30 seconds per orbit at 60fps.
	// Note that if .autoRotate is enabled, you must call .update () in your animation loop.
	SetAutoRotateSpeed(v float64)

	// EnableDamping gets to true to enable damping (inertia), which can be used to give a sense of weight to the controls. Default is false.
	// Note that if this is enabled, you must call .update () in your animation loop.
	EnableDamping() bool

	// SetEnableDamping sets to true to enable damping (inertia), which can be used to give a sense of weight to the controls. Default is false.
	// Note that if this is enabled, you must call .update () in your animation loop.
	SetEnableDamping(e bool)

	// SetEnablePan sets enable or disable camera panning. Default is true.
	SetEnablePan(b bool)

	// SetEnableRotate sets enable or disable horizontal and vertical rotation of the camera. Default is true.
	// Note that it is possible to disable a single axis by setting the min and max of the polar angle or azimuth angle to the same value, which will cause the vertical or horizontal rotation to be fixed at that value.
	SetEnableRotate(b bool)

	// SetEnableZoom sets enable or disable zooming (dollying) of the camera.
	SetEnableZoom(b bool)

	// SetMaxAzimuthAngle sets how far you can orbit horizontally, upper limit. If set, the interval [ min, max ] must be a sub-interval of [ - 2 PI, 2 PI ], with ( max - min < 2 PI ). Default is Infinity.
	SetMaxAzimuthAngle(v float64)

	// SetMaxDistance sets how far you can dolly out ( PerspectiveCamera only ). Default is Infinity.
	SetMaxDistance(v float64)

	// SetMaxPolarAngle sets how far you can orbit vertically, upper limit. Range is 0 to Math.PI radians, and default is Math.PI.
	SetMaxPolarAngle(v float64)

	// SetMaxZoom sets how far you can zoom out ( OrthographicCamera only ). Default is Infinity.
	SetMaxZoom(v float64)

	// SetMinAzimuthAngle sets how far you can orbit horizontally, lower limit. If set, the interval [ min, max ] must be a sub-interval of [ - 2 PI, 2 PI ], with ( max - min < 2 PI ). Default is Infinity.
	SetMinAzimuthAngle(v float64)

	// SetMinDistance sets how far you can dolly in ( PerspectiveCamera only ). Default is 0.
	SetMinDistance(v float64)

	// SetMinPolarAngle sets how far you can orbit vertically, lower limit. Range is 0 to Math.PI radians, and default is 0.
	SetMinPolarAngle(v float64)

	// SetMinZoom sets how far you can zoom in ( OrthographicCamera only ). Default is 0.
	SetMinZoom(v float64)

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

func (c *orbitControlsImp) SetAutoRotate(b bool) {
	c.Set("autoRotate", b)
}

func (c *orbitControlsImp) SetAutoRotateSpeed(v float64) {
	c.Set("autoRotateSpeed", v)
}

// AutoRotate gets to true to automatically rotate around the target.
// Note that if this is enabled, you must call .update () in your animation loop.
func (c *orbitControlsImp) AutoRotate() bool {
	return c.Get("autoRotate").Bool()
}

// AutoRotateSpeed gets how fast to rotate around the target if .autoRotate is true. Default is 2.0, which equates to 30 seconds per orbit at 60fps.
// Note that if .autoRotate is enabled, you must call .update () in your animation loop.
func (c *orbitControlsImp) AutoRotateSpeed() float64 {
	return c.Get("autoRotateSpeed").Float()
}

// SetEnablePan sets enable or disable camera panning. Default is true.
func (c *orbitControlsImp) SetEnablePan(b bool) {
	c.Set("enablePan", b)
}

// SetEnableRotate sets enable or disable horizontal and vertical rotation of the camera. Default is true.
// Note that it is possible to disable a single axis by setting the min and max of the polar angle or azimuth angle to the same value, which will cause the vertical or horizontal rotation to be fixed at that value.
func (c *orbitControlsImp) SetEnableRotate(b bool) {
	c.Set("enableRotate", b)
}

// SetEnableZoom sets enable or disable zooming (dollying) of the camera.
func (c *orbitControlsImp) SetEnableZoom(b bool) {
	c.Set("enableZoom", b)
}

// SetMaxAzimuthAngle sets how far you can orbit horizontally, upper limit. If set, the interval [ min, max ] must be a sub-interval of [ - 2 PI, 2 PI ], with ( max - min < 2 PI ). Default is Infinity.
func (c *orbitControlsImp) SetMaxAzimuthAngle(v float64) {
	c.Set("maxAzimuthAngle", v)
}

// SetMaxDistance sets how far you can dolly out ( PerspectiveCamera only ). Default is Infinity.
func (c *orbitControlsImp) SetMaxDistance(v float64) {
	c.Set("maxDistance", v)
}

// SetMaxPolarAngle sets how far you can orbit vertically, upper limit. Range is 0 to Math.PI radians, and default is Math.PI.
func (c *orbitControlsImp) SetMaxPolarAngle(v float64) {
	c.Set("maxPolarAngle", v)
}

// SetMaxZoom sets how far you can zoom out ( OrthographicCamera only ). Default is Infinity.
func (c *orbitControlsImp) SetMaxZoom(v float64) {
	c.Set("maxZoom", v)
}

// SetMinAzimuthAngle sets how far you can orbit horizontally, lower limit. If set, the interval [ min, max ] must be a sub-interval of [ - 2 PI, 2 PI ], with ( max - min < 2 PI ). Default is Infinity.
func (c *orbitControlsImp) SetMinAzimuthAngle(v float64) {
	c.Set("minAzimuthAngle", v)
}

// SetMinDistance sets how far you can dolly in ( PerspectiveCamera only ). Default is 0.
func (c *orbitControlsImp) SetMinDistance(v float64) {
	c.Set("minDistance", v)
}

// SetMinPolarAngle sets how far you can orbit vertically, lower limit. Range is 0 to Math.PI radians, and default is 0.
func (c *orbitControlsImp) SetMinPolarAngle(v float64) {
	c.Set("minPolarAngle", v)
}

// SetMinZoom sets how far you can zoom in ( OrthographicCamera only ). Default is 0.
func (c *orbitControlsImp) SetMinZoom(v float64) {
	c.Set("minZoom", v)
}
