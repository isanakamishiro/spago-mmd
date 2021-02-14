package threejs

import (
	"syscall/js"
)

// Clock is Object for keeping track of time.
// This uses performance.now if it is available, otherwise it reverts to the less accurate Date.now.
type Clock interface {
	JSValue() js.Value

	AutoStart() bool
	Start()
	Stop()
	Delta() float64
	ElapsedTime() float64
	OldTime() float64
	StartTime() float64
}

// clockImp is default implementation of Clock interface.
type clockImp struct {
	js.Value
}

// NewClock is factory method for Clock.
func NewClock(autoStart bool) Clock {
	return &clockImp{Value: GetJsObject("Clock").New(autoStart)}
}

// JSValue is ...
func (cc *clockImp) JSValue() js.Value {
	return cc.Value
}

// AutoStart is ...
func (cc *clockImp) AutoStart() bool {
	return cc.Get("autoStart").Bool()
}

// Start is ...
func (cc *clockImp) Start() {
	cc.Call("start")
}

// Stop is ...
func (cc *clockImp) Stop() {
	cc.Call("stop")
}

// Delta is ...
func (cc *clockImp) Delta() float64 {
	return cc.Call("getDelta").Float()
}

// ElapsedTime is ...
func (cc *clockImp) ElapsedTime() float64 {
	return cc.Call("getElapsedTime").Float()
}

// OldTime is ...
func (cc *clockImp) OldTime() float64 {
	return cc.Get("oldTime").Float()
}

// StartTime is ...
func (cc *clockImp) StartTime() float64 {
	return cc.Get("startTime").Float()
}

// func (cc *Clock) SetAutoStart(v bool) {
// 	cc.Set("autoStart", v)
// }
// func (cc *Clock) SetElapsedTime(v float64) {
// 	cc.Set("elapsedTime", v)
// }
// func (cc *Clock) SetOldTime(v float64) {
// 	cc.Set("oldTime", v)
// }
// func (cc *Clock) Running() bool {
// 	return cc.Get("running").Bool()
// }
// func (cc *Clock) SetRunning(v bool) {
// 	cc.Set("running", v)
// }
// func (cc *Clock) StartTime() float64 {
// 	return cc.Get("startTime").Float()
// }
// func (cc *Clock) SetStartTime(v float64) {
// 	cc.Set("startTime", v)
// }
