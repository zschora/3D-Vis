package gui

// ControlPanel provides camera control buttons
type ControlPanel struct {
	panel         Panel
	resetButton   Button
	rotateButton  Button
	zoomInButton  Button
	zoomOutButton Button
}

// ControlPanelConfig holds configuration for creating a control panel
type ControlPanelConfig struct {
	X, Y float32
}

// ControlCallbacks holds callback functions for control panel actions
type ControlCallbacks struct {
	OnReset   func()
	OnRotate  func()
	OnZoomIn  func()
	OnZoomOut func()
}

// NewControlPanel creates a new control panel
func NewControlPanel(config ControlPanelConfig, callbacks ControlCallbacks) *ControlPanel {
	panelConfig := DefaultPanelConfig()
	panelConfig.X = config.X
	panelConfig.Y = config.Y
	panelConfig.Width = 200
	panelConfig.Height = 180

	panel := NewPanel(panelConfig).(*panel)

	// Create buttons
	resetButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 10,
		Width:  180,
		Height: 30,
		Text:   "Reset Camera",
	})

	rotateButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 50,
		Width:  180,
		Height: 30,
		Text:   "Toggle Rotate",
	})

	zoomInButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 90,
		Width:  85,
		Height: 30,
		Text:   "Zoom In",
	})

	zoomOutButton := NewButton(ButtonConfig{
		X:      config.X + 105,
		Y:      config.Y + 90,
		Width:  85,
		Height: 30,
		Text:   "Zoom Out",
	})

	// Add buttons to panel
	panel.AddElement(resetButton)
	panel.AddElement(rotateButton)
	panel.AddElement(zoomInButton)
	panel.AddElement(zoomOutButton)

	return &ControlPanel{
		panel:         panel,
		resetButton:   resetButton,
		rotateButton:  rotateButton,
		zoomInButton:  zoomInButton,
		zoomOutButton: zoomOutButton,
	}
}

// Update updates the control panel
func (cp *ControlPanel) Update() {
	cp.panel.Update()
}

// Draw renders the control panel
func (cp *ControlPanel) Draw() {
	cp.panel.Draw()
}

// HandleInput handles button clicks (should be called in update loop)
func (cp *ControlPanel) HandleInput(callbacks ControlCallbacks) {
	if cp.resetButton.IsClicked() && callbacks.OnReset != nil {
		callbacks.OnReset()
	}
	if cp.rotateButton.IsClicked() && callbacks.OnRotate != nil {
		callbacks.OnRotate()
	}
	if cp.zoomInButton.IsClicked() && callbacks.OnZoomIn != nil {
		callbacks.OnZoomIn()
	}
	if cp.zoomOutButton.IsClicked() && callbacks.OnZoomOut != nil {
		callbacks.OnZoomOut()
	}
}

// GetPanel returns the underlying panel
func (cp *ControlPanel) GetPanel() Panel {
	return cp.panel
}
