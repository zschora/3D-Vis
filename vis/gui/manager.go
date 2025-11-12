package gui

// Manager manages all UI elements
type Manager struct {
	elements []UIElement
	enabled  bool
}

// NewManager creates a new GUI manager
func NewManager() *Manager {
	return &Manager{
		elements: make([]UIElement, 0),
		enabled:  true,
	}
}

// AddElement adds a UI element to the manager
func (m *Manager) AddElement(element UIElement) {
	if element != nil {
		m.elements = append(m.elements, element)
	}
}

// RemoveElement removes a UI element from the manager
func (m *Manager) RemoveElement(element UIElement) {
	for i, e := range m.elements {
		if e == element {
			m.elements = append(m.elements[:i], m.elements[i+1:]...)
			return
		}
	}
}

// Clear removes all elements
func (m *Manager) Clear() {
	m.elements = make([]UIElement, 0)
}

// Update updates all UI elements
func (m *Manager) Update() {
	if !m.enabled {
		return
	}

	for _, element := range m.elements {
		element.Update()
	}
}

// Draw renders all UI elements
func (m *Manager) Draw() {
	if !m.enabled {
		return
	}

	for _, element := range m.elements {
		element.Draw()
	}
}

// SetEnabled enables or disables the GUI
func (m *Manager) SetEnabled(enabled bool) {
	m.enabled = enabled
}

// IsEnabled returns whether the GUI is enabled
func (m *Manager) IsEnabled() bool {
	return m.enabled
}

// GetElements returns all elements
func (m *Manager) GetElements() []UIElement {
	result := make([]UIElement, len(m.elements))
	copy(result, m.elements)
	return result
}
