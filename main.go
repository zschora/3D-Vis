package main

import (
	"go4/geom"
	"go4/vis"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Create application with default configuration
	app, err := vis.NewApplicationWithDefaults()
	if err != nil {
		panic(err)
	}

	// Create a scene and add a cube
	scene := vis.NewScene()
	cube := geom.CreateCube(200)
	scene.AddMesh(cube)
	app.AddScene(scene)

	// Set up update function for camera controls
	app.SetUpdateFunction(func(deltaTime time.Duration) {
		deltaSeconds := deltaTime.Seconds()
		camera := app.GetRenderer().GetCamera()

		// Rotate camera with arrow keys
		if rl.IsKeyDown(rl.KeyLeft) {
			camera.RotatePolar(-deltaSeconds)
		}
		if rl.IsKeyDown(rl.KeyRight) {
			camera.RotatePolar(deltaSeconds)
		}
		if rl.IsKeyDown(rl.KeyUp) {
			camera.RotateAzimuth(-deltaSeconds)
		}
		if rl.IsKeyDown(rl.KeyDown) {
			camera.RotateAzimuth(deltaSeconds)
		}

		// Zoom with Q/E keys
		if rl.IsKeyDown(rl.KeyQ) {
			camera.ScaleLinear(-deltaSeconds * 500)
		}
		if rl.IsKeyDown(rl.KeyE) {
			camera.ScaleLinear(deltaSeconds * 500)
		}
	})

	// Run the application
	app.Run()
}
