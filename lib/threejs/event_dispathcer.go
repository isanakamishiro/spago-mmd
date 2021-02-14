package threejs

import "syscall/js"

// EventDispatcher is interface for threejs.EventDispatcher object
type EventDispatcher interface {
	JSValue() js.Value

	AddEventListener(typ string, listener js.Value)
	HasEventListener(typ string, listener js.Value) bool
	RemoveEventListener(typ string, listener js.Value)
	DispatchEvent(event js.Value)
}
