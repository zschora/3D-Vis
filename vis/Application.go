package vis

import (
	"fmt"
	"time"

	"go4/vis/gui"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// ApplicationConfig holds configuration for the application
type ApplicationConfig struct {
	Width         int
	Height        int
	Title         string
	TargetFPS     int
	Camera        CameraConfig
	Renderer      RendererConfig
	LoadTestScene bool             // If true, automatically loads a test scene
	TestScene     *TestSceneConfig // Configuration for test scene (used if LoadTestScene is true)
}

// DefaultApplicationConfig returns default application configuration
func DefaultApplicationConfig() ApplicationConfig {
	return ApplicationConfig{
		Width:         1080,
		Height:        720,
		Title:         "3D Visualization App",
		TargetFPS:     60,
		Camera:        DefaultCameraConfig(),
		Renderer:      DefaultRendererConfig(),
		LoadTestScene: false,
		TestScene:     nil,
	}
}

// Application represents the main application
type Application struct {
	config   ApplicationConfig
	renderer Renderer
	scenes   []Scene
	updateFn func(deltaTime time.Duration)
	gui      *gui.Manager
}

// NewApplication creates a new application with the given configuration
func NewApplication(config ApplicationConfig) (*Application, error) {
	if config.Width <= 0 || config.Height <= 0 {
		return nil, fmt.Errorf("invalid window size: %dx%d", config.Width, config.Height)
	}
	if config.TargetFPS <= 0 {
		return nil, fmt.Errorf("target FPS must be positive, got %d", config.TargetFPS)
	}

	rl.InitWindow(int32(config.Width), int32(config.Height), config.Title)

	camera, err := NewCamera(config.Camera)
	if err != nil {
		rl.CloseWindow()
		return nil, fmt.Errorf("failed to create camera: %w", err)
	}

	renderer := NewRenderer(camera, config.Renderer)

	app := &Application{
		config:   config,
		renderer: renderer,
		scenes:   make([]Scene, 0),
		gui:      gui.NewManager(),
	}

	// Auto-load test scene if configured
	if config.LoadTestScene {
		testConfig := config.TestScene
		if testConfig == nil {
			defaultTestConfig := DefaultTestSceneConfig()
			testConfig = &defaultTestConfig
		}
		app.SetupTestScene(*testConfig)
	}

	return app, nil
}

// NewApplicationWithDefaults creates a new application with default settings
func NewApplicationWithDefaults() (*Application, error) {
	return NewApplication(DefaultApplicationConfig())
}

// AddScene adds a scene to the application
func (app *Application) AddScene(scene Scene) {
	if scene != nil {
		app.scenes = append(app.scenes, scene)
	}
}

// RemoveScene removes a scene from the application by index
func (app *Application) RemoveScene(index int) error {
	if index < 0 || index >= len(app.scenes) {
		return fmt.Errorf("scene index out of bounds: %d (application has %d scenes)", index, len(app.scenes))
	}
	app.scenes = append(app.scenes[:index], app.scenes[index+1:]...)
	return nil
}

// GetScenes returns all scenes in the application
func (app *Application) GetScenes() []Scene {
	result := make([]Scene, len(app.scenes))
	copy(result, app.scenes)
	return result
}

// SetUpdateFunction sets a custom update function that will be called each frame
func (app *Application) SetUpdateFunction(fn func(deltaTime time.Duration)) {
	app.updateFn = fn
}

// GetRenderer returns the renderer
func (app *Application) GetRenderer() Renderer {
	return app.renderer
}

// GetGUI returns the GUI manager
func (app *Application) GetGUI() *gui.Manager {
	return app.gui
}

// Run starts the application main loop
func (app *Application) Run() {
	rl.SetTargetFPS(int32(app.config.TargetFPS))

	for !rl.WindowShouldClose() {
		app.Update()
		app.Render()
	}

	app.Close()
}

// Update updates the application state
func (app *Application) Update() {
	deltaTime := time.Duration(rl.GetFrameTime() * float32(time.Second))

	// Update GUI first (to handle input)
	app.gui.Update()

	// Update application logic
	if app.updateFn != nil {
		app.updateFn(deltaTime)
	}
}

// Render renders all scenes
func (app *Application) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(app.config.Renderer.BackgroundColor)

	// Render 3D scenes
	for _, scene := range app.scenes {
		app.renderer.Render(scene)
	}

	// Render GUI on top
	app.gui.Draw()

	rl.EndDrawing()
}

// Close closes the application and cleans up resources
func (app *Application) Close() {
	rl.CloseWindow()
}
