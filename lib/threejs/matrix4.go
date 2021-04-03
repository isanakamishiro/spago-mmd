package threejs

import (
	"syscall/js"
)

// Matrix4 extend: []
type Matrix4 struct {
	js.Value
}

// NewMatrix4 is factory method for Matrix4.
func NewMatrix4() *Matrix4 {
	return &Matrix4{Value: Threejs("Matrix4").New()}
}

// JSValue is ...
func (mm *Matrix4) JSValue() js.Value {
	return mm.Value
}

// Elements is ...
func (mm *Matrix4) Elements() js.Value {
	return mm.Get("elements")
}

// SetElements is ...
func (mm *Matrix4) SetElements(v js.Value) {
	mm.Set("elements", v)
}

// ApplyToBuffer is ...
// func (mm *Matrix4) ApplyToBuffer(buffer *BufferAttribute, offset int, length int) *BufferAttribute {
// 	return &BufferAttribute{Value: mm.Call("applyToBuffer", buffer, offset, length)}
// }

// ApplyToBufferAttribute is ...
// func (mm *Matrix4) ApplyToBufferAttribute(attribute *BufferAttribute) *BufferAttribute {
// 	return &BufferAttribute{Value: mm.Call("applyToBufferAttribute", attribute)}
// }

// Clone is ...
func (mm *Matrix4) Clone() *Matrix4 {
	return &Matrix4{Value: mm.Call("clone")}
}

// Compose is ...
func (mm *Matrix4) Compose(translation *Vector3, rotation *Quaternion, scale *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("compose", translation, rotation, scale)}
}

// Copy is ...
func (mm *Matrix4) Copy(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("copy", m)}
}

// CopyPosition is ...
func (mm *Matrix4) CopyPosition(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("copyPosition", m)}
}

// CrossVector is ...
func (mm *Matrix4) CrossVector(v js.Value) {
	mm.Call("crossVector", v)
}

// Decompose is ...
func (mm *Matrix4) Decompose(translation *Vector3, rotation *Quaternion, scale *Vector3) js.Value {
	return mm.Call("decompose", translation, rotation, scale)
}

// Determinant is ...
func (mm *Matrix4) Determinant() float64 {
	return mm.Call("determinant").Float()
}

// Equals is ...
func (mm *Matrix4) Equals(matrix *Matrix4) bool {
	return mm.Call("equals", matrix).Bool()
}

// ExtractBasis is ...
func (mm *Matrix4) ExtractBasis(xAxis *Vector3, yAxis *Vector3, zAxis *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("extractBasis", xAxis, yAxis, zAxis)}
}

// ExtractPosition is ...
func (mm *Matrix4) ExtractPosition(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("extractPosition", m)}
}

// ExtractRotation is ...
func (mm *Matrix4) ExtractRotation(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("extractRotation", m)}
}

// FlattenToArrayOffset is ...
func (mm *Matrix4) FlattenToArrayOffset(array js.Value, offset int) js.Value {
	return mm.Call("flattenToArrayOffset", array, offset)
}

// FromArray is ...
func (mm *Matrix4) FromArray(array js.Value, offset int) *Matrix4 {
	return &Matrix4{Value: mm.Call("fromArray", array, offset)}
}

// GetInverse is ...
func (mm *Matrix4) GetInverse(m *Matrix4, throwOnDegeneratee bool) *Matrix4 {
	return &Matrix4{Value: mm.Call("getInverse", m, throwOnDegeneratee)}
}

// GetMaxScaleOnAxis is ...
func (mm *Matrix4) GetMaxScaleOnAxis() float64 {
	return mm.Call("getMaxScaleOnAxis").Float()
}

// Identity is ...
func (mm *Matrix4) Identity() *Matrix4 {
	return &Matrix4{Value: mm.Call("identity")}
}

// LookAt is ...
func (mm *Matrix4) LookAt(eye *Vector3, target *Vector3, up *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("lookAt", eye, target, up)}
}

// MakeBasis is ...
func (mm *Matrix4) MakeBasis(xAxis *Vector3, yAxis *Vector3, zAxis *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeBasis", xAxis, yAxis, zAxis)}
}

// MakeOrthographic is ...
func (mm *Matrix4) MakeOrthographic(left float64, right float64, top float64, bottom float64, near float64, far float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeOrthographic", left, right, top, bottom, near, far)}
}

// MakePerspective is ...
func (mm *Matrix4) MakePerspective(left float64, right float64, bottom float64, top float64, near float64, far float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makePerspective", left, right, bottom, top, near, far)}
}

// MakePerspective2 is ...
func (mm *Matrix4) MakePerspective2(fov float64, aspect float64, near float64, far float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makePerspective", fov, aspect, near, far)}
}

// MakeRotationAxis is ...
func (mm *Matrix4) MakeRotationAxis(axis *Vector3, angle float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationAxis", axis, angle)}
}

// MakeRotationFromEuler is ...
func (mm *Matrix4) MakeRotationFromEuler(euler *Euler) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationFromEuler", euler)}
}

// MakeRotationFromQuaternion is ...
func (mm *Matrix4) MakeRotationFromQuaternion(q *Quaternion) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationFromQuaternion", q)}
}

// MakeRotationX is ...
func (mm *Matrix4) MakeRotationX(theta float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationX", theta)}
}

// MakeRotationY is ...
func (mm *Matrix4) MakeRotationY(theta float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationY", theta)}
}

// MakeRotationZ is ...
func (mm *Matrix4) MakeRotationZ(theta float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeRotationZ", theta)}
}

// MakeScale is ...
func (mm *Matrix4) MakeScale(x float64, y float64, z float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeScale", x, y, z)}
}

// MakeTranslation is ...
func (mm *Matrix4) MakeTranslation(x float64, y float64, z float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("makeTranslation", x, y, z)}
}

// Multiply is ...
func (mm *Matrix4) Multiply(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("multiply", m)}
}

// MultiplyMatrices is ...
func (mm *Matrix4) MultiplyMatrices(a *Matrix4, b *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("multiplyMatrices", a, b)}
}

// MultiplyScalar is ...
func (mm *Matrix4) MultiplyScalar(s float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("multiplyScalar", s)}
}

// MultiplyToArray is ...
func (mm *Matrix4) MultiplyToArray(a *Matrix4, b *Matrix4, r js.Value) *Matrix4 {
	return &Matrix4{Value: mm.Call("multiplyToArray", a, b, r)}
}

// MultiplyVector3 is ...
func (mm *Matrix4) MultiplyVector3(v js.Value) js.Value {
	return mm.Call("multiplyVector3", v)
}

// MultiplyVector3Array is ...
func (mm *Matrix4) MultiplyVector3Array(array js.Value) js.Value {
	return mm.Call("multiplyVector3Array", array)
}

// MultiplyVector4 is ...
func (mm *Matrix4) MultiplyVector4(v js.Value) js.Value {
	return mm.Call("multiplyVector4", v)
}

// Premultiply is ...
func (mm *Matrix4) Premultiply(m *Matrix4) *Matrix4 {
	return &Matrix4{Value: mm.Call("premultiply", m)}
}

// RotateAxis is ...
func (mm *Matrix4) RotateAxis(v js.Value) {
	mm.Call("rotateAxis", v)
}

// Scale is ...
func (mm *Matrix4) Scale(v *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("scale", v)}
}

// Set2 is ...
func (mm *Matrix4) Set2(n11 float64, n12 float64, n13 float64, n14 float64, n21 float64, n22 float64, n23 float64, n24 float64, n31 float64, n32 float64, n33 float64, n34 float64, n41 float64, n42 float64, n43 float64, n44 float64) *Matrix4 {
	return &Matrix4{Value: mm.Call("set", n11, n12, n13, n14, n21, n22, n23, n24, n31, n32, n33, n34, n41, n42, n43, n44)}
}

// SetPosition is ...
func (mm *Matrix4) SetPosition(v *Vector3) *Matrix4 {
	return &Matrix4{Value: mm.Call("setPosition", v)}
}

// SetRotationFromQuaternion is ...
func (mm *Matrix4) SetRotationFromQuaternion(q *Quaternion) *Matrix4 {
	return &Matrix4{Value: mm.Call("setRotationFromQuaternion", q)}
}

// ToArray is ...
func (mm *Matrix4) ToArray() js.Value {
	return mm.Call("toArray")
}

// Transpose is ...
func (mm *Matrix4) Transpose() *Matrix4 {
	return &Matrix4{Value: mm.Call("transpose")}
}
