package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// label is the default implementation of Label interface
type label struct {
	bounds   rl.Rectangle
	text     string
	color    rl.Color
	fontSize int32
}

// LabelConfig holds configuration for creating a label
type LabelConfig struct {
	X, Y     float32
	Text     string
	Color    rl.Color
	FontSize int32
}

// DefaultLabelConfig returns default label configuration
func DefaultLabelConfig() LabelConfig {
	return LabelConfig{
		Color:    rl.White,
		FontSize: 16,
	}
}

// NewLabel creates a new label with the given configuration
func NewLabel(config LabelConfig) Label {
	textWidth := float32(rl.MeasureText(config.Text, config.FontSize))
	return &label{
		bounds:   rl.NewRectangle(config.X, config.Y, textWidth, float32(config.FontSize)),
		text:     config.Text,
		color:    config.Color,
		fontSize: config.FontSize,
	}
}

// Update updates the label state (labels don't need updates, but implement interface)
func (l *label) Update() bool {
	return false
}

// Draw renders the label
func (l *label) Draw() {
	rl.DrawText(l.text, int32(l.bounds.X), int32(l.bounds.Y), l.fontSize, l.color)
}

// GetBounds returns the label's bounding rectangle
func (l *label) GetBounds() rl.Rectangle {
	return l.bounds
}

// SetPosition sets the label's position
func (l *label) SetPosition(x, y float32) {
	l.bounds.X = x
	l.bounds.Y = y
}

// SetText sets the label text
func (l *label) SetText(text string) {
	l.text = text
	// Update bounds based on new text
	l.bounds.Width = float32(rl.MeasureText(l.text, l.fontSize))
}

// GetText returns the label text
func (l *label) GetText() string {
	return l.text
}
