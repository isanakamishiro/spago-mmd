package mmd

import (
	"app/lib/threejs/animation"
	"syscall/js"
)

// AnimationHelperAddOption is option for AnimationHelper's Add method.
type AnimationHelperAddOption func(map[string]interface{}) error

// AnimationClip sets an AnimationClip set to object.
func AnimationClip(clip animation.Clip) AnimationHelperAddOption {
	return func(m map[string]interface{}) error {

		m["animation"] = clip.JSValue()

		return nil
	}
}

// AnimationClips sets an array of AnimationClip set to object.
func AnimationClips(clips []animation.Clip) AnimationHelperAddOption {
	return func(m map[string]interface{}) error {

		if len(clips) == 0 {
			return nil
		}

		// js.Valueに詰め替え
		var a []interface{} = make([]interface{}, len(clips))
		for i, v := range clips {
			a[i] = v
		}
		m["animation"] = js.ValueOf(a)

		return nil
	}
}

// Physics sets A flag whether turn on physics.
func Physics(flag bool) AnimationHelperAddOption {
	return func(m map[string]interface{}) error {

		m["physics"] = flag

		return nil
	}
}

// Warmup sets Physics parameter. Default is 60.
func Warmup(n int) AnimationHelperAddOption {
	return func(m map[string]interface{}) error {

		m["warmup"] = n

		return nil
	}
}

// UnitStep sets Physics parameter. Default is 1 / 65.
func UnitStep(u float64) AnimationHelperAddOption {
	return func(m map[string]interface{}) error {

		m["unitStep"] = u

		return nil
	}
}

// MaxStepNumber sets Physics parameter. Default is 3.
func MaxStepNumber(n int) AnimationHelperAddOption {
	return func(m map[string]interface{}) error {

		m["maxStepNum"] = n

		return nil
	}
}

// Gravity sets Physics parameter. Default is ( 0, - 9.8 * 10, 0 ).
func Gravity(n int) AnimationHelperAddOption {
	return func(m map[string]interface{}) error {

		m["gravity"] = n

		return nil
	}
}

// DelayTime sets Audio. Default is 0.0.
func DelayTime(t float64) AnimationHelperAddOption {
	return func(m map[string]interface{}) error {

		m["delayTime"] = t

		return nil
	}
}
