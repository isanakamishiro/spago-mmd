package mmd

import (
	"app/lib/threejs"
	"app/lib/threejs/animation"
	"errors"
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

// AnimationHelperParameters is ...
type AnimationHelperParameters interface {
}

// AnimationHelper handles animation of MMD assets loaded by MMDLoader with MMD special features as IK, Grant, and Physics. It uses CCDIKSolver and MMDPhysics inside.
type AnimationHelper struct {
	js.Value
}

// NewAnimationHelper creates a new MMDAnimationHelper.
func NewAnimationHelper(params AnimationHelperParameters) *AnimationHelper {
	return &AnimationHelper{
		Value: mmdAnimationHelperModule.New(params),
	}
}

// Update advance mixer time and update the animations of objects added to helper
//
// delta — number in second
func (c *AnimationHelper) Update(delta float64) {
	c.Call("update", delta)
}

// Meshes gets meshes in AnimationHelper.
func (c *AnimationHelper) Meshes() []threejs.SkinnedMesh {

	m := c.Get("meshes")
	l := m.Length()
	var meshes []threejs.SkinnedMesh = make([]threejs.SkinnedMesh, l)
	for i := 0; i < l; i++ {
		meshes[i] = threejs.NewSkinnedMeshFromJSValue(m.Index(i))
	}

	return meshes

}

// AddMesh add an SkinnedMesh to helper and setup animation. The anmation durations of added objects are synched. If camera/audio has already been added, it'll be replaced with a new one.
func (c *AnimationHelper) AddMesh(mesh threejs.Mesh, options ...AnimationHelperAddOption) {

	var param map[string]interface{} = make(map[string]interface{})
	for _, opt := range options {
		opt(param)
	}

	c.Call("add", mesh.JSValue(), param)
}

// RemoveMesh remove mesh.
func (c *AnimationHelper) RemoveMesh(mesh threejs.Mesh) {
	c.Call("remove", mesh.JSValue())
}

// Mixer gets mixer object in animation.
func (c *AnimationHelper) Mixer(mesh threejs.Mesh) (animation.Mixer, error) {
	m := c.Get("objects").Call("get", mesh.JSValue())
	if m.IsNull() || m.IsUndefined() {
		return nil, errors.New("mesh is not registered or nil")
	}

	mixer := m.Get("mixer")
	if mixer.IsNull() || mixer.IsUndefined() {
		return nil, errors.New("mixer is not defined or nil")
	}

	return animation.NewMixerFromJSValue(mixer)

}

// Pose changes the posing of SkinnedMesh as VPD content specifies.
func (c *AnimationHelper) Pose(mesh threejs.Mesh, vpd Vpd, options ...AnimationHelperPoseOption) {
	var param map[string]interface{} = make(map[string]interface{})
	for _, opt := range options {
		opt(param)
	}

	c.Call("pose", mesh.JSValue(), vpd.JSValue(), param)

}
