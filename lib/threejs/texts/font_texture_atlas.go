package texts

import (
	"app/lib/threejs"
	"app/lib/threejs/textures"
	"fmt"
	"math"
)

const (
	defaultCanvasWidth         = 512
	defaultCanvasHeight        = 512
	defaultFontSize            = 40
	defaultFontName            = "serif"
	defaultFontColor           = "white"
	defaultTextureAtlasOffsetX = 1
	defaultTextureAtlasOffsetY = 1

	defaultTextureMagFilter = threejs.LinearFilter
	defaultTextureMinFilter = threejs.LinearMipmapLinearFilter
)

// FontTextureMetrics is ...
type FontTextureMetrics struct {
	Word           string
	U              float64
	V              float64
	Width          float64
	Height         float64
	absoluteWidth  float64
	absoluteHeight float64

	// Canvas   *textures.Canvas
	CanvasTexture textures.CanvasTexture
	// Material      materials.SpriteMaterial
}

// fontTextureMetricsDictionary is ...
type fontTextureMetricsDictionary map[string]*FontTextureMetrics

// FontTextureAtlas is texture atlas for text sprite.
type FontTextureAtlas struct {
	canvasWidth         int
	canvasHeight        int
	fontSize            int
	fontName            string
	fontColor           string
	textureAtlasOffsetX int
	textureAtlasOffsetY int

	textureMagFilter threejs.TextureFilter
	textureMinFilter threejs.TextureFilter

	fontDictionary fontTextureMetricsDictionary
}

// NewFontTextureAtlas creates FontTextureAtlas.
func NewFontTextureAtlas(text string, options ...FontTextureAtlasOption) (*FontTextureAtlas, error) {

	atlas := FontTextureAtlas{
		canvasWidth:         defaultCanvasWidth,
		canvasHeight:        defaultCanvasHeight,
		fontSize:            defaultFontSize,
		fontName:            defaultFontName,
		fontColor:           defaultFontColor,
		textureAtlasOffsetX: defaultTextureAtlasOffsetX,
		textureAtlasOffsetY: defaultTextureAtlasOffsetY,
		textureMagFilter:    defaultTextureMagFilter,
		textureMinFilter:    defaultTextureMinFilter,
	}

	for _, opt := range options {
		opt(&atlas)
	}

	atlas.buildTextureAtlas(text)

	return &atlas, nil
}

// Find gets font metrics with specified word(1 character).
func (c *FontTextureAtlas) Find(word string) (metrics *FontTextureMetrics, ok bool) {
	if v, ok := c.fontDictionary[word]; ok {
		return v, true
	}
	return nil, false
}

func (c *FontTextureAtlas) setCanvasWidth(w int) {
	c.canvasWidth = w
}

func (c *FontTextureAtlas) setCanvasHeight(h int) {
	c.canvasHeight = h
}

// setFontName sets font name.
func (c *FontTextureAtlas) setFontName(fontName string) {
	c.fontName = fontName
}

// setFontSize sets font size.
func (c *FontTextureAtlas) setFontSize(fontSize int) {
	c.fontSize = fontSize
}

func (c *FontTextureAtlas) setFontColor(col string) {
	c.fontColor = col
}

func (c *FontTextureAtlas) setOffsetX(x int) {
	c.textureAtlasOffsetX = x
}

func (c *FontTextureAtlas) setOffsetY(y int) {
	c.textureAtlasOffsetY = y
}

func (c *FontTextureAtlas) setTextureMagFilter(filter threejs.TextureFilter) {
	c.textureMagFilter = filter
}

func (c *FontTextureAtlas) setTextureMinFilter(filter threejs.TextureFilter) {
	c.textureMinFilter = filter
}

// createDefaultCanvasTexture create CanvasTexture object and initialize with default settings.
func (c *FontTextureAtlas) createDefaultCanvasTexture() textures.CanvasTexture {

	cv := textures.NewCanvas()
	// js.Global().Get("document").Call("querySelector", "#dev_canvas").Call("appendChild", cv.Value)
	cv.SetSize(c.canvasWidth, c.canvasHeight)
	cv.Context2D().SetFillStyle("rgba(0, 0, 0, 0.0)")
	cv.Context2D().FillRect(0, 0, c.canvasWidth, c.canvasHeight)
	cv.Context2D().SetFont(fmt.Sprintf("%dpx %s", c.fontSize, c.fontName))
	cv.Context2D().JSValue().Set("textAlign", "left")
	cv.Context2D().JSValue().Set("textBaseline", "ideographic")
	cv.Context2D().SetFillStyle(c.fontColor)

	cvTexture := textures.NewCanvasTexture(cv)
	cvTexture.SetMinFilter(c.textureMinFilter)
	cvTexture.SetMagFilter(c.textureMagFilter)
	cvTexture.SetWrapS(threejs.RepeatWrapping)
	cvTexture.SetWrapT(threejs.RepeatWrapping)

	return cvTexture
}

// BuildTextureAtlas build texture atlas with specified text.
func (c *FontTextureAtlas) buildTextureAtlas(text string) {

	var dictionary fontTextureMetricsDictionary = fontTextureMetricsDictionary{}
	var currentCVT textures.CanvasTexture = c.createDefaultCanvasTexture()

	mx := 0
	my := 0
	maxLineHeight := 0
	xOffset := c.textureAtlasOffsetX
	yOffset := c.textureAtlasOffsetY

	for _, ch := range text {
		word := string([]rune{ch})

		// すでに辞書に登録されていればスキップ
		if _, ok := dictionary[word]; ok {
			continue
		}

		cv := currentCVT.Canvas()
		canvasWidth := cv.Width()
		canvasHeight := cv.Height()
		metrics := cv.Context2D().MeasureText(word)

		w := int(math.Abs(metrics.ActualBoundingBoxLeft)+math.Abs(metrics.ActualBoundingBoxRight)) + 1
		if w < int(metrics.Width) {
			w = int(metrics.Width)
		}
		h := int(math.Abs(metrics.ActualBoundingBoxAscent)+math.Abs(metrics.ActualBoundingBoxDescent)) + 1

		// 同行の最大高さを更新
		if h > maxLineHeight {
			maxLineHeight = h
		}

		// 通常文字なら、レンダリング
		if h > 0 {
			// 横幅がオーバーなら、位置を改行
			if int(mx+w) > canvasWidth {
				mx = 0
				my = my + maxLineHeight + yOffset
				maxLineHeight = 0
			}
			// 描画範囲の高さがオーバーなら、新しいキャンバスを作成
			if int(my+h) > canvasHeight {
				mx = 0
				my = 0
				maxLineHeight = 0

				currentCVT = c.createDefaultCanvasTexture()
			}

			// Canvasへ文字を描画
			cctx := currentCVT.Canvas().Context2D()
			cctx.FillText(
				word,
				mx,
				my+h,
			)
			// cctx.StrokeRect(mx, my, w, h)

			// 文字位置辞書を登録
			dictionary[word] = &FontTextureMetrics{
				Word:           word,
				U:              float64(mx) / float64(canvasWidth),
				V:              1.0 - (float64(my+h) / float64(canvasHeight)),
				Width:          float64(w) / float64(canvasWidth),
				Height:         float64(h) / float64(canvasHeight),
				absoluteWidth:  float64(w),
				absoluteHeight: float64(h),
				CanvasTexture:  currentCVT,
			}

			// 文字位置をずらす
			mx = mx + w + xOffset
		}

	}

	// 結果を格納
	c.fontDictionary = dictionary
}
