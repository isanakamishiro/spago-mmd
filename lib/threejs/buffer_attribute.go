package threejs

import "syscall/js"

// BufferAttribute stores data for an attribute (such as vertex positions, face indices, normals, colors, UVs, and any custom attributes ) associated with a BufferGeometry, which allows for more efficient passing of data to the GPU. See that page for details and a usage example.
//
// Data is stored as vectors of any length (defined by itemSize), and in general in the methods outlined below if passing in an index, this is automatically multiplied by the vector length.
type BufferAttribute struct {
	js.Value
}

// NewFloat32BufferAttribute creates ...
func NewFloat32BufferAttribute(arrayLength int, itemSize int) *BufferAttribute {

	ar := js.Global().Get("Float32Array").New(arrayLength * itemSize)

	return &BufferAttribute{
		Value: GetJsObject("BufferAttribute").New(ar, itemSize),
	}
}

// ItemSize gets the length of vectors that are being stored in the array.
func (c *BufferAttribute) ItemSize() int {
	return c.Get("itemSize").Int()
}

// SetNeedsUpdate sets Flag to indicate that this attribute has changed and should be re-sent to the GPU. Set this to true when you modify the value of the array.
// Setting this to true also increments the version.
func (c *BufferAttribute) SetNeedsUpdate(b bool) {
	c.Set("needsUpdate", b)
}

// X returns the x component of the vector at the given index.
func (c *BufferAttribute) X(index int) {
	c.Call("getX", index)
}

// Y returns the y component of the vector at the given index.
func (c *BufferAttribute) Y(index int) {
	c.Call("getY", index)
}

// Z returns the z component of the vector at the given index.
func (c *BufferAttribute) Z(index int) {
	c.Call("getZ", index)
}

// W returns the w component of the vector at the given index.
func (c *BufferAttribute) W(index int) {
	c.Call("getW", index)
}

// SetXY sets the x and y components of the vector at the given index.
func (c *BufferAttribute) SetXY(index int, x float64, y float64) {
	c.Call("setXY", index, x, y)
}

// SetXYZ sets the x, y and z components of the vector at the given index.
func (c *BufferAttribute) SetXYZ(index int, x float64, y float64, z float64) {
	c.Call("setXY", index, x, y)
}
