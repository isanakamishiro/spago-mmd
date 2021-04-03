package helper

import "app/lib/threejs"

// AxesHelper is an axis object to visualize the 3 axes in a simple way.
// The X axis is red. The Y axis is green. The Z axis is blue.
type AxesHelper struct {
	threejs.Object3D
}

// NewAxesHelper is factory method for AxesHelper.
// size -- (optional) size of the lines representing the axes. Default is 1.
func NewAxesHelper(size float64) *AxesHelper {
	return &AxesHelper{
		threejs.NewObject3DFromJSValue(
			threejs.Threejs("AxesHelper").New(size),
		),
	}
}

// Material is ...
func (c *AxesHelper) Material() threejs.Material {
	return threejs.NewDefaultMaterialFromJSValue(
		c.JSValue().Get("material"),
	)
}
