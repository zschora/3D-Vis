package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// button is the default implementation of Button interface
type button struct {
	bounds      rl.Rectangle
	text        string
	clicked     bool
	hovered     bool
	normalColor rl.Color
	hoverColor  rl.Color
	textColor   rl.Color
	fontSize    int32
}

// ButtonConfig holds configuration for creating a button
type ButtonConfig struct {
	X, Y, Width, Height float32
	Text                string
	NormalColor         rl.Color
	HoverColor          rl.Color
	TextColor           rl.Color
	FontSize            int32
}

// DefaultButtonConfig returns default button configuration
func DefaultButtonConfig() ButtonConfig {
	return ButtonConfig{
		Width:       120,
		Height:      40,
		NormalColor: rl.DarkGray,
		HoverColor:  rl.Gray,
		TextColor:   rl.White,
		FontSize:    16,
	}
}

// NewButton creates a new button with the given configuration
func NewButton(config ButtonConfig) Button {
	return &button{
		bounds:      rl.NewRectangle(config.X, config.Y, config.Width, config.Height),
		text:        config.Text,
		normalColor: config.NormalColor,
		hoverColor:  config.HoverColor,
		textColor:   config.TextColor,
		fontSize:    config.FontSize,
	}
}

// Update updates the button state
func (b *button) Update() bool {
	mousePos := rl.GetMousePosition()
	b.hovered = rl.CheckCollisionPointRec(mousePos, b.bounds)
	b.clicked = b.hovered && rl.IsMouseButtonPressed(rl.MouseLeftButton)
	return b.clicked
}

// Draw renders the button
func (b *button) Draw() {
	color := b.normalColor
	borderColor := rl.NewColor(100, 100, 100, 255) // Gray border

	if b.hovered {
		color = b.hoverColor
		borderColor = rl.NewColor(150, 150, 150, 255) // Lighter border on hover
	}

	// Draw button background
	rl.DrawRectangleRec(b.bounds, color)

	// Draw border
	rl.DrawRectangleLinesEx(b.bounds, 2, borderColor)

	// Draw text centered
	if b.text != "" {
		textWidth := rl.MeasureText(b.text, b.fontSize)
		textX := b.bounds.X + (b.bounds.Width-float32(textWidth))/2
		textY := b.bounds.Y + (b.bounds.Height-float32(b.fontSize))/2
		rl.DrawText(b.text, int32(textX), int32(textY), b.fontSize, b.textColor)
	}
}

// GetBounds returns the button's bounding rectangle
func (b *button) GetBounds() rl.Rectangle {
	return b.bounds
}

// SetPosition sets the button's position
func (b *button) SetPosition(x, y float32) {
	b.bounds.X = x
	b.bounds.Y = y
}

// IsClicked returns true if the button was clicked this frame
func (b *button) IsClicked() bool {
	return b.clicked
}

// SetText sets the button text
func (b *button) SetText(text string) {
	b.text = text
}

// GetText returns the button text
func (b *button) GetText() string {
	return b.text
}
