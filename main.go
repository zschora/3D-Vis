package main

import (
	"go4/geom"
	"go4/vis"
	"go4/vis/gui"
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

	// Setup basic navigation GUI
	setupNavigationGUI(app)

	// Run the application
	app.Run()
}

func setupNavigationGUI(app *vis.Application) {
	guiManager := app.GetGUI()

	// Create info panel
	infoPanel := gui.NewInfoPanel(gui.InfoPanelConfig{
		X: 10,
		Y: 10,
	})
	guiManager.AddElement(infoPanel.GetPanel())

	// Create navigation panel
	camera := app.GetRenderer().GetCamera()
	navPanel := gui.NewNavigationPanel(gui.NavigationPanelConfig{
		X: 10,
		Y: 170,
	}, gui.NavigationCallbacks{
		OnReset: func() {
			cam, _ := vis.NewCamera(vis.DefaultCameraConfig())
			app.GetRenderer().SetCamera(cam)
			camera = cam
		},
		OnZoomIn: func() {
			camera.ScaleLinear(-50)
		},
		OnZoomOut: func() {
			camera.ScaleLinear(50)
		},
	})
	guiManager.AddElement(navPanel.GetPanel())

	// Set up update function for camera controls and GUI
	app.SetUpdateFunction(func(deltaTime time.Duration) {
		deltaSeconds := deltaTime.Seconds()

		// Update info panel
		infoPanel.SetFPS(rl.GetFPS())
		infoPanel.SetCameraInfo(
			camera.GetPolarAngle(),
			camera.GetAzimuth(),
			camera.GetDistanceToScreen(),
		)
		infoPanel.SetSceneCount(len(app.GetScenes()))

		// Handle navigation panel input
		navPanel.HandleInput(gui.NavigationCallbacks{
			OnReset: func() {
				cam, _ := vis.NewCamera(vis.DefaultCameraConfig())
				app.GetRenderer().SetCamera(cam)
				camera = cam
			},
			OnZoomIn: func() {
				camera.ScaleLinear(-50)
			},
			OnZoomOut: func() {
				camera.ScaleLinear(50)
			},
		})

		// Manual camera controls with arrow keys
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
}
