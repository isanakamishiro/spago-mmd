package water

import (
	"app/lib/threejs"
	"app/lib/threejs/geometry"
	"app/lib/threejs/material"
	"app/lib/threejs/shader/fresnel"
	"app/lib/threejs/texture"
)

// Ball is water ball.
type Ball struct {
	threejs.Mesh
}

// NewBall creates water Ball.
func NewBall(tx texture.CubeTexture) *Ball {

	g := geometry.NewSphereGeometry(1, 32, 32)
	shader := fresnel.NewFresnelShader()

	u := shader.Uniforms()
	u.Get("tCube").Set("value", tx.JSValue())

	m := material.NewShaderMaterial(map[string]interface{}{
		"uniforms":       u,
		"vertexShader":   shader.VertexShader(),
		"fragmentShader": shader.FragmentShader(),
	})

	b := &Ball{
		Mesh: threejs.NewMesh(g, m),
	}

	return b
}
