package mathutils

import (
	"app/lib/threejs"
	"syscall/js"
)

// Lerp returns a value linearly interpolated from two known points
// based on the given interval - t = 0 will return x and t = 1 will return y.
//
// x - Start point.
// y - End point.
// t - interpolation factor in the closed interval [0, 1].
func Lerp(x, y, t float64) float64 {
	return getMathUtils().Call("lerp", x, y, t).Float()
}

func getMathUtils() js.Value {
	return threejs.GetJsObject("MathUtils")
}
