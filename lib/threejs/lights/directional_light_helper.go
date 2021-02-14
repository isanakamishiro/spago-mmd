package lights

import "app/lib/threejs"

// DirectionalLightHelper is helper object to assist with
// visualizing a DirectionalLight's effect on the scene.
// This consists of plane and a line representing the light's position and direction.
type DirectionalLightHelper struct {
	threejs.Object3D
}

// NewDirectionalLightHelper creates DirectionalLightHelper.
// light-- The light to be visualized.
func NewDirectionalLightHelper(light DirectionalLight) *DirectionalLightHelper {
	return &DirectionalLightHelper{
		Object3D: threejs.NewObject3DFromJSValue(
			threejs.GetJsObject("DirectionalLightHelper").New(light.JSValue()),
		),
	}
}

// Dispose of the directionalLightHelper.
func (c *DirectionalLightHelper) Dispose() {
	c.JSValue().Call("dispose")
}

// Update updates the helper to match the position
// and direction of the directionalLight being visualized.
func (c *DirectionalLightHelper) Update() {
	c.JSValue().Call("update")
}
