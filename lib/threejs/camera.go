package threejs

// Camera is camera object interface for three.js
type Camera interface {
	Object3D

	// MatrixWorldInverse is the inverse of matrixWorld.
	// MatrixWorld contains the Matrix which has the world transform of the Camera.
	MatrixWorldInverse() *Matrix4

	// ProjectionMatrix is the matrix which contains the projection.
	ProjectionMatrix() *Matrix4

	// ProjectionMatrixInverse is the inverse of projectionMatrix.
	ProjectionMatrixInverse() *Matrix4
	SetProjectionMatrix(v *Matrix4)

	WorldDirection(target *Vector3) *Vector3
	UpdateMatrixWorld(force bool)

	IsCamera() bool
}

// cameraImpl extend: [Object3D]
// type cameraImpl struct {
// 	Object3D
// }
