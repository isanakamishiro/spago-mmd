package animation

import (
	"errors"
	"syscall/js"
)

type Clip interface {
	JSValue() js.Value

	// Duration gets the duration of this clip (in seconds).
	// This can be calculated from the tracks array via resetDuration.
	Duration() float64

	// Name gets a name for this clip.
	// A certain clip can be searched via findByName.
	Name() string

	// UUID gets the UUID of this clip instance.
	// It gets automatically assigned and shouldn't be edited.
	UUID() string

	// Optimize each track by removing equivalent sequential keys (which are common in morph target sequences).
	Optimize()

	// ResetDuration sets the duration of the clip to the duration of its longest KeyframeTrack.
	ResetDuration()

	// Trim trims all tracks to the clip's duration.
	Trim()

	// Validate validates performs minimal validation on each track in the clip. Returns true if all tracks are valid.
	Validate() bool
}

type clipImp struct {
	js.Value
}

func NewClipFromJSValue(v js.Value) (Clip, error) {
	if v.IsNull() || v.IsUndefined() {
		return nil, errors.New("clip is not valid")
	}

	return &clipImp{
		Value: v,
	}, nil
}

// Duration gets the duration of this clip (in seconds).
// This can be calculated from the tracks array via resetDuration.
func (c *clipImp) Duration() float64 {
	return c.Get("duration").Float()
}

// Name gets a name for this clip.
// A certain clip can be searched via findByName.
func (c *clipImp) Name() string {
	return c.Get("name").String()
}

// UUID gets the UUID of this clip instance.
// It gets automatically assigned and shouldn't be edited.
func (c *clipImp) UUID() string {
	return c.Get("uuid").String()
}

// Optimize each track by removing equivalent sequential keys (which are common in morph target sequences).
func (c *clipImp) Optimize() {
	c.Call("optimize")
}

// ResetDuration sets the duration of the clip to the duration of its longest KeyframeTrack.
func (c *clipImp) ResetDuration() {
	c.Call("resetDuration")
}

// Trim trims all tracks to the clip's duration.
func (c *clipImp) Trim() {
	c.Call("trim")
}

// Validate validates performs minimal validation on each track in the clip. Returns true if all tracks are valid.
func (c *clipImp) Validate() bool {
	return c.Call("validate").Bool()
}
