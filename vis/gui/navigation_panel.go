package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// NavigationPanel provides basic navigation controls for 3D viewer
type NavigationPanel struct {
	panel          Panel
	title          Label
	resetButton    Button
	leftButton     Button
	rightButton    Button
	upButton       Button
	downButton     Button
	zoomInButton   Button
	zoomOutButton  Button
	rotationSlider *Slider
	zoomSlider     *Slider
}

// NavigationPanelConfig holds configuration for creating a navigation panel
type NavigationPanelConfig struct {
	X, Y float32
}

// NavigationCallbacks holds callback functions for navigation panel actions
type NavigationCallbacks struct {
	OnReset            func()
	OnRotateHorizontal func(amount float64)
	OnRotateVertical   func(amount float64)
	OnZoom             func(amount float64)
}

// NewNavigationPanel creates a new navigation panel for basic viewer controls
func NewNavigationPanel(config NavigationPanelConfig, callbacks NavigationCallbacks) *NavigationPanel {
	panelConfig := DefaultPanelConfig()
	panelConfig.X = config.X
	panelConfig.Y = config.Y
	panelConfig.Width = 320
	panelConfig.Height = 360

	panel := NewPanel(panelConfig).(*panel)

	title := NewLabel(LabelConfig{
		X:        config.X + 10,
		Y:        config.Y + 10,
		Text:     "Navigation",
		Color:    rl.White,
		FontSize: 18,
	})

	resetButton := NewButton(ButtonConfig{
		X:           config.X + 10,
		Y:           config.Y + 36,
		Width:       240,
		Height:      32,
		Text:        "Reset View",
		NormalColor: rl.NewColor(55, 55, 55, 255),
		HoverColor:  rl.NewColor(80, 80, 80, 255),
		TextColor:   rl.White,
		FontSize:    16,
	})

	// Rotation buttons arranged in a cross layout
	gridX := config.X + 10
	gridY := config.Y + 80
	buttonSize := float32(56)
	gap := float32(8)

	upButton := NewButton(ButtonConfig{
		X:           gridX + buttonSize + gap,
		Y:           gridY,
		Width:       buttonSize,
		Height:      buttonSize,
		Text:        "^",
		NormalColor: rl.NewColor(65, 65, 65, 255),
		HoverColor:  rl.NewColor(95, 95, 95, 255),
		TextColor:   rl.White,
		FontSize:    22,
	})

	leftButton := NewButton(ButtonConfig{
		X:           gridX,
		Y:           gridY + buttonSize + gap,
		Width:       buttonSize,
		Height:      buttonSize,
		Text:        "<",
		NormalColor: rl.NewColor(65, 65, 65, 255),
		HoverColor:  rl.NewColor(95, 95, 95, 255),
		TextColor:   rl.White,
		FontSize:    22,
	})

	rightButton := NewButton(ButtonConfig{
		X:           gridX + 2*(buttonSize+gap),
		Y:           gridY + buttonSize + gap,
		Width:       buttonSize,
		Height:      buttonSize,
		Text:        ">",
		NormalColor: rl.NewColor(65, 65, 65, 255),
		HoverColor:  rl.NewColor(95, 95, 95, 255),
		TextColor:   rl.White,
		FontSize:    22,
	})

	downButton := NewButton(ButtonConfig{
		X:           gridX + buttonSize + gap,
		Y:           gridY + 2*(buttonSize+gap),
		Width:       buttonSize,
		Height:      buttonSize,
		Text:        "v",
		NormalColor: rl.NewColor(65, 65, 65, 255),
		HoverColor:  rl.NewColor(95, 95, 95, 255),
		TextColor:   rl.White,
		FontSize:    22,
	})

	zoomInButton := NewButton(ButtonConfig{
		X:           gridX + 3*(buttonSize+gap) + gap,
		Y:           gridY,
		Width:       48,
		Height:      buttonSize,
		Text:        "+",
		NormalColor: rl.NewColor(55, 55, 55, 255),
		HoverColor:  rl.NewColor(80, 80, 80, 255),
		TextColor:   rl.White,
		FontSize:    22,
	})

	zoomOutButton := NewButton(ButtonConfig{
		X:           gridX + 3*(buttonSize+gap) + gap,
		Y:           gridY + buttonSize + gap,
		Width:       48,
		Height:      buttonSize,
		Text:        "-",
		NormalColor: rl.NewColor(55, 55, 55, 255),
		HoverColor:  rl.NewColor(80, 80, 80, 255),
		TextColor:   rl.White,
		FontSize:    22,
	})

	rotationSlider := NewSlider(SliderConfig{
		X:         config.X + 10,
		Y:         gridY + 3*(buttonSize+gap) + 16,
		Width:     240,
		Label:     "Rotate speed (rad/s)",
		Min:       0.1,
		Max:       2.5,
		Value:     0.8,
		Precision: 2,
	})

	zoomSlider := NewSlider(SliderConfig{
		X:         config.X + 10,
		Y:         rotationSlider.GetBounds().Y + rotationSlider.GetBounds().Height + 24,
		Width:     240,
		Label:     "Zoom speed",
		Min:       30,
		Max:       800,
		Value:     300,
		Precision: 0,
	})

	panel.AddElement(title)
	panel.AddElement(resetButton)
	panel.AddElement(upButton)
	panel.AddElement(leftButton)
	panel.AddElement(rightButton)
	panel.AddElement(downButton)
	panel.AddElement(zoomInButton)
	panel.AddElement(zoomOutButton)
	panel.AddElement(rotationSlider)
	panel.AddElement(zoomSlider)

	return &NavigationPanel{
		panel:          panel,
		title:          title,
		resetButton:    resetButton,
		leftButton:     leftButton,
		rightButton:    rightButton,
		upButton:       upButton,
		downButton:     downButton,
		zoomInButton:   zoomInButton,
		zoomOutButton:  zoomOutButton,
		rotationSlider: rotationSlider,
		zoomSlider:     zoomSlider,
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

// HandleInput handles button interactions (should be called in update loop)
func (np *NavigationPanel) HandleInput(deltaSeconds float64, callbacks NavigationCallbacks) {
	if np.resetButton.IsClicked() && callbacks.OnReset != nil {
		callbacks.OnReset()
	}

	rotationSpeed := np.rotationSlider.Value()
	zoomSpeed := np.zoomSlider.Value()

	if callbacks.OnRotateHorizontal != nil {
		if np.leftButton.IsHeld() {
			callbacks.OnRotateHorizontal(-rotationSpeed * deltaSeconds)
		}
		if np.rightButton.IsHeld() {
			callbacks.OnRotateHorizontal(rotationSpeed * deltaSeconds)
		}
	}

	if callbacks.OnRotateVertical != nil {
		if np.upButton.IsHeld() {
			callbacks.OnRotateVertical(-rotationSpeed * deltaSeconds)
		}
		if np.downButton.IsHeld() {
			callbacks.OnRotateVertical(rotationSpeed * deltaSeconds)
		}
	}

	if callbacks.OnZoom != nil {
		if np.zoomInButton.IsHeld() {
			callbacks.OnZoom(zoomSpeed * deltaSeconds)
		}
		if np.zoomOutButton.IsHeld() {
			callbacks.OnZoom(-zoomSpeed * deltaSeconds)
		}
	}
}

// GetPanel returns the underlying panel
func (np *NavigationPanel) GetPanel() Panel {
	return np.panel
}
