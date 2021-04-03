package camera

import (
	"app/lib/threejs"
	"syscall/js"
)

// OrthographicCamera is camera that uses orthographic projection.
// In this projection mode, an object's size in the rendered image stays constant regardless of its distance from the camera.
// This can be useful for rendering 2D scenes and UI elements, amongst other things.
type OrthographicCamera interface {
	threejs.Camera

	UpdateProjectionMatrix()

	Left() float64
	SetLeft(v float64)
	Right() float64
	SetRight(v float64)
	Top() float64
	SetTop(v float64)
	Bottom() float64
	SetBottom(v float64)
	// Zoom gets the zoom factor of the camera. Default is 1.
	Zoom() float64
	// SetZoom sets the zoom factor of the camera. Default is 1.
	SetZoom(v float64)
}

// orthographicCameraImp is implementation of OrthographicCamera.
type orthographicCameraImp struct {
	threejs.Object3D
}

// NewOrthographicCamera is factory method for OrthographicCamera.
// left — Camera frustum left plane.
// right — Camera frustum right plane.
// top — Camera frustum top plane.
// bottom — Camera frustum bottom plane.
// near — Camera frustum near plane.
// far — Camera frustum far plane.
// Together these define the camera's viewing frustum.
func NewOrthographicCamera(left float64, right float64, top float64, bottom float64, near float64, far float64) OrthographicCamera {
	return &orthographicCameraImp{threejs.NewObject3DFromJSValue(threejs.Threejs("OrthographicCamera").New(left, right, top, bottom, near, far))}
}

// NewOrthographicCameraFromJSValue is ...
func NewOrthographicCameraFromJSValue(value js.Value) OrthographicCamera {
	return &orthographicCameraImp{threejs.NewObject3DFromJSValue(value)}
}

// ProjectionMatrix is ...
func (oc *orthographicCameraImp) ProjectionMatrix() *threejs.Matrix4 {
	return &threejs.Matrix4{Value: oc.JSValue().Get("projectionMatrix")}
}

// SetProjectionMatrix is ...
func (oc *orthographicCameraImp) SetProjectionMatrix(v *threejs.Matrix4) {
	oc.JSValue().Set("projectionMatrix", v.JSValue())
}

// UpdateProjectionMatrix is ...
func (oc *orthographicCameraImp) UpdateProjectionMatrix() {
	oc.JSValue().Call("updateProjectionMatrix")
}

// MatrixWorldInverse is the inverse of matrixWorld.
// MatrixWorld contains the Matrix which has the world transform of the Camera.
func (oc *orthographicCameraImp) MatrixWorldInverse() *threejs.Matrix4 {
	return &threejs.Matrix4{Value: oc.JSValue().Get("matrixWorldInverse")}
}

// ProjectionMatrixInverse is the inverse of projectionMatrix.
func (oc *orthographicCameraImp) ProjectionMatrixInverse() *threejs.Matrix4 {
	return &threejs.Matrix4{Value: oc.JSValue().Get("projectionMatrixInverse")}
}

// UpdateMatrixWorld is ...
func (oc *orthographicCameraImp) UpdateMatrixWorld(force bool) {
	oc.JSValue().Call("updateMatrixWorld", force)
}

// IsCamera is ...
func (oc *orthographicCameraImp) IsCamera() bool {
	return oc.JSValue().Get("isCamera").Bool()
}

// WorldDirection is ...
func (oc *orthographicCameraImp) WorldDirection(target *threejs.Vector3) *threejs.Vector3 {
	return &threejs.Vector3{Value: oc.JSValue().JSValue().Call("getWorldDirection", target)}
}

// Left is ...
func (oc *orthographicCameraImp) Left() float64 {

	return oc.JSValue().Get("left").Float()
}

// SetLeft is ...
func (oc *orthographicCameraImp) SetLeft(v float64) {
	oc.JSValue().Set("left", v)
}

// Right is ...
func (oc *orthographicCameraImp) Right() float64 {
	return oc.JSValue().Get("right").Float()
}

// SetRight is ...
func (oc *orthographicCameraImp) SetRight(v float64) {
	oc.JSValue().Set("right", v)
}

// Top is ...
func (oc *orthographicCameraImp) Top() float64 {
	return oc.JSValue().Get("top").Float()
}

// SetTop is ...
func (oc *orthographicCameraImp) SetTop(v float64) {
	oc.JSValue().Set("top", v)
}

// Bottom is ...
func (oc *orthographicCameraImp) Bottom() float64 {
	return oc.JSValue().Get("bottom").Float()
}

// SetBottom is ...
func (oc *orthographicCameraImp) SetBottom(v float64) {
	oc.JSValue().Set("bottom", v)
}

func (oc *orthographicCameraImp) Zoom() float64 {
	return oc.JSValue().Get("zoom").Float()
}

func (oc *orthographicCameraImp) SetZoom(v float64) {
	oc.JSValue().Set("zoom", v)
}

// func (oc *OrthographicCamera) CastShadow() bool {
// 	return oc.JSValue().Get("castShadow").Bool()
// }
// func (oc *OrthographicCamera) SetCastShadow(v bool) {
// 	oc.JSValue().Set("castShadow", v)
// }
// func (oc *OrthographicCamera) Children() js.Value {
// 	return oc.JSValue().Get("children")
// }
// func (oc *OrthographicCamera) SetChildren(v js.Value) {
// 	oc.JSValue().Set("children", v)
// }
// func (oc *OrthographicCamera) Far() float64 {
// 	return oc.JSValue().Get("far").Float()
// }
// func (oc *OrthographicCamera) SetFar(v float64) {
// 	oc.JSValue().Set("far", v)
// }
// func (oc *OrthographicCamera) FrustumCulled() bool {
// 	return oc.JSValue().Get("frustumCulled").Bool()
// }
// func (oc *OrthographicCamera) SetFrustumCulled(v bool) {
// 	oc.JSValue().Set("frustumCulled", v)
// }
// func (oc *OrthographicCamera) Id() int {
// 	return oc.JSValue().Get("id").Int()
// }
// func (oc *OrthographicCamera) SetId(v int) {
// 	oc.JSValue().Set("id", v)
// }
// func (oc *OrthographicCamera) IsCamera() bool {
// 	return oc.JSValue().Get("isCamera").Bool()
// }
// func (oc *OrthographicCamera) SetIsCamera(v bool) {
// 	oc.JSValue().Set("isCamera", v)
// }
// func (oc *OrthographicCamera) IsObject3D() bool {
// 	return oc.JSValue().Get("isObject3D").Bool()
// }
// func (oc *OrthographicCamera) SetIsObject3D(v bool) {
// 	oc.JSValue().Set("isObject3D", v)
// }
// func (oc *OrthographicCamera) IsOrthographicCamera() bool {
// 	return oc.JSValue().Get("isOrthographicCamera").Bool()
// }
// func (oc *OrthographicCamera) SetIsOrthographicCamera(v bool) {
// 	oc.JSValue().Set("isOrthographicCamera", v)
// }
// func (oc *OrthographicCamera) Layers() *Layers {
// 	return &Layers{Value: oc.JSValue().Get("layers")}
// }
// func (oc *OrthographicCamera) SetLayers(v *Layers) {
// 	oc.JSValue().Set("layers", v.JSValue())
// }
// func (oc *OrthographicCamera) Matrix() *Matrix4 {
// 	return &Matrix4{Value: oc.JSValue().Get("matrix")}
// }
// func (oc *OrthographicCamera) SetMatrix(v *Matrix4) {
// 	oc.JSValue().Set("matrix", v.JSValue())
// }
// func (oc *OrthographicCamera) MatrixAutoUpdate() bool {
// 	return oc.JSValue().Get("matrixAutoUpdate").Bool()
// }
// func (oc *OrthographicCamera) SetMatrixAutoUpdate(v bool) {
// 	oc.JSValue().Set("matrixAutoUpdate", v)
// }
// func (oc *OrthographicCamera) MatrixWorld() *Matrix4 {
// 	return &Matrix4{Value: oc.JSValue().Get("matrixWorld")}
// }
// func (oc *OrthographicCamera) SetMatrixWorld(v *Matrix4) {
// 	oc.JSValue().Set("matrixWorld", v.JSValue())
// }
// func (oc *OrthographicCamera) MatrixWorldInverse() *Matrix4 {
// 	return &Matrix4{Value: oc.JSValue().Get("matrixWorldInverse")}
// }
// func (oc *OrthographicCamera) SetMatrixWorldInverse(v *Matrix4) {
// 	oc.JSValue().Set("matrixWorldInverse", v.JSValue())
// }
// func (oc *OrthographicCamera) MatrixWorldNeedsUpdate() bool {
// 	return oc.JSValue().Get("matrixWorldNeedsUpdate").Bool()
// }
// func (oc *OrthographicCamera) SetMatrixWorldNeedsUpdate(v bool) {
// 	oc.JSValue().Set("matrixWorldNeedsUpdate", v)
// }
// func (oc *OrthographicCamera) ModelViewMatrix() *Matrix4 {
// 	return &Matrix4{Value: oc.JSValue().Get("modelViewMatrix")}
// }
// func (oc *OrthographicCamera) SetModelViewMatrix(v *Matrix4) {
// 	oc.JSValue().Set("modelViewMatrix", v.JSValue())
// }
// func (oc *OrthographicCamera) Name() string {
// 	return oc.JSValue().Get("name").String()
// }
// func (oc *OrthographicCamera) SetName(v string) {
// 	oc.JSValue().Set("name", v)
// }
// func (oc *OrthographicCamera) Near() float64 {
// 	return oc.JSValue().Get("near").Float()
// }
// func (oc *OrthographicCamera) SetNear(v float64) {
// 	oc.JSValue().Set("near", v)
// }
// func (oc *OrthographicCamera) NormalMatrix() *Matrix3 {
// 	return &Matrix3{Value: oc.JSValue().Get("normalMatrix")}
// }
// func (oc *OrthographicCamera) SetNormalMatrix(v *Matrix3) {
// 	oc.JSValue().Set("normalMatrix", v.JSValue())
// }
// func (oc *OrthographicCamera) OnAfterRender() js.Value {
// 	return oc.JSValue().Get("onAfterRender")
// }
// func (oc *OrthographicCamera) SetOnAfterRender(v js.Value) {
// 	oc.JSValue().Set("onAfterRender", v)
// }
// func (oc *OrthographicCamera) OnBeforeRender() js.Value {
// 	return oc.JSValue().Get("onBeforeRender")
// }
// func (oc *OrthographicCamera) SetOnBeforeRender(v js.Value) {
// 	oc.JSValue().Set("onBeforeRender", v)
// }
// func (oc *OrthographicCamera) Parent() *Object3D {
// 	return &Object3D{Value: oc.JSValue().Get("parent")}
// }
// func (oc *OrthographicCamera) SetParent(v *Object3D) {
// 	oc.JSValue().Set("parent", v.JSValue())
// }
// func (oc *OrthographicCamera) Position() *Vector3 {
// 	return &Vector3{Value: oc.JSValue().Get("position")}
// }
// func (oc *OrthographicCamera) SetPosition(v *Vector3) {
// 	oc.JSValue().Set("position", v.JSValue())
// }
// func (oc *OrthographicCamera) ProjectionMatrix() *Matrix4 {
// 	return &Matrix4{Value: oc.JSValue().Get("projectionMatrix")}
// }
// func (oc *OrthographicCamera) SetProjectionMatrix(v *Matrix4) {
// 	oc.JSValue().Set("projectionMatrix", v.JSValue())
// }
// func (oc *OrthographicCamera) Quaternion() *Quaternion {
// 	return &Quaternion{Value: oc.JSValue().Get("quaternion")}
// }
// func (oc *OrthographicCamera) SetQuaternion(v *Quaternion) {
// 	oc.JSValue().Set("quaternion", v.JSValue())
// }
// func (oc *OrthographicCamera) ReceiveShadow() bool {
// 	return oc.JSValue().Get("receiveShadow").Bool()
// }
// func (oc *OrthographicCamera) SetReceiveShadow(v bool) {
// 	oc.JSValue().Set("receiveShadow", v)
// }
// func (oc *OrthographicCamera) RenderOrder() float64 {
// 	return oc.JSValue().Get("renderOrder").Float()
// }
// func (oc *OrthographicCamera) SetRenderOrder(v float64) {
// 	oc.JSValue().Set("renderOrder", v)
// }
// func (oc *OrthographicCamera) Rotation() *Euler {
// 	return &Euler{Value: oc.JSValue().Get("rotation")}
// }
// func (oc *OrthographicCamera) SetRotation(v *Euler) {
// 	oc.JSValue().Set("rotation", v.JSValue())
// }
// func (oc *OrthographicCamera) Scale() *Vector3 {
// 	return &Vector3{Value: oc.JSValue().Get("scale")}
// }
// func (oc *OrthographicCamera) SetScale(v *Vector3) {
// 	oc.JSValue().Set("scale", v.JSValue())
// }
// func (oc *OrthographicCamera) Type() string {
// 	return oc.JSValue().Get("type").String()
// }
// func (oc *OrthographicCamera) SetType(v string) {
// 	oc.JSValue().Set("type", v)
// }
// func (oc *OrthographicCamera) Up() *Vector3 {
// 	return &Vector3{Value: oc.JSValue().Get("up")}
// }
// func (oc *OrthographicCamera) SetUp(v *Vector3) {
// 	oc.JSValue().Set("up", v.JSValue())
// }
// func (oc *OrthographicCamera) UserData() js.Value {
// 	return oc.JSValue().Get("userData")
// }
// func (oc *OrthographicCamera) SetUserData(v js.Value) {
// 	oc.JSValue().Set("userData", v)
// }
// func (oc *OrthographicCamera) Uuid() string {
// 	return oc.JSValue().Get("uuid").String()
// }
// func (oc *OrthographicCamera) SetUuid(v string) {
// 	oc.JSValue().Set("uuid", v)
// }
// func (oc *OrthographicCamera) View() js.Value {
// 	return oc.JSValue().Get("view")
// }
// func (oc *OrthographicCamera) SetView(v js.Value) {
// 	oc.JSValue().Set("view", v)
// }
// func (oc *OrthographicCamera) Visible() bool {
// 	return oc.JSValue().Get("visible").Bool()
// }
// func (oc *OrthographicCamera) SetVisible(v bool) {
// 	oc.JSValue().Set("visible", v)
// }
// func (oc *OrthographicCamera) Zoom() float64 {
// 	return oc.JSValue().Get("zoom").Float()
// }
// func (oc *OrthographicCamera) SetZoom(v float64) {
// 	oc.JSValue().Set("zoom", v)
// }
// func (oc *OrthographicCamera) DefaultMatrixAutoUpdate() bool {
// 	return oc.JSValue().Get("DefaultMatrixAutoUpdate").Bool()
// }
// func (oc *OrthographicCamera) SetDefaultMatrixAutoUpdate(v bool) {
// 	oc.JSValue().Set("DefaultMatrixAutoUpdate", v)
// }
// func (oc *OrthographicCamera) DefaultUp() *Vector3 {
// 	return &Vector3{Value: oc.JSValue().Get("DefaultUp")}
// }
// func (oc *OrthographicCamera) SetDefaultUp(v *Vector3) {
// 	oc.JSValue().Set("DefaultUp", v.JSValue())
// }
// func (oc *OrthographicCamera) Add(object js.Value) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("add", object)}
// }
// func (oc *OrthographicCamera) AddEventListener(typ string, listener js.Value) {
// 	oc.JSValue().Call("addEventListener", typ, listener)
// }
// func (oc *OrthographicCamera) ApplyMatrix(matrix *Matrix4) {
// 	oc.JSValue().Call("applyMatrix", matrix.JSValue())
// }
// func (oc *OrthographicCamera) ApplyQuaternion(quaternion *Quaternion) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("applyQuaternion", quaternion)}
// }
// func (oc *OrthographicCamera) ClearViewOffset() {
// 	oc.JSValue().Call("clearViewOffset")
// }
// func (oc *OrthographicCamera) Clone(recursive bool) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("clone", recursive)}
// }
// func (oc *OrthographicCamera) Copy(source Camera, recursive bool) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("copy", source.JSValue(), recursive)}
// }
// func (oc *OrthographicCamera) DispatchEvent(event js.Value) {
// 	oc.JSValue().Call("dispatchEvent", event)
// }
// func (oc *OrthographicCamera) GetObjectById(id int) *Object3D {
// 	return &Object3D{Value: oc.JSValue().Call("getObjectById", id)}
// }
// func (oc *OrthographicCamera) GetObjectByName(name string) *Object3D {
// 	return &Object3D{Value: oc.JSValue().Call("getObjectByName", name)}
// }
// func (oc *OrthographicCamera) GetObjectByProperty(name string, value string) *Object3D {
// 	return &Object3D{Value: oc.JSValue().Call("getObjectByProperty", name, value)}
// }
// func (oc *OrthographicCamera) GetWorldDirection(target *Vector3) *Vector3 {
// 	return &Vector3{Value: oc.JSValue().Call("getWorldDirection", target)}
// }
// func (oc *OrthographicCamera) GetWorldPosition(target *Vector3) *Vector3 {
// 	return &Vector3{Value: oc.JSValue().Call("getWorldPosition", target)}
// }
// func (oc *OrthographicCamera) GetWorldQuaternion(target *Quaternion) *Quaternion {
// 	return &Quaternion{Value: oc.JSValue().Call("getWorldQuaternion", target)}
// }
// func (oc *OrthographicCamera) GetWorldScale(target *Vector3) *Vector3 {
// 	return &Vector3{Value: oc.JSValue().Call("getWorldScale", target)}
// }
// func (oc *OrthographicCamera) HasEventListener(typ string, listener js.Value) bool {
// 	return oc.JSValue().Call("hasEventListener", typ, listener).Bool()
// }
// func (oc *OrthographicCamera) LocalToWorld(vector *Vector3) *Vector3 {
// 	return &Vector3{Value: oc.JSValue().Call("localToWorld", vector)}
// }
// func (oc *OrthographicCamera) LookAt(vector *Vector3, y float64, z float64) {
// 	oc.JSValue().Call("lookAt", vector, y, z)
// }
// func (oc *OrthographicCamera) Raycast(raycaster *Raycaster, intersects js.Value) {
// 	oc.JSValue().Call("raycast", raycaster.JSValue(), intersects)
// }
// func (oc *OrthographicCamera) Remove(object js.Value) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("remove", object)}
// }
// func (oc *OrthographicCamera) RemoveEventListener(typ string, listener js.Value) {
// 	oc.JSValue().Call("removeEventListener", typ, listener)
// }
// func (oc *OrthographicCamera) RotateOnAxis(axis *Vector3, angle float64) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("rotateOnAxis", axis, angle)}
// }
// func (oc *OrthographicCamera) RotateOnWorldAxis(axis *Vector3, angle float64) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("rotateOnWorldAxis", axis, angle)}
// }
// func (oc *OrthographicCamera) RotateX(angle float64) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("rotateX", angle)}
// }
// func (oc *OrthographicCamera) RotateY(angle float64) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("rotateY", angle)}
// }
// func (oc *OrthographicCamera) RotateZ(angle float64) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("rotateZ", angle)}
// }
// func (oc *OrthographicCamera) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
// 	oc.JSValue().Call("setRotationFromAxisAngle", axis.JSValue(), angle)
// }
// func (oc *OrthographicCamera) SetRotationFromEuler(euler *Euler) {
// 	oc.JSValue().Call("setRotationFromEuler", euler.JSValue())
// }
// func (oc *OrthographicCamera) SetRotationFromMatrix(m *Matrix4) {
// 	oc.JSValue().Call("setRotationFromMatrix", m.JSValue())
// }
// func (oc *OrthographicCamera) SetRotationFromQuaternion(q *Quaternion) {
// 	oc.JSValue().Call("setRotationFromQuaternion", q.JSValue())
// }
// func (oc *OrthographicCamera) SetViewOffset(fullWidth float64, fullHeight float64, offsetX float64, offsetY float64, width float64, height float64) {
// 	oc.JSValue().Call("setViewOffset", fullWidth, fullHeight, offsetX, offsetY, width, height)
// }
// func (oc *OrthographicCamera) ToJSON(meta js.Value) js.Value {
// 	return oc.JSValue().Call("toJSON", meta)
// }
// func (oc *OrthographicCamera) TranslateOnAxis(axis *Vector3, distance float64) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("translateOnAxis", axis, distance)}
// }
// func (oc *OrthographicCamera) TranslateX(distance float64) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("translateX", distance)}
// }
// func (oc *OrthographicCamera) TranslateY(distance float64) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("translateY", distance)}
// }
// func (oc *OrthographicCamera) TranslateZ(distance float64) *OrthographicCamera {
// 	return &OrthographicCamera{Value: oc.JSValue().Call("translateZ", distance)}
// }
// func (oc *OrthographicCamera) Traverse(callback js.Value) {
// 	oc.JSValue().Call("traverse", callback)
// }
// func (oc *OrthographicCamera) TraverseAncestors(callback js.Value) {
// 	oc.JSValue().Call("traverseAncestors", callback)
// }
// func (oc *OrthographicCamera) TraverseVisible(callback js.Value) {
// 	oc.JSValue().Call("traverseVisible", callback)
// }
// func (oc *OrthographicCamera) UpdateMatrix() {
// 	oc.JSValue().Call("updateMatrix")
// }
// func (oc *OrthographicCamera) UpdateMatrixWorld(force bool) {
// 	oc.JSValue().Call("updateMatrixWorld", force)
// }
// func (oc *OrthographicCamera) UpdateProjectionMatrix() {
// 	oc.JSValue().Call("updateProjectionMatrix")
// }
// func (oc *OrthographicCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
// 	oc.JSValue().Call("updateWorldMatrix", updateParents, updateChildren)
// }
// func (oc *OrthographicCamera) WorldToLocal(vector *Vector3) *Vector3 {
// 	return &Vector3{Value: oc.JSValue().Call("worldToLocal", vector)}
// }
