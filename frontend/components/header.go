package components

import (
	"app/frontend/actions"
	"app/frontend/store"
	"syscall/js"

	"github.com/nobonobo/spago"
	"github.com/nobonobo/spago/dispatcher"
)

//go:generate spago generate -c Header -p components header.html

// Header  ...
type Header struct {
	spago.Core

	isActive string
}

// NewHeader creates Header.
func NewHeader() *Header {
	return &Header{
		isActive: "",
	}
}

func (c *Header) toggleMenu(ev js.Value) {

	if c.isActive != "" {
		c.isActive = ""
	} else {
		c.isActive = "is-active"
	}

	dispatcher.Dispatch(actions.Refresh)
}

func (c *Header) changeModelToDiluc(ev js.Value) {

	if store.CurrentModel != store.Diluc {
		store.CurrentModel = store.Diluc

		dispatcher.Dispatch(actions.ChangeModel)
	}

}

func (c *Header) changeModelToLisa(ev js.Value) {

	if store.CurrentModel != store.Lisa {
		store.CurrentModel = store.Lisa

		dispatcher.Dispatch(actions.ChangeModel)
	}

}

func (c *Header) changeModelToMiku(ev js.Value) {

	if store.CurrentModel != store.Miku {
		store.CurrentModel = store.Miku

		dispatcher.Dispatch(actions.ChangeModel)
	}

}

func (c *Header) changeMotionToDance1(ev js.Value) {

	store.CurrentMotion = store.Dance1
	dispatcher.Dispatch(actions.ChangeMotion)
}

func (c *Header) changeMotionToDance2(ev js.Value) {

	store.CurrentMotion = store.Dance2
	dispatcher.Dispatch(actions.ChangeMotion)

}

func (c *Header) changeMotionToDance3(ev js.Value) {

	store.CurrentMotion = store.Dance3
	dispatcher.Dispatch(actions.ChangeMotion)

}
