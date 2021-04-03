package light

import "app/lib/threejs"

// PointLight is a light that gets emitted from a single point in all directions.
// A common use case for this is to replicate the light emitted from a bare lightbulb.
//
// This light can cast shadows - see PointLightShadow page for details.
type PointLight interface {
	threejs.Light

	// Shadow gets a PointLightShadow used to calculate shadows for this light.
	// The lightShadow's camera is set to a PerspectiveCamera with fov of 90, aspect of 1, near clipping plane at 0.5 and far clipping plane at 500.
	Shadow() PointLightShadow

	// Decay gets the amount the light dims along the distance of the light.
	// In physically correct mode, decay = 2 leads to physically realistic light falloff. The default is 1.
	Decay() float64

	// SetDecay sets the amount the light dims along the distance of the light.
	// In physically correct mode, decay = 2 leads to physically realistic light falloff. The default is 1.
	SetDecay(v float64)

	// Distance gets the distance : Default mode — When distance is zero, light does not attenuate.
	// When distance is non-zero, light will attenuate linearly from maximum intensity
	// at the light's position down to zero at this distance from the light.
	Distance() float64

	// SetDistance sets the distance : Default mode — When distance is zero, light does not attenuate.
	// When distance is non-zero, light will attenuate linearly from maximum intensity
	// at the light's position down to zero at this distance from the light.
	SetDistance(v float64)

	// Power gets the light's power.
	// In physically correct mode, the luminous power of the light measured in lumens.
	// Default is 4Math.PI.
	// This is directly related to the intensity in the ratio
	// 'power = intensity * 4pi' and changing this will also change the intensity.
	Power() float64

	// SetPower sets the light's power.
	// In physically correct mode, the luminous power of the light measured in lumens.
	// Default is 4Math.PI.
	// This is directly related to the intensity in the ratio
	// 'power = intensity * 4pi' and changing this will also change the intensity.
	SetPower(v float64)
}

type pointLightImp struct {
	threejs.Light
}

// NewPointLight Creates a new PointLight.
// color - (optional) hexadecimal color of the light. Default is 0xffffff (white).
// intensity - (optional) numeric value of the light's strength/intensity. Default is 1.
// distance - Maximum range of the light. Default is 0 (no limit).
// decay - The amount the light dims along the distance of the light. Default is 1. For physically correct lighting, set this to 2.
func NewPointLight(
	color threejs.ColorValue,
	intensity threejs.LightIntensity,
	distance float64,
	decay float64,
) PointLight {

	return &pointLightImp{
		threejs.NewDefaultLightFromJSValue(
			threejs.Threejs("PointLight").New(
				float64(color),
				float64(intensity),
				distance,
				decay,
			),
		),
	}
}

func (l *pointLightImp) Shadow() PointLightShadow {
	return newPointLightShadowFromJSValue(
		l.JSValue().Get("shadow"),
	)
}

func (l *pointLightImp) Decay() float64 {
	return l.JSValue().Get("decay").Float()
}

func (l *pointLightImp) SetDecay(v float64) {
	l.JSValue().Set("decay", v)
}

func (l *pointLightImp) Distance() float64 {
	return l.JSValue().Get("distance").Float()
}

func (l *pointLightImp) SetDistance(v float64) {
	l.JSValue().Set("distance", v)
}

func (l *pointLightImp) Power() float64 {
	return l.JSValue().Get("power").Float()
}

func (l *pointLightImp) SetPower(v float64) {
	l.JSValue().Set("power", v)
}
