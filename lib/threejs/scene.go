package threejs

import (
	"syscall/js"
)

// Scene is ...
type Scene interface {
	Object3D

	AutoUpdate() bool
	SetAutoUpdate(v bool)
	Background() js.Value
	SetBackground(v js.Value)
	SetBackgroundColor(c Color)

	Add(v Object3D)
	AddLight(v Light)
	AddMesh(v Mesh)

	// Fog gets a fog instance defining the type of fog that affects everything rendered in the scene. Default is null.
	Fog() FogBase
	// SetFog sets a fog instance defining the type of fog that affects everything rendered in the scene. Default is null.
	SetFog(v FogBase)
}

// sceneImpl extend: [Object3D]
type sceneImpl struct {
	Object3D
}

// NewScene is factory method for Scene.
func NewScene() Scene {

	return &sceneImpl{
		NewObject3DFromJSValue(
			Threejs("Scene").New(),
		),
	}
}

// AutoUpdate is ...
func (ss *sceneImpl) AutoUpdate() bool {
	return ss.JSValue().Get("autoUpdate").Bool()
}

// SetAutoUpdate is ...
func (ss *sceneImpl) SetAutoUpdate(v bool) {
	ss.JSValue().Set("autoUpdate", v)
}

// Background is ...
func (ss *sceneImpl) Background() js.Value {
	return ss.JSValue().Get("background")
}

// SetBackground is ...
func (ss *sceneImpl) SetBackground(v js.Value) {
	ss.JSValue().Set("background", v)
}

// SetBackgroundColor is ...
func (ss *sceneImpl) SetBackgroundColor(c Color) {
	ss.JSValue().Set("background", c.JSValue())
}

func (ss *sceneImpl) Add(v Object3D) {
	ss.JSValue().Call("add", v.JSValue())
}

func (ss *sceneImpl) AddLight(v Light) {
	ss.JSValue().Call("add", v.JSValue())
}

func (ss *sceneImpl) AddMesh(v Mesh) {
	ss.JSValue().Call("add", v.JSValue())
}

func (ss *sceneImpl) Fog() FogBase {
	return NewFogBaseFromJSValue(
		ss.JSValue().Get("fog"),
	)
}

func (ss *sceneImpl) SetFog(v FogBase) {
	ss.JSValue().Set("fog", v.JSValue())
}
