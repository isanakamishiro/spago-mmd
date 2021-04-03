package threejs

import (
	"syscall/js"
)

// ColorName is X11 color name - all 140 color names are supported.
// Note the lack of CamelCase in the name
type ColorName string

// ColorValue is numeric value of the RGB component of the color. e.g. 0xffffff.
type ColorValue int

// JSValue return the javascript object.
func (c ColorValue) JSValue() js.Value {
	return js.ValueOf(int(c))
}

// HSL is ...
// type HSL interface {
// }

// Color is ...
type Color interface {
	JSValue() js.Value

	R() float64
	G() float64
	B() float64

	Hex() int
	SetHex(hex int)
	HexString() string

	// SetRGB sets this color from RGB values.
	//
	// r — Red channel value between 0.0 and 1.0.
	// g — Green channel value between 0.0 and 1.0.
	// b — Blue channel value between 0.0 and 1.0.
	SetRGB(r, g, b float64)

	// SetHSL sets color from HSL values.
	//
	// h — hue value between 0.0 and 1.0
	// s — saturation value between 0.0 and 1.0
	// l — lightness value between 0.0 and 1.0
	SetHSL(h, s, l float64)
}

// Color extend: []
type colorImp struct {
	js.Value
}

// NewColor is constructor for color object.
func NewColor() Color {
	return &colorImp{Value: Threejs("Color").New()}
}

// NewColorFromJSValue is constructor for color object.
func NewColorFromJSValue(value js.Value) Color {
	return &colorImp{Value: value}
}

// NewColorFromColorValue is constructor for color object from Hexadecimal value
func NewColorFromColorValue(color ColorValue) Color {
	return &colorImp{Value: Threejs("Color").New(float64(color))}
}

// NewColorFromColorName is constructor for color object from ColorName
func NewColorFromColorName(colorName ColorName) Color {
	return &colorImp{Value: Threejs("Color").New(string(colorName))}
}

// NewColorFromRGB is constructor for color object from separate RGB values between 0 and 1
func NewColorFromRGB(r, g, b float64) Color {
	return &colorImp{Value: Threejs("Color").New(r, g, b)}
}

// JSValue is ...
func (cc *colorImp) JSValue() js.Value {
	return cc.Value
}

// R is ...
func (cc *colorImp) R() float64 {
	return cc.Get("r").Float()
}

// G is ...
func (cc *colorImp) G() float64 {
	return cc.Get("g").Float()
}

// B is ...
func (cc *colorImp) B() float64 {
	return cc.Get("b").Float()
}

// Hex is ...
func (cc *colorImp) Hex() int {
	return cc.Call("getHex").Int()
}

// SetHex is ...
func (cc *colorImp) SetHex(hex int) {
	cc.Call("setHex", hex)
}

// HexString is ...
func (cc *colorImp) HexString() string {
	return cc.Call("getHexString").String()
}

func (cc *colorImp) SetRGB(r, g, b float64) {
	cc.Call("setRGB", r, g, b)
}

func (cc *colorImp) SetHSL(h, s, l float64) {
	cc.Call("setHSL", h, s, l)
}

// func (cc *Color) B() float64 {
// 	return cc.Get("b").Float()
// }
// func (cc *Color) SetB(v float64) {
// 	cc.Set("b", v)
// }
// func (cc *Color) G() float64 {
// 	return cc.Get("g").Float()
// }
// func (cc *Color) SetG(v float64) {
// 	cc.Set("g", v)
// }
// func (cc *Color) R() float64 {
// 	return cc.Get("r").Float()
// }
// func (cc *Color) SetR(v float64) {
// 	cc.Set("r", v)
// }
// func (cc *Color) Add(color *Color) *Color {
// 	return &Color{Value: cc.Call("add", color)}
// }
// func (cc *Color) AddColors(color1 *Color, color2 *Color) *Color {
// 	return &Color{Value: cc.Call("addColors", color1, color2)}
// }
// func (cc *Color) AddScalar(s float64) *Color {
// 	return &Color{Value: cc.Call("addScalar", s)}
// }
// func (cc *Color) Clone() *Color {
// 	return &Color{Value: cc.Call("clone")}
// }
// func (cc *Color) ConvertGammaToLinear() *Color {
// 	return &Color{Value: cc.Call("convertGammaToLinear")}
// }
// func (cc *Color) ConvertLinearToGamma() *Color {
// 	return &Color{Value: cc.Call("convertLinearToGamma")}
// }
// func (cc *Color) Copy(color *Color) *Color {
// 	return &Color{Value: cc.Call("copy", color)}
// }
// func (cc *Color) CopyGammaToLinear(color *Color, gammaFactor float64) *Color {
// 	return &Color{Value: cc.Call("copyGammaToLinear", color, gammaFactor)}
// }
// func (cc *Color) CopyLinearToGamma(color *Color, gammaFactor float64) *Color {
// 	return &Color{Value: cc.Call("copyLinearToGamma", color, gammaFactor)}
// }
// func (cc *Color) Equals(color *Color) bool {
// 	return cc.Call("equals", color).Bool()
// }
// func (cc *Color) FromArray(rgb js.Value, offset int) *Color {
// 	return &Color{Value: cc.Call("fromArray", rgb, offset)}
// }
// func (cc *Color) GetHSL(target HSL) HSL {
// 	return HSL(cc.Call("getHSL", target))
// }
// func (cc *Color) GetHex() int {
// 	return cc.Call("getHex").Int()
// }
// func (cc *Color) GetHexString() string {
// 	return cc.Call("getHexString").String()
// }
// func (cc *Color) GetStyle() string {
// 	return cc.Call("getStyle").String()
// }
// func (cc *Color) Lerp(color *Color, alpha float64) *Color {
// 	return &Color{Value: cc.Call("lerp", color, alpha)}
// }
// func (cc *Color) LerpHSL(color *Color, alpha float64) *Color {
// 	return &Color{Value: cc.Call("lerpHSL", color, alpha)}
// }
// func (cc *Color) Multiply(color *Color) *Color {
// 	return &Color{Value: cc.Call("multiply", color)}
// }
// func (cc *Color) MultiplyScalar(s float64) *Color {
// 	return &Color{Value: cc.Call("multiplyScalar", s)}
// }
// func (cc *Color) OffsetHSL(h float64, s float64, l float64) *Color {
// 	return &Color{Value: cc.Call("offsetHSL", h, s, l)}
// }
// func (cc *Color) Set2(color *Color) *Color {
// 	return &Color{Value: cc.Call("set", color)}
// }
// func (cc *Color) Set3(color int) *Color {
// 	return &Color{Value: cc.Call("set", color)}
// }
// func (cc *Color) Set4(color string) *Color {
// 	return &Color{Value: cc.Call("set", color)}
// }
// func (cc *Color) SetHSL(h float64, s float64, l float64) *Color {
// 	return &Color{Value: cc.Call("setHSL", h, s, l)}
// }
// func (cc *Color) SetHex(hex int) *Color {
// 	return &Color{Value: cc.Call("setHex", hex)}
// }
// func (cc *Color) SetRGB(r float64, g float64, b float64) *Color {
// 	return &Color{Value: cc.Call("setRGB", r, g, b)}
// }
// func (cc *Color) SetScalar(scalar float64) *Color {
// 	return &Color{Value: cc.Call("setScalar", scalar)}
// }
// func (cc *Color) SetStyle(style string) *Color {
// 	return &Color{Value: cc.Call("setStyle", style)}
// }
// func (cc *Color) Sub(color *Color) *Color {
// 	return &Color{Value: cc.Call("sub", color)}
// }
// func (cc *Color) ToArray(array js.Value, offset int) js.Value {
// 	return cc.Call("toArray", array, offset)
// }
// func (cc *Color) ToArray2(xyz js.Value, offset int) js.Value {
// 	return cc.Call("toArray", xyz, offset)
// }
