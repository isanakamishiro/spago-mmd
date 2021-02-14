package mmdloaders

import "syscall/js"

// MMDAnimation is ...
type MMDAnimation interface {
	JSValue() js.Value
}

type mmdAnimationImp struct {
	js.Value
}

func newMMDAnimationFromJSValue(v js.Value) MMDAnimation {
	return &mmdAnimationImp{
		Value: v,
	}
}
