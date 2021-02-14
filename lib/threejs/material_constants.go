package threejs

import "syscall/js"

// Side defines which side of faces will be rendered - front, back or both. Default is FrontSide.
type Side int

var sideMap map[Side]js.Value = make(map[Side]js.Value)

func getSideDictionary() map[Side]js.Value {
	if len(sideMap) == 0 {
		sideMap[FrontSide] = GetJsObject("FrontSide")
		sideMap[BackSide] = GetJsObject("BackSide")
		sideMap[DoubleSide] = GetJsObject("DoubleSide")
	}
	return sideMap
}

const (
	// Undefined is ...
	Undefined Side = iota
	// FrontSide is ...
	FrontSide
	// BackSide is ...
	BackSide
	// DoubleSide is ...
	DoubleSide
)

// JSValue return js.Value for Side
func (c Side) JSValue() js.Value {
	dic := getSideDictionary()
	if v, ok := dic[c]; ok {
		return v
	}
	return js.Null()
}

// SideOf converts js.Value to Side constants.
func SideOf(v js.Value) Side {
	dic := getSideDictionary()
	for key, val := range dic {
		if val.Equal(v) {
			return key
		}
	}
	return Undefined
}
