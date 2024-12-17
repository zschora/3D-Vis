// Scene.go
package vis

import "go4/geom"

type Scene struct {
    meshes []*geom.Mesh
}

func NewScene() *Scene {
    return &Scene{
        meshes: make([]*geom.Mesh, 0),
    }
}

func (s *Scene) AddMesh (m *geom.Mesh) {
    s.meshes = append (s.meshes, m)
}
