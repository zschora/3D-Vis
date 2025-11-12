package gui

// PrimitiveType represents the type of 3D primitive
type PrimitiveType int

const (
	PrimitiveCube PrimitiveType = iota
	PrimitiveTetrahedron
)

// String returns the string representation of the primitive type
func (pt PrimitiveType) String() string {
	switch pt {
	case PrimitiveCube:
		return "Cube"
	case PrimitiveTetrahedron:
		return "Tetrahedron"
	default:
		return "Unknown"
	}
}

// PrimitiveSelector provides buttons to select different primitives
type PrimitiveSelector struct {
	panel        Panel
	cubeButton   Button
	tetraButton  Button
	selectedType PrimitiveType
	onChange     func(PrimitiveType)
}

// PrimitiveSelectorConfig holds configuration for creating a primitive selector
type PrimitiveSelectorConfig struct {
	X, Y float32
}

// NewPrimitiveSelector creates a new primitive selector
func NewPrimitiveSelector(config PrimitiveSelectorConfig, onChange func(PrimitiveType)) *PrimitiveSelector {
	panelConfig := DefaultPanelConfig()
	panelConfig.X = config.X
	panelConfig.Y = config.Y
	panelConfig.Width = 200
	panelConfig.Height = 100

	panel := NewPanel(panelConfig).(*panel)

	// Create buttons
	cubeButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 10,
		Width:  180,
		Height: 30,
		Text:   "Cube",
	})

	tetraButton := NewButton(ButtonConfig{
		X:      config.X + 10,
		Y:      config.Y + 50,
		Width:  180,
		Height: 30,
		Text:   "Tetrahedron",
	})

	// Add buttons to panel
	panel.AddElement(cubeButton)
	panel.AddElement(tetraButton)

	return &PrimitiveSelector{
		panel:        panel,
		cubeButton:   cubeButton,
		tetraButton:  tetraButton,
		selectedType: PrimitiveCube,
		onChange:     onChange,
	}
}

// Update updates the primitive selector
func (ps *PrimitiveSelector) Update() {
	ps.panel.Update()
}

// Draw renders the primitive selector
func (ps *PrimitiveSelector) Draw() {
	ps.panel.Draw()
}

// HandleInput handles button clicks
func (ps *PrimitiveSelector) HandleInput() {
	if ps.cubeButton.IsClicked() {
		ps.selectedType = PrimitiveCube
		if ps.onChange != nil {
			ps.onChange(PrimitiveCube)
		}
	}
	if ps.tetraButton.IsClicked() {
		ps.selectedType = PrimitiveTetrahedron
		if ps.onChange != nil {
			ps.onChange(PrimitiveTetrahedron)
		}
	}
}

// GetSelectedType returns the currently selected primitive type
func (ps *PrimitiveSelector) GetSelectedType() PrimitiveType {
	return ps.selectedType
}

// GetPanel returns the underlying panel
func (ps *PrimitiveSelector) GetPanel() Panel {
	return ps.panel
}
