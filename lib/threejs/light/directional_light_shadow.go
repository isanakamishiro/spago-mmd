package light

import (
	"app/lib/threejs/camera"
	"syscall/js"
)

// DirectionalLightShadow is used internally by DirectionalLights for calculating shadows.
//
// Unlike the other shadow classes, this uses an OrthographicCamera to calculate the shadows,
// rather than a PerspectiveCamera.
// This is because light rays from a DirectionalLight are parallel.
type DirectionalLightShadow interface {
	LightShadow

	Camera() camera.OrthographicCamera
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

func (d *directionalLightShadowImp) Camera() camera.OrthographicCamera {
	return camera.NewOrthographicCameraFromJSValue(
		d.JSValue().Get("camera"),
	)
}
