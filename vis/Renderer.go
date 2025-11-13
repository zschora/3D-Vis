// Renderer.go
package vis

import (
	"go4/geom"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// RendererConfig holds configuration for the renderer
type RendererConfig struct {
	BackgroundColor    rl.Color
	UseRandomFaceColor bool
	FaceColor          rl.Color
	EdgeColor          rl.Color
	AlphaValue         uint8
	DrawFaces          bool
	DrawEdges          bool
	UseBackfaceCulling bool
}

// DefaultRendererConfig returns default renderer configuration
func DefaultRendererConfig() RendererConfig {
	return RendererConfig{
		BackgroundColor:    rl.LightGray,
		AlphaValue:         220,
		UseRandomFaceColor: false,
		FaceColor:          rl.Gray,
		EdgeColor:          rl.Black,
		DrawFaces:          true,
		DrawEdges:          true,
		UseBackfaceCulling: true,
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
	cameraPosition := r.cameraPosition()
	for i := 0; i < faceNumber; i++ {
		v1, err1 := mesh.VertexInFace(i, 0)
		v2, err2 := mesh.VertexInFace(i, 1)
		v3, err3 := mesh.VertexInFace(i, 2)

		// Skip face if any vertex access fails
		if err1 != nil || err2 != nil || err3 != nil {
			continue
		}

		r.renderFace(v1, v2, v3, cameraPosition)
	}
}

const (
	maxColorValue = 255.0
)

// renderFace performs backface culling before rendering the triangle and its edges.
// This prevents invisible (back-facing) faces and lines from being drawn.
func (r *renderer) renderFace(v1, v2, v3 geom.Vertex, cameraPosition geom.Vector) {

	if r.config.UseBackfaceCulling && !r.checkVisibility(v1, v2, v3, cameraPosition) {
		return
	}

	// Transform vertices to screen space
	v2d1, v2d2, v2d3 := r.convertTo2D(v1), r.convertTo2D(v2), r.convertTo2D(v3)

	if r.config.DrawFaces {
		rl.DrawTriangle(v2d1, v2d2, v2d3, r.getFaceColor(v2d1, v2d2, v2d3))
	}

	if r.config.DrawEdges {
		rl.DrawTriangleLines(v2d1, v2d2, v2d3, r.config.EdgeColor)
	}
}

func (r *renderer) getFaceColor(v1, v2, v3 rl.Vector2) rl.Color {
	faceColor := r.config.FaceColor
	faceColor.A = r.config.AlphaValue
	colorScaleFactor := float64(r.screenHeight) + float64(r.screenWidth)
	if r.config.UseRandomFaceColor {
		faceColor = rl.Color{
			R: uint8((float64(v1.X) + float64(v1.Y)) / colorScaleFactor * maxColorValue),
			G: uint8((float64(v2.Y) + float64(v2.X)) / colorScaleFactor * maxColorValue),
			B: uint8((float64(v3.X) + float64(v3.Y)) / colorScaleFactor * maxColorValue),
			A: r.config.AlphaValue,
		}
	}

	return faceColor
}

func (r *renderer) checkVisibility(v1, v2, v3 geom.Vertex, cameraPosition geom.Vector) bool {
	// Compute normal using original 3D vertices
	// The camera looks toward -Z by convention
	edge1 := geom.NewVectorFromVertices(v1, v2)
	edge2 := geom.NewVectorFromVertices(v1, v3)
	normal := edge1.Cross(edge2)
	normal.Normalize() // outward normal

	// Calculate vector from camera to face position
	cameraToFace := geom.NewVectorFromVertex(v1)
	cameraToFace = cameraToFace.Subtracted(cameraPosition)

	return normal.Dot(cameraToFace) < 0
}

func (r *renderer) cameraPosition() geom.Vector {
	radius := r.camera.GetRadius()
	polar := r.camera.GetPolarAngle()
	azimuth := r.camera.GetAzimuth()

	sinAzimuth := math.Sin(azimuth)
	cosAzimuth := math.Cos(azimuth)
	sinPolar := math.Sin(polar)
	cosPolar := math.Cos(polar)

	x := radius * sinAzimuth * cosPolar
	y := radius * sinAzimuth * sinPolar
	z := radius * cosAzimuth

	return geom.NewVector(x, y, z)
}

// convertTo2D transforms a 3D vertex to 2D screen coordinates
func (r *renderer) convertTo2D(v geom.Vertex) rl.Vector2 {
	vTransformed := r.camera.Transform(v, r.screenWidth, r.screenHeight)
	return rl.Vector2{
		X: float32(vTransformed.X()),
		Y: float32(vTransformed.Y()),
	}
}

// SetConfig replaces renderer configuration at runtime.
func (r *renderer) SetConfig(config RendererConfig) {
	r.config = config
}

// GetConfig returns a copy of the current renderer configuration.
func (r *renderer) GetConfig() RendererConfig {
	return r.config
}
