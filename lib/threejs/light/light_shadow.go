package light

import (
	"app/lib/threejs"
	"syscall/js"
)

// LightShadow serves as a base class for the other shadow classes.
type LightShadow interface {
	JSValue() js.Value

	// Bias is shadow map bias, how much to add or subtract from the normalized depth when deciding whether a surface is in shadow.
	// The default is 0. Very tiny adjustments here (in the order of 0.0001) may help reduce artefacts in shadows
	Bias() float64

	// SetBias set shadow map bias, how much to add or subtract from the normalized depth when deciding whether a surface is in shadow.
	// The default is 0. Very tiny adjustments here (in the order of 0.0001) may help reduce artefacts in shadows
	SetBias(bias float64)

	// NormalBias() float64
	// SetNormalBias(bias float64)

	// Radius() float64
	// SetRadius(radius float64)

	// MapSize is a Vector2 defining the width and height of the shadow map.
	//
	// Higher values give better quality shadows at the cost of computation time.
	// Values must be powers of 2, up to the WebGLRenderer.capabilities.maxTextureSize for a given device,
	// although the width and height don't have to be the same (so, for example, (512, 1024) is valid).
	// The default is ( 512, 512 ).
	MapSize() *threejs.Vector2

	// SetMapSize(v *threejs.Vector2)

	// AutoUpdate() bool
	// needsUpdate() bool
}

type defaultLightShadowImp struct {
	js.Value
}

// NewDefaultLightShadow is constructor
func NewDefaultLightShadow(value js.Value) LightShadow {
	return &defaultLightShadowImp{value}
}

// Bias is ...
func (d *defaultLightShadowImp) Bias() float64 {
	return d.Get("bias").Float()
}

// SetBias is ...
func (d *defaultLightShadowImp) SetBias(bias float64) {
	d.Set("bias", bias)
}

// MapSize is ...
func (d *defaultLightShadowImp) MapSize() *threejs.Vector2 {
	return &threejs.Vector2{Value: d.Get("mapSize")}
}
