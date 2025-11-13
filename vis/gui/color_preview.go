package gui

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	defaultPreviewWidth  = 32
	defaultPreviewHeight = 20
)

// ColorPreview renders a small colored rectangle.
type ColorPreview struct {
	bounds rl.Rectangle
	color  rl.Color
}

// ColorPreviewConfig configures the color preview element.
type ColorPreviewConfig struct {
	X, Y          float32
	Width, Height float32
	Color         rl.Color
}

// NewColorPreview creates a color preview element.
func NewColorPreview(config ColorPreviewConfig) *ColorPreview {
	width := config.Width
	if width == 0 {
		width = defaultPreviewWidth
	}

	height := config.Height
	if height == 0 {
		height = defaultPreviewHeight
	}

	return &ColorPreview{
		bounds: rl.NewRectangle(config.X, config.Y, width, height),
		color:  config.Color,
	}
}

// Update is a no-op for the preview (required by interface).
func (p *ColorPreview) Update() bool {
	return false
}

// Draw renders the color preview.
func (p *ColorPreview) Draw() {
	rl.DrawRectangleRec(p.bounds, p.color)
	rl.DrawRectangleLinesEx(p.bounds, 2, rl.NewColor(25, 25, 25, 255))
}

// GetBounds returns the preview bounds.
func (p *ColorPreview) GetBounds() rl.Rectangle {
	return p.bounds
}

// SetPosition moves the preview to a new location.
func (p *ColorPreview) SetPosition(x, y float32) {
	p.bounds.X = x
	p.bounds.Y = y
}

// SetColor updates the preview color.
func (p *ColorPreview) SetColor(color rl.Color) {
	p.color = color
}
