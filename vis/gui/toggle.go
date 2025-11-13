package gui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	defaultToggleWidth  = 240
	defaultToggleHeight = 28
	switchWidth         = 50
	switchHeight        = 22
)

// Toggle provides a labeled on/off switch.
type Toggle struct {
	bounds     rl.Rectangle
	label      string
	value      bool
	changed    bool
	fontSize   int32
	labelColor rl.Color
}

// ToggleConfig configures a toggle element.
type ToggleConfig struct {
	X, Y     float32
	Width    float32
	Label    string
	Initial  bool
	FontSize int32
	Color    rl.Color
}

// NewToggle creates a toggle switch with the provided configuration.
func NewToggle(config ToggleConfig) *Toggle {
	width := config.Width
	if width == 0 {
		width = defaultToggleWidth
	}

	fontSize := config.FontSize
	if fontSize == 0 {
		fontSize = 16
	}

	color := config.Color
	if color.A == 0 {
		color = rl.White
	}

	return &Toggle{
		bounds:     rl.NewRectangle(config.X, config.Y, width, defaultToggleHeight),
		label:      config.Label,
		value:      config.Initial,
		fontSize:   fontSize,
		labelColor: color,
	}
}

// Update processes user input for the toggle.
func (t *Toggle) Update() bool {
	t.changed = false

	mousePos := rl.GetMousePosition()
	if rl.CheckCollisionPointRec(mousePos, t.bounds) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		t.value = !t.value
		t.changed = true
	}

	return t.changed
}

// Draw renders the toggle.
func (t *Toggle) Draw() {
	labelY := t.bounds.Y + (t.bounds.Height-float32(t.fontSize))/2
	rl.DrawText(t.label, int32(t.bounds.X), int32(labelY), t.fontSize, t.labelColor)

	switchX := t.bounds.X + t.bounds.Width - switchWidth
	switchY := t.bounds.Y + (t.bounds.Height-switchHeight)/2
	switchRect := rl.NewRectangle(switchX, switchY, switchWidth, switchHeight)

	background := rl.NewColor(70, 70, 70, 255)
	if t.value {
		background = rl.NewColor(30, 144, 255, 255) // Dodger blue
	}

	rl.DrawRectangleRounded(switchRect, 0.5, 8, background)
	rl.DrawRectangleRoundedLinesEx(switchRect, 0.5, 8, 2, rl.NewColor(20, 20, 20, 200))

	knobRadius := float32(switchHeight)/2 - 3
	knobX := switchX + knobRadius + 3
	if t.value {
		knobX = switchX + float32(switchWidth) - knobRadius - 3
	}
	rl.DrawCircle(int32(knobX), int32(switchY+float32(switchHeight)/2), knobRadius, rl.White)
}

// GetBounds returns the toggle bounds.
func (t *Toggle) GetBounds() rl.Rectangle {
	return t.bounds
}

// SetPosition moves the toggle to a new position.
func (t *Toggle) SetPosition(x, y float32) {
	t.bounds.X = x
	t.bounds.Y = y
}

// Value returns the current toggle state.
func (t *Toggle) Value() bool {
	return t.value
}

// SetValue sets the toggle state without triggering change event.
func (t *Toggle) SetValue(value bool) {
	t.value = value
}
