package geometries

import "app/lib/threejs"

// BoxBufferGeometry is the BufferGeometry port of BoxGeometry.
type BoxBufferGeometry interface {
	threejs.BufferGeometry
}

type boxBufferGeometryImp struct {
	threejs.BufferGeometry
}

// NewBoxBufferGeometry is constructor.
// width — Width; that is, the length of the edges parallel to the X axis. Optional; defaults to 1.
// height — Height; that is, the length of the edges parallel to the Y axis. Optional; defaults to 1.
// depth — Depth; that is, the length of the edges parallel to the Z axis. Optional; defaults to 1.
// widthSegments — Number of segmented rectangular faces along the width of the sides. Optional; defaults to 1.
// heightSegments — Number of segmented rectangular faces along the height of the sides. Optional; defaults to 1.
// depthSegments — Number of segmented rectangular faces along the depth of the sides. Optional; defaults to 1.
func NewBoxBufferGeometry(
	width float64, height float64, depth float64,
	widthSegments int, heightSegments int, depthSegments int,
) PlaneBufferGeometry {

	return &boxBufferGeometryImp{
		threejs.NewDefaultBufferGeometryFromJSValue(
			threejs.GetJsObject("BoxBufferGeometry").New(
				width, height, depth, widthSegments, heightSegments, depthSegments),
		),
	}
}
