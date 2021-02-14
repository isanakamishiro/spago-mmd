package texts

import (
	"app/lib/threejs"
	"app/lib/threejs/materials"
	"app/lib/threejs/objects"
)

const (
	defaultCharacterPositionOffsetX = 0.0
	defaultCharacterPositionOffsetY = 0.0
	defaultRenderFontScale          = 0.001
	defaultRenderAutoWrap           = false
	defaultRenderWrapWidth          = 3.0
	defaultRenderColor              = threejs.ColorValue(0xffffff)
)

// CharacterMeshFactoryMethod is factory method to create character mesh.
type CharacterMeshFactoryMethod func(threejs.Color, *FontTextureMetrics) (CharacterMesh, error)

// CharacterMesh is character mesh.
type CharacterMesh interface {
	threejs.Object3D

	Material() threejs.Material
}

// TextMesh is ...
type TextMesh struct {
	threejs.Object3D

	textureAtlas *FontTextureAtlas
	renderText   string

	meshWidth  float64
	meshHeight float64
	bottom     float64 // 文字描画時の下端

	characterPositionOffsetX float64
	characterPositionOffsetY float64
	renderFontScale          float64
	renderAutoWrap           bool
	renderWrapWidth          float64
	renderColor              threejs.Color

	characterMeshList []CharacterMesh
	// objectList objectGroup
}

// NewTextMesh is ...
func NewTextMesh(text string, f *FontTextureAtlas, factoryMethod CharacterMeshFactoryMethod, options []TextMeshOption) *TextMesh {

	mesh := &TextMesh{
		Object3D:     threejs.NewObject3D(),
		textureAtlas: f,

		meshWidth:  1,
		meshHeight: 1,

		characterPositionOffsetX: defaultCharacterPositionOffsetX,
		characterPositionOffsetY: defaultCharacterPositionOffsetY,
		renderFontScale:          defaultRenderFontScale,
		renderAutoWrap:           defaultRenderAutoWrap,
		renderWrapWidth:          defaultRenderWrapWidth,
		renderColor:              threejs.NewColorFromColorValue(defaultRenderColor),
		renderText:               text,
	}

	for _, opt := range options {
		opt(mesh)
	}

	mesh.buildTextMesh(text, factoryMethod)

	return mesh
}

// NewTextSprite creates TextSprite object and initialize it.
func NewTextSprite(text string, f *FontTextureAtlas, options ...TextMeshOption) *TextMesh {

	var cmfm CharacterMeshFactoryMethod = func(c threejs.Color, ftm *FontTextureMetrics) (CharacterMesh, error) {
		// Material
		material := materials.NewSpriteMaterial(map[string]interface{}{
			"map":   ftm.CanvasTexture,
			"color": c.Hex(),
		})

		// UV mapping buffer attribute
		ba := threejs.NewFloat32BufferAttribute(4, 3)
		ba.SetXY(0, ftm.U, ftm.V)
		ba.SetXY(1, ftm.U+ftm.Width, ftm.V)
		ba.SetXY(2, ftm.U+ftm.Width, ftm.V+ftm.Height)
		ba.SetXY(3, ftm.U, ftm.V+ftm.Height)
		ba.SetNeedsUpdate(true)

		sprite := objects.NewSprite(material)
		sprite.CloneGeometryToOwn()
		sprite.Geometry().SetAttribute("uv", ba)

		return sprite, nil
	}

	return NewTextMesh(text, f, cmfm, options)
}

// Bottom gets botton position for text.
func (c *TextMesh) Bottom() float64 {
	return c.bottom
}

// createTextMesh is ...
func (c *TextMesh) buildTextMesh(text string, meshFactoryMethod CharacterMeshFactoryMethod) error {

	base := c // set delegate object
	outputText := text

	fontScale := c.renderFontScale
	color := c.renderColor

	// coord managed variables
	startPosX := -(c.meshWidth / 2.0)
	startPosY := (c.meshHeight / 2.0)
	renderPosX := startPosX
	lineWidth := 0.0
	renderOffsetX := c.characterPositionOffsetX
	renderOffsetY := c.characterPositionOffsetY

	// auto wrapping setting
	wrap := c.renderAutoWrap
	renderWidthLimit := c.meshWidth

	type objectGroup []threejs.Object3D
	var currentLineGroupRef *objectGroup = &objectGroup{}
	var LineGroupArray []*objectGroup
	LineGroupArray = append(LineGroupArray, currentLineGroupRef)

	// 改行時の変数初期化関数
	newline := func() {
		renderPosX = startPosX
		lineWidth = 0.0
		currentLineGroupRef = &objectGroup{}
		LineGroupArray = append(LineGroupArray, currentLineGroupRef)
	}

	// Text Render
	for _, wd := range outputText {

		word := string([]rune{wd})

		// 改行文字処理
		if word == "\n" {
			newline()

			continue
		}

		// 辞書なし
		wm, ok := c.textureAtlas.Find(word)
		if !ok {
			continue
		}

		// 文字サイズ設定
		fontWidth := fontScale * wm.absoluteWidth
		fontHeight := fontScale * wm.absoluteHeight

		// auto wrapping
		if wrap && ((lineWidth + fontWidth) > renderWidthLimit) {
			newline()
		}

		// Mesh
		obj, err := meshFactoryMethod(color, wm)
		if err != nil {
			continue
		}

		// 座標変換用のオブジェクト
		graph := threejs.NewObject3D()
		graph.Scale().Set2(
			fontWidth,
			fontHeight,
			1,
		)
		graph.Position().SetX(fontWidth/2 + renderPosX)
		graph.Position().SetY(0)

		graph.Add(obj)
		base.Add(graph)

		// add obj to line group / mesh list
		*currentLineGroupRef = append(*currentLineGroupRef, graph)
		c.characterMeshList = append(c.characterMeshList, obj)

		// 次の文字の表示位置調整
		renderPosX += (fontWidth + renderOffsetX)
		lineWidth += (fontWidth + renderOffsetX)
	}

	// 縦位置調整
	renderPosY := startPosY
	for _, group := range LineGroupArray {

		// 行の最大高さを計算
		maxHeight := 0.0
		for _, s := range *group {
			h := s.Scale().Y()
			if maxHeight < h {
				maxHeight = h
			}
		}

		// 縦位置設定
		for _, s := range *group {
			h := s.Scale().Y()
			s.Position().SetY((h / 2) + renderPosY - maxHeight)
		}

		// 次の行位置調整
		renderPosY -= (maxHeight + renderOffsetY)
	}

	// 下端位置を保存
	c.bottom = renderPosY

	return nil
}
