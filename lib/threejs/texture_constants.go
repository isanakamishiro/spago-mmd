package threejs

import "syscall/js"

// Mapping defines the texture's mapping mode.
// UVMapping is the default, and maps the texture using the mesh's UV coordinates.
type Mapping int

// Wrapping defines the texture's wrapS and wrapT properties,
// which define horizontal and vertical texture wrapping.
type Wrapping int

// TextureFilter is for use with a texture's minFilter/magFilter property,
// these define the texture minifying/magnification function that is used whenever the pixel being textured maps to an area greater than one texture element (texel).
type TextureFilter int

const (
	// UVMapping is ...
	UVMapping Mapping = iota
	// CubeReflectionMapping is ...
	CubeReflectionMapping
	// CubeRefractionMapping is ...
	CubeRefractionMapping
	// EquirectangularReflectionMapping is ...
	EquirectangularReflectionMapping
	// EquirectangularRefractionMapping is ...
	EquirectangularRefractionMapping
	// CubeUVReflectionMapping is ...
	CubeUVReflectionMapping
	// CubeUVRefractionMapping is ...
	CubeUVRefractionMapping
)

const (
	// RepeatWrapping is ...
	RepeatWrapping Wrapping = iota
	// ClampToEdgeWrapping is ...
	ClampToEdgeWrapping
	// MirroredRepeatWrapping is ...
	MirroredRepeatWrapping
)

const (
	// NearestFilter is ...
	NearestFilter TextureFilter = iota
	// NearestMipmapNearestFilter is ...
	NearestMipmapNearestFilter
	// NearestMipmapLinearFilter is ...
	NearestMipmapLinearFilter
	// LinearFilter is ...
	LinearFilter
	// LinearMipmapNearestFilter is ...
	LinearMipmapNearestFilter
	// LinearMipmapLinearFilter is ...
	LinearMipmapLinearFilter
)

var mappingDic map[Mapping]js.Value = make(map[Mapping]js.Value)
var wrappingDic map[Wrapping]js.Value = make(map[Wrapping]js.Value)
var textureFilterDic map[TextureFilter]js.Value = make(map[TextureFilter]js.Value)

func getMappingDictionary() map[Mapping]js.Value {
	if len(mappingDic) == 0 {
		mappingDic[UVMapping] = GetJsObject("UVMapping")
		mappingDic[CubeReflectionMapping] = GetJsObject("CubeReflectionMapping")
		mappingDic[CubeRefractionMapping] = GetJsObject("CubeRefractionMapping")
		mappingDic[EquirectangularReflectionMapping] = GetJsObject("EquirectangularReflectionMapping")
		mappingDic[EquirectangularRefractionMapping] = GetJsObject("EquirectangularRefractionMapping")
		mappingDic[CubeUVReflectionMapping] = GetJsObject("CubeUVReflectionMapping")
		mappingDic[CubeUVRefractionMapping] = GetJsObject("CubeUVRefractionMapping")
	}
	return mappingDic
}

func getWrappingDictionary() map[Wrapping]js.Value {
	if len(wrappingDic) == 0 {
		wrappingDic[RepeatWrapping] = GetJsObject("RepeatWrapping")
		wrappingDic[ClampToEdgeWrapping] = GetJsObject("ClampToEdgeWrapping")
		wrappingDic[MirroredRepeatWrapping] = GetJsObject("MirroredRepeatWrapping")
	}
	return wrappingDic
}

func getTextureFilterDictionary() map[TextureFilter]js.Value {
	if len(textureFilterDic) == 0 {
		textureFilterDic[NearestFilter] = GetJsObject("NearestFilter")
		textureFilterDic[NearestMipmapNearestFilter] = GetJsObject("NearestMipmapNearestFilter")
		textureFilterDic[NearestMipmapLinearFilter] = GetJsObject("NearestMipmapLinearFilter")

		textureFilterDic[LinearFilter] = GetJsObject("LinearFilter")
		textureFilterDic[LinearMipmapNearestFilter] = GetJsObject("LinearMipmapNearestFilter")
		textureFilterDic[LinearMipmapLinearFilter] = GetJsObject("LinearMipmapLinearFilter")
	}
	return textureFilterDic
}

// JSValue return js.Value for Mapping
func (c Mapping) JSValue() js.Value {
	dic := getMappingDictionary()
	if v, ok := dic[c]; ok {
		return v
	}
	return js.Null()
}

// JSValue return js.Value for Wrapping
func (c Wrapping) JSValue() js.Value {
	dic := getWrappingDictionary()
	if v, ok := dic[c]; ok {
		return v
	}
	return js.Null()
}

// JSValue return js.Value for TextureFilter
func (c TextureFilter) JSValue() js.Value {
	dic := getTextureFilterDictionary()
	if v, ok := dic[c]; ok {
		return v
	}
	return js.Null()
}
