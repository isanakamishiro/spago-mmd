package light

import "app/lib/threejs"

// PointLightHelper is This displays a helper object consisting of a spherical Mesh for visualizing a PointLight.
type PointLightHelper struct {
	threejs.Object3D
}

// NewPointLightHelper creates PointLightHelper.
// light-- The light to be visualized.
func NewPointLightHelper(light PointLight) *PointLightHelper {
	return &PointLightHelper{
		Object3D: threejs.NewObject3DFromJSValue(
			threejs.Threejs("PointLightHelper").New(light.JSValue()),
		),
	}
}

// Dispose of the PointLightHelper.
func (c *PointLightHelper) Dispose() {
	c.JSValue().Call("dispose")
}

// Update updates the helper to match the position of the .light.
func (c *PointLightHelper) Update() {
	c.JSValue().Call("update")
}
