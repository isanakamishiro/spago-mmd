package threejs

import (
	"errors"
	"syscall/js"
)

// MaterialParameters is ...
type MaterialParameters interface {
}

// Material is the appearance of objects.
// They are defined in a (mostly) renderer-independent way, so you don't have to rewrite materials if you decide to use a different renderer.
// The following properties and methods are inherited by all other material types (although they may have different defaults).
type Material interface {
	JSValue() js.Value

	DepthTest() bool
	SetDepthTest(b bool)

	// FlatShading gets whether the material is rendered with flat shading. Default is false.
	FlatShading() bool

	// SetFlatShading sets whether the material is rendered with flat shading. Default is false.
	SetFlatShading(b bool)

	// NeedsUpdate gets that the material needs to be recompiled.
	NeedsUpdate() bool

	// SetFlatShading sets whether the material is rendered with flat shading. Default is false.
	SetNeedsUpdate(b bool)

	// Side gets which side of faces will be rendered - front, back or both.
	// Default is THREE.FrontSide. Other options are THREE.BackSide and THREE.DoubleSide.
	Side() Side

	// SetSide sets which side of faces will be rendered - front, back or both.
	// Default is THREE.FrontSide. Other options are THREE.BackSide and THREE.DoubleSide.
	SetSide(v Side)

	// Opacity gets float in the range of 0.0 - 1.0 indicating how transparent the material is. A value of 0.0 indicates fully transparent, 1.0 is fully opaque.
	// If the material's transparent property is not set to true, the material will remain fully opaque and this value will only affect its color.
	// Default is 1.0.
	Opacity() float64

	// SetOpacity sets float in the range of 0.0 - 1.0 indicating how transparent the material is. A value of 0.0 indicates fully transparent, 1.0 is fully opaque.
	// If the material's transparent property is not set to true, the material will remain fully opaque and this value will only affect its color.
	// Default is 1.0.
	SetOpacity(v float64)

	// Disposes the material. Textures of a material don't get disposed. These needs to be disposed by Texture.
	Dispose()

	// Map gets The color map. Default is null.
	Map() (Texture, error)

	// EnvMap gets The environment map. To ensure a physically correct rendering, you should only add environment maps which were preprocessed by PMREMGenerator. Default is null.
	EnvMap() (Texture, error)

	// GradientMap gets Gradient map for toon shading. It's required to set Texture.minFilter and Texture.magFilter to THREE.NearestFilter when using this type of texture. Default is null.
	GradientMap() (Texture, error)

	// AlphaMap gets The alpha map is a grayscale texture that controls the opacity across the surface (black: fully transparent; white: fully opaque). Default is null.
	AlphaMap() (Texture, error)

	// AOMap gets The red channel of this texture is used as the ambient occlusion map. Default is null. The aoMap requires a second set of UVs.
	AOMap() (Texture, error)

	// BumpMap gets The texture to create a bump map. The black and white values map to the perceived depth in relation to the lights. Bump doesn't actually affect the geometry of the object, only the lighting. If a normal map is defined this will be ignored.
	BumpMap() (Texture, error)

	// EmissiveMap gets emisssive (glow) map. Default is null. The emissive map color is modulated by the emissive color and the emissive intensity. If you have an emissive map, be sure to set the emissive color to something other than black.
	EmissiveMap() (Texture, error)

	// LightMap gets The light map. Default is null. The lightMap requires a second set of UVs.
	LightMap() (Texture, error)

	// NormalMap gets The texture to create a normal map.
	// The RGB values affect the surface normal for each pixel fragment and change the way the color is lit.
	// Normal maps do not change the actual shape of the surface, only the lighting.
	// In case the material has a normal map authored using the left handed convention, the y component of normalScale should be negated to compensate for the different handedness.
	NormalMap() (Texture, error)
}

// MaterialImpl extend: [EventDispatcher]
type defaultMaterialImpl struct {
	js.Value
}

// NewDefaultMaterialFromJSValue is ...
func NewDefaultMaterialFromJSValue(value js.Value) Material {
	return &defaultMaterialImpl{Value: value}
}

// JSValue is ...
func (m *defaultMaterialImpl) JSValue() js.Value {
	return m.Value
}

// DepthTest is ...
func (m *defaultMaterialImpl) DepthTest() bool {
	return m.Get("depthTest").Bool()
}

// SetDepthTest is ...
func (m *defaultMaterialImpl) SetDepthTest(b bool) {
	m.Set("depthTest", b)
}

// FlatShading gets whether the material is rendered with flat shading. Default is false.
func (m *defaultMaterialImpl) FlatShading() bool {
	return m.Get("flatShading").Bool()
}

// SetFlatShading sets whether the material is rendered with flat shading. Default is false.
func (m *defaultMaterialImpl) SetFlatShading(b bool) {
	m.Set("flatShading", b)
}

// NeedsUpdate gets that the material needs to be recompiled.
func (m *defaultMaterialImpl) NeedsUpdate() bool {
	return m.Get("needsUpdate").Bool()
}

// SetNeedsUpdate sets that the material needs to be recompiled.
func (m *defaultMaterialImpl) SetNeedsUpdate(b bool) {
	m.Set("needsUpdate", b)
}

// Side gets which side of faces will be rendered - front, back or both.
// Default is THREE.FrontSide. Other options are THREE.BackSide and THREE.DoubleSide.
func (m *defaultMaterialImpl) Side() Side {
	return SideOf(m.Get("side"))
}

// SetSide sets which side of faces will be rendered - front, back or both.
// Default is THREE.FrontSide. Other options are THREE.BackSide and THREE.DoubleSide.
func (m *defaultMaterialImpl) SetSide(v Side) {
	m.Set("side", v.JSValue())
}

func (m *defaultMaterialImpl) Opacity() float64 {
	return m.Get("opacity").Float()
}

func (m *defaultMaterialImpl) SetOpacity(v float64) {
	m.Set("opacity", v)
}

func (m *defaultMaterialImpl) Dispose() {
	m.Call("dispose")
	// t := m.Get("dispose")
	// log.Println(t)
}

func (m *defaultMaterialImpl) Map() (Texture, error) {
	tx := m.Get("map")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("map is null or undefined")
	}
	return NewDefaultTextureFromJSValue(tx), nil
}

func (m *defaultMaterialImpl) EnvMap() (Texture, error) {
	tx := m.Get("envMap")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("env map is null or undefined")
	}
	return NewDefaultTextureFromJSValue(tx), nil
}

func (m *defaultMaterialImpl) GradientMap() (Texture, error) {
	tx := m.Get("gradientMap")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("gradient map is null or undefined")
	}
	return NewDefaultTextureFromJSValue(tx), nil
}

func (m *defaultMaterialImpl) AlphaMap() (Texture, error) {
	tx := m.Get("alphaMap")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("alpha map is null or undefined")
	}
	return NewDefaultTextureFromJSValue(tx), nil
}

func (m *defaultMaterialImpl) AOMap() (Texture, error) {
	tx := m.Get("alphaMap")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("ao map is null or undefined")
	}
	return NewDefaultTextureFromJSValue(tx), nil
}

func (m *defaultMaterialImpl) BumpMap() (Texture, error) {
	tx := m.Get("alphaMap")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("bump map is null or undefined")
	}
	return NewDefaultTextureFromJSValue(tx), nil
}

func (m *defaultMaterialImpl) EmissiveMap() (Texture, error) {
	tx := m.Get("alphaMap")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("emissiv map is null or undefined")
	}
	return NewDefaultTextureFromJSValue(tx), nil
}

func (m *defaultMaterialImpl) LightMap() (Texture, error) {
	tx := m.Get("alphaMap")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("light map is null or undefined")
	}
	return NewDefaultTextureFromJSValue(tx), nil
}

func (m *defaultMaterialImpl) NormalMap() (Texture, error) {
	tx := m.Get("alphaMap")
	if tx.IsUndefined() || tx.IsNull() {
		return nil, errors.New("normal map is null or undefined")
	}
	return NewDefaultTextureFromJSValue(tx), nil
}
