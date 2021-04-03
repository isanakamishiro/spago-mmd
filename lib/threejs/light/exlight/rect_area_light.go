package exlight

import (
	"app/lib/threejs"
	"log"
	"syscall/js"
)

const rectAreaLightModulePath = "./assets/threejs/ex/jsm/lights/RectAreaLightUniformsLib.js"

var rectAreaLightModule js.Value

func init() {

	m := threejs.LoadModule([]string{"RectAreaLightUniformsLib"}, rectAreaLightModulePath)
	if len(m) == 0 {
		log.Fatal("RectAreaLightUniformsLib module could not be loaded.")
	}
	rectAreaLightModule = m[0]

	rectAreaLightModule.Call("init")
}

// RectAreaLight emits light uniformly across the face a rectangular plane.
// This light type can be used to simulate light sources such as bright windows or strip lighting.
//
// Important Notes:
//
// There is no shadow support.
// Only MeshStandardMaterial and MeshPhysicalMaterial are supported.
// You have to include RectAreaLightUniformsLib into your scene and call init().
type RectAreaLight interface {
	threejs.Light
}

type rectAreaLightImp struct {
	threejs.Light
}

// NewRectAreaLight Creates a new RectAreaLight.
// color - (optional) hexadecimal color of the light. Default is 0xffffff (white).
// intensity - (optional) the light's intensity, or brightness. Default is 1.
// width - (optional) width of the light. Default is 10.
// height - (optional) height of the light. Default is 10.
func NewRectAreaLight(
	color threejs.ColorValue,
	intensity threejs.LightIntensity,
	width float64,
	height float64,
) RectAreaLight {

	return &rectAreaLightImp{
		threejs.NewDefaultLightFromJSValue(
			threejs.Threejs("RectAreaLight").New(
				float64(color),
				float64(intensity),
				width,
				height,
			),
		),
	}
}
