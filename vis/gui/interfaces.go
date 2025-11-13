package gui

import rl "github.com/gen2brain/raylib-go/raylib"

// UIElement represents a basic UI element
type UIElement interface {
	// Update updates the element state and returns true if it was interacted with
	Update() bool

	// Draw renders the element
	Draw()

	// GetBounds returns the element's bounding rectangle
	GetBounds() rl.Rectangle

	// SetPosition sets the element's position
	SetPosition(x, y float32)
}

// Button represents a clickable button
type Button interface {
	UIElement
	// IsClicked returns true if the button was clicked this frame
	IsClicked() bool
	// IsHeld returns true while the button is kept pressed
	IsHeld() bool
	// SetText sets the button text
	SetText(text string)
	// GetText returns the button text
	GetText() string
	// SetColors updates button colors for normal and hover states
	SetColors(normal, hover rl.Color)
	// SetTextColor updates the button text color
	SetTextColor(color rl.Color)
}

// Label represents a text label
type Label interface {
	UIElement
	// SetText sets the label text
	SetText(text string)
	// GetText returns the label text
	GetText() string
}

// Panel represents a container for other UI elements
type Panel interface {
	UIElement
	// AddElement adds a UI element to the panel
	AddElement(element UIElement)
	// GetElements returns all elements in the panel
	GetElements() []UIElement
}
