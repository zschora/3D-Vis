package gui

import rl "github.com/gen2brain/raylib-go/raylib"

// Scenario describes a demo configuration option.
type Scenario struct {
	Name        string
	Description string
}

// ScenarioPanel presents a list of scenarios to choose from.
type ScenarioPanel struct {
	panel         Panel
	title         Label
	buttons       []Button
	description   Label
	selectedIndex int
	scenarios     []Scenario
	onSelect      func(int)
}

// ScenarioPanelConfig defines the layout for the panel.
type ScenarioPanelConfig struct {
	X, Y float32
}

// NewScenarioPanel creates a new panel listing scenarios.
func NewScenarioPanel(config ScenarioPanelConfig, scenarios []Scenario, onSelect func(int)) *ScenarioPanel {
	panelConfig := DefaultPanelConfig()
	panelConfig.X = config.X
	panelConfig.Y = config.Y
	panelConfig.Width = 300
	panelConfig.Height = 360

	panel := NewPanel(panelConfig).(*panel)

	title := NewLabel(LabelConfig{
		X:        config.X + 10,
		Y:        config.Y + 10,
		Text:     "Test Scenarios",
		FontSize: 18,
	})

	buttonY := config.Y + 48
	buttonHeight := float32(32)
	buttonGap := float32(6)
	buttons := make([]Button, len(scenarios))

	for i := range scenarios {
		btn := NewButton(ButtonConfig{
			X:           config.X + 10,
			Y:           buttonY + float32(i)*(buttonHeight+buttonGap),
			Width:       280,
			Height:      buttonHeight,
			Text:        scenarios[i].Name,
			NormalColor: rl.NewColor(55, 55, 55, 255),
			HoverColor:  rl.NewColor(80, 80, 80, 255),
			TextColor:   rl.White,
			FontSize:    16,
		})
		buttons[i] = btn
		panel.AddElement(btn)
	}

	descriptionLabel := NewLabel(LabelConfig{
		X:        config.X + 10,
		Y:        buttonY + float32(len(scenarios))*(buttonHeight+buttonGap) + 16,
		Text:     "",
		Color:    rl.LightGray,
		FontSize: 14,
	})

	panel.AddElement(title)
	panel.AddElement(descriptionLabel)

	sp := &ScenarioPanel{
		panel:         panel,
		title:         title,
		buttons:       buttons,
		description:   descriptionLabel,
		selectedIndex: -1,
		scenarios:     scenarios,
		onSelect:      onSelect,
	}

	sp.setSelected(0)
	return sp
}

// Update processes interactions with the scenario list.
func (sp *ScenarioPanel) Update() {
	for i, btn := range sp.buttons {
		if btn.IsClicked() {
			sp.setSelected(i)
			break
		}
	}
}

// Draw delegates drawing to the panel.
func (sp *ScenarioPanel) Draw() {
	sp.panel.Draw()
}

// Panel returns the underlying panel.
func (sp *ScenarioPanel) Panel() Panel {
	return sp.panel
}

// Selected returns the currently active scenario index.
func (sp *ScenarioPanel) Selected() int {
	return sp.selectedIndex
}

func (sp *ScenarioPanel) setSelected(index int) {
	if index < 0 || index >= len(sp.scenarios) {
		return
	}

	if sp.selectedIndex == index {
		return
	}

	highlight := rl.NewColor(30, 144, 255, 255)
	highlightHover := rl.NewColor(60, 164, 255, 255)
	defaultColor := rl.NewColor(55, 55, 55, 255)
	defaultHover := rl.NewColor(80, 80, 80, 255)

	for i, btn := range sp.buttons {
		if i == index {
			btn.SetColors(highlight, highlightHover)
			btn.SetTextColor(rl.White)
		} else {
			btn.SetColors(defaultColor, defaultHover)
			btn.SetTextColor(rl.White)
		}
	}

	sp.selectedIndex = index
	if index >= 0 {
		desc := sp.scenarios[index].Description
		if clipper, ok := sp.description.(interface{ SetTextClipped(string, float32) }); ok {
			clipper.SetTextClipped(desc, 280)
		} else {
			sp.description.SetText(desc)
		}
		if sp.onSelect != nil {
			sp.onSelect(index)
		}
	}
}
