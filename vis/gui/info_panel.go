package gui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// InfoPanel displays application information (FPS, camera info, etc.)
type InfoPanel struct {
	panel         Panel
	fpsLabel      Label
	cameraLabel   Label
	sceneLabel    Label
	scenarioLabel Label
}

// InfoPanelConfig holds configuration for creating an info panel
type InfoPanelConfig struct {
	X, Y float32
}

// NewInfoPanel creates a new info panel
func NewInfoPanel(config InfoPanelConfig) *InfoPanel {
	panelConfig := DefaultPanelConfig()
	panelConfig.X = config.X
	panelConfig.Y = config.Y
	panelConfig.Width = 280
	panelConfig.Height = 190

	panel := NewPanel(panelConfig).(*panel)

	fpsLabel := NewLabel(LabelConfig{
		X:        config.X + 10,
		Y:        config.Y + 10,
		Text:     "FPS: 0",
		FontSize: 14,
		Color:    rl.Green,
	})

	cameraLabel := NewLabel(LabelConfig{
		X:        config.X + 10,
		Y:        config.Y + 35,
		Text:     "Camera: N/A",
		FontSize: 12,
		Color:    rl.White,
	})

	sceneLabel := NewLabel(LabelConfig{
		X:        config.X + 10,
		Y:        config.Y + 55,
		Text:     "Scenes: 0",
		FontSize: 12,
		Color:    rl.White,
	})

	scenarioLabel := NewLabel(LabelConfig{
		X:        config.X + 10,
		Y:        config.Y + 80,
		Text:     "Scenario: n/a",
		FontSize: 12,
		Color:    rl.LightGray,
	})

	panel.AddElement(fpsLabel)
	panel.AddElement(cameraLabel)
	panel.AddElement(sceneLabel)
	panel.AddElement(scenarioLabel)

	return &InfoPanel{
		panel:         panel,
		fpsLabel:      fpsLabel,
		cameraLabel:   cameraLabel,
		sceneLabel:    sceneLabel,
		scenarioLabel: scenarioLabel,
	}
}

// Update updates the info panel
func (ip *InfoPanel) Update() {
	ip.panel.Update()
}

// Draw renders the info panel
func (ip *InfoPanel) Draw() {
	ip.panel.Draw()
}

// SetFPS updates the FPS display
func (ip *InfoPanel) SetFPS(fps int32) {
	ip.fpsLabel.SetText(fmt.Sprintf("FPS: %d", fps))
}

// SetCameraInfo updates the camera information display
func (ip *InfoPanel) SetCameraInfo(polarAngle, azimuth, distance float64) {
	text := fmt.Sprintf("Camera: θ=%.1f° φ=%.1f° d=%.0f",
		polarAngle*180/3.14159, azimuth*180/3.14159, distance)
	if clipper, ok := ip.cameraLabel.(interface{ SetTextClipped(string, float32) }); ok {
		clipper.SetTextClipped(text, 260)
	} else {
		ip.cameraLabel.SetText(text)
	}
}

// SetSceneCount updates the scene count display
func (ip *InfoPanel) SetSceneCount(count int) {
	ip.sceneLabel.SetText(fmt.Sprintf("Scenes: %d", count))
}

// SetActiveScenario updates the scenario display
func (ip *InfoPanel) SetActiveScenario(name string) {
	text := fmt.Sprintf("Scenario: %s", name)
	if clipper, ok := ip.scenarioLabel.(interface{ SetTextClipped(string, float32) }); ok {
		clipper.SetTextClipped(text, 260)
	} else {
		ip.scenarioLabel.SetText(text)
	}
}

// GetPanel returns the underlying panel
func (ip *InfoPanel) GetPanel() Panel {
	return ip.panel
}
