package threejs

import (
	"syscall/js"
)

// Matrix is ...
type Matrix interface {
	Clone() *Matrix
	Copy(m *Matrix) *Matrix
	Determinant() float64
	GetInverse(matrix Matrix, throwOnInvertible bool) Matrix
	Identity() Matrix
	MultiplyScalar(s float64) Matrix
	Transpose() Matrix
}

// Matrix3 extend: []
type Matrix3 struct {
	js.Value
}

// NewMatrix3 is factory method for Matrix3.
func NewMatrix3() *Matrix3 {
	return &Matrix3{Value: GetJsObject("Matrix3").New()}
}

// JSValue is ...
func (mm *Matrix3) JSValue() js.Value {
	return mm.Value
}

// Elements is ...
func (mm *Matrix3) Elements() js.Value {
	return mm.Get("elements")
}

// SetElements is ...
func (mm *Matrix3) SetElements(v js.Value) {
	mm.Set("elements", v)
}

// ApplyToBuffer is ...
// func (mm *Matrix3) ApplyToBuffer(buffer *BufferAttribute, offset int, length int) *BufferAttribute {
// 	return &BufferAttribute{Value: mm.Call("applyToBuffer", buffer, offset, length)}
// }

// ApplyToBufferAttribute is ...
// func (mm *Matrix3) ApplyToBufferAttribute(attribute *BufferAttribute) *BufferAttribute {
// 	return &BufferAttribute{Value: mm.Call("applyToBufferAttribute", attribute)}
// }

// Clone is ...
func (mm *Matrix3) Clone() *Matrix3 {
	return &Matrix3{Value: mm.Call("clone")}
}

// Copy is ...
func (mm *Matrix3) Copy(m *Matrix3) *Matrix3 {
	return &Matrix3{Value: mm.Call("copy", m)}
}

// Determinant is ...
func (mm *Matrix3) Determinant() float64 {
	return mm.Call("determinant").Float()
}

// FlattenToArrayOffset is ...
func (mm *Matrix3) FlattenToArrayOffset(array js.Value, offset int) js.Value {
	return mm.Call("flattenToArrayOffset", array, offset)
}

// FromArray is ...
func (mm *Matrix3) FromArray(array js.Value, offset int) *Matrix3 {
	return &Matrix3{Value: mm.Call("fromArray", array, offset)}
}

// GetInverse is ...
func (mm *Matrix3) GetInverse(matrix *Matrix3, throwOnDegenerate bool) *Matrix3 {
	return &Matrix3{Value: mm.Call("getInverse", matrix, throwOnDegenerate)}
}

// GetInverse2 is ...
func (mm *Matrix3) GetInverse2(matrix *Matrix4, throwOnDegenerate bool) *Matrix3 {
	return &Matrix3{Value: mm.Call("getInverse", matrix, throwOnDegenerate)}
}

// GetNormalMatrix is ...
func (mm *Matrix3) GetNormalMatrix(matrix4 *Matrix4) *Matrix3 {
	return &Matrix3{Value: mm.Call("getNormalMatrix", matrix4)}
}

// Identity is ...
func (mm *Matrix3) Identity() *Matrix3 {
	return &Matrix3{Value: mm.Call("identity")}
}

// Multiply is ...
func (mm *Matrix3) Multiply(m *Matrix3) *Matrix3 {
	return &Matrix3{Value: mm.Call("multiply", m)}
}

// MultiplyMatrices is ...
func (mm *Matrix3) MultiplyMatrices(a *Matrix3, b *Matrix3) *Matrix3 {
	return &Matrix3{Value: mm.Call("multiplyMatrices", a, b)}
}

// MultiplyScalar is ...
func (mm *Matrix3) MultiplyScalar(s float64) *Matrix3 {
	return &Matrix3{Value: mm.Call("multiplyScalar", s)}
}

// MultiplyVector3 is ...
func (mm *Matrix3) MultiplyVector3(vector *Vector3) js.Value {
	return mm.Call("multiplyVector3", vector)
}

// MultiplyVector3Array is ...
func (mm *Matrix3) MultiplyVector3Array(a js.Value) js.Value {
	return mm.Call("multiplyVector3Array", a)
}

// Premultiply is ...
func (mm *Matrix3) Premultiply(m *Matrix3) *Matrix3 {
	return &Matrix3{Value: mm.Call("premultiply", m)}
}

// Set2 is ...
func (mm *Matrix3) Set2(n11 float64, n12 float64, n13 float64, n21 float64, n22 float64, n23 float64, n31 float64, n32 float64, n33 float64) *Matrix3 {
	return &Matrix3{Value: mm.Call("set", n11, n12, n13, n21, n22, n23, n31, n32, n33)}
}

// SetFromMatrix4 is ...
func (mm *Matrix3) SetFromMatrix4(m *Matrix4) *Matrix3 {
	return &Matrix3{Value: mm.Call("setFromMatrix4", m)}
}

// ToArray is ...
func (mm *Matrix3) ToArray() js.Value {
	return mm.Call("toArray")
}

// Transpose is ...
func (mm *Matrix3) Transpose() *Matrix3 {
	return &Matrix3{Value: mm.Call("transpose")}
}

// TransposeIntoArray is ...
func (mm *Matrix3) TransposeIntoArray(r js.Value) js.Value {
	return mm.Call("transposeIntoArray", r)
}
