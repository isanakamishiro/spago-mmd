package light

import "app/lib/threejs"

// AmbientLight is a light globally illuminates all objects in the scene equally.
// This light cannot be used to cast shadows as it does not have a direction.
type AmbientLight interface {
	threejs.Light
}

type ambientLightImp struct {
	threejs.Light
}

// NewAmbientLight is factory method for AmbientLigh
// color - Numeric value of the RGB component of the color. Default is 0xffffff.
// intensity - Numeric value of the light's strength/intensity. Default is 1.
func NewAmbientLight(color threejs.ColorValue, intensity threejs.LightIntensity) AmbientLight {
	return &ambientLightImp{
		threejs.NewDefaultLightFromJSValue(
			threejs.Threejs("AmbientLight").New(float64(color), float64(intensity)),
		),
	}
}
