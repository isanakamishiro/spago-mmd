package lights

import (
	"app/lib/threejs/cameras"
	"syscall/js"
)

// DirectionalLightShadow is used internally by DirectionalLights for calculating shadows.
//
// Unlike the other shadow classes, this uses an OrthographicCamera to calculate the shadows,
// rather than a PerspectiveCamera.
// This is because light rays from a DirectionalLight are parallel.
type DirectionalLightShadow interface {
	LightShadow

	Camera() cameras.OrthographicCamera
}

type directionalLightShadowImp struct {
	LightShadow
}

// newDirectionalLightShadowFromJSValue is constructor
func newDirectionalLightShadowFromJSValue(value js.Value) DirectionalLightShadow {
	return &directionalLightShadowImp{
		NewDefaultLightShadow(value),
	}
}

func (d *directionalLightShadowImp) Camera() cameras.OrthographicCamera {
	return cameras.NewOrthographicCameraFromJSValue(
		d.JSValue().Get("camera"),
	)
}
