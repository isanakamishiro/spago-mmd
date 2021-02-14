package lights

import (
	"app/lib/threejs/cameras"
	"syscall/js"
)

// SpotLightShadow is used internally by SpotLights for calculating shadows.
type SpotLightShadow interface {
	LightShadow

	Camera() cameras.PerspectiveCamera

	// Focus gets that used to focus the shadow camera. The camera's field of view is set as a percentage of the spotlight's field-of-view. Range is [0, 1]. Default is 1.0.
	Focus() float64

	// SetFocus sets that used to focus the shadow camera. The camera's field of view is set as a percentage of the spotlight's field-of-view. Range is [0, 1]. Default is 1.0.
	SetFocus(v float64)
}

type spotLightShadowImp struct {
	LightShadow
}

// newSpotLightShadowFromJSValue is constructor
func newSpotLightShadowFromJSValue(value js.Value) SpotLightShadow {
	return &spotLightShadowImp{
		NewDefaultLightShadow(value),
	}
}

func (d *spotLightShadowImp) Camera() cameras.PerspectiveCamera {
	return cameras.NewPerspectiveCameraFromJSValue(
		d.JSValue().Get("camera"),
	)
}

func (d *spotLightShadowImp) Focus() float64 {
	return d.JSValue().Get("focus").Float()
}

func (d *spotLightShadowImp) SetFocus(v float64) {
	d.JSValue().Set("focus", v)
}
