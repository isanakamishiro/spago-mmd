package mmd

import (
	"app/lib/threejs"
	"app/lib/threejs/animation"
)

type Future interface {
	// Loaded gets loaded bytes.
	Loaded() uint
	// Total gets estimated total bytes.
	Total() uint
	// Err returns error. If error is not happened, return nil.
	Err() error
}

type FutureMesh interface {
	Future

	// Mesh gets loaded mesh.
	Mesh() threejs.SkinnedMesh
}

type FutureClip interface {
	Future

	// Clip gets loaded clip.
	Clip() animation.Clip
}

type FutureVpd interface {
	Future

	// Vpd gets loaded vpd file.
	Vpd() Vpd
}

type futureImp struct {
	loaded uint
	total  uint
	err    error
}

type futureMeshImp struct {
	futureImp

	mesh threejs.SkinnedMesh
}

type futureClipImp struct {
	futureImp

	clip animation.Clip
}

type futureVpdImp struct {
	futureImp

	vpd Vpd
}

// NewFutureMesh creates FutureMesh.
func NewFutureMesh(mesh threejs.SkinnedMesh, loaded uint, total uint, err error) FutureMesh {
	return &futureMeshImp{
		mesh: mesh,
		futureImp: futureImp{
			loaded: loaded,
			total:  total,
			err:    err,
		},
	}
}

// NewFutureClip creates FutureClip.
func NewFutureClip(clip animation.Clip, loaded uint, total uint, err error) FutureClip {
	return &futureClipImp{
		clip: clip,
		futureImp: futureImp{
			loaded: loaded,
			total:  total,
			err:    err,
		},
	}
}

// NewFutureVpd creates FutureVpd.
func NewFutureVpd(vpd Vpd, loaded uint, total uint, err error) FutureVpd {
	return &futureVpdImp{
		vpd: vpd,
		futureImp: futureImp{
			loaded: loaded,
			total:  total,
			err:    err,
		},
	}
}

func (c *futureImp) Loaded() uint {
	return c.loaded
}

func (c *futureImp) Total() uint {
	return c.total
}

func (c *futureImp) Err() error {
	return c.err
}

func (c *futureMeshImp) Mesh() threejs.SkinnedMesh {
	return c.mesh
}

func (c *futureClipImp) Clip() animation.Clip {
	return c.clip
}

func (c *futureVpdImp) Vpd() Vpd {
	return c.vpd
}
