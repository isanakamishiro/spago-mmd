package cameras

import (
	"app/lib/threejs"
)

// CameraHelper helps with visualizing what a camera contains in its frustum.
// It visualizes the frustum of a camera using a LineSegments.
type CameraHelper struct {
	threejs.Object3D
}

// NewCameraHelper create a new CameraHelper for the specified camera.
// camera -- The camera to visualize.
func NewCameraHelper(camera threejs.Camera) *CameraHelper {
	return &CameraHelper{
		Object3D: threejs.NewObject3DFromJSValue(
			threejs.GetJsObject("CameraHelper").New(camera.JSValue()),
		),
	}
}

// Update updates the helper based on the projectionMatrix of the camera.
func (c *CameraHelper) Update() {
	c.JSValue().Call("update")
}
