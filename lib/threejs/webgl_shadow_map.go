package threejs

import (
	"syscall/js"
)

// WebGLShadowMap extend: []
type WebGLShadowMap struct {
	js.Value
}

// func NewWebGLShadowMap(_renderer *WebGLRenderer, _lights js.Value, _objects js.Value, capabilities js.Value) *WebGLShadowMap {
// 	return &WebGLShadowMap{Value: get("WebGLShadowMap").New(_renderer.JSValue(), _lights, _objects, capabilities)}
// }

// JSValue is ..
func (wglsm *WebGLShadowMap) JSValue() js.Value {
	return wglsm.Value
}

// AutoUpdate is ..
func (wglsm *WebGLShadowMap) AutoUpdate() bool {
	return wglsm.Get("autoUpdate").Bool()
}

// SetAutoUpdate is ..
func (wglsm *WebGLShadowMap) SetAutoUpdate(v bool) {
	wglsm.Set("autoUpdate", v)
}

// CullFace is ..
func (wglsm *WebGLShadowMap) CullFace() js.Value {
	return wglsm.Get("cullFace")
}

// SetCullFace is ..
func (wglsm *WebGLShadowMap) SetCullFace(v js.Value) {
	wglsm.Set("cullFace", v)
}

// Enabled is ..
func (wglsm *WebGLShadowMap) Enabled() bool {
	return wglsm.Get("enabled").Bool()
}

// SetEnabled is ..
func (wglsm *WebGLShadowMap) SetEnabled(v bool) {
	wglsm.Set("enabled", v)
}

// NeedsUpdate is ..
func (wglsm *WebGLShadowMap) NeedsUpdate() bool {
	return wglsm.Get("needsUpdate").Bool()
}

// SetNeedsUpdate is ...
func (wglsm *WebGLShadowMap) SetNeedsUpdate(v bool) {
	wglsm.Set("needsUpdate", v)
}

// SetRenderSinglesSided is ...
func (wglsm *WebGLShadowMap) SetRenderSinglesSided(v bool) {
	wglsm.Set("renderSingleSided", v)
}

// SetRenderReverseSided is ...
func (wglsm *WebGLShadowMap) SetRenderReverseSided(v bool) {
	wglsm.Set("renderReverseSided", v)
}

// //
// func (wglsm *WebGLShadowMap) Type() ShadowMapType {
// 	return ShadowMapType(wglsm.Get("type"))
// }
// func (wglsm *WebGLShadowMap) SetType(v ShadowMapType) {
// 	wglsm.Set("type", v)
// }
// func (wglsm *WebGLShadowMap) Render(scene Scene, camera Camera) {
// 	wglsm.Call("render", scene.JSValue(), camera.JSValue())
// }
