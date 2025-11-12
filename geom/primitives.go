// Package geom provides geometric primitives for 2D and 3D modeling and visualization.
package geom

// CreateCube creates a cube mesh with the specified size
func CreateCube(size float64) *Mesh {
	mesh := &Mesh{}

	halfSize := size / 2
	v1 := mesh.AddVertex(NewVertex(halfSize, -halfSize, -halfSize))
	v2 := mesh.AddVertex(NewVertex(-halfSize, -halfSize, -halfSize))
	v3 := mesh.AddVertex(NewVertex(-halfSize, halfSize, -halfSize))
	v4 := mesh.AddVertex(NewVertex(halfSize, halfSize, -halfSize))
	v5 := mesh.AddVertex(NewVertex(halfSize, -halfSize, halfSize))
	v6 := mesh.AddVertex(NewVertex(-halfSize, -halfSize, halfSize))
	v7 := mesh.AddVertex(NewVertex(-halfSize, halfSize, halfSize))
	v8 := mesh.AddVertex(NewVertex(halfSize, halfSize, halfSize))

	// Front face
	_, _ = mesh.AddFace(v5, v1, v8)
	_, _ = mesh.AddFace(v4, v8, v1)

	// Back face
	_, _ = mesh.AddFace(v7, v8, v3)
	_, _ = mesh.AddFace(v4, v3, v8)

	// Right face
	_, _ = mesh.AddFace(v7, v6, v8)
	_, _ = mesh.AddFace(v5, v8, v6)

	// Left face
	_, _ = mesh.AddFace(v5, v6, v1)
	_, _ = mesh.AddFace(v2, v1, v6)

	// Top face
	_, _ = mesh.AddFace(v6, v7, v2)
	_, _ = mesh.AddFace(v3, v2, v7)

	// Bottom face
	_, _ = mesh.AddFace(v2, v3, v1)
	_, _ = mesh.AddFace(v4, v1, v3)

	return mesh
}

// CreateTetrahedron creates a tetrahedron mesh
func CreateTetrahedron(size float64) *Mesh {
	mesh := &Mesh{}

	halfSize := size / 2
	v1 := mesh.AddVertex(NewVertex(-halfSize, -halfSize, -halfSize))
	v2 := mesh.AddVertex(NewVertex(halfSize, 0, 0))
	v3 := mesh.AddVertex(NewVertex(0, halfSize, 0))
	v4 := mesh.AddVertex(NewVertex(0, 0, halfSize))

	_, _ = mesh.AddFace(v2, v3, v4)
	_, _ = mesh.AddFace(v1, v2, v4)
	_, _ = mesh.AddFace(v1, v3, v2)
	_, _ = mesh.AddFace(v1, v4, v3)

	return mesh
}
