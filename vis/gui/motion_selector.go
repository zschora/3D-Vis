package gui

// MotionType represents the type of camera motion
type MotionType int

const (
	MotionNone MotionType = iota
	MotionRotate
	MotionZoom
	MotionRotateAndZoom
	MotionOrbit
)

// String returns the string representation of the motion type
func (mt MotionType) String() string {
	switch mt {
	case MotionNone:
		return "None"
	case MotionRotate:
		return "Rotate"
	case MotionZoom:
		return "Zoom"
	case MotionRotateAndZoom:
		return "Rotate+Zoom"
	case MotionOrbit:
		return "Orbit"
	default:
		return "Unknown"
	}
}

// MotionSelector provides buttons to select different motion types
type MotionSelector struct {
	panel            Panel
	noneButton       Button
	rotateButton     Button
	zoomButton       Button
	rotateZoomButton Button
	orbitButton      Button
	selectedMotion   MotionType
	onChange         func(MotionType)
}

// MotionSelectorConfig holds configuration for creating a motion selector
type MotionSelectorConfig struct {
	X, Y float32
}

// NewMotionSelector creates a new motion selector
func NewMotionSelector(config MotionSelectorConfig, onChange func(MotionType)) *MotionSelector {
	panelConfig := DefaultPanelConfig()
	panelConfig.X = config.X
	panelConfig.Y = config.Y
	panelConfig.Width = 200
	panelConfig.Height = 200

	panel := NewPanel(panelConfig).(*panel)

	// Create buttons
	noneButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 10,
		Width:  180,
		Height: 30,
		Text:   "None",
	})

	rotateButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 50,
		Width:  180,
		Height: 30,
		Text:   "Rotate",
	})

	zoomButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 90,
		Width:  180,
		Height: 30,
		Text:   "Zoom",
	})

	rotateZoomButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 130,
		Width:  180,
		Height: 30,
		Text:   "Rotate+Zoom",
	})

	orbitButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 170,
		Width:  180,
		Height: 30,
		Text:   "Orbit",
	})

	// Add buttons to panel
	panel.AddElement(noneButton)
	panel.AddElement(rotateButton)
	panel.AddElement(zoomButton)
	panel.AddElement(rotateZoomButton)
	panel.AddElement(orbitButton)

	return &MotionSelector{
		panel:            panel,
		noneButton:       noneButton,
		rotateButton:     rotateButton,
		zoomButton:       zoomButton,
		rotateZoomButton: rotateZoomButton,
		orbitButton:      orbitButton,
		selectedMotion:   MotionRotate,
		onChange:         onChange,
	}
}

// Update updates the motion selector
func (ms *MotionSelector) Update() {
	ms.panel.Update()
}

// Draw renders the motion selector
func (ms *MotionSelector) Draw() {
	ms.panel.Draw()
}

// HandleInput handles button clicks
func (ms *MotionSelector) HandleInput() {
	if ms.noneButton.IsClicked() {
		ms.selectedMotion = MotionNone
		if ms.onChange != nil {
			ms.onChange(MotionNone)
		}
	}
	if ms.rotateButton.IsClicked() {
		ms.selectedMotion = MotionRotate
		if ms.onChange != nil {
			ms.onChange(MotionRotate)
		}
	}
	if ms.zoomButton.IsClicked() {
		ms.selectedMotion = MotionZoom
		if ms.onChange != nil {
			ms.onChange(MotionZoom)
		}
	}
	if ms.rotateZoomButton.IsClicked() {
		ms.selectedMotion = MotionRotateAndZoom
		if ms.onChange != nil {
			ms.onChange(MotionRotateAndZoom)
		}
	}
	if ms.orbitButton.IsClicked() {
		ms.selectedMotion = MotionOrbit
		if ms.onChange != nil {
			ms.onChange(MotionOrbit)
		}
	}
}

// GetSelectedMotion returns the currently selected motion type
func (ms *MotionSelector) GetSelectedMotion() MotionType {
	return ms.selectedMotion
}

// GetPanel returns the underlying panel
func (ms *MotionSelector) GetPanel() Panel {
	return ms.panel
}
