# Testing and Development Guide

## Quick Testing During Development

### Option 1: Demo Application (Recommended)

The easiest way to test during development is to use the demo application:

```bash
go run ./cmd/demo
```

This will:
- Load a test scene with a cube
- Automatically rotate the camera
- Display the 3D visualization

### Option 2: Enable Test Scene in Your Application

You can enable the test scene in your own application:

```go
package main

import "go4/vis"

func main() {
	config := vis.DefaultApplicationConfig()
	config.LoadTestScene = true  // Enable auto-loading test scene
	
	app, err := vis.NewApplication(config)
	if err != nil {
		panic(err)
	}
	
	app.Run()
}
```

### Option 3: Manual Test Scene Setup

For more control, you can manually set up a test scene:

```go
package main

import "go4/vis"

func main() {
	app, err := vis.NewApplicationWithDefaults()
	if err != nil {
		panic(err)
	}
	
	// Setup test scene with custom configuration
	testConfig := vis.TestSceneConfig{
		AutoRotate: true,  // Enable auto-rotation
		AutoZoom:   false, // Disable auto-zoom
		MeshSize:   200,   // Cube size
	}
	app.SetupTestScene(testConfig)
	
	app.Run()
}
```

### Option 4: Quick Test Scene Helper

You can also use the helper function:

```go
app.SetupTestSceneWithDefaults()  // Uses default settings
```

## Test Scene Configuration

The `TestSceneConfig` struct allows you to customize the test scene:

- **AutoRotate**: Automatically rotates the camera around the object
- **AutoZoom**: Automatically zooms in/out (placeholder for future implementation)
- **MeshSize**: Size of the test cube mesh

## Controls (when using manual camera control)

- **Arrow Keys**: Rotate camera
  - Left/Right: Rotate around polar axis
  - Up/Down: Rotate around azimuth axis
- **Q/E Keys**: Zoom in/out

## Building Test Executables

Build the demo application:

```bash
go build -o demo.exe ./cmd/demo
```

Then run:
```bash
./demo.exe
```

## Unit Testing

For unit tests of geometric operations:

```bash
go test ./geom -v
```

## Tips for Development

1. **Quick Visual Testing**: Use `cmd/demo` for instant visual feedback
2. **Custom Scenes**: Create your own test scenes using `vis.NewScene()` and `geom.CreateCube()` or `geom.CreateTetrahedron()`
3. **Multiple Meshes**: Add multiple meshes to test complex scenes
4. **Camera Testing**: Use `app.GetRenderer().GetCamera()` to access and test camera controls

