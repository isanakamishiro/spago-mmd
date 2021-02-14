package threejs

import "syscall/js"

// Rectangle is rectangle object insted of DOMRect
type Rectangle struct {
	js.Value
}

// NewRectangleFromJSValue creates Rectangle based on js.Value
func NewRectangleFromJSValue(v js.Value) *Rectangle {
	return &Rectangle{
		Value: v,
	}
}

// Left is ...
func (c *Rectangle) Left() float64 {
	return c.Get("left").Float()
}

// Top is ...
func (c *Rectangle) Top() float64 {
	return c.Get("top").Float()
}

// Right is ...
func (c *Rectangle) Right() float64 {
	return c.Get("right").Float()
}

// Bottom is ...
func (c *Rectangle) Bottom() float64 {
	return c.Get("bottom").Float()
}

// Width is ...
func (c *Rectangle) Width() float64 {
	return c.Get("width").Float()
}

// Height is ...
func (c *Rectangle) Height() float64 {
	return c.Get("height").Float()
}
