package mmd

import "syscall/js"

type Vpd interface {
	JSValue() js.Value
}

type vpdImp struct {
	js.Value
}

func NewVpdFromJSValue(v js.Value) Vpd {
	return &vpdImp{
		Value: v,
	}
}
