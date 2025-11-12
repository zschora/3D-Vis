// Renderer.go
package vis

import (
	"go4/geom"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// RendererConfig holds configuration for the renderer
type RendererConfig struct {
	BackgroundColor  rl.Color
	ColorScaleFactor float64
	AlphaValue       uint8
}

// DefaultRendererConfig returns default renderer configuration
func DefaultRendererConfig() RendererConfig {
	return RendererConfig{
		BackgroundColor:  rl.LightGray,
		ColorScaleFactor: 1500.0,
		AlphaValue:       200,
	}
}

// renderer is the default implementation of Renderer interface
type renderer struct {
	camera       Camera
	screenWidth  int
	screenHeight int
	config       RendererConfig
}

// NewRenderer creates a new renderer with the given camera and configuration
func NewRenderer(camera Camera, config RendererConfig) Renderer {
	return &renderer{
		camera: camera,
		config: config,
	}
}

// NewRendererWithDefaults creates a new renderer with default settings
func NewRendererWithDefaults(camera Camera) Renderer {
	return NewRenderer(camera, DefaultRendererConfig())
}

// Render renders a scene
func (r *renderer) Render(scene Scene) {
	r.screenWidth = rl.GetScreenWidth()
	r.screenHeight = rl.GetScreenHeight()

	for _, mesh := range scene.GetMeshes() {
		r.RenderMesh(mesh)
	}
}

// SetCamera sets the camera for rendering
func (r *renderer) SetCamera(camera Camera) {
	r.camera = camera
}

// GetCamera returns the current camera
func (r *renderer) GetCamera() Camera {
	return r.camera
}

// RenderMesh renders all faces of the mesh
func (r *renderer) RenderMesh(mesh *geom.Mesh) {
	faceNumber := mesh.FaceNumber()
	for i := 0; i < faceNumber; i++ {
		v1, err1 := mesh.VertexInFace(i, 0)
		v2, err2 := mesh.VertexInFace(i, 1)
		v3, err3 := mesh.VertexInFace(i, 2)

		// Skip face if any vertex access fails
		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}

		r.RenderFace(v1, v2, v3)
	}
}

const (
	maxColorValue = 255.0
)

func (r *renderer) RenderFace(v1, v2, v3 geom.Vertex) {
	v2d1, v2d2, v2d3 := r.convertTo2D(v1), r.convertTo2D(v2), r.convertTo2D(v3)

	color := rl.Color{
		R: uint8((float64(v2d1.X) + float64(v2d1.Y)) / r.config.ColorScaleFactor * maxColorValue),
		G: uint8((float64(v2d2.Y) + float64(v2d2.X)) / r.config.ColorScaleFactor * maxColorValue),
		B: uint8((float64(v2d3.X) + float64(v2d3.Y)) / r.config.ColorScaleFactor * maxColorValue),
		A: r.config.AlphaValue,
	}

	rl.DrawTriangle(v2d1, v2d2, v2d3, color)
}

// convertTo2D transforms a 3D vertex to 2D screen coordinates
func (r *renderer) convertTo2D(v geom.Vertex) rl.Vector2 {
	vTransformed := r.camera.Transform(v, r.screenWidth, r.screenHeight)
	return rl.Vector2{
		X: float32(vTransformed.X()),
		Y: float32(vTransformed.Y()),
	}
}
