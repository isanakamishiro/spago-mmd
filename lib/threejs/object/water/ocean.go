package water

import (
	"app/lib/threejs"
	"app/lib/threejs/geometry"
	"log"
	"syscall/js"
)

const waterModulePath = "./assets/threejs/ex/jsm/objects/Water.js"

var waterModule js.Value

func init() {

	m := threejs.LoadModule([]string{"Water"}, waterModulePath)
	if len(m) == 0 {
		log.Fatal("water module could not be loaded.")
	}
	waterModule = m[0]
}

// Ocean is water object
type Ocean struct {
	threejs.Mesh
}

// NewOcean creates water object.
func NewOcean(width float64, height float64, options ...OceanOption) *Ocean {

	geom := geometry.NewPlaneGeometry(width, height, 1, 1)
	// geom := geometries.NewBoxBufferGeometry(width, height, 1000, 1, 1, 1)
	// geom := geometries.NewSphereBufferGeometry(width/100, 100, 100)

	var param map[string]interface{} = make(map[string]interface{})
	for _, opt := range options {
		opt(param)
	}

	w := &Ocean{
		Mesh: threejs.NewMeshFromJSValue(waterModule.New(
			geom.JSValue(),
			param,
		)),
	}

	return w
}

// Time gets the time for ocean status.
func (c *Ocean) Time() float64 {
	obj := c.JSValue().Get("material").Get("uniforms")
	return obj.Get("time").Get("value").Float()
}

// SetTime sets the time for ocean status.
func (c *Ocean) SetTime(v float64) {
	obj := c.JSValue().Get("material").Get("uniforms")
	obj.Get("time").Set("value", v)
}

// SetSize sets size of ocean.
func (c *Ocean) SetSize(v float64) {
	obj := c.JSValue().Get("material").Get("uniforms")
	obj.Get("size").Set("value", v)
}
