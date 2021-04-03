package texture

import (
	"app/lib/threejs"
	"syscall/js"
)

// CanvasTexture is a texture from a canvas element.
// This is almost the same as the base Texture class, except that it sets needsUpdate to true immediately.
type CanvasTexture interface {
	threejs.Texture

	Canvas() *Canvas
}

// Canvas is ...
type Canvas struct {
	js.Value

	context *CanvasContext
}

// CanvasContext is ...
type CanvasContext struct {
	js.Value
}

// CanvasTextMetrics is ...
type CanvasTextMetrics struct {
	Width                    float64
	ActualBoundingBoxLeft    float64
	ActualBoundingBoxRight   float64
	ActualBoundingBoxAscent  float64
	ActualBoundingBoxDescent float64
	// FontBoundingBoxAscent    float64
	// FontBoundingBoxDescent   float64
	// EmHeightAscent           float64
	// EmHeightDescent          float64
	// HangingBaseline          float64
	// AlphabeticBaseline       float64
	// IdeographicBaseline      float64
}

type canvasTextureImp struct {
	threejs.Texture

	canvas *Canvas
}

// constructor

// NewCanvasTexture creates CanvasTexture from Canvas object.
func NewCanvasTexture(c *Canvas) CanvasTexture {
	return &canvasTextureImp{
		Texture: threejs.NewDefaultTextureFromJSValue(
			threejs.Threejs("CanvasTexture").New(c.JSValue()),
		),
		canvas: c,
	}
}

// NewCanvas creates Canvas object.
func NewCanvas() *Canvas {

	jsCanvas := js.Global().Get("document").Call("createElement", "canvas")
	jsCanvasContext := jsCanvas.Call("getContext", "2d")

	return &Canvas{
		Value: jsCanvas,
		context: &CanvasContext{
			Value: jsCanvasContext,
		},
	}
}

// --- Methods for Canvas
func (c *canvasTextureImp) Canvas() *Canvas {
	return c.canvas
}

// Context2D gets canvas 2D context.
func (c *Canvas) Context2D() *CanvasContext {
	return c.context
}

// SetWidth is ...
func (c *Canvas) SetWidth(w int) {
	c.Set("width", w)
}

// Width is ...
func (c *Canvas) Width() int {
	return c.Get("width").Int()
}

// SetHeight is ...
func (c *Canvas) SetHeight(h int) {
	c.Set("height", h)
}

// Height is ...
func (c *Canvas) Height() int {
	return c.Get("height").Int()
}

// SetSize is shortcut for SetWidth/Height.
func (c *Canvas) SetSize(w, h int) {
	c.SetWidth(w)
	c.SetHeight(h)
}

// --- Methods for CanvasContext

// SetFillStyle is ...
//
// color parameter can be set the following format:
// 	キーワードの使用 (blue や transparent など)
// 	RGB 立方体座標方式の使用 (#+16進数値や、rgb() や rgba() の関数記法によって)
// 	HSL 円柱座標方式 の使用 (hsl() や hsla() の関数記法によって)
func (c *CanvasContext) SetFillStyle(color string) {
	c.Set("fillStyle", color)
}

// SetContextSize is ...
// func (c *CanvasContext) SetContextSize(w, h int) {
// 	c.Get("canvas").Set("width", w)
// 	c.Get("canvas").Set("height", h)
// }

// FillRect is ...
func (c *CanvasContext) FillRect(x, y, width, height int) {
	c.Call("fillRect", x, y, width, height)
}

// StrokeRect is ...
func (c *CanvasContext) StrokeRect(x, y, width, height int) {
	c.Call("strokeRect", x, y, width, height)
}

// FillText is ...
func (c *CanvasContext) FillText(text string, x, y int) {
	c.Call("fillText", text, x, y)
}

// StrokeText is ...
func (c *CanvasContext) StrokeText(text string, x, y int) {
	c.Call("strokeText", text, x, y)
}

// MeasureText is ...
func (c *CanvasContext) MeasureText(text string) CanvasTextMetrics {
	metrics := c.Call("measureText", text)

	return CanvasTextMetrics{
		Width:                  metrics.Get("width").Float(),
		ActualBoundingBoxLeft:  metrics.Get("actualBoundingBoxLeft").Float(),
		ActualBoundingBoxRight: metrics.Get("actualBoundingBoxRight").Float(),
		// FontBoundingBoxAscent:    metrics.Get("fontBoundingBoxAscent").Float(),
		// FontBoundingBoxDescent:   metrics.Get("fontBoundingBoxDescent").Float(),
		ActualBoundingBoxAscent:  metrics.Get("actualBoundingBoxAscent").Float(),
		ActualBoundingBoxDescent: metrics.Get("actualBoundingBoxDescent").Float(),
		// EmHeightAscent:           metrics.Get("emHeightAscent").Float(),
		// EmHeightDescent:          metrics.Get("emHeightDescent").Float(),
		// HangingBaseline:          metrics.Get("hangingBaseline").Float(),
		// AlphabeticBaseline:       metrics.Get("alphabeticBaseline").Float(),
		// IdeographicBaseline:      metrics.Get("ideographicBaseline").Float(),
	}
}

// SetFont is ...
// e.g. '10px serif'
func (c *CanvasContext) SetFont(font string) {
	c.Set("font", font)
}
