package threejs

import (
	"syscall/js"
)

// Layers is a Layers object assigns an Object3D to 1 or more of 32 layers numbered 0 to 31 -
// internally the layers are stored as a bit mask, and by default all Object3Ds are a member of layer 0.
// This can be used to control visibility - an object must share a layer with a camera to be visible when that camera's view is renderered.
// All classes that inherit from Object3D have an Object3D.layers property which is an instance of this class.
type Layers struct {
	js.Value
}

// NewLayers create a new Layers object, with membership initially set to layer 0.
func NewLayers() *Layers {
	return &Layers{Value: GetJsObject("Layers").New()}
}

// JSValue ...
func (ll *Layers) JSValue() js.Value {
	return ll.Value
}

// Mask ...
func (ll *Layers) Mask() float64 {
	return ll.Get("mask").Float()
}

// SetMask ...
func (ll *Layers) SetMask(v float64) {
	ll.Set("mask", v)
}

// Disable ...
func (ll *Layers) Disable(channel float64) {
	ll.Call("disable", channel)
}

// Enable ...
func (ll *Layers) Enable(channel float64) {
	ll.Call("enable", channel)
}

// Set2 ...
func (ll *Layers) Set2(channel float64) {
	ll.Call("set", channel)
}

// Test ...
func (ll *Layers) Test(layers *Layers) bool {
	return ll.Call("test", layers).Bool()
}

// Toggle ...
func (ll *Layers) Toggle(channel float64) {
	ll.Call("toggle", channel)
}
