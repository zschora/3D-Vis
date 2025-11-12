package vis

import (
	"go4/geom"
	"time"
)

// TestSceneConfig holds configuration for creating a test scene
type TestSceneConfig struct {
	AutoRotate bool
	AutoZoom   bool
	MeshSize   float64
}

// DefaultTestSceneConfig returns default test scene configuration
func DefaultTestSceneConfig() TestSceneConfig {
	return TestSceneConfig{
		AutoRotate: true,
		AutoZoom:   true,
		MeshSize:   200,
	}
}

// CreateTestScene creates a test scene with a cube for development/testing
func CreateTestScene(config TestSceneConfig) Scene {
	scene := NewScene()
	cube := geom.CreateCube(config.MeshSize)
	scene.AddMesh(cube)
	return scene
}

// CreateTestSceneWithDefaults creates a test scene with default settings
func CreateTestSceneWithDefaults() Scene {
	return CreateTestScene(DefaultTestSceneConfig())
}

// SetupTestScene adds a test scene to the application and sets up auto-rotation/zoom
func (app *Application) SetupTestScene(config TestSceneConfig) {
	scene := CreateTestScene(config)
	app.AddScene(scene)

	if config.AutoRotate || config.AutoZoom {
		app.SetUpdateFunction(func(deltaTime time.Duration) {
			deltaSeconds := deltaTime.Seconds()
			camera := app.GetRenderer().GetCamera()

			if config.AutoRotate {
				// Автоматическое вращение камеры
				camera.RotatePolar(deltaSeconds * 0.5) // Медленное вращение
			}

			if config.AutoZoom {
				// Автоматическое приближение/отдаление (синусоидальное движение)
				// Можно добавить логику для синусоидального движения в будущем
				_ = deltaSeconds // Placeholder for future implementation
			}
		})
	}
}

// SetupTestSceneWithDefaults sets up a test scene with default settings
func (app *Application) SetupTestSceneWithDefaults() {
	app.SetupTestScene(DefaultTestSceneConfig())
}
