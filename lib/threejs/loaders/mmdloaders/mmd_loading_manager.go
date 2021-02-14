package mmdloaders

import "syscall/js"

// MMDLoadingManager is ...
type MMDLoadingManager interface {
	JSValue() js.Value
}
