package sky

import (
	"app/lib/threejs"
	"log"
	"math"
	"syscall/js"
)

const skyModulePath = "./assets/threejs/ex/jsm/objects/Sky.js"

var skyModule js.Value

func init() {

	m := threejs.LoadModule([]string{"Sky"}, skyModulePath)
	if len(m) == 0 {
		log.Fatal("sky module could not be loaded.")
	}
	skyModule = m[0]
}

// Sky is sky object
type Sky struct {
	threejs.Mesh

	sun         *threejs.Vector3
	inclination float64
	azimuth     float64
}

// NewSky creates Sky object.
func NewSky() *Sky {
	s := &Sky{
		Mesh:        threejs.NewMeshFromJSValue(skyModule.New()),
		sun:         threejs.NewVector3(0, 0, 0),
		inclination: 0.49,
		azimuth:     0.4,
	}
	s.resetSunPosition()

	return s
}

// SetTurbidity sets turbidity in the sky.
func (c *Sky) SetTurbidity(v float64) {

	obj := c.JSValue().Get("material").Get("uniforms")
	obj.Get("turbidity").Set("value", v)

}

// SetRayleigh sets rayleigh in the sky.
func (c *Sky) SetRayleigh(v float64) {

	obj := c.JSValue().Get("material").Get("uniforms")
	obj.Get("rayleigh").Set("value", v)

}

// SetMieCoefficient sets mieCoefficient in the sky.
func (c *Sky) SetMieCoefficient(v float64) {

	obj := c.JSValue().Get("material").Get("uniforms")
	obj.Get("mieCoefficient").Set("value", v)

}

// SetMieDirectionalG sets mieCoefficient in the sky.
func (c *Sky) SetMieDirectionalG(v float64) {

	obj := c.JSValue().Get("material").Get("uniforms")
	obj.Get("mieDirectionalG").Set("value", v)

}

// SetInclination sets inclination in the sky.
func (c *Sky) SetInclination(v float64) {

	c.inclination = v
	c.resetSunPosition()
}

// SetAzimuth sets azimuth in the sky.
func (c *Sky) SetAzimuth(v float64) {

	c.azimuth = v
	c.resetSunPosition()
}

func (c *Sky) resetSunPosition() {

	theta := math.Pi * (c.inclination - 0.5)
	phi := 2 * math.Pi * (c.azimuth - 0.5)

	c.sun.SetX(math.Cos(phi))
	c.sun.SetY(math.Sin(phi) * math.Sin(theta))
	c.sun.SetZ(math.Sin(phi) * math.Cos(theta))

	obj := c.JSValue().Get("material").Get("uniforms")
	obj.Get("sunPosition").Set("value", c.sun)

}
