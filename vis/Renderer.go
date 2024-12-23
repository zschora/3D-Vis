// Renderer.go
package vis

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"go4/geom"
	"math/rand"
)

type Renderer struct {
	MyCamera     *Camera
	screenWidth  int
	screenHeight int
}

func NewRenderer(camera *Camera) *Renderer {
	//rl.DisableBackfaceCulling()
	return &Renderer{MyCamera: camera}
}

func (r *Renderer) Render(sc *Scene) {
	r.screenWidth = rl.GetScreenWidth()
	r.screenHeight = rl.GetScreenHeight()

	for _, mesh := range sc.meshes {
		r.RenderMesh(mesh)
	}
}

// Метод для отрисовки меша
func (r *Renderer) RenderMesh(mesh *geom.Mesh) {
	rand.Seed(0)
	faceNumber := mesh.FaceNumber()
	for i := 0; i < faceNumber; i++ {
		//if i != 2 && i != 3 {
		//    continue
		//}
		v1, v2, v3 := mesh.VertexInFace(i, 0), mesh.VertexInFace(i, 1), mesh.VertexInFace(i, 2)
		r.RenderFace(v1, v2, v3)
		if i == 0 {
			//println(v1.X(), v1.Y(), v1.Z())
		}
	}
}

func (r *Renderer) RenderFace(v1, v2, v3 geom.Vertex) {
	v2d1, v2d2, v2d3 := r.convertTo2D(v1), r.convertTo2D(v2), r.convertTo2D(v3)

	color := rl.Color{
		R: uint8((v2d1.X + v2d1.Y) / 1500 * 255),
		G: uint8((v2d2.Y + v2d2.X) / 1500 * 255),
		B: uint8((v2d3.X + v2d3.Y) / 1500 * 255),
		A: 200}

	//color := rl.Red

	rl.DrawTriangle(v2d1, v2d2, v2d3, color)
}

// Метод преобразования 3D в 2D (пример)
func (r *Renderer) convertTo2D(v geom.Vertex) rl.Vector2 {
	// Просто пример, в реальной ситуации может быть использована проекция
	vTransformed := r.MyCamera.Transform(v, r.screenWidth, r.screenHeight)

	return rl.Vector2{float32(vTransformed.X()), float32(vTransformed.Y())}
}
