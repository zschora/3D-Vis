package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// panel is the default implementation of Panel interface
type panel struct {
	bounds          rl.Rectangle
	elements        []UIElement
	backgroundColor rl.Color
	borderColor     rl.Color
	showBorder      bool
}

// PanelConfig holds configuration for creating a panel
type PanelConfig struct {
	X, Y, Width, Height float32
	BackgroundColor     rl.Color
	BorderColor         rl.Color
	ShowBorder          bool
}

// DefaultPanelConfig returns default panel configuration
func DefaultPanelConfig() PanelConfig {
	return PanelConfig{
		Width:           300,
		Height:          200,
		BackgroundColor: rl.NewColor(40, 40, 40, 200), // Semi-transparent dark gray
		BorderColor:     rl.White,
		ShowBorder:      true,
	}
}

// NewPanel creates a new panel with the given configuration
func NewPanel(config PanelConfig) Panel {
	return &panel{
		bounds:          rl.NewRectangle(config.X, config.Y, config.Width, config.Height),
		elements:        make([]UIElement, 0),
		backgroundColor: config.BackgroundColor,
		borderColor:     config.BorderColor,
		showBorder:      config.ShowBorder,
	}
}

// Update updates all elements in the panel
func (p *panel) Update() bool {
	interacted := false
	for _, element := range p.elements {
		if element.Update() {
			interacted = true
		}
	}
	return interacted
}

// Draw renders the panel and all its elements
func (p *panel) Draw() {
	// Draw background
	rl.DrawRectangleRec(p.bounds, p.backgroundColor)

	// Draw border
	if p.showBorder {
		rl.DrawRectangleLinesEx(p.bounds, 2, p.borderColor)
	}

	// Draw all elements
	for _, element := range p.elements {
		element.Draw()
	}
}

// GetBounds returns the panel's bounding rectangle
func (p *panel) GetBounds() rl.Rectangle {
	return p.bounds
}

// SetPosition sets the panel's position
func (p *panel) SetPosition(x, y float32) {
	// Calculate offset
	offsetX := x - p.bounds.X
	offsetY := y - p.bounds.Y

	// Update panel position
	p.bounds.X = x
	p.bounds.Y = y

	// Update all elements relative to panel
	for _, element := range p.elements {
		elemBounds := element.GetBounds()
		element.SetPosition(elemBounds.X+offsetX, elemBounds.Y+offsetY)
	}
}

// AddElement adds a UI element to the panel
func (p *panel) AddElement(element UIElement) {
	if element != nil {
		p.elements = append(p.elements, element)
	}
}

// GetElements returns all elements in the panel
func (p *panel) GetElements() []UIElement {
	result := make([]UIElement, len(p.elements))
	copy(result, p.elements)
	return result
}
