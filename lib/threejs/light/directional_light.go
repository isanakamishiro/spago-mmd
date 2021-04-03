package light

import "app/lib/threejs"

// DirectionalLight is a light that gets emitted in a specific direction.
// This light will behave as though it is infinitely far away and the rays produced from it are all parallel.
// The common use case for this is to simulate daylight;
// the sun is far enough away that its position can be considered to be infinite, and all light rays coming from it are parallel.
//
// This light can cast shadows - see the DirectionalLightShadow page for details.
type DirectionalLight interface {
	threejs.Light

	SetPosition(v *threejs.Vector3)
	Shadow() DirectionalLightShadow

	Target() threejs.Object3D
	SetTarget(v threejs.Object3D)
}

type directionalLightImp struct {
	threejs.Light
}

// NewDirectionalLight is factory method for DirectionalLight
// color - Numeric value of the RGB component of the color. Default is 0xffffff.
// intensity - Numeric value of the light's strength/intensity. Default is 1.
func NewDirectionalLight(color threejs.ColorValue, intensity threejs.LightIntensity) DirectionalLight {
	return &directionalLightImp{
		threejs.NewDefaultLightFromJSValue(
			threejs.Threejs("DirectionalLight").New(float64(color), float64(intensity)),
		),
	}
}

// Shadow is a DirectionalLightShadow used to calculate shadows for this light.
func (l *directionalLightImp) Shadow() DirectionalLightShadow {
	return newDirectionalLightShadowFromJSValue(l.JSValue().Get("shadow"))
}

func (l *directionalLightImp) SetPosition(v *threejs.Vector3) {
	l.JSValue().Set("position", v.JSValue())
}

// Target gets the DirectionalLight points from its position to target.position.
// The default position of the target is (0, 0, 0).
// Note: For the target's position to be changed to anything other than the default,
// it must be added to the scene using scene.add( light.target );
func (l *directionalLightImp) Target() threejs.Object3D {
	return threejs.NewObject3DFromJSValue(
		l.JSValue().Get("target"),
	)
}

// SetTarget sets the DirectionalLight points from its position to target.position.
// It is also possible to set the target to be another object in the scene
// (anything with a position property).
func (l *directionalLightImp) SetTarget(v threejs.Object3D) {
	l.JSValue().Set("target", v.JSValue())
}
