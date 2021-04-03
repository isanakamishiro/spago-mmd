package animation

import (
	"app/lib/threejs"
	"errors"
	"syscall/js"
)

type Action interface {
	JSValue() js.Value

	// ClampWhenFinished returnes if clampWhenFinished is set to true the animation will automatically be paused on its last frame.
	// If clampWhenFinished is set to false, enabled will automatically be switched to false when the last loop of the action has finished, so that this action has no further impact.
	// Default is false.
	// Note: clampWhenFinished has no impact if the action is interrupted (it has only an effect if its last loop has really finished).
	ClampWhenFinished() bool

	// Enabled gets Setting enabled to false disables this action, so that it has no impact. Default is true.
	Enabled() bool
	// SetEnabled sets Setting enabled to false disables this action, so that it has no impact. Default is true.
	SetEnabled(b bool)

	// Loop gets the looping mode (can be changed with setLoop).
	// Default is THREE.LoopRepeat (with an infinite number of repetitions)
	Loop() Looping
	// SetLoop sets the loop mode and the number of repetitions.
	SetLoop(l Looping, repetitions float64)

	// Paused gets Setting paused to true pauses the execution of the action by setting the effective time scale to 0. Default is false.
	Paused() bool
	// SetPaused sets Setting paused to true pauses the execution of the action by setting the effective time scale to 0. Default is false.
	SetPaused(b bool)

	// Repetitions gets the number of repetitions of the performed AnimationClip over the course of this action. Can be set via setLoop. Default is Infinity.
	// Setting this number has no effect, if the loop mode is set to THREE.LoopOnce.
	Repetitions() float64

	// Time gets the local time of this action (in seconds, starting with 0).
	// The value gets clamped or wrapped to 0...clip.duration (according to the loop state).
	// It can be scaled relativly to the global mixer time by changing timeScale (using setEffectiveTimeScale or setDuration).
	Time() float64
	// SetTime sets the local time of this action (in seconds, starting with 0).
	// The value gets clamped or wrapped to 0...clip.duration (according to the loop state).
	// It can be scaled relativly to the global mixer time by changing timeScale (using setEffectiveTimeScale or setDuration).
	SetTime(t float64)

	// TimeScale gets Scaling factor for the time. A value of 0 causes the animation to pause.
	// Negative values cause the animation to play backwards. Default is 1.
	TimeScale() float64
	// SetTimeScale sets Scaling factor for the time. A value of 0 causes the animation to pause.
	// Negative values cause the animation to play backwards. Default is 1.
	SetTimeScale(s float64)

	// Weight gets The degree of influence of this action (in the interval [0, 1]).
	// Values between 0 (no impact) and 1 (full impact) can be used to blend between several actions. Default is 1.
	Weight() float64
	// SetWeight sets The degree of influence of this action (in the interval [0, 1]).
	// Values between 0 (no impact) and 1 (full impact) can be used to blend between several actions. Default is 1.
	SetWeight(w float64)

	// ZeroSlopeAtEnd gets smooth interpolation without separate clips for start, loop and end. Default is true.
	ZeroSlopeAtEnd() bool
	// SetZeroSlopeAtEnd sets smooth interpolation without separate clips for start, loop and end. Default is true.
	SetZeroSlopeAtEnd(b bool)

	// ZeroSlopeAtStart gets smooth interpolation without separate clips for start, loop and end. Default is true.
	ZeroSlopeAtStart() bool
	// SetZeroSlopeAtStart sets smooth interpolation without separate clips for start, loop and end. Default is true.
	SetZeroSlopeAtStart(b bool)

	// CrossFadeFrom Causes this action to fade in, fading out another action simultaneously, within the passed time interval.
	// If warpBoolean is true, additional warping (gradually changes of the time scales) will be applied.
	CrossFadeFrom(fadeOutAction Action, durationInSeconds float64, warp bool)

	// CrossFadeTo Causes this action to fade out, fading in another action simultaneously, within the passed time interval.
	// If warpBoolean is true, additional warping (gradually changes of the time scales) will be applied.
	CrossFadeTo(fadeOutAction Action, durationInSeconds float64, warp bool)

	// FadeIn increases the weight of this action gradually from 0 to 1, within the passed time interval.
	FadeIn(durationInSeconds float64)

	// FadeOut decreases the weight of this action gradually from 1 to 0, within the passed time interval.
	FadeOut(durationInSeconds float64)

	// EffectiveTimeScale gets the effective time scale (considering the current states of warping and paused).
	EffectiveTimeScale() float64

	// SetEffectiveTimeScale sets Sets the timeScale and stops any scheduled warping.
	// If paused is false, the effective time scale (an internal property) will also be set to this value;
	// otherwise the effective time scale (directly affecting the animation at this moment) will be set to 0.
	SetEffectiveTimeScale(s float64)

	// EffectiveWeight gets the effective weight (considering the current states of fading and enabled).
	EffectiveWeight() float64
	// SetEffectiveWeight sets Sets the weight and stops any scheduled fading.
	// If enabled is true, the effective weight (an internal property) will also be set to this value;
	// otherwise the effective weight (directly affecting the animation at this moment) will be set to 0.
	SetEffectiveWeight(w float64)

	// Clip Returns the clip which holds the animation data for this action.
	Clip() (Clip, error)

	// Mixer Returns the mixer which is responsible for playing this action.
	Mixer() (Mixer, error)

	// Root Returns the root object on which this action is performed.
	Root() (threejs.Object3D, error)

	// Halt Decelerates this animation's speed to 0 by decreasing timeScale gradually (starting from its current value), within the passed time interval.
	Halt(durationInSeconds float64)

	// Running returns true if the action’s time is currently running.
	// In addition to being activated in the mixer (see isScheduled) the following conditions must be fulfilled:
	// paused is equal to false, enabled is equal to true, timeScale is different from 0, and there is no scheduling for a delayed start (startAt).
	Running() bool

	// Scheduled returns true, if this action is activated in the mixer.
	Scheduled() bool

	// Play tells the mixer to activate the action.
	Play()

	// Reset Resets the action.
	// This method sets paused to false, enabled to true, time to 0, interrupts any scheduled fading and warping,
	// and removes the internal loop count and scheduling for delayed starting.
	Reset()

	// SetDuration sets the duration for a single loop of this action (by adjusting timeScale and stopping any scheduled warping).
	SetDuration(d float64)

	// StartAt defines the time for a delayed start (usually passed as AnimationMixer.time + deltaTimeInSeconds).
	StartAt(startTimeInSeconds float64)

	// Stop tells the mixer to deactivate this action.
	// The action will be immediately stopped and completely reset.
	Stop()

	// StopFading stops any scheduled fading which is applied to this action.
	StopFading()

	// StopWarping stops any scheduled warping which is applied to this action.
	StopWarping()

	// SyncWith synchronizes this action with the passed other action.
	// Synchronizing is done by setting this action’s time and timeScale values to the corresponding values of the other action (stopping any scheduled warping).
	SyncWith(action Action)

	// Warp changes the playback speed, within the passed time interval, by modifying timeScale gradually from startTimeScale to endTimeScale.
	Warp(startTimeScale float64, endTimeScale float64, durationInSeconds float64)
}

type actionImp struct {
	js.Value
}

// NewActionFromJSValue creates Action from js.Value.
func NewActionFromJSValue(v js.Value) (Action, error) {
	if v.IsNull() || v.IsUndefined() {
		return nil, errors.New("action is not valid")
	}

	return &actionImp{
		Value: v,
	}, nil
}

// ClampWhenFinished returnes if clampWhenFinished is set to true the animation will automatically be paused on its last frame.
// If clampWhenFinished is set to false, enabled will automatically be switched to false when the last loop of the action has finished, so that this action has no further impact.
// Default is false.
// Note: clampWhenFinished has no impact if the action is interrupted (it has only an effect if its last loop has really finished).
func (c *actionImp) ClampWhenFinished() bool {
	return c.Get("clampWhenFinished").Bool()
}

// Enabled gets Setting enabled to false disables this action, so that it has no impact. Default is true.
func (c *actionImp) Enabled() bool {
	return c.Get("enabled").Bool()
}

// SetEnabled sets Setting enabled to false disables this action, so that it has no impact. Default is true.
func (c *actionImp) SetEnabled(b bool) {
	c.Set("enabled", b)
}

// Loop gets the looping mode (can be changed with setLoop).
// Default is THREE.LoopRepeat (with an infinite number of repetitions)
func (c *actionImp) Loop() Looping {
	l := c.Get("loop")
	return convertLooping(l)
}

// SetLoop sets the loop mode and the number of repetitions.
func (c *actionImp) SetLoop(l Looping, repetitions float64) {
	c.Call("setLoop", l.val(), repetitions)
}

// Paused gets Setting paused to true pauses the execution of the action by setting the effective time scale to 0. Default is false.
func (c *actionImp) Paused() bool {
	return c.Get("paused").Bool()
}

// SetPaused sets Setting paused to true pauses the execution of the action by setting the effective time scale to 0. Default is false.
func (c *actionImp) SetPaused(b bool) {
	c.Set("paused", b)
}

// Repetitions gets the number of repetitions of the performed AnimationClip over the course of this action. Can be set via setLoop. Default is Infinity.
// Setting this number has no effect, if the loop mode is set to THREE.LoopOnce.
func (c *actionImp) Repetitions() float64 {
	return c.Get("repetitions").Float()
}

// Time gets the local time of this action (in seconds, starting with 0).
// The value gets clamped or wrapped to 0...clip.duration (according to the loop state).
// It can be scaled relativly to the global mixer time by changing timeScale (using setEffectiveTimeScale or setDuration).
func (c *actionImp) Time() float64 {
	return c.Get("time").Float()
}

// SetTime sets the local time of this action (in seconds, starting with 0).
// The value gets clamped or wrapped to 0...clip.duration (according to the loop state).
// It can be scaled relativly to the global mixer time by changing timeScale (using setEffectiveTimeScale or setDuration).
func (c *actionImp) SetTime(t float64) {
	c.Set("time", t)
}

// TimeScale gets Scaling factor for the time. A value of 0 causes the animation to pause.
// Negative values cause the animation to play backwards. Default is 1.
func (c *actionImp) TimeScale() float64 {
	return c.Get("timeScale").Float()
}

// SetTimeScale sets Scaling factor for the time. A value of 0 causes the animation to pause.
// Negative values cause the animation to play backwards. Default is 1.
func (c *actionImp) SetTimeScale(s float64) {
	c.Set("timeScale", s)
}

// Weight gets The degree of influence of this action (in the interval [0, 1]).
// Values between 0 (no impact) and 1 (full impact) can be used to blend between several actions. Default is 1.
func (c *actionImp) Weight() float64 {
	return c.Get("weight").Float()
}

// SetWeight sets The degree of influence of this action (in the interval [0, 1]).
// Values between 0 (no impact) and 1 (full impact) can be used to blend between several actions. Default is 1.
func (c *actionImp) SetWeight(w float64) {
	c.Set("weight", w)
}

// ZeroSlopeAtEnd gets smooth interpolation without separate clips for start, loop and end. Default is true.
func (c *actionImp) ZeroSlopeAtEnd() bool {
	return c.Get("zeroSlopeAtEnd").Bool()
}

// SetZeroSlopeAtEnd sets smooth interpolation without separate clips for start, loop and end. Default is true.
func (c *actionImp) SetZeroSlopeAtEnd(b bool) {
	c.Set("zeroSlopeAtEnd", b)
}

// ZeroSlopeAtStart gets smooth interpolation without separate clips for start, loop and end. Default is true.
func (c *actionImp) ZeroSlopeAtStart() bool {
	return c.Get("zeroSlopeAtStart").Bool()
}

// SetZeroSlopeAtStart sets smooth interpolation without separate clips for start, loop and end. Default is true.
func (c *actionImp) SetZeroSlopeAtStart(b bool) {
	c.Set("zeroSlopeAtStart", b)
}

// CrossFadeFrom Causes this action to fade in, fading out another action simultaneously, within the passed time interval.
// If warpBoolean is true, additional warping (gradually changes of the time scales) will be applied.
func (c *actionImp) CrossFadeFrom(fadeOutAction Action, durationInSeconds float64, warp bool) {
	c.Call("crossFadeFrom", fadeOutAction.JSValue(), durationInSeconds, warp)
}

// CrossFadeTo Causes this action to fade out, fading in another action simultaneously, within the passed time interval.
// If warpBoolean is true, additional warping (gradually changes of the time scales) will be applied.
func (c *actionImp) CrossFadeTo(fadeOutAction Action, durationInSeconds float64, warp bool) {
	c.Call("crossFadeTo", fadeOutAction.JSValue(), durationInSeconds, warp)
}

// FadeIn increases the weight of this action gradually from 0 to 1, within the passed time interval.
func (c *actionImp) FadeIn(durationInSeconds float64) {
	c.Call("fadeIn", durationInSeconds)
}

// FadeOut decreases the weight of this action gradually from 1 to 0, within the passed time interval.
func (c *actionImp) FadeOut(durationInSeconds float64) {
	c.Call("fadeOut", durationInSeconds)
}

// EffectiveTimeScale gets the effective time scale (considering the current states of warping and paused).
func (c *actionImp) EffectiveTimeScale() float64 {
	return c.Call("getEffectiveTimeScale").Float()
}

// SetEffectiveTimeScale sets Sets the timeScale and stops any scheduled warping.
// If paused is false, the effective time scale (an internal property) will also be set to this value;
// otherwise the effective time scale (directly affecting the animation at this moment) will be set to 0.
func (c *actionImp) SetEffectiveTimeScale(s float64) {
	c.Call("setEffectiveTimeScale", s)
}

// EffectiveWeight gets the effective weight (considering the current states of fading and enabled).
func (c *actionImp) EffectiveWeight() float64 {
	return c.Call("getEffectiveWeight").Float()
}

// SetEffectiveWeight sets Sets the weight and stops any scheduled fading.
// If enabled is true, the effective weight (an internal property) will also be set to this value;
// otherwise the effective weight (directly affecting the animation at this moment) will be set to 0.
func (c *actionImp) SetEffectiveWeight(w float64) {
	c.Call("setEffectiveWeight", w)
}

// Clip Returns the clip which holds the animation data for this action.
func (c *actionImp) Clip() (Clip, error) {
	v := c.Call("getClip")
	return NewClipFromJSValue(v)
}

// Mixer Returns the mixer which is responsible for playing this action.
func (c *actionImp) Mixer() (Mixer, error) {
	v := c.Call("getMixer")
	return NewMixerFromJSValue(v)
}

// Root Returns the root object on which this action is performed.
func (c *actionImp) Root() (threejs.Object3D, error) {
	v := c.Call("getRoot")
	if v.IsNull() || v.IsUndefined() {
		return nil, errors.New("root is nothing")
	}

	return threejs.NewObject3DFromJSValue(v), nil
}

// Halt Decelerates this animation's speed to 0 by decreasing timeScale gradually (starting from its current value), within the passed time interval.
func (c *actionImp) Halt(durationInSeconds float64) {
	c.Call("halt")
}

// Running returns true if the action’s time is currently running.
// In addition to being activated in the mixer (see isScheduled) the following conditions must be fulfilled:
// paused is equal to false, enabled is equal to true, timeScale is different from 0, and there is no scheduling for a delayed start (startAt).
func (c *actionImp) Running() bool {
	return c.Call("isRunning").Bool()
}

// Scheduled returns true, if this action is activated in the mixer.
func (c *actionImp) Scheduled() bool {
	return c.Call("isScheduled").Bool()
}

// Play tells the mixer to activate the action.
func (c *actionImp) Play() {
	c.Call("play")
}

// Reset Resets the action.
// This method sets paused to false, enabled to true, time to 0, interrupts any scheduled fading and warping,
// and removes the internal loop count and scheduling for delayed starting.
func (c *actionImp) Reset() {
	c.Call("reset")
}

// SetDuration sets the duration for a single loop of this action (by adjusting timeScale and stopping any scheduled warping).
func (c *actionImp) SetDuration(d float64) {
	c.Call("setDuration", d)
}

// StartAt defines the time for a delayed start (usually passed as AnimationMixer.time + deltaTimeInSeconds).
func (c *actionImp) StartAt(startTimeInSeconds float64) {
	c.Call("startAt", startTimeInSeconds)
}

// Stop tells the mixer to deactivate this action.
// The action will be immediately stopped and completely reset.
func (c *actionImp) Stop() {
	c.Call("stop")
}

// StopFading stops any scheduled fading which is applied to this action.
func (c *actionImp) StopFading() {
	c.Call("stopFading")
}

// StopWarping stops any scheduled warping which is applied to this action.
func (c *actionImp) StopWarping() {
	c.Call("stopWarping")
}

// SyncWith synchronizes this action with the passed other action.
// Synchronizing is done by setting this action’s time and timeScale values to the corresponding values of the other action (stopping any scheduled warping).
func (c *actionImp) SyncWith(action Action) {
	c.Call("syncWith", action.JSValue())
}

// Warp changes the playback speed, within the passed time interval, by modifying timeScale gradually from startTimeScale to endTimeScale.
func (c *actionImp) Warp(startTimeScale float64, endTimeScale float64, durationInSeconds float64) {
	c.Call("warp", startTimeScale, endTimeScale, durationInSeconds)
}
