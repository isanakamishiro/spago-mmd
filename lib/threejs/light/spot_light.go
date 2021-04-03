package light

import "app/lib/threejs"

// SpotLight gets emitted from a single point in one direction,
// along a cone that increases in size the further from the light it gets.
//
// This light can cast shadows - see the SpotLightShadow page for details.
type SpotLight interface {
	threejs.Light

	// Shadow gets a SpotLightShadow used to calculate shadows for this light.
	Shadow() SpotLightShadow

	// Angle gets maximum extent of the spotlight, in radians, from its direction.
	// Should be no more than Math.PI/2. The default is Math.PI/3.
	Angle() float64
	// SetAngle sets maximum extent of the spotlight, in radians, from its direction.
	// Should be no more than Math.PI/2. The default is Math.PI/3.
	SetAngle(v float64)

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

	// Penumbra gets percent of the spotlight cone that is attenuated due to penumbra.
	// Takes values between zero and 1. The default is 0.0.
	Penumbra() float64

	// SetPenumbra gets percent of the spotlight cone that is attenuated due to penumbra.
	// Takes values between zero and 1. The default is 0.0.
	SetPenumbra(v float64)

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

	// Target gets the DirectionalLight points from its position to target.position.
	// The default position of the target is (0, 0, 0).
	// Note: For the target's position to be changed to anything other than the default,
	// it must be added to the scene using scene.add( light.target );
	Target() threejs.Object3D

	// SetTarget sets the DirectionalLight points from its position to target.position.
	// It is also possible to set the target to be another object in the scene
	// (anything with a position property).
	SetTarget(v threejs.Object3D)
}

type spotLightImp struct {
	threejs.Light
}

// NewSpotLight Creates a new SpotLight.
// color - (optional) hexadecimal color of the light. Default is 0xffffff (white).
// intensity - (optional) numeric value of the light's strength/intensity. Default is 1.
func NewSpotLight(
	color threejs.ColorValue,
	intensity threejs.LightIntensity,
) SpotLight {

	return &spotLightImp{
		threejs.NewDefaultLightFromJSValue(
			threejs.Threejs("SpotLight").New(
				float64(color),
				float64(intensity),
			),
		),
	}
}

func (l *spotLightImp) Shadow() SpotLightShadow {
	return newSpotLightShadowFromJSValue(
		l.JSValue().Get("shadow"),
	)
}

func (l *spotLightImp) Angle() float64 {
	return l.JSValue().Get("angle").Float()
}

func (l *spotLightImp) SetAngle(v float64) {
	l.JSValue().Set("angle", v)
}

func (l *spotLightImp) Decay() float64 {
	return l.JSValue().Get("decay").Float()
}

func (l *spotLightImp) SetDecay(v float64) {
	l.JSValue().Set("decay", v)
}

func (l *spotLightImp) Distance() float64 {
	return l.JSValue().Get("distance").Float()
}

func (l *spotLightImp) SetDistance(v float64) {
	l.JSValue().Set("distance", v)
}

func (l *spotLightImp) Penumbra() float64 {
	return l.JSValue().Get("penumbra").Float()
}

func (l *spotLightImp) SetPenumbra(v float64) {
	l.JSValue().Set("penumbra", v)
}

func (l *spotLightImp) Power() float64 {
	return l.JSValue().Get("power").Float()
}

func (l *spotLightImp) SetPower(v float64) {
	l.JSValue().Set("power", v)
}

func (l *spotLightImp) Target() threejs.Object3D {
	return threejs.NewObject3DFromJSValue(
		l.JSValue().Get("target"),
	)
}

func (l *spotLightImp) SetTarget(v threejs.Object3D) {
	l.JSValue().Set("target", v.JSValue())
}
