package water

import (
	"app/lib/threejs"
	"app/lib/threejs/geometries"
	"app/lib/threejs/materials"
	"app/lib/threejs/shaders/fresnel"
	"app/lib/threejs/textures"
)

// Ball is water ball.
type Ball struct {
	threejs.Mesh
}

// NewBall creates water Ball.
func NewBall(tx textures.CubeTexture) *Ball {

	g := geometries.NewSphereGeometry(1, 32, 32)
	shader := fresnel.NewFresnelShader()

	u := shader.Uniforms()
	u.Get("tCube").Set("value", tx.JSValue())

	m := materials.NewShaderMaterial(map[string]interface{}{
		"uniforms":       u,
		"vertexShader":   shader.VertexShader(),
		"fragmentShader": shader.FragmentShader(),
	})

	b := &Ball{
		Mesh: threejs.NewMesh(g, m),
	}

	return b
}
