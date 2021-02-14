package mmdloaders

import (
	"app/lib/threejs"
	"log"
	"syscall/js"
)

const (
	mmdAnimationHelperModulePath = "./assets/threejs/ex/jsm/animation/MMDAnimationHelper.js"
)

var (
	mmdAnimationHelperModule js.Value
)

func init() {

	m := threejs.LoadModule([]string{"MMDAnimationHelper"}, mmdAnimationHelperModulePath)
	if len(m) == 0 {
		log.Fatal("MMDAnimationHelper module could not be loaded.")
	}
	mmdAnimationHelperModule = m[0]

}

// MMDAnimationHelperParameters is ...
type MMDAnimationHelperParameters interface {
}

// MMDAnimationHelperAddParameter is ...
type MMDAnimationHelperAddParameter interface {
}

// MMDAnimationHelper handles animation of MMD assets loaded by MMDLoader with MMD special features as IK, Grant, and Physics. It uses CCDIKSolver and MMDPhysics inside.
type MMDAnimationHelper struct {
	js.Value
}

// NewMMDAnimationHelper creates a new MMDAnimationHelper.
func NewMMDAnimationHelper(params MMDAnimationHelperParameters) *MMDAnimationHelper {
	return &MMDAnimationHelper{
		Value: mmdAnimationHelperModule.New(params),
	}
}

// Add an SkinnedMesh, Camera, or Audio to helper and setup animation. The anmation durations of added objects are synched. If camera/audio has already been added, it'll be replaced with a new one.
//
// params are like below.
//
// 	animation - an AnimationClip or an array of AnimationClip set to object. Only for SkinnedMesh and Camera. Default is undefined.
// 	physics - Only for SkinnedMesh. A flag whether turn on physics. Default is true.
// 	warmup - Only for SkinnedMesh and physics is true. Physics parameter. Default is 60.
// 	unitStep - Only for SkinnedMesh and physics is true. Physics parameter. Default is 1 / 65.
// 	maxStepNum - Only for SkinnedMesh and physics is true. Physics parameter. Default is 3.
// 	gravity - Only for SkinnedMesh and physics is true. Physics parameter. Default is ( 0, - 9.8 * 10, 0 ).
// 	delayTime - Only for Audio. Default is 0.0.
func (c *MMDAnimationHelper) Add(obj threejs.Object3D, params MMDAnimationHelperAddParameter) {
	c.Call("add", obj.JSValue(), params)
}

// Update advance mixer time and update the animations of objects added to helper
//
// delta — number in second
func (c *MMDAnimationHelper) Update(delta float64) {
	c.Call("update", delta)
}
