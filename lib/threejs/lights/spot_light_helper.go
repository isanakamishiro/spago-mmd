package lights

import "app/lib/threejs"

// SpotLightHelper displays a cone shaped helper object for a SpotLight.
type SpotLightHelper struct {
	threejs.Object3D
}

// NewSpotLightHelper creates SpotLightHelper.
// light-- The light to be visualized.
func NewSpotLightHelper(light PointLight) *SpotLightHelper {
	return &SpotLightHelper{
		Object3D: threejs.NewObject3DFromJSValue(
			threejs.GetJsObject("SpotLightHelper").New(light.JSValue()),
		),
	}
}

// Dispose of the light helper.
func (c *SpotLightHelper) Dispose() {
	c.JSValue().Call("dispose")
}

// Update updates the light helper.
func (c *SpotLightHelper) Update() {
	c.JSValue().Call("update")
}
