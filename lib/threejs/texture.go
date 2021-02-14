package threejs

import "syscall/js"

// Texture is ...
type Texture interface {
	EventDispatcher

	SetWrapS(s Wrapping)
	SetWrapT(t Wrapping)
	SetMagFilter(v TextureFilter)
	SetMinFilter(v TextureFilter)

	// Offset gets how much a single repetition of the texture is offset from the beginning,
	// in each direction U and V. Typical range is 0.0 to 1.0.
	// _Note:_ The offset property is a convenience modifier and only affects the Texture's application to the first set of UVs on a model. If the Texture is used as a map requiring additional UV sets (e.g. the aoMap or lightMap of most stock materials), those UVs must be manually assigned to achieve the desired offset.
	Offset() *Vector2

	// Repeat gets how many times the texture is repeated across the surface,
	// in each direction U and V. If repeat is set greater than 1 in either direction,
	// the corresponding Wrap parameter should also be set to THREE.RepeatWrapping or
	// THREE.MirroredRepeatWrapping to achieve the desired tiling effect.
	// _Note:_ The repeat property is a convenience modifier and only affects the Texture's application to the first set of UVs on a model. If the Texture is used as a map requiring additional UV sets (e.g. the aoMap or lightMap of most stock materials), those UVs must be manually assigned to achieve the desired repetiton.
	Repeat() *Vector2

	// SetNeedsUpdate set this to true to trigger an update next time the texture is used. Particularly important for setting the wrap mode.
	SetNeedsUpdate(b bool)
}

type textureImp struct {
	js.Value
}

// NewDefaultTextureFromJSValue creates Texture from js.Value
func NewDefaultTextureFromJSValue(v js.Value) Texture {
	return &textureImp{
		Value: v,
	}
}

// AddEventListener is ...
func (c *textureImp) AddEventListener(typ string, listener js.Value) {
	c.Call("addEventListener", typ, listener)
}

// RemoveEventListener is ...
func (c *textureImp) RemoveEventListener(typ string, listener js.Value) {
	c.Call("removeEventListener", typ, listener)
}

// HasEventListener is ...
func (c *textureImp) HasEventListener(typ string, listener js.Value) bool {
	return c.Call("hasEventListener", typ, listener).Bool()
}

// DispatchEvent is ...
func (c *textureImp) DispatchEvent(event js.Value) {
	c.Call("dispatchEvent", event)
}

// SetWrapS is ...
func (c *textureImp) SetWrapS(s Wrapping) {
	c.Set("wrapS", s.JSValue())
}

// SetWrapT is ...
func (c *textureImp) SetWrapT(t Wrapping) {
	c.Set("wrapT", t.JSValue())
}

// SetMagFilter is ...
func (c *textureImp) SetMagFilter(v TextureFilter) {
	c.Set("magFilter", v.JSValue())
}

// SetMinFilter is ...
func (c *textureImp) SetMinFilter(v TextureFilter) {
	c.Set("minFilter", v.JSValue())
}

// Offset is ...
func (c *textureImp) Offset() *Vector2 {
	return &Vector2{
		Value: c.Get("offset"),
	}
}

// Repeat is ...
func (c *textureImp) Repeat() *Vector2 {
	return &Vector2{
		Value: c.Get("repeat"),
	}
}

// SetNeedsUpdate is ...
func (c *textureImp) SetNeedsUpdate(b bool) {
	c.Set("needsUpdate", b)
}
