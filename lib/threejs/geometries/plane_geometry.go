package geometries

import "app/lib/threejs"

// PlaneBufferGeometry extend: [BufferGeometry]
type PlaneBufferGeometry interface {
	threejs.BufferGeometry
}

type planeBufferGeometryImp struct {
	threejs.BufferGeometry
}

// NewPlaneBufferGeometry is constructor.
// width — Width along the X axis. Default is 1.
// height — Height along the Y axis. Default is 1.
// widthSegments — Optional. Default is 1.
// heightSegments — Optional. Default is 1.
func NewPlaneBufferGeometry(width float64, height float64, widthSegments int, heightSegments int) PlaneBufferGeometry {
	return planeBufferGeometryImp{
		threejs.NewDefaultBufferGeometryFromJSValue(
			threejs.GetJsObject("PlaneBufferGeometry").New(width, height, widthSegments, heightSegments),
		),
	}
}

// PlaneGeometry extend: [BufferGeometry]
type PlaneGeometry interface {
	threejs.Geometry
}

type planeGeometryImp struct {
	threejs.Geometry
}

// NewPlaneGeometry is constructor.
// width — Width along the X axis. Default is 1.
// height — Height along the Y axis. Default is 1.
// widthSegments — Optional. Default is 1.
// heightSegments — Optional. Default is 1.
func NewPlaneGeometry(width float64, height float64, widthSegments int, heightSegments int) PlaneGeometry {
	return &planeGeometryImp{
		threejs.NewDefaultGeometryFromJSValue(
			threejs.GetJsObject("PlaneGeometry").New(width, height, widthSegments, heightSegments),
		),
	}
}
