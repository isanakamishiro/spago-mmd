package geometries

import "app/lib/threejs"

// PlaneBufferGeometry extend: [BufferGeometry]
type PlaneBufferGeometry interface {
	threejs.BufferGeometry
}

type planeBufferGeometryImp struct {
	threejs.BufferGeometry
}

// NewPlaneGeometry is constructor.
// width — Width along the X axis. Default is 1.
// height — Height along the Y axis. Default is 1.
// widthSegments — Optional. Default is 1.
// heightSegments — Optional. Default is 1.
func NewPlaneGeometry(width float64, height float64, widthSegments int, heightSegments int) PlaneBufferGeometry {
	return planeBufferGeometryImp{
		threejs.NewDefaultBufferGeometryFromJSValue(
			threejs.GetJsObject("PlaneGeometry").New(width, height, widthSegments, heightSegments),
		),
	}
}
