package geom

import "testing"

func TestMesh_AddFace_OutOfBounds(t *testing.T) {
	mesh := &Mesh{}

	// Добавляем одну вершину
	mesh.AddVertex(NewVertex(1, 0, 0))

	// Пытаемся добавить грань с несуществующими индексами
	_, err := mesh.AddFace(0, 1, 2)
	if err == nil {
		t.Error("Expected error for out of bounds vertex indices, got nil")
	}
}

func TestMesh_Vertex_OutOfBounds(t *testing.T) {
	mesh := &Mesh{}
	mesh.AddVertex(NewVertex(1, 0, 0))

	_, err := mesh.Vertex(1)
	if err == nil {
		t.Error("Expected error for out of bounds vertex index, got nil")
	}

	_, err = mesh.Vertex(-1)
	if err == nil {
		t.Error("Expected error for negative vertex index, got nil")
	}
}

func TestMesh_VertexInFace_OutOfBounds(t *testing.T) {
	mesh := &Mesh{}
	v1 := mesh.AddVertex(NewVertex(1, 0, 0))
	v2 := mesh.AddVertex(NewVertex(0, 1, 0))
	v3 := mesh.AddVertex(NewVertex(0, 0, 1))

	faceIndex, _ := mesh.AddFace(v1, v2, v3)

	// Тест несуществующей грани
	_, err := mesh.VertexInFace(999, 0)
	if err == nil {
		t.Error("Expected error for out of bounds face index, got nil")
	}

	// Тест несуществующего индекса вершины в грани
	_, err = mesh.VertexInFace(faceIndex, 3)
	if err == nil {
		t.Error("Expected error for out of bounds vertex index in face, got nil")
	}
}

func TestMesh_Normal_OutOfBounds(t *testing.T) {
	mesh := &Mesh{}
	v1 := mesh.AddVertex(NewVertex(1, 0, 0))
	v2 := mesh.AddVertex(NewVertex(0, 1, 0))
	v3 := mesh.AddVertex(NewVertex(0, 0, 1))

	mesh.AddFace(v1, v2, v3)

	_, err := mesh.Normal(999)
	if err == nil {
		t.Error("Expected error for out of bounds face index, got nil")
	}
}

func TestMesh_SetFaceNormal_OutOfBounds(t *testing.T) {
	mesh := &Mesh{}
	v1 := mesh.AddVertex(NewVertex(1, 0, 0))
	v2 := mesh.AddVertex(NewVertex(0, 1, 0))
	v3 := mesh.AddVertex(NewVertex(0, 0, 1))

	mesh.AddFace(v1, v2, v3)

	normal := NewVector(0, 0, 1)
	err := mesh.SetFaceNormal(999, normal)
	if err == nil {
		t.Error("Expected error for out of bounds face index, got nil")
	}
}
