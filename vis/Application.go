package vis

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "go4/geom"
    "math"
)

type Application struct {
    width    int
    height   int
    renderer *Renderer
    scenes   []*Scene
}

func CreateCube() *geom.Mesh {
	mesh := &geom.Mesh{}

	v1 := mesh.AddVertex (geom.NewVertex (200, 0, 0))
	v2 := mesh.AddVertex (geom.NewVertex (0, 0, 0))
	v3 := mesh.AddVertex (geom.NewVertex (0, 200, 0))
	v4 := mesh.AddVertex (geom.NewVertex (200, 200, 0))
	v5 := mesh.AddVertex (geom.NewVertex (200, 0, 200))
	v6 := mesh.AddVertex (geom.NewVertex (0, 0, 200))
	v7 := mesh.AddVertex (geom.NewVertex (0, 200, 200))
	v8 := mesh.AddVertex (geom.NewVertex (200, 200, 200))

	mesh.AddFace (v5, v1, v8)
	mesh.AddFace (v4, v8, v1)
	mesh.AddFace (v7, v8, v3)
	mesh.AddFace (v4, v3, v8)
	mesh.AddFace (v7, v6, v8)
	mesh.AddFace (v5, v8, v6)
	mesh.AddFace (v5, v6, v1)
	mesh.AddFace (v2, v1, v6)
	mesh.AddFace (v6, v7, v2)
	mesh.AddFace (v3, v2, v7)
	mesh.AddFace (v2, v3, v1)
	mesh.AddFace (v4, v1, v3)

	return mesh
}

func CreateTetrahedron() *geom.Mesh {
	mesh := &geom.Mesh{}

	v1 := mesh.AddVertex (geom.NewVertex (-50, -50, -50))
	v2 := mesh.AddVertex (geom.NewVertex (200, 0, 0))
	v3 := mesh.AddVertex (geom.NewVertex (0, 200, 0))
	v4 := mesh.AddVertex (geom.NewVertex (0, 0, 200))

	mesh.AddFace (v2, v3, v4)
	mesh.AddFace (v1, v2, v4)
	mesh.AddFace (v1, v3, v2)
	mesh.AddFace (v1, v4, v3)

	return mesh
}

func NewApplication(width, height int) *Application {
    rl.InitWindow(int32(width), int32(height), "3D Visualization App")

    app := &Application{
        width:    width,
        height:   height,
        renderer: NewRenderer(NewCamera(1000, math.Pi / 4, math.Pi / 4, 500)),
        scenes:   make([]*Scene, 0),
    }

    // сцена с тестовом мешем
    sc := NewScene()
    sc.AddMesh(CreateCube())
    app.scenes = append(app.scenes, sc)


    return app
}

func (app *Application) Run() {
    rl.SetTargetFPS(60)

    for !rl.WindowShouldClose() {
        app.Update()
        app.Render()
    }

    app.Close()
}

func (app *Application) Update() {
    deltaTime := float64(rl.GetFrameTime())  // время между кадрами
    //app.renderer.MyCamera.RotateAzimuth (deltaTime)
    //app.renderer.MyCamera.RotatePolar (deltaTime)
    app.renderer.MyCamera.ScaleLinear (+deltaTime*500)
}

func (app *Application) Render() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.LightGray)

    for _, scene := range app.scenes {
        app.renderer.Render(scene)
    }

    rl.EndDrawing()
}

func (app *Application) Close() {
    rl.CloseWindow()
}
