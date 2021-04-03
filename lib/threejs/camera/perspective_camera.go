package camera

import (
	"app/lib/threejs"
	"syscall/js"
)

// PerspectiveCamera is Camera that uses perspective projection.
// This projection mode is designed to mimic the way the human eye sees.
// It is the most common projection mode used for rendering a 3D scene.
type PerspectiveCamera interface {
	threejs.Camera

	UpdateProjectionMatrix()

	Aspect() float64
	SetAspect(v float64)

	// Far gets camera frustum far plane. Default is 2000.
	// Must be greater than the current value of near plane.
	Far() float64

	// SetFar sets camera frustum far plane. Default is 2000.
	// Must be greater than the current value of near plane.
	SetFar(v float64)

	// Fov gets Camera frustum vertical field of view,
	// from bottom to top of view, in degrees. Default is 50.
	Fov() float64

	// SetFov sets Camera frustum vertical field of view,
	// from bottom to top of view, in degrees. Default is 50.
	SetFov(v float64)

	// Near gets camera frustum near plane. Default is 0.1.
	// The valid range is greater than 0 and less than the current value of the far plane.
	// Note that, unlike for the OrthographicCamera, 0 is not a valid value for a PerspectiveCamera's near plane.
	Near() float64

	// SetNear sets camera frustum near plane. Default is 0.1.
	// The valid range is greater than 0 and less than the current value of the far plane.
	// Note that, unlike for the OrthographicCamera, 0 is not a valid value for a PerspectiveCamera's near plane.
	SetNear(v float64)
}

// perspectiveCameraImp is implementation of PerspectiveCamera
type perspectiveCameraImp struct {
	threejs.Object3D
	// js.Value
}

// NewPerspectiveCamera is factory method for PerspectiveCamera.
func NewPerspectiveCamera(fov float64, aspect float64, near float64, far float64) PerspectiveCamera {
	return &perspectiveCameraImp{threejs.NewObject3DFromJSValue(threejs.Threejs("PerspectiveCamera").New(fov, aspect, near, far))}
}

// NewPerspectiveCameraFromJSValue is ...
func NewPerspectiveCameraFromJSValue(value js.Value) PerspectiveCamera {
	return &perspectiveCameraImp{threejs.NewObject3DFromJSValue(value)}
}

// Aspect is ...
func (pc *perspectiveCameraImp) Aspect() float64 {
	return pc.JSValue().Get("aspect").Float()
}

// SetAspect is ...
func (pc *perspectiveCameraImp) SetAspect(v float64) {
	pc.JSValue().Set("aspect", v)
}

// ProjectionMatrix is ...
func (pc *perspectiveCameraImp) ProjectionMatrix() *threejs.Matrix4 {
	return &threejs.Matrix4{Value: pc.JSValue().Get("projectionMatrix")}
}

// SetProjectionMatrix is ...
func (pc *perspectiveCameraImp) SetProjectionMatrix(v *threejs.Matrix4) {
	pc.JSValue().Set("projectionMatrix", v.JSValue())
}

// UpdateProjectionMatrix is ...
func (pc *perspectiveCameraImp) UpdateProjectionMatrix() {
	pc.JSValue().Call("updateProjectionMatrix")
}

// MatrixWorldInverse is the inverse of matrixWorld.
// MatrixWorld contains the Matrix which has the world transform of the Camera.
func (pc *perspectiveCameraImp) MatrixWorldInverse() *threejs.Matrix4 {
	return &threejs.Matrix4{Value: pc.JSValue().Get("matrixWorldInverse")}
}

// ProjectionMatrixInverse is the inverse of projectionMatrix.
func (pc *perspectiveCameraImp) ProjectionMatrixInverse() *threejs.Matrix4 {
	return &threejs.Matrix4{Value: pc.JSValue().Get("projectionMatrixInverse")}
}

// UpdateMatrixWorld is ...
func (pc *perspectiveCameraImp) UpdateMatrixWorld(force bool) {
	pc.JSValue().Call("updateMatrixWorld", force)
}

// IsCamera is ...
func (pc *perspectiveCameraImp) IsCamera() bool {
	return pc.JSValue().Get("isCamera").Bool()
}

// WorldDirection is ...
func (pc *perspectiveCameraImp) WorldDirection(target *threejs.Vector3) *threejs.Vector3 {
	return &threejs.Vector3{Value: pc.JSValue().JSValue().Call("getWorldDirection", target)}
}

func (pc *perspectiveCameraImp) Far() float64 {
	return pc.JSValue().Get("far").Float()
}

func (pc *perspectiveCameraImp) SetFar(v float64) {
	pc.JSValue().Set("far", v)
}

func (pc *perspectiveCameraImp) Fov() float64 {
	return pc.JSValue().Get("fov").Float()
}

func (pc *perspectiveCameraImp) SetFov(v float64) {
	pc.JSValue().Set("fov", v)
}

func (pc *perspectiveCameraImp) Near() float64 {
	return pc.JSValue().Get("near").Float()
}

func (pc *perspectiveCameraImp) SetNear(v float64) {
	pc.JSValue().Set("near", v)
}
