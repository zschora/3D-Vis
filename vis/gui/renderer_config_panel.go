package gui

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var rendererColorPalette = []rl.Color{
	rl.White,
	rl.Gray,
	rl.DarkGray,
	rl.Black,
	rl.Red,
	rl.Orange,
	rl.Gold,
	rl.Green,
	rl.Blue,
	rl.DarkBlue,
	rl.Purple,
	rl.SkyBlue,
}

// RendererConfigData mirrors renderer visual configuration without introducing package cycles.
type RendererConfigData struct {
	BackgroundColor    rl.Color
	UseRandomFaceColor bool
	FaceColor          rl.Color
	EdgeColor          rl.Color
	AlphaValue         uint8
	DrawFaces          bool
	DrawEdges          bool
	UseBackfaceCulling bool
}

// RendererConfigPanel provides interactive controls for renderer configuration.
type RendererConfigPanel struct {
	panel             Panel
	title             Label
	drawFaces         *Toggle
	drawEdges         *Toggle
	randomColors      *Toggle
	backfaceCull      *Toggle
	alphaSlider       *Slider
	faceLabel         Label
	edgeLabel         Label
	backgroundLabel   Label
	facePreview       *ColorPreview
	edgePreview       *ColorPreview
	backgroundPreview *ColorPreview
	faceButton        Button
	edgeButton        Button
	backgroundButton  Button
	state             RendererConfigData
	onChange          func(RendererConfigData)
}

// RendererConfigPanelConfig configures layout for the panel.
type RendererConfigPanelConfig struct {
	X, Y float32
}

// NewRendererConfigPanel constructs the panel using an initial state and change callback.
func NewRendererConfigPanel(layout RendererConfigPanelConfig, initial RendererConfigData, onChange func(RendererConfigData)) *RendererConfigPanel {
	panelConfig := DefaultPanelConfig()
	panelConfig.X = layout.X
	panelConfig.Y = layout.Y
	panelConfig.Width = 340
	panelConfig.Height = 380

	panel := NewPanel(panelConfig).(*panel)

	title := NewLabel(LabelConfig{
		X:        layout.X + 10,
		Y:        layout.Y + 10,
		Text:     "Renderer Config",
		FontSize: 18,
	})

	toggleY := layout.Y + 46
	drawFaces := NewToggle(ToggleConfig{
		X:       layout.X + 10,
		Y:       toggleY,
		Label:   "Draw faces",
		Initial: initial.DrawFaces,
	})

	drawEdges := NewToggle(ToggleConfig{
		X:       layout.X + 10,
		Y:       toggleY + 32,
		Label:   "Draw edges",
		Initial: initial.DrawEdges,
	})

	randomColors := NewToggle(ToggleConfig{
		X:       layout.X + 10,
		Y:       toggleY + 64,
		Label:   "Use random face color",
		Initial: initial.UseRandomFaceColor,
	})

	backfaceCull := NewToggle(ToggleConfig{
		X:       layout.X + 10,
		Y:       toggleY + 96,
		Label:   "Backface culling",
		Initial: initial.UseBackfaceCulling,
	})

	colorSectionY := toggleY + 140
	labelWidth := float32(160)
	previewX := layout.X + 10 + labelWidth + 8
	buttonX := layout.X + 10 + labelWidth + 8 + 40 + 8

	faceLabel := newColorLabel(layout.X+10, colorSectionY, "Face color", initial.FaceColor, initial.AlphaValue)
	facePreview := NewColorPreview(ColorPreviewConfig{
		X:      previewX,
		Y:      colorSectionY,
		Width:  40,
		Height: 20,
		Color:  initial.FaceColor,
	})
	faceButton := NewButton(ButtonConfig{
		X:           buttonX,
		Y:           colorSectionY - 2,
		Width:       60,
		Height:      24,
		Text:        "Next",
		NormalColor: rl.NewColor(60, 60, 60, 255),
		HoverColor:  rl.NewColor(90, 90, 90, 255),
		TextColor:   rl.White,
		FontSize:    14,
	})

	edgeLabel := newColorLabel(layout.X+10, colorSectionY+32, "Edge color", initial.EdgeColor, 255)
	edgePreview := NewColorPreview(ColorPreviewConfig{
		X:      previewX,
		Y:      colorSectionY + 32,
		Width:  40,
		Height: 20,
		Color:  initial.EdgeColor,
	})
	edgeButton := NewButton(ButtonConfig{
		X:           buttonX,
		Y:           colorSectionY + 30,
		Width:       60,
		Height:      24,
		Text:        "Next",
		NormalColor: rl.NewColor(60, 60, 60, 255),
		HoverColor:  rl.NewColor(90, 90, 90, 255),
		TextColor:   rl.White,
		FontSize:    14,
	})

	backgroundLabel := newColorLabel(layout.X+10, colorSectionY+64, "Background", initial.BackgroundColor, 255)
	backgroundPreview := NewColorPreview(ColorPreviewConfig{
		X:      previewX,
		Y:      colorSectionY + 64,
		Width:  40,
		Height: 20,
		Color:  initial.BackgroundColor,
	})
	backgroundButton := NewButton(ButtonConfig{
		X:           buttonX,
		Y:           colorSectionY + 62,
		Width:       60,
		Height:      24,
		Text:        "Next",
		NormalColor: rl.NewColor(60, 60, 60, 255),
		HoverColor:  rl.NewColor(90, 90, 90, 255),
		TextColor:   rl.White,
		FontSize:    14,
	})

	alphaSlider := NewSlider(SliderConfig{
		X:         layout.X + 10,
		Y:         colorSectionY + 108,
		Width:     320,
		Label:     "Face alpha",
		Min:       0,
		Max:       255,
		Value:     float64(initial.AlphaValue),
		Precision: 0,
	})

	panel.AddElement(title)
	panel.AddElement(drawFaces)
	panel.AddElement(drawEdges)
	panel.AddElement(randomColors)
	panel.AddElement(backfaceCull)
	panel.AddElement(faceLabel)
	panel.AddElement(facePreview)
	panel.AddElement(faceButton)
	panel.AddElement(edgeLabel)
	panel.AddElement(edgePreview)
	panel.AddElement(edgeButton)
	panel.AddElement(backgroundLabel)
	panel.AddElement(backgroundPreview)
	panel.AddElement(backgroundButton)
	panel.AddElement(alphaSlider)

	return &RendererConfigPanel{
		panel:             panel,
		title:             title,
		drawFaces:         drawFaces,
		drawEdges:         drawEdges,
		randomColors:      randomColors,
		backfaceCull:      backfaceCull,
		alphaSlider:       alphaSlider,
		faceLabel:         faceLabel,
		edgeLabel:         edgeLabel,
		backgroundLabel:   backgroundLabel,
		facePreview:       facePreview,
		edgePreview:       edgePreview,
		backgroundPreview: backgroundPreview,
		faceButton:        faceButton,
		edgeButton:        edgeButton,
		backgroundButton:  backgroundButton,
		state:             initial,
		onChange:          onChange,
	}
}

// Update refreshes UI and emits configuration changes.
func (rcp *RendererConfigPanel) Update() {
	prev := rcp.state

	if rcp.faceButton.IsClicked() {
		rcp.state.FaceColor = nextPaletteColor(rcp.state.FaceColor)
		rcp.facePreview.SetColor(rcp.state.FaceColor)
	}
	if rcp.edgeButton.IsClicked() {
		rcp.state.EdgeColor = nextPaletteColor(rcp.state.EdgeColor)
		rcp.edgePreview.SetColor(rcp.state.EdgeColor)
	}
	if rcp.backgroundButton.IsClicked() {
		rcp.state.BackgroundColor = nextPaletteColor(rcp.state.BackgroundColor)
		rcp.backgroundPreview.SetColor(rcp.state.BackgroundColor)
	}

	rcp.state.DrawFaces = rcp.drawFaces.Value()
	rcp.state.DrawEdges = rcp.drawEdges.Value()
	rcp.state.UseRandomFaceColor = rcp.randomColors.Value()
	rcp.state.UseBackfaceCulling = rcp.backfaceCull.Value()

	rcp.state.AlphaValue = uint8(math.Round(rcp.alphaSlider.Value()))

	// Update labels to reflect color values
	rcp.updateColorLabelClipped(rcp.faceLabel, "Face color", rcp.state.FaceColor, rcp.state.AlphaValue)
	rcp.updateColorLabelClipped(rcp.edgeLabel, "Edge color", rcp.state.EdgeColor, 255)
	rcp.updateColorLabelClipped(rcp.backgroundLabel, "Background", rcp.state.BackgroundColor, 255)

	if !rendererConfigEqual(prev, rcp.state) && rcp.onChange != nil {
		rcp.onChange(rcp.state)
	}
}

// Draw delegates draw call to the underlying panel.
func (rcp *RendererConfigPanel) Draw() {
	rcp.panel.Draw()
}

// Panel returns the root panel.
func (rcp *RendererConfigPanel) Panel() Panel {
	return rcp.panel
}

// State returns the current renderer configuration.
func (rcp *RendererConfigPanel) State() RendererConfigData {
	return rcp.state
}

func rendererConfigEqual(a, b RendererConfigData) bool {
	return a.BackgroundColor == b.BackgroundColor &&
		a.FaceColor == b.FaceColor &&
		a.EdgeColor == b.EdgeColor &&
		a.UseRandomFaceColor == b.UseRandomFaceColor &&
		a.AlphaValue == b.AlphaValue &&
		a.DrawFaces == b.DrawFaces &&
		a.DrawEdges == b.DrawEdges &&
		a.UseBackfaceCulling == b.UseBackfaceCulling
}

func nextPaletteColor(current rl.Color) rl.Color {
	if len(rendererColorPalette) == 0 {
		return current
	}
	for i, color := range rendererColorPalette {
		if color == current {
			return rendererColorPalette[(i+1)%len(rendererColorPalette)]
		}
	}
	return rendererColorPalette[0]
}

func newColorLabel(x, y float32, prefix string, color rl.Color, alpha uint8) Label {
	return NewLabel(LabelConfig{
		X:        x,
		Y:        y,
		Text:     colorLabelText(prefix, color, alpha),
		FontSize: 14,
		Color:    rl.LightGray,
	})
}

func updateColorLabel(label Label, prefix string, color rl.Color, alpha uint8) {
	label.SetText(colorLabelText(prefix, color, alpha))
}

func colorLabelText(prefix string, color rl.Color, alpha uint8) string {
	return fmt.Sprintf("%s: R%3d G%3d B%3d A%3d", prefix, color.R, color.G, color.B, alpha)
}

func (rcp *RendererConfigPanel) updateColorLabelClipped(label Label, prefix string, color rl.Color, alpha uint8) {
	text := colorLabelText(prefix, color, alpha)
	if clipper, ok := label.(interface{ SetTextClipped(string, float32) }); ok {
		clipper.SetTextClipped(text, 160)
	} else {
		label.SetText(text)
	}
}
