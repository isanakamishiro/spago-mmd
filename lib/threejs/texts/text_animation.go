package texts

import "app/lib/threejs/mathutils"

// Animation is ...
type Animation func(mesh CharacterMesh, frame float64)

// FadeIn is ...
func FadeIn() Animation {
	return func(m CharacterMesh, frame float64) {
		m.Material().SetOpacity(frame)
	}
}

// ScaleWidthIn is ...
func ScaleWidthIn(w float64) Animation {
	return func(m CharacterMesh, frame float64) {
		scale := mathutils.Lerp(w, 1.0, frame)
		m.Scale().SetX(scale)
	}
}

// ScaleHeightIn is ...
func ScaleHeightIn(h float64) Animation {
	return func(m CharacterMesh, frame float64) {
		scale := mathutils.Lerp(h, 1.0, frame)
		m.Scale().SetY(scale)
	}
}
