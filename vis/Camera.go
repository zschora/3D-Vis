// Camera.go
package vis

import (
	"go4/geom"
	"math"
)

type Camera struct {
	radius, polarAngle, azimuth, distanceToScreen float64
}

func NewCamera(radius, polarAngle, azimuth, distanceToScreen float64) *Camera {
	return &Camera{radius, polarAngle, azimuth, distanceToScreen}
}

func (c *Camera) RotatePolar(angle float64) {
	c.polarAngle += angle
}

func (c *Camera) RotateAzimuth(angle float64) {
	c.azimuth += angle
}

func (c *Camera) ScaleLinear(distance float64) {
	if c.distanceToScreen+distance > c.radius {
		//c.distanceToScreen = c.radius
		//return
	} else if c.distanceToScreen+distance < 0 {
		c.distanceToScreen = 0
		return
	}
	c.distanceToScreen += distance
}

func (c *Camera) fromWorldToView(v geom.Vertex) geom.Vertex {
	x := -v.X()*math.Sin(c.polarAngle) + v.Y()*math.Cos(c.polarAngle)
	y := -v.X()*math.Cos(c.azimuth)*math.Cos(c.polarAngle) - v.Y()*math.Cos(c.azimuth)*math.Sin(c.polarAngle) +
		v.Z()*math.Sin(c.azimuth)
	z := -v.X()*math.Sin(c.azimuth)*math.Cos(c.polarAngle) - v.Y()*math.Sin(c.azimuth)*math.Sin(c.polarAngle) -
		v.Z()*math.Cos(c.azimuth) + c.radius

	return geom.NewVertex(x, y, z)
}

func (c *Camera) fromViewToScreen(v geom.Vertex, screenWidth, screenHeight int) geom.Vector2d {
	var x, y float64
	if false {
		// non perspective projection
		x = v.X()
		y = v.Y()
	} else {
		//scale
		x = c.distanceToScreen * v.X() / v.Z()
		y = c.distanceToScreen * v.Y() / v.Z()

	}
	// move
	x += float64(screenWidth) / 2
	y = float64(screenHeight)/2 - y

	return geom.NewVector2d(x, y)
}

// from world to view coords
func (c *Camera) Transform(v geom.Vertex, screenWidth, screenHeight int) geom.Vertex2d {
	view_v := c.fromWorldToView(v)
	screen_v := c.fromViewToScreen(view_v, screenWidth, screenHeight)

	return geom.NewVertex2d(screen_v.X(), screen_v.Y())
}
