package text

import "app/lib/threejs"

// FontTextureAtlasOption is functional parameter option interface.
type FontTextureAtlasOption func(*FontTextureAtlas) error

// FontSize is option for setting font size.
func FontSize(s int) FontTextureAtlasOption {
	return func(f *FontTextureAtlas) error {
		f.setFontSize(s)
		return nil
	}
}

// FontName is option for setting font name.
func FontName(name string) FontTextureAtlasOption {
	return func(f *FontTextureAtlas) error {
		f.setFontName(name)
		return nil
	}
}

// FontColor is option for setting font color.
func FontColor(col threejs.Color) FontTextureAtlasOption {
	return func(f *FontTextureAtlas) error {
		c := "#" + col.HexString()
		f.setFontColor(c)
		return nil
	}
}

// CanvasWidth is option for setting canvas width.
func CanvasWidth(w int) FontTextureAtlasOption {
	return func(f *FontTextureAtlas) error {
		f.setCanvasWidth(w)
		return nil
	}
}

// CanvasHeight is option for setting canvas height.
func CanvasHeight(h int) FontTextureAtlasOption {
	return func(f *FontTextureAtlas) error {
		f.setCanvasHeight(h)
		return nil
	}
}

// OffsetX is option for setting canvas offset x.
func OffsetX(x int) FontTextureAtlasOption {
	return func(f *FontTextureAtlas) error {
		f.setOffsetX(x)
		return nil
	}
}

// OffsetY is option for setting canvas offset x.
func OffsetY(y int) FontTextureAtlasOption {
	return func(f *FontTextureAtlas) error {
		f.setOffsetY(y)
		return nil
	}
}
