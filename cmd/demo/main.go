// Package main provides a demo application for testing 3D visualization
// This is a simple way to quickly test the visualization during development
package main

import (
	"go4/geom"
	"go4/vis"
	"go4/vis/gui"
	"math"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Create application without auto-loading test scene (we'll create it manually)
	config := vis.DefaultApplicationConfig()
	config.LoadTestScene = false
	config.Title = "3D Visualization Demo"

	app, err := vis.NewApplication(config)
	if err != nil {
		panic(err)
	}

	// Setup GUI and scene
	setupGUI(app)

	// Run the application
	app.Run()
}

func setupGUI(app *vis.Application) {
	guiManager := app.GetGUI()

	// Create scene
	scene := vis.NewScene()
	currentMesh := geom.CreateCube(200)
	scene.AddMesh(currentMesh)
	app.AddScene(scene)

	// Create info panel
	infoPanel := gui.NewInfoPanel(gui.InfoPanelConfig{
		X: 10,
		Y: 10,
	})
	guiManager.AddElement(infoPanel.GetPanel())

	// Create primitive selector
	primitiveSelector := gui.NewPrimitiveSelector(gui.PrimitiveSelectorConfig{
		X: 10,
		Y: 170,
	}, func(pt gui.PrimitiveType) {
		// Change primitive
		scene.Clear()
		var newMesh *geom.Mesh
		if pt == gui.PrimitiveCube {
			newMesh = geom.CreateCube(200)
		} else {
			newMesh = geom.CreateTetrahedron(200)
		}
		scene.AddMesh(newMesh)
		currentMesh = newMesh
	})
	guiManager.AddElement(primitiveSelector.GetPanel())

	// Create motion selector
	currentMotion := gui.MotionRotate
	motionSelector := gui.NewMotionSelector(gui.MotionSelectorConfig{
		X: 10,
		Y: 280,
	}, func(mt gui.MotionType) {
		currentMotion = mt
	})
	guiManager.AddElement(motionSelector.GetPanel())

	// Create control panel
	camera := app.GetRenderer().GetCamera()
	controlPanel := gui.NewControlPanel(gui.ControlPanelConfig{
		X: 10,
		Y: 490,
	}, gui.ControlCallbacks{}) // Empty callbacks, will be handled in update
	guiManager.AddElement(controlPanel.GetPanel())

	// Motion state
	motionTime := 0.0

	// Update function to refresh GUI and handle motion
	app.SetUpdateFunction(func(deltaTime time.Duration) {
		deltaSeconds := deltaTime.Seconds()
		motionTime += deltaSeconds

		// Update info panel
		infoPanel.SetFPS(rl.GetFPS())
		infoPanel.SetCameraInfo(
			camera.GetPolarAngle(),
			camera.GetAzimuth(),
			camera.GetDistanceToScreen(),
		)
		infoPanel.SetSceneCount(len(app.GetScenes()))

		// Handle primitive selector
		primitiveSelector.HandleInput()

		// Handle motion selector
		motionSelector.HandleInput()
		currentMotion = motionSelector.GetSelectedMotion()

		// Handle control panel input
		controlPanel.HandleInput(gui.ControlCallbacks{
			OnReset: func() {
				cam, _ := vis.NewCamera(vis.DefaultCameraConfig())
				app.GetRenderer().SetCamera(cam)
				camera = cam
				motionTime = 0
			},
			OnRotate: func() {
				// Toggle motion (not used, motion is controlled by motion selector)
			},
			OnZoomIn: func() {
				camera.ScaleLinear(-50)
			},
			OnZoomOut: func() {
				camera.ScaleLinear(50)
			},
		})

		// Apply selected motion
		switch currentMotion {
		case gui.MotionRotate:
			camera.RotatePolar(deltaSeconds * 0.5)
		case gui.MotionZoom:
			// Sinusoidal zoom
			zoomAmount := math.Sin(motionTime) * deltaSeconds * 200
			camera.ScaleLinear(zoomAmount)
		case gui.MotionRotateAndZoom:
			camera.RotatePolar(deltaSeconds * 0.5)
			zoomAmount := math.Sin(motionTime) * deltaSeconds * 200
			camera.ScaleLinear(zoomAmount)
		case gui.MotionOrbit:
			// Orbit: rotate both polar and azimuth
			camera.RotatePolar(deltaSeconds * 0.5)
			camera.RotateAzimuth(deltaSeconds * 0.3)
		case gui.MotionNone:
			// No automatic motion
		}

		// Manual camera controls
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
		if rl.IsKeyDown(rl.KeyQ) {
			camera.ScaleLinear(-deltaSeconds * 500)
		}
		if rl.IsKeyDown(rl.KeyE) {
			camera.ScaleLinear(deltaSeconds * 500)
		}
	})
}
