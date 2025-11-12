// Package vis provides visualization components for rendering 3D scenes.
//
// The package uses an interface-based architecture for flexibility and testability:
//   - Application: Main application loop and window management with configurable settings
//   - Camera interface: 3D camera with perspective projection (implemented by camera)
//   - Renderer interface: Renders meshes to screen using raylib (implemented by renderer)
//   - Scene interface: Container for 3D meshes (implemented by scene)
//
// All components can be configured through Config structs and support dependency injection
// through interfaces, making the codebase flexible and easy to test.
package vis
