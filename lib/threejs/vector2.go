package threejs

import (
	"syscall/js"
)

// Vector is Vector interface in Three.js
type Vector interface {
	Add(v Vector, w Vector) *Vector
	AddScalar(scalar float64) *Vector
	AddScaledVector(vector Vector, scale float64) *Vector
	AddVectors(a Vector, b Vector) *Vector
	Clone() *Vector
	Copy(v Vector) *Vector
	DistanceTo(v Vector) float64
	DistanceToSquared(v Vector) float64
	DivideScalar(s float64) *Vector
	Dot(v Vector) float64
	Equals(v Vector) bool
	GetComponent(index int) float64
	Length() float64
	LengthSq() float64
	Lerp(v Vector, alpha float64) *Vector
	MultiplyScalar(s float64) *Vector
	Negate() *Vector
	Normalize() *Vector
	Set(args js.Value) *Vector
	SetComponent(index int, value float64) *Vector
	SetLength(l float64) *Vector
	SetScalar(scalar float64) *Vector
	Sub(v Vector) *Vector
	SubVectors(a Vector, b Vector) *Vector
}

// Vector2 extend: []
type Vector2 struct {
	js.Value
}

// NewVector2 is ...
func NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{Value: GetJsObject("Vector2").New(x, y)}
}

// JSValue is ...
func (vv *Vector2) JSValue() js.Value {
	return vv.Value
}

// Height is ...
func (vv *Vector2) Height() float64 {
	return vv.Get("height").Float()
}

// SetHeight is ...
func (vv *Vector2) SetHeight(v float64) {
	vv.Set("height", v)
}

// IsVector2 is ...
func (vv *Vector2) IsVector2() bool {
	return vv.Get("isVector2").Bool()
}

// SetIsVector2 is ...
func (vv *Vector2) SetIsVector2(v bool) {
	vv.Set("isVector2", v)
}

// Width is ...
func (vv *Vector2) Width() float64 {
	return vv.Get("width").Float()
}

// SetWidth is ...
func (vv *Vector2) SetWidth(v float64) {
	vv.Set("width", v)
}

// X is ...
func (vv *Vector2) X() float64 {
	return vv.Get("x").Float()
}

// SetX is ...
func (vv *Vector2) SetX(v float64) {
	vv.Set("x", v)
}

// Y is ...
func (vv *Vector2) Y() float64 {
	return vv.Get("y").Float()
}

// SetY is ...
func (vv *Vector2) SetY(v float64) {
	vv.Set("y", v)
}

// Add is ...
func (vv *Vector2) Add(v *Vector2, w *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("add", v, w)}
}

// AddScalar is ...
func (vv *Vector2) AddScalar(s float64) *Vector2 {
	return &Vector2{Value: vv.Call("addScalar", s)}
}

// AddScaledVector is ...
func (vv *Vector2) AddScaledVector(v *Vector2, s float64) *Vector2 {
	return &Vector2{Value: vv.Call("addScaledVector", v, s)}
}

// AddVectors is ...
func (vv *Vector2) AddVectors(a *Vector2, b *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("addVectors", a, b)}
}

// Angle is ...
func (vv *Vector2) Angle() float64 {
	return vv.Call("angle").Float()
}

// ApplyMatrix3 is ...
func (vv *Vector2) ApplyMatrix3(m *Matrix3) *Vector2 {
	return &Vector2{Value: vv.Call("applyMatrix3", m)}
}

// Ceil is ...
func (vv *Vector2) Ceil() *Vector2 {
	return &Vector2{Value: vv.Call("ceil")}
}

// Clamp is ...
func (vv *Vector2) Clamp(min *Vector2, max *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("clamp", min, max)}
}

// ClampLength is ...
func (vv *Vector2) ClampLength(min float64, max float64) *Vector2 {
	return &Vector2{Value: vv.Call("clampLength", min, max)}
}

// ClampScalar is ...
func (vv *Vector2) ClampScalar(min float64, max float64) *Vector2 {
	return &Vector2{Value: vv.Call("clampScalar", min, max)}
}

// Clone is ...
func (vv *Vector2) Clone() *Vector2 {
	return &Vector2{Value: vv.Call("clone")}
}

// Copy is ...
func (vv *Vector2) Copy(v *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("copy", v)}
}

// Cross is ...
func (vv *Vector2) Cross(v *Vector2) int {
	return vv.Call("cross", v).Int()
}

// DistanceTo is ...
func (vv *Vector2) DistanceTo(v *Vector2) float64 {
	return vv.Call("distanceTo", v).Float()
}

// DistanceToManhattan is ...
func (vv *Vector2) DistanceToManhattan(v *Vector2) float64 {
	return vv.Call("distanceToManhattan", v).Float()
}

// DistanceToSquared is ...
func (vv *Vector2) DistanceToSquared(v *Vector2) float64 {
	return vv.Call("distanceToSquared", v).Float()
}

// Divide is ...
func (vv *Vector2) Divide(v *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("divide", v)}
}

// DivideScalar is ...
func (vv *Vector2) DivideScalar(s float64) *Vector2 {
	return &Vector2{Value: vv.Call("divideScalar", s)}
}

// Dot is ...
func (vv *Vector2) Dot(v *Vector2) float64 {
	return vv.Call("dot", v).Float()
}

// Equals is ...
func (vv *Vector2) Equals(v *Vector2) bool {
	return vv.Call("equals", v).Bool()
}

// Floor is ...
func (vv *Vector2) Floor() *Vector2 {
	return &Vector2{Value: vv.Call("floor")}
}

// FromArray is ...
func (vv *Vector2) FromArray(array js.Value, offset int) *Vector2 {
	return &Vector2{Value: vv.Call("fromArray", array, offset)}
}

// FromBufferAttribute is ...
// func (vv *Vector2) FromBufferAttribute(attribute *BufferAttribute, index int) *Vector2 {
// 	return &Vector2{Value: vv.Call("fromBufferAttribute", attribute, index)}
// }

// GetComponent is ...
func (vv *Vector2) GetComponent(index int) float64 {
	return vv.Call("getComponent", index).Float()
}

// Length is ...
func (vv *Vector2) Length() float64 {
	return vv.Call("length").Float()
}

// LengthManhattan is ...
func (vv *Vector2) LengthManhattan() float64 {
	return vv.Call("lengthManhattan").Float()
}

// LengthSq is ...
func (vv *Vector2) LengthSq() float64 {
	return vv.Call("lengthSq").Float()
}

// Lerp is ...
func (vv *Vector2) Lerp(v *Vector2, alpha float64) *Vector2 {
	return &Vector2{Value: vv.Call("lerp", v, alpha)}
}

// LerpVectors is ...
func (vv *Vector2) LerpVectors(v1 *Vector2, v2 *Vector2, alpha float64) *Vector2 {
	return &Vector2{Value: vv.Call("lerpVectors", v1, v2, alpha)}
}

// ManhattanDistanceTo is ...
func (vv *Vector2) ManhattanDistanceTo(v *Vector2) float64 {
	return vv.Call("manhattanDistanceTo", v).Float()
}

// ManhattanDistanceTo2 is ...
func (vv *Vector2) ManhattanDistanceTo2(v *Vector2) float64 {
	return vv.Call("manhattanDistanceTo", v).Float()
}

// ManhattanLength is ...
func (vv *Vector2) ManhattanLength() float64 {
	return vv.Call("manhattanLength").Float()
}

// ManhattanLength2 is ...
func (vv *Vector2) ManhattanLength2() float64 {
	return vv.Call("manhattanLength").Float()
}

// Max is ...
func (vv *Vector2) Max(v *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("max", v)}
}

// Min is ...
func (vv *Vector2) Min(v *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("min", v)}
}

// Multiply is ...
func (vv *Vector2) Multiply(v *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("multiply", v)}
}

// MultiplyScalar is ...
func (vv *Vector2) MultiplyScalar(scalar float64) *Vector2 {
	return &Vector2{Value: vv.Call("multiplyScalar", scalar)}
}

// Negate is ...
func (vv *Vector2) Negate() *Vector2 {
	return &Vector2{Value: vv.Call("negate")}
}

// Normalize is ...
func (vv *Vector2) Normalize() *Vector2 {
	return &Vector2{Value: vv.Call("normalize")}
}

// RotateAround is ...
func (vv *Vector2) RotateAround(center *Vector2, angle float64) *Vector2 {
	return &Vector2{Value: vv.Call("rotateAround", center, angle)}
}

// Round is ...
func (vv *Vector2) Round() *Vector2 {
	return &Vector2{Value: vv.Call("round")}
}

// RoundToZero is ...
func (vv *Vector2) RoundToZero() *Vector2 {
	return &Vector2{Value: vv.Call("roundToZero")}
}

// Set2 is ...
func (vv *Vector2) Set2(x float64, y float64) *Vector2 {
	return &Vector2{Value: vv.Call("set", x, y)}
}

// SetComponent is ...
func (vv *Vector2) SetComponent(index int, value float64) *Vector2 {
	return &Vector2{Value: vv.Call("setComponent", index, value)}
}

// SetLength is ...
func (vv *Vector2) SetLength(length float64) *Vector2 {
	return &Vector2{Value: vv.Call("setLength", length)}
}

// SetScalar is ...
func (vv *Vector2) SetScalar(scalar float64) *Vector2 {
	return &Vector2{Value: vv.Call("setScalar", scalar)}
}

// SetX2 is ...
func (vv *Vector2) SetX2(x float64) *Vector2 {
	return &Vector2{Value: vv.Call("setX", x)}
}

// SetY2 is ...
func (vv *Vector2) SetY2(y float64) *Vector2 {
	return &Vector2{Value: vv.Call("setY", y)}
}

// Sub is ...
func (vv *Vector2) Sub(v *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("sub", v)}
}

// SubScalar is ...
func (vv *Vector2) SubScalar(s float64) *Vector2 {
	return &Vector2{Value: vv.Call("subScalar", s)}
}

// SubVectors is ...
func (vv *Vector2) SubVectors(a *Vector2, b *Vector2) *Vector2 {
	return &Vector2{Value: vv.Call("subVectors", a, b)}
}

// ToArray is ...
func (vv *Vector2) ToArray(array js.Value, offset int) js.Value {
	return vv.Call("toArray", array, offset)
}

// ToArray2 is ...
func (vv *Vector2) ToArray2(array js.Value, offset int) js.Value {
	return vv.Call("toArray", array, offset)
}
