package threejs

import (
	"syscall/js"
)

// CallbackRepository is simple repository of js.Func object.
type CallbackRepository interface {
	Register(fn func(this js.Value, args []js.Value) interface{}) js.Func
	ReleaseAll()
}

type callbackRepositoryImp struct {
	list []js.Func
}

// NewCallbackRepository creates CallbackRepository of default implementation.
func NewCallbackRepository() CallbackRepository {
	return &callbackRepositoryImp{}
}

func (c *callbackRepositoryImp) Register(fn func(this js.Value, args []js.Value) interface{}) js.Func {

	var cb js.Func
	cb = js.FuncOf(fn)

	c.list = append(c.list, cb)

	return cb
}

func (c *callbackRepositoryImp) ReleaseAll() {
	for _, cb := range c.list {

		cb.Release()
	}

	c.list = make([]js.Func, 0)

}
