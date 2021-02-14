package fogs

import (
	"app/lib/threejs"
)

// Fog contains the parameters that define linear fog, i.e., that grows linearly denser with the distance.
type Fog interface {
	threejs.FogBase

	//Near gets the minimum distance to start applying fog. Objects that are less than 'near' units from the active camera won't be affected by fog.Default is 1.
	Near() float64
	//SetNear sets the minimum distance to start applying fog. Objects that are less than 'near' units from the active camera won't be affected by fog.Default is 1.
	SetNear(v float64)

	// Far gets the maximum distance at which fog stops being calculated and applied. Objects that are more than 'far' units away from the active camera won't be affected by fog.
	// Default is 1000.
	Far() float64

	// SetFar sets the maximum distance at which fog stops being calculated and applied. Objects that are more than 'far' units away from the active camera won't be affected by fog.
	// Default is 1000.
	SetFar(v float64)
}

type fogImp struct {
	threejs.FogBase
}

// NewFog creates Fog.
func NewFog(color threejs.Color, near float64, far float64) Fog {
	return &fogImp{
		FogBase: threejs.NewFogBaseFromJSValue(
			threejs.GetJsObject("Fog").New(color.Hex(), near, far),
		),
	}
}

//Near gets the minimum distance to start applying fog. Objects that are less than 'near' units from the active camera won't be affected by fog.Default is 1.
func (c *fogImp) Near() float64 {
	return c.JSValue().Get("near").Float()
}

//SetNear sets the minimum distance to start applying fog. Objects that are less than 'near' units from the active camera won't be affected by fog.Default is 1.
func (c *fogImp) SetNear(v float64) {
	c.JSValue().Set("near", v)
}

// Far gets the maximum distance at which fog stops being calculated and applied. Objects that are more than 'far' units away from the active camera won't be affected by fog.
// Default is 1000.
func (c *fogImp) Far() float64 {
	return c.JSValue().Get("far").Float()
}

// SetFar sets the maximum distance at which fog stops being calculated and applied. Objects that are more than 'far' units away from the active camera won't be affected by fog.
// Default is 1000.
func (c *fogImp) SetFar(v float64) {
	c.JSValue().Set("far", v)
}
