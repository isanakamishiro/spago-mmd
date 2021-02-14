package texts

// import (
// 	"app/lib/threejs"
// 	"app/frontend/lib/threejs/materials"
// 	"app/frontend/lib/threejs/mathutils"
// 	"app/frontend/lib/threejs/objects"
// 	"app/frontend/lib/threejs/textures"
// )

// const (
// 	defaultRenderOffsetX   = 0.0
// 	defaultRenderOffsetY   = 0.0
// 	defaultRenderFontScale = 0.005
// 	defaultRenderAutoWrap  = true
// 	defaultRenderWrapWidth = 3.0
// 	defaultRenderColor     = threejs.ColorValue(0xffffff)

// 	defaultAnimateSpeed = 10.0
// )

// type wordDictionary map[string]*FontTextureMetrics
// type spriteGroup []objects.Sprite

// // TextSprite is ...
// type TextSprite struct {
// 	threejs.Object3D

// 	textureAtlas *FontTextureAtlas

// 	renderOffsetX   float64
// 	renderOffsetY   float64
// 	renderFontScale float64
// 	renderAutoWrap  bool
// 	renderWrapWidth float64
// 	renderColor     threejs.Color
// 	renderText      string

// 	dictionary wordDictionary
// 	spriteList spriteGroup

// 	animateFrame float64
// 	animateSpeed float64
// }

// // NewTextSprite creates TextSprite.
// func NewTextSprite(f *FontTextureAtlas) *TextSprite {
// 	return &TextSprite{
// 		Object3D:        threejs.NewObject3D(),
// 		textureAtlas:    f,
// 		renderOffsetX:   defaultRenderOffsetX,
// 		renderOffsetY:   defaultRenderOffsetY,
// 		renderFontScale: defaultRenderFontScale,
// 		renderAutoWrap:  defaultRenderAutoWrap,
// 		renderWrapWidth: defaultRenderWrapWidth,
// 		renderColor:     threejs.NewColorFromColorValue(defaultRenderColor),
// 		renderText:      "",
// 		animateSpeed:    defaultAnimateSpeed,
// 	}
// }

// // Color gets color
// func (c *TextSprite) Color() threejs.Color {
// 	return c.renderColor
// }

// // SetColor set text color.
// func (c *TextSprite) SetColor(col threejs.Color) {
// 	c.renderColor = col
// }

// // AnimationSpeed gets animation speed
// func (c *TextSprite) AnimationSpeed() float64 {
// 	return c.animateSpeed
// }

// // SetAnimationSpeed sets animation speed.
// func (c *TextSprite) SetAnimationSpeed(speed float64) {
// 	c.animateSpeed = speed
// }

// // WrapWidth gets wrap width
// func (c *TextSprite) WrapWidth() float64 {
// 	return c.renderWrapWidth
// }

// // SetWrapWidth sets animation speed.
// func (c *TextSprite) SetWrapWidth(w float64) {
// 	c.renderWrapWidth = w
// }

// // AutoWrap gets auto wrap.
// func (c *TextSprite) AutoWrap() bool {
// 	return c.renderAutoWrap
// }

// // SetAutoWrap sets auto wrap
// func (c *TextSprite) SetAutoWrap(b bool) {
// 	c.renderAutoWrap = b
// }

// // BuildText is ...
// func (c *TextSprite) BuildText(text string) {

// 	c.ClearText()

// 	c.renderText = text
// 	// c.dictionary = c.createTextTextureAtlas(text)
// 	c.spriteList = c.createTextMesh(text, c.dictionary)

// 	// for debug.
// 	// for _, cvt := range c.canvasList {
// 	// 	js.Global().Get("document").
// 	// 		Call("querySelector", "#dev_canvas").
// 	// 		Call("appendChild", cvt.Canvas().JSValue())

// 	// 	break
// 	// }

// }

// // ClearText is ...
// func (c *TextSprite) ClearText() {

// 	c.animateFrame = 0
// 	c.renderText = ""

// 	c.dictionary = nil
// 	c.spriteList = nil

// 	c.Clear()

// }

// // Restart is ...
// func (c *TextSprite) Restart() {
// 	c.animateFrame = 0
// }

// // Update is ...
// func (c *TextSprite) Update(delta float64) {

// 	maxFrame := float64(len(c.spriteList))
// 	c.animateFrame += (delta * c.animateSpeed)
// 	if c.animateFrame > maxFrame {
// 		c.animateFrame = maxFrame
// 	}

// 	// Spriteにアニメーションをかける
// 	for n, sp := range c.spriteList {

// 		offset := float64(n)
// 		opacity := c.animateFrame - offset
// 		if opacity < 0.0 {
// 			opacity = 0.0
// 		} else if opacity > 1.0 {
// 			opacity = 1.0
// 		}
// 		sp.Material().SetOpacity(opacity)

// 		p := sp.Parent()
// 		scaleX := mathutils.Lerp(10.0, 1.0, opacity)
// 		scaleY := mathutils.Lerp(0.1, 1.0, opacity)
// 		p.Scale().Set2(scaleX, scaleY, 1.0)
// 		// rot := mathutils.Lerp(math.Pi, 0.0, opacity)
// 		// p.RotateZ(rot)

// 		visible := (opacity > 0)
// 		p.SetVisible(visible)

// 	}
// }

// // createSpriteMaterial creates new SpriteMaterial from CanvasTexture.
// func (c *TextSprite) createSpriteMaterial(cvt textures.CanvasTexture) materials.SpriteMaterial {
// 	return materials.NewSpriteMaterial(map[string]interface{}{
// 		"map":   cvt,
// 		"color": c.renderColor.Hex(),
// 		// "opacity": 0.5,
// 		// "transparent": false,
// 		// "premultipliedAlpha": true,
// 		// "alphaTest": 0.2,
// 	})
// }

// // createTextMesh is ...
// func (c *TextSprite) createTextMesh(text string, dictionary wordDictionary) spriteGroup {

// 	base := c
// 	outputText := text

// 	renderPosX := 0.0

// 	fontScale := c.renderFontScale
// 	renderOffsetX := c.renderOffsetX
// 	renderOffsetY := c.renderOffsetY
// 	wrap := c.renderAutoWrap
// 	renderWidthLimit := c.renderWrapWidth

// 	var currentLineGroupRef *spriteGroup = &spriteGroup{}
// 	var LineGroupArray []*spriteGroup
// 	LineGroupArray = append(LineGroupArray, currentLineGroupRef)

// 	var totalSprites spriteGroup = spriteGroup{}

// 	// Text Render
// 	for _, wd := range outputText {

// 		word := string([]rune{wd})

// 		// 改行文字処理
// 		if word == "\n" {
// 			renderPosX = 0.0
// 			currentLineGroupRef = &spriteGroup{}
// 			LineGroupArray = append(LineGroupArray, currentLineGroupRef)

// 			continue
// 		}

// 		// 辞書なし
// 		wm, ok := c.textureAtlas.Find(word)
// 		if !ok {
// 			continue
// 		}

// 		fontWidth := fontScale * wm.absoluteWidth
// 		fontHeight := fontScale * wm.absoluteHeight

// 		// auto wrap
// 		if wrap && ((renderPosX + fontWidth) > renderWidthLimit) {
// 			renderPosX = 0.0
// 			currentLineGroupRef = &spriteGroup{}
// 			LineGroupArray = append(LineGroupArray, currentLineGroupRef)
// 		}

// 		// 非表示の特殊文字
// 		if word == "\n" {
// 			continue
// 		}

// 		// Texture
// 		// cvTexture.Offset().Set2(wm.U, wm.V)
// 		// cvTexture.Repeat().Set2(wm.Width, wm.Height)
// 		// cvTexture.SetNeedsUpdate(true)

// 		// UV mapping
// 		ba := threejs.NewFloat32BufferAttribute(4, 3)
// 		ba.SetXY(0, wm.U, wm.V)
// 		ba.SetXY(1, wm.U+wm.Width, wm.V)
// 		ba.SetXY(2, wm.U+wm.Width, wm.V+wm.Height)
// 		ba.SetXY(3, wm.U, wm.V+wm.Height)
// 		ba.SetNeedsUpdate(true)

// 		// Material
// 		material := c.createSpriteMaterial(wm.CanvasTexture)

// 		// Sprite Mesh
// 		// sprite := objects.NewSprite(wm.Material)
// 		sprite := objects.NewSprite(material)
// 		// Geometry setting
// 		sprite.CloneGeometryToOwn()
// 		sprite.Geometry().SetAttribute("uv", ba)

// 		sprite.Scale().Set2(
// 			fontWidth,
// 			fontHeight,
// 			1,
// 		)
// 		// sprite.Position().Set2(
// 		// 	fontWidth/2, //+renderPosX,
// 		// 	fontHeight/2,
// 		// 	0,
// 		// )

// 		graph := threejs.NewObject3D()
// 		// graph.Position().SetX(renderPosX)
// 		graph.Position().SetX(fontWidth/2 + renderPosX)
// 		graph.Position().SetY(fontHeight / 2)

// 		graph.Add(sprite)

// 		base.Add(graph)

// 		totalSprites = append(totalSprites, sprite)

// 		*currentLineGroupRef = append(*currentLineGroupRef, sprite)

// 		// 次の文字の表示位置調整
// 		renderPosX = renderPosX + fontWidth + renderOffsetX
// 	}

// 	// 縦位置調整

// 	renderPosY := 0.0
// 	for _, group := range LineGroupArray {

// 		// 行の最大高さを計算
// 		maxHeight := 0.0
// 		for _, s := range *group {
// 			h := s.Scale().Y()
// 			if maxHeight < h {
// 				maxHeight = h
// 			}
// 		}

// 		// 縦位置設定
// 		for _, s := range *group {
// 			pa := s.Parent()
// 			// s.Position().SetY(s.Position().Y() + renderPosY - maxHeight)
// 			pa.Position().SetY(pa.Position().Y() + renderPosY - maxHeight)
// 		}

// 		// 次の行位置調整
// 		renderPosY = renderPosY - maxHeight - renderOffsetY
// 	}

// 	return totalSprites

// }
