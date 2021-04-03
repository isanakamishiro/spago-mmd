package water

import "app/lib/threejs"

// OceanOption is functional parameter option for water.
type OceanOption func(map[string]interface{}) error

// TextureSize sets texture size(width and height).
func TextureSize(w, h float64) OceanOption {
	return func(m map[string]interface{}) error {

		m["textureWidth"] = w
		m["textureHeight"] = h

		return nil
	}
}

// Alpha sets alpha.
func Alpha(a float64) OceanOption {
	return func(m map[string]interface{}) error {

		m["alpha"] = a

		return nil
	}
}

// OceanColor sets ocean color.
func OceanColor(c threejs.Color) OceanOption {
	return func(m map[string]interface{}) error {

		m["waterColor"] = c.Hex()

		return nil
	}
}

// SunColor sets ...
func SunColor(c threejs.Color) OceanOption {
	return func(m map[string]interface{}) error {

		m["sunColor"] = c.Hex()

		return nil
	}
}

// SunDirection sets ...
func SunDirection(v threejs.Vector3) OceanOption {
	return func(m map[string]interface{}) error {

		m["sunDirection"] = v.JSValue()

		return nil
	}
}

// DistortionScale sets distortionScale.
func DistortionScale(s float64) OceanOption {
	return func(m map[string]interface{}) error {

		m["distortionScale"] = s

		return nil
	}
}

// Fog sets ...
func Fog(b bool) OceanOption {
	return func(m map[string]interface{}) error {

		m["fog"] = b

		return nil
	}
}

// NormalizeTexture sets normal texture for water.
func NormalizeTexture(tx threejs.Texture) OceanOption {
	return func(m map[string]interface{}) error {

		m["waterNormals"] = tx.JSValue()

		return nil
	}
}
