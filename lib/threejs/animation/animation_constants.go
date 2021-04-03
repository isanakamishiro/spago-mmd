package animation

import (
	"app/lib/threejs"
	"syscall/js"
)

type Looping uint8

const (
	LoopOnce Looping = iota + 1
	LoopRepeat
	LoopPingPong
)

var (
	_looping = [4]js.Value{
		js.Null(),
		threejs.Threejs("LoopOnce"),
		threejs.Threejs("LoopRepeat"),
		threejs.Threejs("LoopPingPong"),
	}
)

// val returns js.Value from constants.
func (v Looping) val() js.Value {
	return _looping[v]
}

// convertLooping gets Looping from js.Value.
func convertLooping(loop js.Value) Looping {
	for i, l := range _looping {
		if l.Equal(loop) {
			return Looping(i)
		}
	}
	return 0
}
