package text

import "app/lib/threejs"

// MeshOption is functional parameter option for NewTextMesh.
type MeshOption func(*Mesh) error

// WrapWidth sets wrap width and enable auto-wrap.
func WrapWidth(w float64) MeshOption {
	return func(t *Mesh) error {
		t.renderAutoWrap = true
		t.renderWrapWidth = w

		return nil
	}
}

// PositionOffsetX sets character position offset x.
func PositionOffsetX(x float64) MeshOption {
	return func(t *Mesh) error {
		t.characterPositionOffsetX = x

		return nil
	}
}

// PositionOffsetY sets character position offset y.
func PositionOffsetY(y float64) MeshOption {
	return func(t *Mesh) error {
		t.characterPositionOffsetX = y

		return nil
	}
}

// FontScale sets font scale ratio.
func FontScale(scale float64) MeshOption {
	return func(t *Mesh) error {
		t.renderFontScale = scale

		return nil
	}
}

// MaterialColor sets material color.
func MaterialColor(col threejs.Color) MeshOption {
	return func(t *Mesh) error {
		t.renderColor = col

		return nil
	}
}
