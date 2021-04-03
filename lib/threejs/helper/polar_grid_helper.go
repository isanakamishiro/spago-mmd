package helper

import "app/lib/threejs"

// PolarGridHelper extend: []
type PolarGridHelper struct {
	threejs.Object3D
}

// NewPolarGridHelper is factory method for PolarGridHelper.
func NewPolarGridHelper(radius float64, radials float64) *PolarGridHelper {
	return &PolarGridHelper{
		threejs.NewObject3DFromJSValue(
			threejs.Threejs("PolarGridHelper").New(radius, radials),
		),
	}
}
