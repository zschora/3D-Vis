package gui

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	defaultSliderWidth  = 240
	defaultSliderHeight = 36
	sliderTrackHeight   = 6
	knobRadius          = 10
)

// Slider provides a horizontal slider for numeric values.
type Slider struct {
	bounds     rl.Rectangle
	label      string
	min        float64
	max        float64
	value      float64
	dragging   bool
	fontSize   int32
	precision  int
	labelColor rl.Color
	valueColor rl.Color
}

// SliderConfig configures a slider element.
type SliderConfig struct {
	X, Y       float32
	Width      float32
	Label      string
	Min        float64
	Max        float64
	Value      float64
	Precision  int
	FontSize   int32
	Color      rl.Color
	ValueColor rl.Color
}

// NewSlider creates a slider with the provided configuration.
func NewSlider(config SliderConfig) *Slider {
	width := config.Width
	if width == 0 {
		width = defaultSliderWidth
	}

	fontSize := config.FontSize
	if fontSize == 0 {
		fontSize = 16
	}

	precision := config.Precision
	if precision < 0 {
		precision = 0
	}

	labelColor := config.Color
	if labelColor.A == 0 {
		labelColor = rl.White
	}

	valueColor := config.ValueColor
	if valueColor.A == 0 {
		valueColor = rl.LightGray
	}

	min := config.Min
	max := config.Max
	if max <= min {
		max = min + 1
	}

	value := clamp(config.Value, min, max)

	return &Slider{
		bounds:     rl.NewRectangle(config.X, config.Y, width, defaultSliderHeight),
		label:      config.Label,
		min:        min,
		max:        max,
		value:      value,
		fontSize:   fontSize,
		precision:  precision,
		labelColor: labelColor,
		valueColor: valueColor,
	}
}

// Update processes input for the slider.
func (s *Slider) Update() bool {
	mousePos := rl.GetMousePosition()
	trackRect := s.trackRect()

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) &&
		(rl.CheckCollisionPointRec(mousePos, trackRect) || s.pointOnKnob(mousePos)) {
		s.dragging = true
	}

	if s.dragging {
		if rl.IsMouseButtonDown(rl.MouseLeftButton) {
			s.updateValueFromMouse(mousePos.X)
		} else {
			s.dragging = false
		}
	}

	return s.dragging
}

// Draw renders the slider.
func (s *Slider) Draw() {
	labelY := s.bounds.Y + 2
	valueText := s.formatValue()
	valueTextWidth := float32(rl.MeasureText(valueText, s.fontSize))

	// Ensure label doesn't overlap with value
	maxLabelWidth := s.bounds.Width - valueTextWidth - 8
	labelText := s.label
	labelWidth := float32(rl.MeasureText(labelText, s.fontSize))
	if labelWidth > maxLabelWidth {
		// Clip label if too long
		for len(labelText) > 0 {
			labelText = labelText[:len(labelText)-1]
			labelWidth = float32(rl.MeasureText(labelText+"...", s.fontSize))
			if labelWidth <= maxLabelWidth {
				labelText += "..."
				break
			}
		}
	}

	rl.DrawText(labelText, int32(s.bounds.X), int32(labelY), s.fontSize, s.labelColor)
	rl.DrawText(valueText, int32(s.bounds.X+s.bounds.Width-valueTextWidth), int32(labelY), s.fontSize, s.valueColor)

	trackRect := s.trackRect()
	rl.DrawRectangleRounded(trackRect, 0.5, 8, rl.NewColor(80, 80, 80, 255))

	fillWidth := float32((s.value - s.min) / (s.max - s.min))
	fillRect := rl.NewRectangle(trackRect.X, trackRect.Y, trackRect.Width*fillWidth, trackRect.Height)
	rl.DrawRectangleRounded(fillRect, 0.5, 8, rl.NewColor(30, 144, 255, 255))

	knobCenterX := trackRect.X + trackRect.Width*float32(fillWidth)
	knobCenterY := trackRect.Y + trackRect.Height/2
	rl.DrawCircle(int32(knobCenterX), int32(knobCenterY), knobRadius, rl.White)
	rl.DrawCircleLines(int32(knobCenterX), int32(knobCenterY), knobRadius, rl.NewColor(40, 40, 40, 255))
}

// GetBounds returns the slider bounds.
func (s *Slider) GetBounds() rl.Rectangle {
	return s.bounds
}

// SetPosition moves the slider to a new position.
func (s *Slider) SetPosition(x, y float32) {
	s.bounds.X = x
	s.bounds.Y = y
}

// Value returns the current slider value.
func (s *Slider) Value() float64 {
	return s.value
}

// SetValue sets the slider value within its range.
func (s *Slider) SetValue(value float64) {
	s.value = clamp(value, s.min, s.max)
}

// trackRect returns the rectangle representing the slider track.
func (s *Slider) trackRect() rl.Rectangle {
	trackY := s.bounds.Y + s.bounds.Height - sliderTrackHeight - 6
	return rl.NewRectangle(s.bounds.X, trackY, s.bounds.Width, sliderTrackHeight)
}

func (s *Slider) updateValueFromMouse(mouseX float32) {
	track := s.trackRect()
	t := clampFloat((mouseX-track.X)/track.Width, 0, 1)
	s.value = s.min + float64(t)*(s.max-s.min)
}

func (s *Slider) pointOnKnob(point rl.Vector2) bool {
	track := s.trackRect()
	fillWidth := float32((s.value - s.min) / (s.max - s.min))
	knobCenterX := track.X + track.Width*fillWidth
	knobCenterY := track.Y + track.Height/2
	dx := point.X - knobCenterX
	dy := point.Y - knobCenterY
	return dx*dx+dy*dy <= knobRadius*knobRadius
}

func (s *Slider) formatValue() string {
	format := fmt.Sprintf("%%.%df", s.precision)
	return fmt.Sprintf(format, s.value)
}

func clamp(value, min, max float64) float64 {
	return math.Min(math.Max(value, min), max)
}

func clampFloat(value, min, max float32) float32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
