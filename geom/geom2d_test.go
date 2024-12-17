package geom

import (
	"math"
	"testing"
)

func TestCoords2d_Equals(t *testing.T) {
	c1 := Coords2d{X: 1.0, Y: 2.0}
	c2 := Coords2d{X: 1.0, Y: 2.0}
	c3 := Coords2d{X: 1.0, Y: 2.000001}

	if !c1.Equals(c2) {
		t.Errorf("Expected Coords2d %v to equal %v", c1, c2)
	}

	if c1.Equals(c3) {
		t.Errorf("Expected Coords2d %v to not equal %v (within tolerance)", c1, c3)
	}
}

func TestCoords2d_Length(t *testing.T) {
	c := Coords2d{X: 3.0, Y: 4.0}
	expectedLength := 5.0 // 3-4-5 triangle

	if math.Abs(c.Length()-expectedLength) > 1e-8 {
		t.Errorf("Expected length to be %v, got %v", expectedLength, c.Length())
	}
}

func TestVertex2d_Subtract(t *testing.T) {
	v1 := NewVertex2d(3.0, 4.0)
	v2 := NewVertex2d(1.0, 2.0)

	v1.Subtract(v2)
	expected := NewVertex2d(2.0, 2.0)

	if v1.X() != expected.X() || v1.Y() != expected.Y() {
		t.Errorf("Subtract failed: expected %v, got %v", expected, v1)
	}
}

func TestMesh2d_AddVertexAndFace(t *testing.T) {
	mesh := Mesh2d{}

	// Добавляем вершины
	v1 := mesh.AddVertex(NewVertex2d(0, 0))
	v2 := mesh.AddVertex(NewVertex2d(1, 0))
	v3 := mesh.AddVertex(NewVertex2d(0, 1))

	if mesh.VertexNumber() != 3 {
		t.Errorf("Expected 3 vertices, got %v", mesh.VertexNumber())
	}

	// Добавляем грань
	mesh.AddFace(v1, v2, v3)

	if mesh.FaceNumber() != 1 {
		t.Errorf("Expected 1 face, got %v", mesh.FaceNumber())
	}
}

func TestMesh2d_VertexInFace(t *testing.T) {
	mesh := Mesh2d{}

	// Добавляем вершины
	v1 := mesh.AddVertex(NewVertex2d(0, 0))
	v2 := mesh.AddVertex(NewVertex2d(1, 0))
	v3 := mesh.AddVertex(NewVertex2d(0, 1))

	// Добавляем грань
	mesh.AddFace(v1, v2, v3)

	// Проверяем вершины в грани
	vInFace1 := mesh.VertexInFace(0, 0)
	vInFace2 := mesh.VertexInFace(0, 1)
	vInFace3 := mesh.VertexInFace(0, 2)

	if vInFace1.X() != 0 || vInFace1.Y() != 0 {
		t.Errorf("VertexInFace failed: expected (0,0), got (%v,%v)", vInFace1.X(), vInFace1.Y())
	}
	if vInFace2.X() != 1 || vInFace2.Y() != 0 {
		t.Errorf("VertexInFace failed: expected (1,0), got (%v,%v)", vInFace2.X(), vInFace2.Y())
	}
	if vInFace3.X() != 0 || vInFace3.Y() != 1 {
		t.Errorf("VertexInFace failed: expected (0,1), got (%v,%v)", vInFace3.X(), vInFace3.Y())
	}
}
