package gui

import rl "github.com/gen2brain/raylib-go/raylib"

// NavigationPanel provides basic navigation controls for 3D viewer
type NavigationPanel struct {
	panel         Panel
	resetButton   Button
	zoomInButton  Button
	zoomOutButton Button
}

// NavigationPanelConfig holds configuration for creating a navigation panel
type NavigationPanelConfig struct {
	X, Y float32
}

// NavigationCallbacks holds callback functions for navigation panel actions
type NavigationCallbacks struct {
	OnReset   func()
	OnZoomIn  func()
	OnZoomOut func()
}

// NewNavigationPanel creates a new navigation panel for basic viewer controls
func NewNavigationPanel(config NavigationPanelConfig, callbacks NavigationCallbacks) *NavigationPanel {
	panelConfig := DefaultPanelConfig()
	panelConfig.X = config.X
	panelConfig.Y = config.Y
	panelConfig.Width = 200
	panelConfig.Height = 120

	panel := NewPanel(panelConfig).(*panel)

	// Create buttons with better styling
	resetButton := NewButton(ButtonConfig{
		X:           config.X + 10,
		Y:           config.Y + 10,
		Width:       180,
		Height:      30,
		Text:        "Reset View",
		NormalColor: rl.NewColor(60, 60, 60, 255), // Dark gray
		HoverColor:  rl.NewColor(80, 80, 80, 255), // Lighter gray
		TextColor:   rl.White,
		FontSize:    14,
	})

	zoomInButton := NewButton(ButtonConfig{
		X:           config.X + 10,
		Y:           config.Y + 50,
		Width:       85,
		Height:      30,
		Text:        "Zoom In",
		NormalColor: rl.NewColor(60, 60, 60, 255),
		HoverColor:  rl.NewColor(80, 80, 80, 255),
		TextColor:   rl.White,
		FontSize:    14,
	})

	zoomOutButton := NewButton(ButtonConfig{
		X:           config.X + 105,
		Y:           config.Y + 50,
		Width:       85,
		Height:      30,
		Text:        "Zoom Out",
		NormalColor: rl.NewColor(60, 60, 60, 255),
		HoverColor:  rl.NewColor(80, 80, 80, 255),
		TextColor:   rl.White,
		FontSize:    14,
	})

	// Add buttons to panel
	panel.AddElement(resetButton)
	panel.AddElement(zoomInButton)
	panel.AddElement(zoomOutButton)

	return &NavigationPanel{
		panel:         panel,
		resetButton:   resetButton,
		zoomInButton:  zoomInButton,
		zoomOutButton: zoomOutButton,
	}
}

// Update updates the navigation panel
func (np *NavigationPanel) Update() {
	np.panel.Update()
}

// Draw renders the navigation panel
func (np *NavigationPanel) Draw() {
	np.panel.Draw()
}

// HandleInput handles button clicks (should be called in update loop)
func (np *NavigationPanel) HandleInput(callbacks NavigationCallbacks) {
	if np.resetButton.IsClicked() && callbacks.OnReset != nil {
		callbacks.OnReset()
	}
	if np.zoomInButton.IsClicked() && callbacks.OnZoomIn != nil {
		callbacks.OnZoomIn()
	}
	if np.zoomOutButton.IsClicked() && callbacks.OnZoomOut != nil {
		callbacks.OnZoomOut()
	}
}

// GetPanel returns the underlying panel
func (np *NavigationPanel) GetPanel() Panel {
	return np.panel
}
