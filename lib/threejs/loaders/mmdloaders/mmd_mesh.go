package mmdloaders

import (
	"app/lib/threejs"
	"syscall/js"
)

// MMDMesh ...
type MMDMesh interface {
	threejs.Mesh
}

type mmdMeshImp struct {
	threejs.Mesh
}

func newMMDMeshFromJSValue(v js.Value) MMDMesh {
	return &mmdMeshImp{
		Mesh: threejs.NewSkinnedMeshFromJSValue(v),
	}
}
