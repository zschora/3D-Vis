// vector_test.go
package geom

import (
	"fmt"
	"math"
	"testing"
)

// Constructors

func TestNewVertex(t *testing.T) {
	// Test NewVertex
	v1 := NewVertex(1, 2, 3)
	if v1.myCoords.X != 1 || v1.myCoords.Y != 2 || v1.myCoords.Z != 3 {
		t.Errorf("NewVertex failed, got: %v", v1.myCoords)
	}
}

func TestNewVector(t *testing.T) {
	// Test NewVector
	v1 := NewVector(1, 2, 3)
	if v1.myCoords.X != 1 || v1.myCoords.Y != 2 || v1.myCoords.Z != 3 {
		t.Errorf("NewVector failed, got: %v", v1.myCoords)
	}

	// Test NewVectorFromVertex
	vertex := Vertex{myCoords: Coords3d{4, 5, 6}}
	v2 := NewVectorFromVertex(vertex)
	if v2.myCoords.X != 4 || v2.myCoords.Y != 5 || v2.myCoords.Z != 6 {
		t.Errorf("NewVectorFromVertex failed, got: %v", v2.myCoords)
	}

	// Test NewVectorFromVertices
	begin := Vertex{myCoords: Coords3d{1, 1, 1}}
	end := Vertex{myCoords: Coords3d{4, 5, 6}}
	v3 := NewVectorFromVertices(begin, end)
	if v3.myCoords.X != 3 || v3.myCoords.Y != 4 || v3.myCoords.Z != 5 {
		t.Errorf("NewVectorFromVertices failed, got: %v", v3.myCoords)
	}
}

// Operations

func TestLength(t *testing.T) {
	cases := []struct {
		vector   Vector
		expected float64
	}{
		{NewVector(1, 2, 3), 3.7416573867739413},
		{NewVector(0, 0, 0), 0},
		{NewVector(1, 2, -3), 3.7416573867739413},
	}

	for _, test := range cases {
		t.Run(
			fmt.Sprintf("%v", test.vector),
			func(t *testing.T) {
				if test.vector.Length() != test.expected {
					t.Errorf("Expected %f but got %f", test.expected, test.vector.Length())
				}
			},
		)
	}
}

func TestNormalize(t *testing.T) {
	cases := []struct {
		vector   Vector
		expected Vector
	}{
		{NewVector(0, 6, 8), NewVector(0, 0.6, 0.8)},
		{NewVector(0, 0, 0), NewVector(0, 0, 0)},
	}

	for _, test := range cases {
		t.Run(
			fmt.Sprintf("%v", test.vector),
			func(t *testing.T) {
				test.vector.Normalize()
				if !test.vector.Equals(test.expected) {
					t.Errorf("Expected %v but got %v", test.expected, test.vector)
				}
			},
		)
	}
}

func TestDot(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(4, 5, 6)
	v3 := NewVector(0, 0, 0)

	tests := []struct {
		v1, v2   Vector
		expected float64
	}{
		{v1, v2, 32},
		{v1, v3, 0},
	}

	for _, test := range tests {
		t.Run("Dot", func(t *testing.T) {
			result := test.v1.Dot(test.v2)
			if math.Abs(result-test.expected) > 1e-9 {
				t.Errorf("Expected %f, but got %f", test.expected, result)
			}
		})
	}
}

func TestCross(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(4, 5, 6)
	v3 := NewVector(0, 0, 0)

	tests := []struct {
		v1, v2   Vector
		expected Vector
	}{
		{v1, v2, NewVector(-3, 6, -3)},
		{v1, v3, NewVector(0, 0, 0)},
	}

	for _, test := range tests {
		t.Run("Cross", func(t *testing.T) {
			result := test.v1.Cross(test.v2)
			if !result.Equals(test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func TestAngle(t *testing.T) {
	v1 := NewVector(1, 0, 0)
	v2 := NewVector(0, 1, 0)
	v3 := NewVector(0, 0, 0)

	tests := []struct {
		v1, v2   Vector
		expected float64
	}{
		{v1, v2, math.Pi / 2}, // Угол 90 градусов
		{v1, v1, 0},           // Угол 0 градусов (одинаковые направления)
		{v1, v3, math.NaN()},  // Угол с нулевым вектором
	}

	for _, test := range tests {
		t.Run("Angle", func(t *testing.T) {
			result := test.v1.Angle(test.v2)
			if math.IsNaN(result) && !math.IsNaN(test.expected) {
				t.Errorf("Expected NaN, but got %f", result)
			} else if !math.IsNaN(result) && math.Abs(result-test.expected) > 1e-9 {
				t.Errorf("Expected %f, but got %f", test.expected, result)
			}
		})
	}
}

// Mesh

func TestMesh_AddVertex(t *testing.T) {
	mesh := &Mesh{}

	// Добавляем вершины
	v1 := NewVertex(1, 0, 0)
	v2 := NewVertex(0, 1, 0)
	v3 := NewVertex(0, 0, 1)

	// Проверка добавления вершин
	index1 := mesh.AddVertex(v1)
	index2 := mesh.AddVertex(v2)
	index3 := mesh.AddVertex(v3)

	// Проверяем количество вершин в модели
	if mesh.VertexNumber() != 3 {
		t.Errorf("Expected 3 vertices, got %d", mesh.VertexNumber())
	}

	// Проверяем, что индексы совпадают
	if index1 != 0 {
		t.Errorf("Expected index 0 for first vertex, got %d", index1)
	}
	if index2 != 1 {
		t.Errorf("Expected index 1 for second vertex, got %d", index2)
	}
	if index3 != 2 {
		t.Errorf("Expected index 2 for third vertex, got %d", index3)
	}
}

func TestMesh_AddFace(t *testing.T) {
	mesh := &Mesh{}

	// Добавляем вершины
	v1 := NewVertex(1, 0, 0)
	v2 := NewVertex(0, 1, 0)
	v3 := NewVertex(0, 0, 1)

	index1 := mesh.AddVertex(v1)
	index2 := mesh.AddVertex(v2)
	index3 := mesh.AddVertex(v3)

	// Добавляем грань
	_, err := mesh.AddFace(index1, index2, index3)
	if err != nil {
		t.Errorf("AddFace failed: %v", err)
	}

	// Проверка количества граней
	if mesh.FaceNumber() != 1 {
		t.Errorf("Expected 1 face, got %d", mesh.FaceNumber())
	}
}

func TestMesh_SetFaceNormal(t *testing.T) {
	mesh := &Mesh{}

	// Добавляем вершины
	v1 := NewVertex(1, 0, 0)
	v2 := NewVertex(0, 1, 0)
	v3 := NewVertex(0, 0, 1)

	index1 := mesh.AddVertex(v1)
	index2 := mesh.AddVertex(v2)
	index3 := mesh.AddVertex(v3)

	// Добавляем грань
	faceIndex, err := mesh.AddFace(index1, index2, index3)
	if err != nil {
		t.Errorf("AddFace failed: %v", err)
	}

	// Устанавливаем нормаль для грани
	normal := NewVector(0, 0, 1)
	err = mesh.SetFaceNormal(faceIndex, normal)
	if err != nil {
		t.Errorf("SetFaceNormal failed: %v", err)
	}

	// Проверяем нормаль
	meshNormal, err := mesh.Normal(faceIndex)
	if err != nil {
		t.Errorf("Normal failed: %v", err)
	}
	if !normal.Equals(meshNormal) {
		t.Errorf("Expected normal %v, got %v", normal, meshNormal)
	}
}

func TestMesh_Vertex(t *testing.T) {
	mesh := &Mesh{}

	// Добавляем вершины
	v1 := NewVertex(1, 0, 0)
	v2 := NewVertex(0, 1, 0)
	v3 := NewVertex(0, 0, 1)

	index1 := mesh.AddVertex(v1)
	index2 := mesh.AddVertex(v2)
	index3 := mesh.AddVertex(v3)

	// Добавляем грань
	faceIndex, err := mesh.AddFace(index1, index2, index3)
	if err != nil {
		t.Errorf("AddFace failed: %v", err)
	}

	// Проверяем, что вершин можно правильно получить
	vertex1, err1 := mesh.VertexInFace(faceIndex, 0)
	vertex2, err2 := mesh.VertexInFace(faceIndex, 1)
	vertex3, err3 := mesh.VertexInFace(faceIndex, 2)

	if err1 != nil || err2 != nil || err3 != nil {
		t.Errorf("VertexInFace failed: %v, %v, %v", err1, err2, err3)
	}

	// Проверка значений координат вершин
	if vertex1.myCoords != v1.myCoords {
		t.Errorf("Expected vertex1 to be %v, got %v", v1.myCoords, vertex1.myCoords)
	}
	if vertex2.myCoords != v2.myCoords {
		t.Errorf("Expected vertex2 to be %v, got %v", v2.myCoords, vertex2.myCoords)
	}
	if vertex3.myCoords != v3.myCoords {
		t.Errorf("Expected vertex3 to be %v, got %v", v3.myCoords, vertex3.myCoords)
	}
}
