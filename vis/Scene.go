// Scene.go
package vis

import (
	"fmt"
	"go4/geom"
)

// scene is the default implementation of Scene interface
type scene struct {
	meshes []*geom.Mesh
}

// NewScene creates a new empty scene
func NewScene() Scene {
	return &scene{
		meshes: make([]*geom.Mesh, 0),
	}
}

// AddMesh adds a mesh to the scene
func (s *scene) AddMesh(m *geom.Mesh) {
	if m == nil {
		return
	}
	s.meshes = append(s.meshes, m)
}

// RemoveMesh removes a mesh from the scene by index
func (s *scene) RemoveMesh(index int) error {
	if index < 0 || index >= len(s.meshes) {
		return fmt.Errorf("mesh index out of bounds: %d (scene has %d meshes)", index, len(s.meshes))
	}
	
	// Remove element by creating new slice without it
	s.meshes = append(s.meshes[:index], s.meshes[index+1:]...)
	return nil
}

// GetMeshes returns all meshes in the scene
func (s *scene) GetMeshes() []*geom.Mesh {
	// Return a copy to prevent external modification
	result := make([]*geom.Mesh, len(s.meshes))
	copy(result, s.meshes)
	return result
}

// Clear removes all meshes from the scene
func (s *scene) Clear() {
	s.meshes = make([]*geom.Mesh, 0)
}

// MeshCount returns the number of meshes in the scene
func (s *scene) MeshCount() int {
	return len(s.meshes)
}
