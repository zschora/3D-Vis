# 3D Viewer Application

A Go-based application for rendering and interacting with 3D models.
The project provides a clean architecture for 3D visualization with a focus on extensibility and ease of use.

---

## Features

- **geom library**: Geometric entities for 3D modeling (vertices, vectors, meshes, primitives).
- **3D Model Rendering**: Displays 3D meshes with correctly oriented faces and back-face culling.
- **Camera Control**: Polar camera system with rotation, zoom, and perspective controls.
- **Flexible Architecture**: Interface-based design for easy testing and extension.
- **Test Scene**: Built-in test scene with auto-rotation for quick development testing.
- **GUI System**: Simple GUI with buttons, labels, and panels for application control.
  - Navigation panel with reset view and zoom controls
  - Info panel displaying FPS, camera parameters, and scene information
  - Demo application with primitive selection and motion type controls

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/zschora/3D-Vis.git
   ```
2. Navigate to the project directory:
   ```bash
   cd 3D-Vis
   ```
3. Ensure you have Go installed (version 1.20 or later).
4. Install dependencies:
   ```bash
   go mod tidy
   ```
5. Run the application:
   ```bash
   go run main.go
   ```
   
   Or run the demo with auto-rotating test scene:
   ```bash
   go run ./cmd/demo
   ```

---

## Testing

For quick testing during development, see [TESTING.md](TESTING.md).

Quick start:
```bash
go run ./cmd/demo
```

This will launch a demo with an auto-rotating test scene.

---

## Controls

### Keyboard Controls
- **Arrow Keys**: Rotate camera (Left/Right: polar rotation, Up/Down: azimuth rotation)
- **Q/E**: Zoom in/out
- **ESC**: Close application

### Mouse Controls
- GUI buttons for navigation (Reset View, Zoom In/Out)

---

## Roadmap

- [ ] Add support for loading external 3D files (e.g., STL, OBJ).
- [ ] Improve renderer functionality (better face sorting, depth buffer).
- [ ] Implement lighting and shading.
- [ ] Improve camera controls (mouse drag rotation, smooth zoom).
- [x] Introduce a GUI for easier user interaction.
- [ ] Add mesh editing capabilities (vertex manipulation, face editing).
- [ ] Support for multiple scenes and scene management.
- [ ] Export functionality (save rendered images, export meshes).

---

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature/new-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "feat: add new feature"
   ```
4. Push your branch:
   ```bash
   git push origin feature/new-feature
   ```
5. Create a Pull Request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

## Acknowledgements

- Built using [Raylib-go](https://github.com/gen2brain/raylib-go) for rendering.
- Inspired by the desire to create a simple and efficient 3D viewer.

---

## Screenshots

![Screenshot of the 3D Viewer](sc0.png)
