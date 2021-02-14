package threejs

import (
	"syscall/js"
)

// Vector3 extend: []
type Vector3 struct {
	js.Value
}

// NewVector3 is ...
func NewVector3(x float64, y float64, z float64) *Vector3 {
	return &Vector3{Value: GetJsObject("Vector3").New(x, y, z)}
}

// NewVector3FromJSValue is ...
func NewVector3FromJSValue(v js.Value) *Vector3 {
	return &Vector3{Value: v}
}

// JSValue is ...
func (vv *Vector3) JSValue() js.Value {
	return vv.Value
}

// IsVector3 is ...
func (vv *Vector3) IsVector3() bool {
	return vv.Get("isVector3").Bool()
}

// SetIsVector3 is ...
func (vv *Vector3) SetIsVector3(v bool) {
	vv.Set("isVector3", v)
}

// X is ...
func (vv *Vector3) X() float64 {
	return vv.Get("x").Float()
}

// SetX is ...
func (vv *Vector3) SetX(v float64) {
	vv.Set("x", v)
}

// Y is ...
func (vv *Vector3) Y() float64 {
	return vv.Get("y").Float()
}

// SetY is ...
func (vv *Vector3) SetY(v float64) {
	vv.Set("y", v)
}

// Z is ...
func (vv *Vector3) Z() float64 {
	return vv.Get("z").Float()
}

// SetZ is ...
func (vv *Vector3) SetZ(v float64) {
	vv.Set("z", v)
}

// Add is ...
func (vv *Vector3) Add(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("add", a, b)}
}

// AddScalar is ...
func (vv *Vector3) AddScalar(s float64) *Vector3 {
	return &Vector3{Value: vv.Call("addScalar", s)}
}

// AddScaledVector is ...
func (vv *Vector3) AddScaledVector(v *Vector3, s float64) *Vector3 {
	return &Vector3{Value: vv.Call("addScaledVector", v, s)}
}

// AddVectors is ...
func (vv *Vector3) AddVectors(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("addVectors", a, b)}
}

// AngleTo is ...
func (vv *Vector3) AngleTo(v *Vector3) float64 {
	return vv.Call("angleTo", v).Float()
}

// ApplyAxisAngle is ...
func (vv *Vector3) ApplyAxisAngle(axis *Vector3, angle float64) *Vector3 {
	return &Vector3{Value: vv.Call("applyAxisAngle", axis, angle)}
}

// ApplyEuler is ...
func (vv *Vector3) ApplyEuler(euler *Euler) *Vector3 {
	return &Vector3{Value: vv.Call("applyEuler", euler)}
}

// ApplyMatrix3 is ...
func (vv *Vector3) ApplyMatrix3(m *Matrix3) *Vector3 {
	return &Vector3{Value: vv.Call("applyMatrix3", m)}
}

// ApplyMatrix4 is ...
func (vv *Vector3) ApplyMatrix4(m *Matrix4) *Vector3 {
	return &Vector3{Value: vv.Call("applyMatrix4", m)}
}

// ApplyQuaternion is ...
func (vv *Vector3) ApplyQuaternion(q *Quaternion) *Vector3 {
	return &Vector3{Value: vv.Call("applyQuaternion", q)}
}

// Ceil is ...
func (vv *Vector3) Ceil() *Vector3 {
	return &Vector3{Value: vv.Call("ceil")}
}

// Clamp is ...
func (vv *Vector3) Clamp(min *Vector3, max *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("clamp", min, max)}
}

// ClampLength is ...
func (vv *Vector3) ClampLength(min float64, max float64) *Vector3 {
	return &Vector3{Value: vv.Call("clampLength", min, max)}
}

// ClampScalar is ...
func (vv *Vector3) ClampScalar(min float64, max float64) *Vector3 {
	return &Vector3{Value: vv.Call("clampScalar", min, max)}
}

// Clone is ...
func (vv *Vector3) Clone() *Vector3 {
	return &Vector3{Value: vv.Call("clone")}
}

// Copy is ...
func (vv *Vector3) Copy(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("copy", v)}
}

// Cross is ...
func (vv *Vector3) Cross(a *Vector3, w *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("cross", a, w)}
}

// CrossVectors is ...
func (vv *Vector3) CrossVectors(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("crossVectors", a, b)}
}

// DistanceTo is ...
func (vv *Vector3) DistanceTo(v *Vector3) float64 {
	return vv.Call("distanceTo", v).Float()
}

// DistanceToManhattan is ...
func (vv *Vector3) DistanceToManhattan(v *Vector3) float64 {
	return vv.Call("distanceToManhattan", v).Float()
}

// DistanceToSquared is ...
func (vv *Vector3) DistanceToSquared(v *Vector3) float64 {
	return vv.Call("distanceToSquared", v).Float()
}

// Divide is ...
func (vv *Vector3) Divide(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("divide", v)}
}

// DivideScalar is ...
func (vv *Vector3) DivideScalar(s float64) *Vector3 {
	return &Vector3{Value: vv.Call("divideScalar", s)}
}

// Dot is ...
func (vv *Vector3) Dot(v *Vector3) float64 {
	return vv.Call("dot", v).Float()
}

// Equals is ...
func (vv *Vector3) Equals(v *Vector3) bool {
	return vv.Call("equals", v).Bool()
}

// Floor is ...
func (vv *Vector3) Floor() *Vector3 {
	return &Vector3{Value: vv.Call("floor")}
}

// FromArray is ...
func (vv *Vector3) FromArray(xyz js.Value, offset int) *Vector3 {
	return &Vector3{Value: vv.Call("fromArray", xyz, offset)}
}

// FromBufferAttribute is ...
// func (vv *Vector3) FromBufferAttribute(attribute *BufferAttribute, index int, offset int) *Vector3 {
// 	return &Vector3{Value: vv.Call("fromBufferAttribute", attribute, index, offset)}
// }

// GetComponent is ...
func (vv *Vector3) GetComponent(index int) float64 {
	return vv.Call("getComponent", index).Float()
}

// Length is ...
func (vv *Vector3) Length() float64 {
	return vv.Call("length").Float()
}

// LengthManhattan is ...
func (vv *Vector3) LengthManhattan() float64 {
	return vv.Call("lengthManhattan").Float()
}

// LengthSq is ...
func (vv *Vector3) LengthSq() float64 {
	return vv.Call("lengthSq").Float()
}

// Lerp is ...
func (vv *Vector3) Lerp(v *Vector3, alpha float64) *Vector3 {
	return &Vector3{Value: vv.Call("lerp", v, alpha)}
}

// LerpVectors is ...
func (vv *Vector3) LerpVectors(v1 *Vector3, v2 *Vector3, alpha float64) *Vector3 {
	return &Vector3{Value: vv.Call("lerpVectors", v1, v2, alpha)}
}

// ManhattanDistanceTo is ...
func (vv *Vector3) ManhattanDistanceTo(v *Vector3) float64 {
	return vv.Call("manhattanDistanceTo", v).Float()
}

// ManhattanLength is ...
func (vv *Vector3) ManhattanLength() float64 {
	return vv.Call("manhattanLength").Float()
}

// Max is ...
func (vv *Vector3) Max(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("max", v)}
}

// Min is ...
func (vv *Vector3) Min(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("min", v)}
}

// Multiply is ...
func (vv *Vector3) Multiply(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("multiply", v)}
}

// MultiplyScalar is ...
func (vv *Vector3) MultiplyScalar(s float64) *Vector3 {
	return &Vector3{Value: vv.Call("multiplyScalar", s)}
}

// MultiplyVectors is ...
func (vv *Vector3) MultiplyVectors(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("multiplyVectors", a, b)}
}

// Negate is ...
func (vv *Vector3) Negate() *Vector3 {
	return &Vector3{Value: vv.Call("negate")}
}

// Normalize is ...
func (vv *Vector3) Normalize() *Vector3 {
	return &Vector3{Value: vv.Call("normalize")}
}

// Project is ...
// func (vv *Vector3) Project(camera cameras.Camera) *Vector3 {
// 	return &Vector3{Value: vv.Call("project", camera.JSValue())}
// }

// ProjectOnPlane is ...
func (vv *Vector3) ProjectOnPlane(planeNormal *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("projectOnPlane", planeNormal)}
}

// ProjectOnVector is ...
func (vv *Vector3) ProjectOnVector(v *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("projectOnVector", v)}
}

// Reflect is ...
func (vv *Vector3) Reflect(vector *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("reflect", vector)}
}

// Round is ...
func (vv *Vector3) Round() *Vector3 {
	return &Vector3{Value: vv.Call("round")}
}

// RoundToZero is ...
func (vv *Vector3) RoundToZero() *Vector3 {
	return &Vector3{Value: vv.Call("roundToZero")}
}

// Set2 is ...
func (vv *Vector3) Set2(x float64, y float64, z float64) *Vector3 {
	return &Vector3{Value: vv.Call("set", x, y, z)}
}

// SetComponent is ...
func (vv *Vector3) SetComponent(index int, value float64) *Vector3 {
	return &Vector3{Value: vv.Call("setComponent", index, value)}
}

// SetFromCylindrical is ...
// func (vv *Vector3) SetFromCylindrical(s *Cylindrical) *Vector3 {
// 	return &Vector3{Value: vv.Call("setFromCylindrical", s)}
// }

// SetFromMatrixColumn is ...
func (vv *Vector3) SetFromMatrixColumn(matrix *Matrix4, index int) *Vector3 {
	return &Vector3{Value: vv.Call("setFromMatrixColumn", matrix, index)}
}

// SetFromMatrixPosition is ...
func (vv *Vector3) SetFromMatrixPosition(m *Matrix4) *Vector3 {
	return &Vector3{Value: vv.Call("setFromMatrixPosition", m)}
}

// SetFromMatrixScale is ...
func (vv *Vector3) SetFromMatrixScale(m *Matrix4) *Vector3 {
	return &Vector3{Value: vv.Call("setFromMatrixScale", m)}
}

// SetFromSpherical is ...
// func (vv *Vector3) SetFromSpherical(s *Spherical) *Vector3 {
// 	return &Vector3{Value: vv.Call("setFromSpherical", s)}
// }

// SetLength is ...
func (vv *Vector3) SetLength(l float64) *Vector3 {
	return &Vector3{Value: vv.Call("setLength", l)}
}

// SetScalar is ...
func (vv *Vector3) SetScalar(scalar float64) *Vector3 {
	return &Vector3{Value: vv.Call("setScalar", scalar)}
}

// SetX2 is ...
func (vv *Vector3) SetX2(x float64) *Vector3 {
	return &Vector3{Value: vv.Call("setX", x)}
}

// SetY2 is ...
func (vv *Vector3) SetY2(y float64) *Vector3 {
	return &Vector3{Value: vv.Call("setY", y)}
}

// SetZ2 is ...
func (vv *Vector3) SetZ2(z float64) *Vector3 {
	return &Vector3{Value: vv.Call("setZ", z)}
}

// Sub is ...
func (vv *Vector3) Sub(a *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("sub", a)}
}

// SubScalar is ...
func (vv *Vector3) SubScalar(s float64) *Vector3 {
	return &Vector3{Value: vv.Call("subScalar", s)}
}

// SubVectors is ...
func (vv *Vector3) SubVectors(a *Vector3, b *Vector3) *Vector3 {
	return &Vector3{Value: vv.Call("subVectors", a, b)}
}

// ToArray is ...
func (vv *Vector3) ToArray(xyz js.Value, offset int) js.Value {
	return vv.Call("toArray", xyz, offset)
}

// ToArray2 is ...
func (vv *Vector3) ToArray2(xyz js.Value, offset int) js.Value {
	return vv.Call("toArray", xyz, offset)
}

// TransformDirection is ...
func (vv *Vector3) TransformDirection(m *Matrix4) *Vector3 {
	return &Vector3{Value: vv.Call("transformDirection", m)}
}

// Unproject is ...
// func (vv *Vector3) Unproject(camera cameras.Camera) *Vector3 {
// 	return &Vector3{Value: vv.Call("unproject", camera.JSValue())}
// }
