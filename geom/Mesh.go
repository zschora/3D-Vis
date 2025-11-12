// Mesh.go
package geom

import "fmt"

type Triangle struct {
	myVertexIndices [3]int
	myNormal        Vector
}

func (t *Triangle) SetNormal(normal Vector) {
	t.myNormal = normal
}

type Mesh struct {
	myVertices []Vertex
	myFaces    []Triangle
}

func (m *Mesh) VertexNumber() int {
	return len(m.myVertices)
}

func (m *Mesh) FaceNumber() int {
	return len(m.myFaces)
}

func (m *Mesh) AddVertex(v Vertex) int {
	m.myVertices = append(m.myVertices, v)
	return len(m.myVertices) - 1
}

// ComputeNormal calculates the normal vector for a triangle defined by three vertices
func ComputeNormal(v1, v2, v3 Vertex) Vector {
	side1 := NewVectorFromVertices(v1, v2)
	side2 := NewVectorFromVertices(v1, v3)
	normal := side1.Cross(side2)
	normal.Normalize()

	return normal
}

func (m *Mesh) AddFace(v1, v2, v3 int) (int, error) {
	// Validate vertex indices
	if v1 < 0 || v1 >= len(m.myVertices) {
		return -1, fmt.Errorf("vertex index v1 out of bounds: %d (mesh has %d vertices)", v1, len(m.myVertices))
	}
	if v2 < 0 || v2 >= len(m.myVertices) {
		return -1, fmt.Errorf("vertex index v2 out of bounds: %d (mesh has %d vertices)", v2, len(m.myVertices))
	}
	if v3 < 0 || v3 >= len(m.myVertices) {
		return -1, fmt.Errorf("vertex index v3 out of bounds: %d (mesh has %d vertices)", v3, len(m.myVertices))
	}

	normal := ComputeNormal(m.myVertices[v1], m.myVertices[v2], m.myVertices[v3])
	face := Triangle{
		myVertexIndices: [3]int{v1, v2, v3},
		myNormal:        normal,
	}
	m.myFaces = append(m.myFaces, face)

	return len(m.myFaces) - 1, nil
}

func (m *Mesh) SetFaceNormal(faceIndex int, normal Vector) error {
	if faceIndex < 0 || faceIndex >= len(m.myFaces) {
		return fmt.Errorf("face index out of bounds: %d (mesh has %d faces)", faceIndex, len(m.myFaces))
	}
	m.myFaces[faceIndex].SetNormal(normal)
	return nil
}

func (m *Mesh) Vertex(vertexIndex int) (Vertex, error) {
	if vertexIndex < 0 || vertexIndex >= len(m.myVertices) {
		return Vertex{}, fmt.Errorf("vertex index out of bounds: %d (mesh has %d vertices)", vertexIndex, len(m.myVertices))
	}
	return m.myVertices[vertexIndex], nil
}

func (m *Mesh) VertexInFace(faceIndex, vertexIndex int) (Vertex, error) {
	if faceIndex < 0 || faceIndex >= len(m.myFaces) {
		return Vertex{}, fmt.Errorf("face index out of bounds: %d (mesh has %d faces)", faceIndex, len(m.myFaces))
	}
	if vertexIndex < 0 || vertexIndex >= 3 {
		return Vertex{}, fmt.Errorf("vertex index in face out of bounds: %d (must be 0, 1, or 2)", vertexIndex)
	}
	realVertexIndex := m.myFaces[faceIndex].myVertexIndices[vertexIndex]
	return m.Vertex(realVertexIndex)
}

func (m *Mesh) Normal(faceIndex int) (Vector, error) {
	if faceIndex < 0 || faceIndex >= len(m.myFaces) {
		return Vector{}, fmt.Errorf("face index out of bounds: %d (mesh has %d faces)", faceIndex, len(m.myFaces))
	}
	return m.myFaces[faceIndex].myNormal, nil
}
