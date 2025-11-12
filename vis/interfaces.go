package vis

import "go4/geom"

// Camera defines the interface for camera operations
type Camera interface {
	// Transform converts a world-space vertex to screen-space coordinates
	Transform(v geom.Vertex, screenWidth, screenHeight int) geom.Vertex2d

	// RotatePolar rotates the camera around the polar axis
	RotatePolar(angle float64)

	// RotateAzimuth rotates the camera around the azimuth axis
	RotateAzimuth(angle float64)

	// ScaleLinear adjusts the camera distance
	ScaleLinear(distance float64)

	// GetRadius returns the camera radius
	GetRadius() float64

	// GetPolarAngle returns the current polar angle
	GetPolarAngle() float64

	// GetAzimuth returns the current azimuth angle
	GetAzimuth() float64

	// GetDistanceToScreen returns the distance to screen
	GetDistanceToScreen() float64
}

// Renderer defines the interface for rendering operations
type Renderer interface {
	// Render renders a scene
	Render(scene Scene)

	// SetCamera sets the camera for rendering
	SetCamera(camera Camera)

	// GetCamera returns the current camera
	GetCamera() Camera
}

// Scene defines the interface for scene management
type Scene interface {
	// AddMesh adds a mesh to the scene
	AddMesh(mesh *geom.Mesh)

	// RemoveMesh removes a mesh from the scene by index
	RemoveMesh(index int) error

	// GetMeshes returns all meshes in the scene
	GetMeshes() []*geom.Mesh

	// Clear removes all meshes from the scene
	Clear()

	// MeshCount returns the number of meshes in the scene
	MeshCount() int
}
