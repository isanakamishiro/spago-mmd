package lights

import (
	"app/lib/threejs/cameras"
	"syscall/js"
)

// PointLightShadow is used internally by PointLights for calculating shadows.
type PointLightShadow interface {
	LightShadow

	Camera() cameras.PerspectiveCamera
}

type pointLightShadowImp struct {
	LightShadow
}

// newPointLightShadowFromJSValue is constructor
func newPointLightShadowFromJSValue(value js.Value) PointLightShadow {
	return &pointLightShadowImp{
		NewDefaultLightShadow(value),
	}
}

func (d *pointLightShadowImp) Camera() cameras.PerspectiveCamera {
	return cameras.NewPerspectiveCameraFromJSValue(
		d.JSValue().Get("camera"),
	)
}
