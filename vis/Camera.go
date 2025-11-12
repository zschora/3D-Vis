// Camera.go
package vis

import (
	"fmt"
	"go4/geom"
	"math"
)

// camera is the default implementation of Camera interface
type camera struct {
	radius           float64
	polarAngle       float64
	azimuth          float64
	distanceToScreen float64
}

// CameraConfig holds configuration for creating a camera
type CameraConfig struct {
	Radius           float64
	PolarAngle       float64
	Azimuth          float64
	DistanceToScreen float64
}

// DefaultCameraConfig returns a default camera configuration
func DefaultCameraConfig() CameraConfig {
	return CameraConfig{
		Radius:           1000,
		PolarAngle:       math.Pi / 4,
		Azimuth:          math.Pi / 4,
		DistanceToScreen: 500,
	}
}

// NewCamera creates a new camera with the given configuration
func NewCamera(config CameraConfig) (Camera, error) {
	if config.Radius <= 0 {
		return nil, fmt.Errorf("camera radius must be positive, got %f", config.Radius)
	}
	if config.DistanceToScreen < 0 {
		return nil, fmt.Errorf("distance to screen must be non-negative, got %f", config.DistanceToScreen)
	}
	if config.DistanceToScreen > config.Radius {
		return nil, fmt.Errorf("distance to screen (%f) cannot exceed radius (%f)", config.DistanceToScreen, config.Radius)
	}

	return &camera{
		radius:           config.Radius,
		polarAngle:       config.PolarAngle,
		azimuth:          config.Azimuth,
		distanceToScreen: config.DistanceToScreen,
	}, nil
}

// NewCameraWithDefaults creates a new camera with default settings
func NewCameraWithDefaults() Camera {
	cam, _ := NewCamera(DefaultCameraConfig())
	return cam
}

// RotatePolar rotates the camera around the polar axis
func (c *camera) RotatePolar(angle float64) {
	c.polarAngle += angle
	// Normalize angle to [0, 2π)
	c.polarAngle = math.Mod(c.polarAngle, 2*math.Pi)
	if c.polarAngle < 0 {
		c.polarAngle += 2 * math.Pi
	}
}

// RotateAzimuth rotates the camera around the azimuth axis
func (c *camera) RotateAzimuth(angle float64) {
	c.azimuth += angle
	// Clamp azimuth to [0, π] to prevent flipping
	c.azimuth = math.Max(0, math.Min(math.Pi, c.azimuth))
}

// ScaleLinear adjusts the camera distance
func (c *camera) ScaleLinear(distance float64) {
	newDistance := c.distanceToScreen + distance
	if newDistance < 0 {
		c.distanceToScreen = 0
		return
	}
	if newDistance > c.radius {
		c.distanceToScreen = c.radius
		return
	}
	c.distanceToScreen = newDistance
}

// GetRadius returns the camera radius
func (c *camera) GetRadius() float64 {
	return c.radius
}

// GetPolarAngle returns the current polar angle
func (c *camera) GetPolarAngle() float64 {
	return c.polarAngle
}

// GetAzimuth returns the current azimuth angle
func (c *camera) GetAzimuth() float64 {
	return c.azimuth
}

// GetDistanceToScreen returns the distance to screen
func (c *camera) GetDistanceToScreen() float64 {
	return c.distanceToScreen
}

func (c *camera) fromWorldToView(v geom.Vertex) geom.Vertex {
	x := -v.X()*math.Sin(c.polarAngle) + v.Y()*math.Cos(c.polarAngle)
	y := -v.X()*math.Cos(c.azimuth)*math.Cos(c.polarAngle) - v.Y()*math.Cos(c.azimuth)*math.Sin(c.polarAngle) +
		v.Z()*math.Sin(c.azimuth)
	z := -v.X()*math.Sin(c.azimuth)*math.Cos(c.polarAngle) - v.Y()*math.Sin(c.azimuth)*math.Sin(c.polarAngle) -
		v.Z()*math.Cos(c.azimuth) + c.radius

	return geom.NewVertex(x, y, z)
}

const (
	minZDistance = 1e-6 // Минимальное расстояние для избежания деления на ноль
)

func (c *camera) fromViewToScreen(v geom.Vertex, screenWidth, screenHeight int) geom.Vector2d {
	var x, y float64

	// Perspective projection
	z := v.Z()
	if math.Abs(z) < minZDistance {
		// Если точка слишком близка к камере, используем ортогональную проекцию
		x = v.X()
		y = v.Y()
	} else {
		// Perspective projection
		x = c.distanceToScreen * v.X() / z
		y = c.distanceToScreen * v.Y() / z
	}

	// Центрирование на экране
	x += float64(screenWidth) / 2
	y = float64(screenHeight)/2 - y

	return geom.NewVector2d(x, y)
}

// Transform converts a world-space vertex to screen-space coordinates
func (c *camera) Transform(v geom.Vertex, screenWidth, screenHeight int) geom.Vertex2d {
	view_v := c.fromWorldToView(v)
	screen_v := c.fromViewToScreen(view_v, screenWidth, screenHeight)

	return geom.NewVertex2d(screen_v.X(), screen_v.Y())
}
