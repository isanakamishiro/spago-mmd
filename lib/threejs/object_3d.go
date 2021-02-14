package threejs

import (
	"syscall/js"
)

// Object3D is interface for threejs.Object3D object.
// This is the base class for most objects in three.js and provides a set of properties
// and methods for manipulating objects in 3D space.
//
// Note that this can be used for grouping objects via the .add( object ) method
// which adds the object as a child, however it is better to use Group for this.
type Object3D interface {
	EventDispatcher

	JSValue() js.Value

	Add(o Object3D)

	ID() int
	UUID() string
	Name() string
	SetName(v string)
	Type() string

	Children() js.Value
	SetChildren(v js.Value)
	Parent() Object3D
	SetParent(v Object3D)

	IsObject3D() bool
	SetIsObject3D(v bool)
	Layers() *Layers
	SetLayers(v *Layers)

	Up() *Vector3
	SetUp(v *Vector3)
	Position() *Vector3
	Rotation() *Euler
	Quaternion() *Quaternion
	Scale() *Vector3
	ModelViewMatrix() *Matrix4
	NormalMatrix() *Matrix3
	Matrix() *Matrix4
	SetMatrix(v *Matrix4)
	MatrixWorld() *Matrix4
	SetMatrixWorld(v *Matrix4)
	MatrixAutoUpdate() bool
	SetMatrixAutoUpdate(v bool)
	MatrixWorldNeedsUpdate() bool
	SetMatrixWorldNeedsUpdate(v bool)

	CastShadow() bool
	SetCastShadow(v bool)
	ReceiveShadow() bool
	SetReceiveShadow(v bool)
	FrustumCulled() bool
	SetFrustumCulled(v bool)
	RenderOrder() float64
	SetRenderOrder(v float64)

	OnAfterRender() js.Value
	SetOnAfterRender(v js.Value)
	OnBeforeRender() js.Value
	SetOnBeforeRender(v js.Value)

	UserData() js.Value
	SetUserData(v js.Value)
	Visible() bool
	SetVisible(v bool)

	Attach(object Object3D) Object3D
	Remove(object Object3D) Object3D
	Clone(recursive bool) Object3D
	Clear() Object3D

	ApplyMatrix(matrix *Matrix4)
	ApplyQuaternion(quaternion *Quaternion) Object3D

	LookAt(vector *Vector3)
	LookAtXYZ(x float64, y float64, z float64)

	Raycast(raycaster *Raycaster, intersects js.Value) js.Value

	RotateOnAxis(axis *Vector3, angle float64) Object3D
	RotateOnWorldAxis(axis *Vector3, angle float64) Object3D

	RotateX(angle float64) Object3D
	RotateY(angle float64) Object3D
	RotateZ(angle float64) Object3D

	SetRotationFromAxisAngle(axis *Vector3, angle float64)
	SetRotationFromEuler(euler *Euler)
	SetRotationFromMatrix(m *Matrix4)
	SetRotationFromQuaternion(q *Quaternion)

	ToJSON(meta js.Value) js.Value

	TranslateX(distance float64) Object3D
	TranslateY(distance float64) Object3D
	TranslateZ(distance float64) Object3D

	Traverse(callback js.Value)
	TraverseAncestors(callback js.Value)
	TraverseVisible(callback js.Value)

	// UpdateMatrix updates the local transform.
	UpdateMatrix()

	// UpdateMatrixWorld updates the global transform of the object and its descendants.
	UpdateMatrixWorld(force bool)

	// UpdateWorldMatrix updates the global transform of the object.
	// updateParents - recursively updates global transform of ancestors.
	// updateChildren - recursively updates global transform of descendants.
	UpdateWorldMatrix(updateParents bool, updateChildren bool)
}

// defaultObject3D extend: [EventDispatcher]
type defaultObject3D struct {
	js.Value
}

// NewObject3D is factory method for Object3D.
func NewObject3D() Object3D {
	return NewObject3DFromJSValue(GetJsObject("Object3D").New())
}

// NewObject3DFromJSValue is factory method for Object3D.
func NewObject3DFromJSValue(delegater js.Value) Object3D {
	return &defaultObject3D{Value: delegater}
}

// JSValue is ...
func (od *defaultObject3D) JSValue() js.Value {
	return od.Value
}

// Add adds object as child of this object.
// An arbitrary number of objects may be added.
// Any current parent on an object passed in here will be removed,
// since an object can have at most one parent.
//
// See Group for info on manually grouping objects.
func (od *defaultObject3D) Add(o Object3D) {
	od.Call("add", o.JSValue())
}

// CastShadow is ...
func (od *defaultObject3D) CastShadow() bool {
	return od.Get("castShadow").Bool()
}

// SetCastShadow is ...
func (od *defaultObject3D) SetCastShadow(v bool) {
	od.Set("castShadow", v)
}

// Children is ...
func (od *defaultObject3D) Children() js.Value {
	return od.Get("children")
}

// SetChildren is ...
func (od *defaultObject3D) SetChildren(v js.Value) {
	od.Set("children", v)
}

// FrustumCulled is ...
func (od *defaultObject3D) FrustumCulled() bool {
	return od.Get("frustumCulled").Bool()
}

// SetFrustumCulled is ...
func (od *defaultObject3D) SetFrustumCulled(v bool) {
	od.Set("frustumCulled", v)
}

// ID is ...
func (od *defaultObject3D) ID() int {
	return od.Get("id").Int()
}

// SetID is ...
func (od *defaultObject3D) SetID(v int) {
	od.Set("id", v)
}

// IsObject3D is ...
func (od *defaultObject3D) IsObject3D() bool {
	return od.Get("isObject3D").Bool()
}

// SetIsObject3D is ...
func (od *defaultObject3D) SetIsObject3D(v bool) {
	od.Set("isObject3D", v)
}

// Layers is ...
func (od *defaultObject3D) Layers() *Layers {
	return &Layers{Value: od.Get("layers")}
}

// SetLayers is ...
func (od *defaultObject3D) SetLayers(v *Layers) {
	od.Set("layers", v.JSValue())
}

// Matrix is ...
func (od *defaultObject3D) Matrix() *Matrix4 {
	return &Matrix4{Value: od.Get("matrix")}
}

// SetMatrix is ...
func (od *defaultObject3D) SetMatrix(v *Matrix4) {
	od.Set("matrix", v.JSValue())
}

// MatrixAutoUpdate is ...
func (od *defaultObject3D) MatrixAutoUpdate() bool {
	return od.Get("matrixAutoUpdate").Bool()
}

// SetMatrixAutoUpdate is ...
func (od *defaultObject3D) SetMatrixAutoUpdate(v bool) {
	od.Set("matrixAutoUpdate", v)
}

// MatrixWorld is ...
func (od *defaultObject3D) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: od.Get("matrixWorld")}
}

// SetMatrixWorld is ...
func (od *defaultObject3D) SetMatrixWorld(v *Matrix4) {
	od.Set("matrixWorld", v.JSValue())
}

// MatrixWorldNeedsUpdate is ...
func (od *defaultObject3D) MatrixWorldNeedsUpdate() bool {
	return od.Get("matrixWorldNeedsUpdate").Bool()
}

// SetMatrixWorldNeedsUpdate is ...
func (od *defaultObject3D) SetMatrixWorldNeedsUpdate(v bool) {
	od.Set("matrixWorldNeedsUpdate", v)
}

// ModelViewMatrix is ...
func (od *defaultObject3D) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: od.Get("modelViewMatrix")}
}

// SetModelViewMatrix is ...
func (od *defaultObject3D) SetModelViewMatrix(v *Matrix4) {
	od.Set("modelViewMatrix", v.JSValue())
}

// Name is ...
func (od *defaultObject3D) Name() string {
	return od.Get("name").String()
}

// SetName is ...
func (od *defaultObject3D) SetName(v string) {
	od.Set("name", v)
}

// NormalMatrix is ...
func (od *defaultObject3D) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: od.Get("normalMatrix")}
}

// SetNormalMatrix is ...
func (od *defaultObject3D) SetNormalMatrix(v *Matrix3) {
	od.Set("normalMatrix", v.JSValue())
}

// OnAfterRender is ...
func (od *defaultObject3D) OnAfterRender() js.Value {
	return od.Get("onAfterRender")
}

// SetOnAfterRender is ...
func (od *defaultObject3D) SetOnAfterRender(v js.Value) {
	od.Set("onAfterRender", v)
}

// OnBeforeRender is ...
func (od *defaultObject3D) OnBeforeRender() js.Value {
	return od.Get("onBeforeRender")
}

// SetOnBeforeRender is ...
func (od *defaultObject3D) SetOnBeforeRender(v js.Value) {
	od.Set("onBeforeRender", v)
}

// Parent is ...
func (od *defaultObject3D) Parent() Object3D {
	return &defaultObject3D{Value: od.Get("parent")}
}

// SetParent is ...
func (od *defaultObject3D) SetParent(v Object3D) {
	od.Set("parent", v.JSValue())
}

// Position is ...
func (od *defaultObject3D) Position() *Vector3 {
	return &Vector3{Value: od.Get("position")}
}

// SetPosition is ...
func (od *defaultObject3D) SetPosition(v *Vector3) {
	od.Set("position", v.JSValue())
}

// Quaternion is ...
func (od *defaultObject3D) Quaternion() *Quaternion {
	return &Quaternion{Value: od.Get("quaternion")}
}

// SetQuaternion is ...
func (od *defaultObject3D) SetQuaternion(v *Quaternion) {
	od.Set("quaternion", v.JSValue())
}

// ReceiveShadow is ...
func (od *defaultObject3D) ReceiveShadow() bool {
	return od.Get("receiveShadow").Bool()
}

// SetReceiveShadow is ...
func (od *defaultObject3D) SetReceiveShadow(v bool) {
	od.Set("receiveShadow", v)
}

// RenderOrder is ...
func (od *defaultObject3D) RenderOrder() float64 {
	return od.Get("renderOrder").Float()
}

// SetRenderOrder is ...
func (od *defaultObject3D) SetRenderOrder(v float64) {
	od.Set("renderOrder", v)
}

// Rotation is ...
func (od *defaultObject3D) Rotation() *Euler {
	return &Euler{Value: od.Get("rotation")}
}

// SetRotation is ...
func (od *defaultObject3D) SetRotation(v *Euler) {
	od.Set("rotation", v.JSValue())
}

// Scale is ...
func (od *defaultObject3D) Scale() *Vector3 {
	return &Vector3{Value: od.Get("scale")}
}

// SetScale is ...
func (od *defaultObject3D) SetScale(v *Vector3) {
	od.Set("scale", v.JSValue())
}

// Type is ...
func (od *defaultObject3D) Type() string {
	return od.Get("type").String()
}

// SetType is ...
func (od *defaultObject3D) SetType(v string) {
	od.Set("type", v)
}

// Up is ...
func (od *defaultObject3D) Up() *Vector3 {
	return &Vector3{Value: od.Get("up")}
}

// SetUp is ...
func (od *defaultObject3D) SetUp(v *Vector3) {
	od.Set("up", v.JSValue())
}

// UserData is ...
func (od *defaultObject3D) UserData() js.Value {
	return od.Get("userData")
}

// SetUserData is ...
func (od *defaultObject3D) SetUserData(v js.Value) {
	od.Set("userData", v)
}

// UUID is ...
func (od *defaultObject3D) UUID() string {
	return od.Get("uuid").String()
}

// SetUUID is ...
func (od *defaultObject3D) SetUUID(v string) {
	od.Set("uuid", v)
}

// Visible is ...
func (od *defaultObject3D) Visible() bool {
	return od.Get("visible").Bool()
}

// SetVisible is ...
func (od *defaultObject3D) SetVisible(v bool) {
	od.Set("visible", v)
}

// DefaultMatrixAutoUpdate is ...
func (od *defaultObject3D) DefaultMatrixAutoUpdate() bool {
	return od.Get("DefaultMatrixAutoUpdate").Bool()
}

// SetDefaultMatrixAutoUpdate is ...
func (od *defaultObject3D) SetDefaultMatrixAutoUpdate(v bool) {
	od.Set("DefaultMatrixAutoUpdate", v)
}

// DefaultUp is ...
func (od *defaultObject3D) DefaultUp() *Vector3 {
	return &Vector3{Value: od.Get("DefaultUp")}
}

// SetDefaultUp is ...
func (od *defaultObject3D) SetDefaultUp(v *Vector3) {
	od.Set("DefaultUp", v.JSValue())
}

// Attach is ...
func (od *defaultObject3D) Attach(object Object3D) Object3D {
	return &defaultObject3D{Value: od.Call("add", object)}
}

// AddEventListener is ...
func (od *defaultObject3D) AddEventListener(typ string, listener js.Value) {
	od.Call("addEventListener", typ, listener)
}

// ApplyMatrix is ...
func (od *defaultObject3D) ApplyMatrix(matrix *Matrix4) {
	od.Call("applyMatrix", matrix.JSValue())
}

// ApplyQuaternion is ...
func (od *defaultObject3D) ApplyQuaternion(quaternion *Quaternion) Object3D {
	return &defaultObject3D{Value: od.Call("applyQuaternion", quaternion)}
}

// Clone is ...
func (od *defaultObject3D) Clone(recursive bool) Object3D {
	return &defaultObject3D{Value: od.Call("clone", recursive)}
}

// Copy is ...
func (od *defaultObject3D) Copy(source *defaultObject3D, recursive bool) Object3D {
	return &defaultObject3D{Value: od.Call("copy", source, recursive)}
}

// DispatchEvent is ...
func (od *defaultObject3D) DispatchEvent(event js.Value) {
	od.Call("dispatchEvent", event)
}

// GetObjectByID is ...
func (od *defaultObject3D) GetObjectByID(id int) Object3D {
	return &defaultObject3D{Value: od.Call("getObjectById", id)}
}

// GetObjectByName is ...
func (od *defaultObject3D) GetObjectByName(name string) Object3D {
	return &defaultObject3D{Value: od.Call("getObjectByName", name)}
}

// GetObjectByProperty is ...
func (od *defaultObject3D) GetObjectByProperty(name string, value string) Object3D {
	return &defaultObject3D{Value: od.Call("getObjectByProperty", name, value)}
}

// GetWorldDirection is ...
func (od *defaultObject3D) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("getWorldDirection", target)}
}

// GetWorldPosition is ...
func (od *defaultObject3D) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("getWorldPosition", target)}
}

// GetWorldQuaternion is ...
func (od *defaultObject3D) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: od.Call("getWorldQuaternion", target)}
}

// GetWorldScale is ...
func (od *defaultObject3D) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("getWorldScale", target)}
}

// HasEventListener is ...
func (od *defaultObject3D) HasEventListener(typ string, listener js.Value) bool {
	return od.Call("hasEventListener", typ, listener).Bool()
}

// LocalToWorld is ...
func (od *defaultObject3D) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("localToWorld", vector)}
}

// LookAt is ...
func (od *defaultObject3D) LookAt(vector *Vector3) {
	od.Call("lookAt", vector.JSValue())
}

// LookAtXYZ is ...
func (od *defaultObject3D) LookAtXYZ(x float64, y float64, z float64) {
	od.Call("lookAt", x, y, z)
}

// Raycast is ...
func (od *defaultObject3D) Raycast(raycaster *Raycaster, intersects js.Value) js.Value {
	return od.Call("raycast", raycaster.JSValue(), intersects)
}

// Remove is ...
func (od *defaultObject3D) Remove(object Object3D) Object3D {
	return &defaultObject3D{Value: od.Call("remove", object)}
}

// RemoveEventListener is ...
func (od *defaultObject3D) RemoveEventListener(typ string, listener js.Value) {
	od.Call("removeEventListener", typ, listener)
}

// RotateOnAxis is ...
func (od *defaultObject3D) RotateOnAxis(axis *Vector3, angle float64) Object3D {
	return &defaultObject3D{Value: od.Call("rotateOnAxis", axis, angle)}
}

// RotateOnWorldAxis is ...
func (od *defaultObject3D) RotateOnWorldAxis(axis *Vector3, angle float64) Object3D {
	return &defaultObject3D{Value: od.Call("rotateOnWorldAxis", axis, angle)}
}

// RotateX is ...
func (od *defaultObject3D) RotateX(angle float64) Object3D {
	return &defaultObject3D{Value: od.Call("rotateX", angle)}
}

// RotateY is ...
func (od *defaultObject3D) RotateY(angle float64) Object3D {
	return &defaultObject3D{Value: od.Call("rotateY", angle)}
}

// RotateZ is ...
func (od *defaultObject3D) RotateZ(angle float64) Object3D {
	return &defaultObject3D{Value: od.Call("rotateZ", angle)}
}

// SetRotationFromAxisAngle is ...
func (od *defaultObject3D) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	od.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}

// SetRotationFromEuler is ...
func (od *defaultObject3D) SetRotationFromEuler(euler *Euler) {
	od.Call("setRotationFromEuler", euler.JSValue())
}

// SetRotationFromMatrix is ...
func (od *defaultObject3D) SetRotationFromMatrix(m *Matrix4) {
	od.Call("setRotationFromMatrix", m.JSValue())
}

// SetRotationFromQuaternion is ...
func (od *defaultObject3D) SetRotationFromQuaternion(q *Quaternion) {
	od.Call("setRotationFromQuaternion", q.JSValue())
}

// ToJSON is ...
func (od *defaultObject3D) ToJSON(meta js.Value) js.Value {
	return od.Call("toJSON", meta)
}

// TranslateOnAxis is ...
func (od *defaultObject3D) TranslateOnAxis(axis *Vector3, distance float64) Object3D {
	return &defaultObject3D{Value: od.Call("translateOnAxis", axis, distance)}
}

// TranslateX is ...
func (od *defaultObject3D) TranslateX(distance float64) Object3D {
	return &defaultObject3D{Value: od.Call("translateX", distance)}
}

// TranslateY is ...
func (od *defaultObject3D) TranslateY(distance float64) Object3D {
	return &defaultObject3D{Value: od.Call("translateY", distance)}
}

// TranslateZ is ...
func (od *defaultObject3D) TranslateZ(distance float64) Object3D {
	return &defaultObject3D{Value: od.Call("translateZ", distance)}
}

// Traverse is ...
func (od *defaultObject3D) Traverse(callback js.Value) {
	od.Call("traverse", callback)
}

// TraverseAncestors is ...
func (od *defaultObject3D) TraverseAncestors(callback js.Value) {
	od.Call("traverseAncestors", callback)
}

// TraverseVisible is ...
func (od *defaultObject3D) TraverseVisible(callback js.Value) {
	od.Call("traverseVisible", callback)
}

// UpdateMatrix is ...
func (od *defaultObject3D) UpdateMatrix() {
	od.Call("updateMatrix")
}

// UpdateMatrixWorld is ...
func (od *defaultObject3D) UpdateMatrixWorld(force bool) {
	od.Call("updateMatrixWorld", force)
}

// UpdateWorldMatrix is ...
func (od *defaultObject3D) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	od.Call("updateWorldMatrix", updateParents, updateChildren)
}

// WorldToLocal is ...
func (od *defaultObject3D) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("worldToLocal", vector)}
}

// Clear is ...
func (od *defaultObject3D) Clear() Object3D {
	return &defaultObject3D{Value: od.Call("clear")}
}
