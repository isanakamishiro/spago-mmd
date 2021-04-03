package helper

import "app/lib/threejs"

// GridHelper is an The GridHelper is an object to define grids.
// Grids are two-dimensional arrays of lines.
type GridHelper struct {
	threejs.Object3D
}

// NewGridHelper is factory method for GridHelper.
// size -- The size of the grid. Default is 10.
// divisions -- The number of divisions across the grid. Default is 10.
// colorCenterLine -- The color of the centerline. This can be a Color, a hexadecimal value and an CSS-Color name. Default is 0x444444
// colorGrid -- The color of the lines of the grid. This can be a Color, a hexadecimal value and an CSS-Color name. Default is 0x888888
//
// Creates a new GridHelper of size 'size' and divided into 'divisions' segments per side.
// Colors are optional.
func NewGridHelper(size float64, divisions float64) *GridHelper {
	return &GridHelper{
		threejs.NewObject3DFromJSValue(
			threejs.Threejs("GridHelper").New(size, divisions),
		),
	}
}

// Material is ...
func (c *GridHelper) Material() threejs.Material {
	return threejs.NewDefaultMaterialFromJSValue(
		c.JSValue().Get("material"),
	)
}
