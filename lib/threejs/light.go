package threejs

import (
	"syscall/js"
)

// LightIntensity is numeric value of the light's strength/intensity. Default is 1.
type LightIntensity float64

// Light is interface for lights - all other light types inherit the properties and methods described here.
type Light interface {
	Object3D

	Color() Color
	SetColor(v Color)
	Intensity() float64
	SetIntensity(v float64)
}

// defaultLightImp is implementation of Lights
type defaultLightImp struct {
	Object3D
}

// NewDefaultLightFromJSValue creates a new Light.
// Note that this is not intended to be called directly (use one of derived classes instead).
func NewDefaultLightFromJSValue(value js.Value) Light {
	return &defaultLightImp{
		NewObject3DFromJSValue(value),
	}
}

// Color of the light. Defaults to a new Color set to white, if not passed in the constructor.
func (l *defaultLightImp) Color() Color {
	return NewColorFromJSValue(l.JSValue().Get("color"))
}

// Color of the light. Defaults to a new Color set to white, if not passed in the constructor.
func (l *defaultLightImp) SetColor(v Color) {
	l.JSValue().Set("color", v.JSValue())
}

// Intensity get the light's intensity, or strength.
// In physically correct mode, the product of color * intensity is interpreted as luminous intensity measured in candela.
// Default - 1.0.
func (l *defaultLightImp) Intensity() float64 {
	return l.JSValue().Get("intensity").Float()
}

// SetIntensity set the light's intensity, or strength.
// In physically correct mode, the product of color * intensity is interpreted as luminous intensity measured in candela.
// Default - 1.0.
func (l *defaultLightImp) SetIntensity(v float64) {
	l.JSValue().Set("intensity", v)
}

// ToJSON is ...
// meta -- object containing metadata such as materials, textures for objects.
// Convert the light to three.js JSON Object/Scene format.
func (l *defaultLightImp) ToJSON(meta js.Value) js.Value {
	return l.JSValue().Call("toJSON", meta)
}
