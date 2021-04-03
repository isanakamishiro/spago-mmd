package animation

import (
	"app/lib/threejs"
	"errors"
	"syscall/js"
)

// Mixer is a player for animations on a particular object in the scene.
// When multiple objects in the scene are animated independently, one AnimationMixer may be used for each object.
// For an overview of the different elements of the three.js animation system see the "Animation System" article in the "Next Steps" section of the manual.
type Mixer interface {
	JSValue() js.Value

	// Time gets The global mixer time (in seconds; starting with 0 on the mixer's creation).
	Time() float64
	// SetTime sets The global mixer time (in seconds; starting with 0 on the mixer's creation).
	SetTime(t float64)

	// TimeScale gets A scaling factor for the global mixer time.
	// Note: Setting the mixer's timeScale to 0 and later back to 1 is a possibility to pause/unpause all actions that are controlled by this mixer.
	TimeScale() float64

	// SetTimeScale sets A scaling factor for the global mixer time.
	// Note: Setting the mixer's timeScale to 0 and later back to 1 is a possibility to pause/unpause all actions that are controlled by this mixer.
	SetTimeScale(s float64)

	// ClipAction returns an AnimationAction for the passed clip, optionally using a root object different from the mixer's default root.
	// The first parameter can be either an AnimationClip object or the name of an AnimationClip.
	// If an action fitting the clip and root parameters doesn't yet exist, it will be created by this method.
	// Calling this method several times with the same clip and root parameters always returns the same clip instance.
	ClipAction(clip Clip) (Action, error)

	// ExistingAction returns an existing AnimationAction for the passed clip,
	// optionally using a root object different from the mixer's default root.
	// The first parameter can be either an AnimationClip object or the name of an AnimationClip.
	ExistingAction(clip Clip) (Action, error)

	// Root returns this mixer's root object.
	Root() threejs.Object3D

	// StopAllAction deactivates all previously scheduled actions on this mixer.
	StopAllAction()

	// Update advances the global mixer time and updates the animation.
	// This is usually done in the render loop, passing clock.getDelta scaled by the mixer's timeScale).
	Update(delta float64)

	// UncacheClip deallocates all memory resources for a clip.
	UncacheClip(clip Clip)

	// UncacheAction deallocates all memory resources for an action.
	UncacheAction(clip Clip)
}

type mixerImp struct {
	js.Value
}

// NewMixerFromJSValue creates Mixer from js.Value.
func NewMixerFromJSValue(v js.Value) (Mixer, error) {
	if v.IsNull() || v.IsUndefined() {
		return nil, errors.New("parameter for mixer is not valid")
	}

	return &mixerImp{
		Value: v,
	}, nil
}

// Time gets The global mixer time (in seconds; starting with 0 on the mixer's creation).
func (c *mixerImp) Time() float64 {
	return c.Get("time").Float()
}

// SetTime sets The global mixer time (in seconds; starting with 0 on the mixer's creation).
func (c *mixerImp) SetTime(t float64) {
	c.Call("setTime", t)
}

// TimeScale gets A scaling factor for the global mixer time.
// Note: Setting the mixer's timeScale to 0 and later back to 1 is a possibility to pause/unpause all actions that are controlled by this mixer.
func (c *mixerImp) TimeScale() float64 {
	return c.Get("timeScale").Float()
}

// SetTimeScale sets A scaling factor for the global mixer time.
// Note: Setting the mixer's timeScale to 0 and later back to 1 is a possibility to pause/unpause all actions that are controlled by this mixer.
func (c *mixerImp) SetTimeScale(s float64) {
	c.Set("timeScale", s)
}

// ClipAction returns an AnimationAction for the passed clip, optionally using a root object different from the mixer's default root.
// The first parameter can be either an AnimationClip object or the name of an AnimationClip.
// If an action fitting the clip and root parameters doesn't yet exist, it will be created by this method.
// Calling this method several times with the same clip and root parameters always returns the same clip instance.
func (c *mixerImp) ClipAction(clip Clip) (Action, error) {
	action := c.Call("clipAction", clip.JSValue())
	return NewActionFromJSValue(action)
}

// ExistingAction returns an existing AnimationAction for the passed clip,
// optionally using a root object different from the mixer's default root.
// The first parameter can be either an AnimationClip object or the name of an AnimationClip.
func (c *mixerImp) ExistingAction(clip Clip) (Action, error) {
	action := c.Call("existingAction", clip.JSValue())
	return NewActionFromJSValue(action)
}

// Root returns this mixer's root object.
func (c *mixerImp) Root() threejs.Object3D {
	root := c.Call("getRoot")
	return threejs.NewObject3DFromJSValue(root)
}

// StopAllAction deactivates all previously scheduled actions on this mixer.
func (c *mixerImp) StopAllAction() {
	c.Call("stopAllAction")
}

// Update advances the global mixer time and updates the animation.
// This is usually done in the render loop, passing clock.getDelta scaled by the mixer's timeScale).
func (c *mixerImp) Update(delta float64) {
	c.Call("update", delta)
}

// UncacheClip deallocates all memory resources for a clip.
func (c *mixerImp) UncacheClip(clip Clip) {
	c.Call("uncacheClip", clip.JSValue())
}

// UncacheAction deallocates all memory resources for an action.
func (c *mixerImp) UncacheAction(clip Clip) {
	c.Call("uncacheAction", clip.JSValue())
}
