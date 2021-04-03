package exlight

import (
	"app/lib/threejs"
	"log"
	"syscall/js"
)

const rectAreaLightHelperModulePath = "./assets/threejs/ex/jsm/helpers/RectAreaLightHelper.js"

var rectAreaLightHelperModule js.Value

func init() {

	m := threejs.LoadModule([]string{"RectAreaLightHelper"}, rectAreaLightHelperModulePath)
	if len(m) == 0 {
		log.Fatal("RectAreaLightHelper module could not be loaded.")
	}
	rectAreaLightHelperModule = m[0]

}

// RectAreaLightHelper Creates a visual aid for a RectAreaLight.
type RectAreaLightHelper struct {
	threejs.Object3D
}

// NewRectAreaLightHelper creates RectAreaLightHelper.
// light-- The light to be visualized.
func NewRectAreaLightHelper(light RectAreaLight) *RectAreaLightHelper {
	return &RectAreaLightHelper{
		Object3D: threejs.NewObject3DFromJSValue(
			rectAreaLightHelperModule.New(light.JSValue()),
		),
	}
}

// Dispose of the light helper.
func (c *RectAreaLightHelper) Dispose() {
	c.JSValue().Call("dispose")
}

// Update Updates the helper to match the position and direction of the .light.
func (c *RectAreaLightHelper) Update() {
	c.JSValue().Call("update")
}
