package geometry

import "app/lib/threejs"

// BoxGeometry is the BufferGeometry port of BoxGeometry.
type BoxGeometry interface {
	threejs.BufferGeometry
}

type boxGeometryImp struct {
	threejs.BufferGeometry
}

// NewBoxGeometry is constructor.
// width — Width; that is, the length of the edges parallel to the X axis. Optional; defaults to 1.
// height — Height; that is, the length of the edges parallel to the Y axis. Optional; defaults to 1.
// depth — Depth; that is, the length of the edges parallel to the Z axis. Optional; defaults to 1.
// widthSegments — Number of segmented rectangular faces along the width of the sides. Optional; defaults to 1.
// heightSegments — Number of segmented rectangular faces along the height of the sides. Optional; defaults to 1.
// depthSegments — Number of segmented rectangular faces along the depth of the sides. Optional; defaults to 1.
func NewBoxGeometry(
	width float64, height float64, depth float64,
	widthSegments int, heightSegments int, depthSegments int,
) PlaneBufferGeometry {

	return &boxGeometryImp{
		threejs.NewBufferGeometryFromJSValue(
			threejs.Threejs("BoxGeometry").New(
				width, height, depth, widthSegments, heightSegments, depthSegments),
		),
	}
}
