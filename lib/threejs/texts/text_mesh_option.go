package texts

import "app/lib/threejs"

// TextMeshOption is functional parameter option for NewTextMesh.
type TextMeshOption func(*TextMesh) error

// WrapWidth sets wrap width and enable auto-wrap.
func WrapWidth(w float64) TextMeshOption {
	return func(t *TextMesh) error {
		t.renderAutoWrap = true
		t.renderWrapWidth = w

		return nil
	}
}

// PositionOffsetX sets character position offset x.
func PositionOffsetX(x float64) TextMeshOption {
	return func(t *TextMesh) error {
		t.characterPositionOffsetX = x

		return nil
	}
}

// PositionOffsetY sets character position offset y.
func PositionOffsetY(y float64) TextMeshOption {
	return func(t *TextMesh) error {
		t.characterPositionOffsetX = y

		return nil
	}
}

// FontScale sets font scale ratio.
func FontScale(scale float64) TextMeshOption {
	return func(t *TextMesh) error {
		t.renderFontScale = scale

		return nil
	}
}

// MaterialColor sets material color.
func MaterialColor(col threejs.Color) TextMeshOption {
	return func(t *TextMesh) error {
		t.renderColor = col

		return nil
	}
}
