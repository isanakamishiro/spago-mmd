package light

import "app/lib/threejs"

// HemisphereLight is a light source positioned directly above the scene,
// with color fading from the sky color to the ground color.
// This light cannot be used to cast shadows.
type HemisphereLight interface {
	threejs.Light

	GroundColor() threejs.Color
	SetGroundColor(color threejs.Color)
}

type hemisphereLightImp struct {
	threejs.Light
}

// NewHemisphereLight creates a new HemisphereLight.
// skyColor - hexadecimal color of the sky. Default is 0xffffff.
// groundColor - hexadecimal color of the ground. Default is 0xffffff.
// intensity - numeric value of the light's strength/intensity. Default is 1.
func NewHemisphereLight(skyColor threejs.ColorValue, groundColor threejs.ColorValue, intensity threejs.LightIntensity) HemisphereLight {
	return &hemisphereLightImp{
		threejs.NewDefaultLightFromJSValue(
			threejs.Threejs("HemisphereLight").New(float64(skyColor), float64(groundColor), float64(intensity)),
		),
	}
}

// GroundColor gets the light's ground color, as passed in the constructor.
// Default is a new Color set to white (0xffffff).
func (c *hemisphereLightImp) GroundColor() threejs.Color {
	return threejs.NewColorFromJSValue(c.JSValue().Get("groundColor"))
}

// SetGroundColor sets the light's ground color, as passed in the constructor.
// Default is a new Color set to white (0xffffff).
func (c *hemisphereLightImp) SetGroundColor(color threejs.Color) {
	c.JSValue().Set("groundColor", color.JSValue())
}
