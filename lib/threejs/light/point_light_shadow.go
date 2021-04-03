package light

import (
	"app/lib/threejs/camera"
	"syscall/js"
)

// PointLightShadow is used internally by PointLights for calculating shadows.
type PointLightShadow interface {
	LightShadow

	Camera() camera.PerspectiveCamera
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

func (d *pointLightShadowImp) Camera() camera.PerspectiveCamera {
	return camera.NewPerspectiveCameraFromJSValue(
		d.JSValue().Get("camera"),
	)
}
