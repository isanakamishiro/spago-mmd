package effect

import (
	"app/lib/threejs"
	"log"
	"syscall/js"
)

const outlineEffectModulePath = "./assets/threejs/ex/jsm/effects/OutlineEffect.js"

var outlineEffectModule js.Value

func init() {

	m := threejs.LoadModule([]string{"OutlineEffect"}, outlineEffectModulePath)
	if len(m) == 0 {
		log.Fatal("outlineEffect module could not be loaded.")
	}
	outlineEffectModule = m[0]
}

// OutlineEffect is ...
type OutlineEffect struct {
	js.Value
}

// NewOutlineEffect is ...
func NewOutlineEffect(renderer threejs.Renderer) *OutlineEffect {
	return &OutlineEffect{Value: outlineEffectModule.New(renderer.JSValue())}
}

// Render is ...
func (oe *OutlineEffect) Render(scene threejs.Scene, camera threejs.Camera) {
	oe.Call("render", scene.JSValue(), camera.JSValue())
}

// SetSize is ...
func (oe *OutlineEffect) SetSize(width float64, height float64, updateStyle bool) {
	oe.Call("setSize", width, height, updateStyle)
}
