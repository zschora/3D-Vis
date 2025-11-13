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

type scenarioEntry struct {
	data  gui.Scenario
	setup func()
}

func main() {
	config := vis.DefaultApplicationConfig()
	config.LoadTestScene = false
	config.Title = "3D Visualization Developer Panel"
	config.Width = 1400
	config.Height = 920

	app, err := vis.NewApplication(config)
	if err != nil {
		panic(err)
	}

	setupGUI(app, config)
	app.Run()
}

func setupGUI(app *vis.Application, appConfig vis.ApplicationConfig) {
	ui := newDevPanelUI(app, appConfig)
	app.SetUpdateFunction(ui.update)
}

const (
	tabRendererID   = "renderer"
	tabNavigationID = "navigation"

	margin          = float32(20)
	tabHeight       = float32(72)
	sectionSpacing  = float32(12)
	scenarioSpacing = float32(15)
)

type devPanelUI struct {
	app      *vis.Application
	config   vis.ApplicationConfig
	scene    vis.Scene
	gui      *gui.Manager
	renderer vis.Renderer
	camera   vis.Camera

	infoPanel      *gui.InfoPanel
	infoPanelPanel gui.Panel

	scenarioPanel   *gui.ScenarioPanel
	scenarioPanelUI gui.Panel

	rendererPanel   *gui.RendererConfigPanel
	rendererPanelUI gui.Panel

	navigationPanel   *gui.NavigationPanel
	navigationPanelUI gui.Panel

	tabPanel            gui.Panel
	rendererTabButton   gui.Button
	navigationTabButton gui.Button
	activeTab           string

	scenarios       []scenarioEntry
	currentScenario int
	autoUpdate      func(delta float64)
	motionTime      float64

	lastScreenWidth  int
	lastScreenHeight int
	rightPanelWidth  float32
}

func newDevPanelUI(app *vis.Application, appConfig vis.ApplicationConfig) *devPanelUI {
	renderer := app.GetRenderer()

	ui := &devPanelUI{
		app:      app,
		config:   appConfig,
		scene:    vis.NewScene(),
		gui:      app.GetGUI(),
		renderer: renderer,
		camera:   renderer.GetCamera(),
	}

	app.AddScene(ui.scene)

	ui.infoPanel = gui.NewInfoPanel(gui.InfoPanelConfig{
		X: margin,
		Y: margin,
	})
	ui.infoPanelPanel = ui.infoPanel.GetPanel()
	ui.gui.AddElement(ui.infoPanelPanel)

	ui.scenarios = ui.buildScenarios()

	ui.scenarioPanel = gui.NewScenarioPanel(gui.ScenarioPanelConfig{
		X: margin,
		Y: margin + 160,
	}, ui.extractScenarioData(), ui.onScenarioSelected)
	ui.scenarioPanelUI = ui.scenarioPanel.Panel()
	ui.gui.AddElement(ui.scenarioPanelUI)

	ui.rendererPanel = gui.NewRendererConfigPanel(gui.RendererConfigPanelConfig{
		X: margin,
		Y: margin + tabHeight + sectionSpacing,
	}, ui.toRendererConfigData(), ui.onRendererConfigChanged)
	ui.rendererPanelUI = ui.rendererPanel.Panel()

	ui.navigationPanel = gui.NewNavigationPanel(gui.NavigationPanelConfig{
		X: margin,
		Y: margin + tabHeight + sectionSpacing,
	}, gui.NavigationCallbacks{})
	ui.navigationPanelUI = ui.navigationPanel.GetPanel()

	tabConfig := gui.DefaultPanelConfig()
	tabConfig.X = margin
	tabConfig.Y = margin
	tabConfig.Width = 320
	tabConfig.Height = tabHeight
	tabConfig.ShowBorder = false
	tabConfig.BackgroundColor = rl.NewColor(35, 35, 35, 220)
	ui.tabPanel = gui.NewPanel(tabConfig)
	ui.gui.AddElement(ui.tabPanel)
	ui.rightPanelWidth = tabConfig.Width

	ui.rendererTabButton = gui.NewButton(gui.ButtonConfig{
		X:           tabConfig.X + 10,
		Y:           tabConfig.Y + 16,
		Width:       140,
		Height:      36,
		Text:        "Renderer",
		NormalColor: rl.NewColor(55, 55, 55, 255),
		HoverColor:  rl.NewColor(80, 80, 80, 255),
		TextColor:   rl.White,
		FontSize:    16,
	})
	ui.navigationTabButton = gui.NewButton(gui.ButtonConfig{
		X:           tabConfig.X + 160,
		Y:           tabConfig.Y + 16,
		Width:       140,
		Height:      36,
		Text:        "Navigation",
		NormalColor: rl.NewColor(55, 55, 55, 255),
		HoverColor:  rl.NewColor(80, 80, 80, 255),
		TextColor:   rl.White,
		FontSize:    16,
	})

	ui.tabPanel.AddElement(ui.rendererTabButton)
	ui.tabPanel.AddElement(ui.navigationTabButton)

	ui.activeTab = ""
	ui.lastScreenWidth = rl.GetScreenWidth()
	ui.lastScreenHeight = rl.GetScreenHeight()

	ui.activateTab(tabRendererID)
	ui.layout(ui.lastScreenWidth, ui.lastScreenHeight)

	return ui
}

func (ui *devPanelUI) update(deltaTime time.Duration) {
	delta := deltaTime.Seconds()
	ui.motionTime += delta

	width := rl.GetScreenWidth()
	height := rl.GetScreenHeight()
	if width != ui.lastScreenWidth || height != ui.lastScreenHeight {
		ui.layout(width, height)
	}

	if ui.rendererTabButton.IsClicked() {
		ui.activateTab(tabRendererID)
	}
	if ui.navigationTabButton.IsClicked() {
		ui.activateTab(tabNavigationID)
	}

	ui.infoPanel.SetFPS(rl.GetFPS())
	ui.infoPanel.SetCameraInfo(
		ui.camera.GetPolarAngle(),
		ui.camera.GetAzimuth(),
		ui.camera.GetDistanceToScreen(),
	)
	ui.infoPanel.SetSceneCount(len(ui.app.GetScenes()))

	ui.scenarioPanel.Update()

	switch ui.activeTab {
	case tabRendererID:
		ui.rendererPanel.Update()
	case tabNavigationID:
		ui.navigationPanel.HandleInput(delta, gui.NavigationCallbacks{
			OnReset: func() {
				ui.setScenario(ui.currentScenario)
			},
			OnRotateHorizontal: func(amount float64) {
				ui.camera.RotatePolar(amount)
			},
			OnRotateVertical: func(amount float64) {
				ui.camera.RotateAzimuth(amount)
			},
			OnZoom: func(amount float64) {
				ui.camera.ScaleLinear(-amount)
			},
		})
	}

	if ui.autoUpdate != nil {
		ui.autoUpdate(delta)
	}

	ui.handleKeyboard(delta)
}

func (ui *devPanelUI) layout(width, height int) {
	ui.lastScreenWidth = width
	ui.lastScreenHeight = height

	rightX := float32(width) - ui.rightPanelWidth - margin

	ui.infoPanelPanel.SetPosition(margin, margin)
	infoBounds := ui.infoPanelPanel.GetBounds()

	scenarioY := margin + infoBounds.Height + scenarioSpacing
	ui.scenarioPanelUI.SetPosition(margin, scenarioY)

	ui.tabPanel.SetPosition(rightX, margin)

	tabBounds := ui.tabPanel.GetBounds()
	contentY := margin + tabBounds.Height + sectionSpacing
	ui.rendererPanelUI.SetPosition(rightX, contentY)
	ui.navigationPanelUI.SetPosition(rightX, contentY)
}

func (ui *devPanelUI) activateTab(id string) {
	if id == ui.activeTab {
		return
	}

	switch ui.activeTab {
	case tabRendererID:
		ui.gui.RemoveElement(ui.rendererPanelUI)
	case tabNavigationID:
		ui.gui.RemoveElement(ui.navigationPanelUI)
	}

	switch id {
	case tabRendererID:
		ui.gui.AddElement(ui.rendererPanelUI)
	case tabNavigationID:
		ui.gui.AddElement(ui.navigationPanelUI)
	default:
		return
	}

	ui.activeTab = id
	ui.updateTabStyles()
	if ui.lastScreenWidth > 0 && ui.lastScreenHeight > 0 {
		ui.layout(ui.lastScreenWidth, ui.lastScreenHeight)
	}
}

func (ui *devPanelUI) updateTabStyles() {
	activeNormal := rl.NewColor(30, 144, 255, 255)
	activeHover := rl.NewColor(60, 164, 255, 255)
	inactiveNormal := rl.NewColor(55, 55, 55, 255)
	inactiveHover := rl.NewColor(80, 80, 80, 255)

	if ui.activeTab == tabRendererID {
		ui.rendererTabButton.SetColors(activeNormal, activeHover)
		ui.navigationTabButton.SetColors(inactiveNormal, inactiveHover)
	} else {
		ui.rendererTabButton.SetColors(inactiveNormal, inactiveHover)
		ui.navigationTabButton.SetColors(activeNormal, activeHover)
	}
}

func (ui *devPanelUI) onScenarioSelected(index int) {
	ui.setScenario(index)
}

func (ui *devPanelUI) onRendererConfigChanged(data gui.RendererConfigData) {
	ui.applyRendererConfig(data)
}

func (ui *devPanelUI) buildScenarios() []scenarioEntry {
	return []scenarioEntry{
		{
			data: gui.Scenario{
				Name:        "Static Cube",
				Description: "Baseline cube without automation; perfect for manual navigation tweaks.",
			},
			setup: func() {
				ui.resetCamera()
				ui.setSceneMesh(geom.CreateCube(220))
				ui.autoUpdate = nil
			},
		},
		{
			data: gui.Scenario{
				Name:        "Static Tetrahedron",
				Description: "Simple tetrahedron for checking wireframe rendering and face visibility.",
			},
			setup: func() {
				ui.resetCamera()
				ui.setSceneMesh(geom.CreateTetrahedron(260))
				ui.autoUpdate = nil
			},
		},
		{
			data: gui.Scenario{
				Name:        "Spin Cube",
				Description: "Constant horizontal spin to validate rotation smoothing.",
			},
			setup: func() {
				ui.resetCamera()
				ui.setSceneMesh(geom.CreateCube(220))
				ui.autoUpdate = func(delta float64) {
					ui.camera.RotatePolar(delta * 0.8)
				}
			},
		},
		{
			data: gui.Scenario{
				Name:        "Orbit Tetrahedron",
				Description: "Camera orbits to test backface culling and orientation.",
			},
			setup: func() {
				ui.resetCamera()
				ui.setSceneMesh(geom.CreateTetrahedron(240))
				ui.autoUpdate = func(delta float64) {
					ui.camera.RotatePolar(delta * 0.6)
					ui.camera.RotateAzimuth(delta * 0.45)
				}
			},
		},
		{
			data: gui.Scenario{
				Name:        "Breathing Cube",
				Description: "Sinusoid zoom in/out for smooth zoom evaluation.",
			},
			setup: func() {
				ui.resetCamera()
				ui.setSceneMesh(geom.CreateCube(200))
				ui.autoUpdate = func(delta float64) {
					zoom := math.Sin(ui.motionTime) * delta * 280
					ui.camera.ScaleLinear(zoom)
				}
			},
		},
	}
}

func (ui *devPanelUI) extractScenarioData() []gui.Scenario {
	data := make([]gui.Scenario, len(ui.scenarios))
	for i, entry := range ui.scenarios {
		data[i] = entry.data
	}
	return data
}

func (ui *devPanelUI) setScenario(index int) {
	if index < 0 || index >= len(ui.scenarios) {
		return
	}

	ui.currentScenario = index
	ui.motionTime = 0
	ui.autoUpdate = nil
	ui.scenarios[index].setup()
	ui.infoPanel.SetActiveScenario(ui.scenarios[index].data.Name)
}

func (ui *devPanelUI) resetCamera() {
	newCam, err := vis.NewCamera(ui.config.Camera)
	if err != nil {
		return
	}
	ui.renderer.SetCamera(newCam)
	ui.camera = newCam
}

func (ui *devPanelUI) setSceneMesh(mesh *geom.Mesh) {
	ui.scene.Clear()
	if mesh != nil {
		ui.scene.AddMesh(mesh)
	}
}

func (ui *devPanelUI) toRendererConfigData() gui.RendererConfigData {
	config := ui.app.GetRendererConfig()
	faceColor := config.FaceColor
	if config.AlphaValue != 0 {
		faceColor.A = config.AlphaValue
	}
	return gui.RendererConfigData{
		BackgroundColor:    config.BackgroundColor,
		UseRandomFaceColor: config.UseRandomFaceColor,
		FaceColor:          faceColor,
		EdgeColor:          config.EdgeColor,
		AlphaValue:         config.AlphaValue,
		DrawFaces:          config.DrawFaces,
		DrawEdges:          config.DrawEdges,
		UseBackfaceCulling: config.UseBackfaceCulling,
	}
}

func (ui *devPanelUI) applyRendererConfig(data gui.RendererConfigData) {
	config := ui.app.GetRendererConfig()
	config.BackgroundColor = data.BackgroundColor
	config.UseRandomFaceColor = data.UseRandomFaceColor
	config.FaceColor = data.FaceColor
	config.FaceColor.A = data.AlphaValue
	config.EdgeColor = data.EdgeColor
	config.AlphaValue = data.AlphaValue
	config.DrawFaces = data.DrawFaces
	config.DrawEdges = data.DrawEdges
	config.UseBackfaceCulling = data.UseBackfaceCulling

	ui.app.SetRendererConfig(config)
}

func (ui *devPanelUI) handleKeyboard(delta float64) {
	const rotateSpeed = 1.2
	const zoomSpeed = 420.0

	if rl.IsKeyDown(rl.KeyLeft) || rl.IsKeyDown(rl.KeyA) {
		ui.camera.RotatePolar(-delta * rotateSpeed)
	}
	if rl.IsKeyDown(rl.KeyRight) || rl.IsKeyDown(rl.KeyD) {
		ui.camera.RotatePolar(delta * rotateSpeed)
	}
	if rl.IsKeyDown(rl.KeyUp) || rl.IsKeyDown(rl.KeyW) {
		ui.camera.RotateAzimuth(-delta * rotateSpeed)
	}
	if rl.IsKeyDown(rl.KeyDown) || rl.IsKeyDown(rl.KeyS) {
		ui.camera.RotateAzimuth(delta * rotateSpeed)
	}
	if rl.IsKeyDown(rl.KeyQ) || rl.IsKeyDown(rl.KeyPageUp) {
		ui.camera.ScaleLinear(-delta * zoomSpeed)
	}
	if rl.IsKeyDown(rl.KeyE) || rl.IsKeyDown(rl.KeyPageDown) {
		ui.camera.ScaleLinear(delta * zoomSpeed)
	}
	if rl.IsKeyDown(rl.KeyR) {
		ui.camera.ScaleLinear(-delta * zoomSpeed * 0.5)
		ui.camera.RotatePolar(-delta * rotateSpeed * 0.5)
	}
}
