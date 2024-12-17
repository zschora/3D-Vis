// Mesh.go
package geom

type Triangle struct {
    myVertexIndices [3]int
    myNormal        Vector
}

func (t *Triangle) SetNormal (normal Vector) {
    t.myNormal = normal
}

type Mesh struct {
    myVertices  []Vertex
    myFaces     []Triangle
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

func ComputeNormal(v1, v2, v3 Vertex) Vector {
    side1 := NewVectorFromVertices(v1, v2)
    side2 := NewVectorFromVertices(v1, v3)
    normal := side1.Cross(side2)
    normal.Normalize()
    
    return normal
}

func (m *Mesh) AddFace(v1, v2, v3 int) int {
    normal := ComputeNormal(m.Vertex (v1), m.Vertex(v2), m.Vertex(v3))
    face := Triangle{
        myVertexIndices:    [3]int{v1, v2, v3},
        myNormal:           normal,
    }
    m.myFaces = append(m.myFaces, face)
    
    return len(m.myFaces) - 1
}

func (m *Mesh) SetFaceNormal (faceIndex int, normal Vector) {
    m.myFaces[faceIndex].SetNormal (normal)
}

func (m *Mesh) Vertex (vertexIndex int) Vertex {
    return m.myVertices[vertexIndex]
}

func (m *Mesh) VertexInFace (faceIndex, vertexIndex int) Vertex {
    realVertexIndex := m.myFaces[faceIndex].myVertexIndices[vertexIndex]
    return m.Vertex(realVertexIndex)
}

func (m *Mesh) Normal (faceIndex int) Vector {
    return m.myFaces[faceIndex].myNormal
}
