// Mesh2d.go
package geom

type Triangle2d struct {
	myVertexIndices [3]int
	myIsVisible     bool // may be not need it
}

type Mesh2d struct {
	myVertices []Vertex2d
	myFaces    []Triangle2d
}

func (m *Mesh2d) VertexNumber() int {
	return len(m.myVertices)
}

func (m *Mesh2d) FaceNumber() int {
	return len(m.myFaces)
}

func (m *Mesh2d) AddVertex(v Vertex2d) int {
	m.myVertices = append(m.myVertices, v)
	return len(m.myVertices) - 1
}

func (m *Mesh2d) AddFace(v1, v2, v3 int) int {
	face := Triangle2d{
		myVertexIndices: [3]int{v1, v2, v3},
		myIsVisible:     true,
	}
	m.myFaces = append(m.myFaces, face)

	return len(m.myFaces) - 1
}

func (m *Mesh2d) Vertex(vertexIndex int) Vertex2d {
	return m.myVertices[vertexIndex]
}

func (m *Mesh2d) VertexInFace(faceIndex, vertexIndex int) Vertex2d {
	realVertexIndex := m.myFaces[faceIndex].myVertexIndices[vertexIndex]
	return m.Vertex(realVertexIndex)
}

func (m *Mesh2d) IsVisible(faceIndex int) bool {
	return m.myFaces[faceIndex].myIsVisible
}
