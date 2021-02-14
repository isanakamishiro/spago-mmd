package threejs

import "syscall/js"

// FogBase contains the parameters that define linear fog, i.e., that grows linearly denser with the distance.
type FogBase interface {
	JSValue() js.Value

	// Name gets optional name of the object (doesn't need to be unique). Default is an empty string.
	Name() string
	// SetName sets optional name of the object (doesn't need to be unique). Default is an empty string.
	SetName(v string)

	// Color gets fog color. Example: If set to black, far away objects will be rendered black.
	Color() Color

	// SetColor sets fog color. Example: If set to black, far away objects will be rendered black.
	SetColor(v Color)
}

type fogBaseImp struct {
	js.Value
}

// NewFogBaseFromJSValue creates FogBase from javascript object.
func NewFogBaseFromJSValue(v js.Value) FogBase {
	return &fogBaseImp{
		Value: v,
	}
}

func (c *fogBaseImp) JSValue() js.Value {
	return c.Value
}

// Name gets optional name of the object (doesn't need to be unique). Default is an empty string.
func (c *fogBaseImp) Name() string {
	return c.Get("name").String()
}

// SetName sets optional name of the object (doesn't need to be unique). Default is an empty string.
func (c *fogBaseImp) SetName(v string) {
	c.Set("name", v)
}

// Color gets fog color. Example: If set to black, far away objects will be rendered black.
func (c *fogBaseImp) Color() Color {
	return NewColorFromJSValue(
		c.Get("color"),
	)
}

// SetColor sets fog color. Example: If set to black, far away objects will be rendered black.
func (c *fogBaseImp) SetColor(v Color) {
	c.Set("color", v.JSValue())
}
