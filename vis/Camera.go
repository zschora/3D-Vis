// Camera.go
package vis

import (
    "go4/geom"
    "math"
)

type Camera struct {
    radius, polarAngle, azimuth, distanceToScreen float64
}

func NewCamera (radius, polarAngle, azimuth, distanceToScreen float64) *Camera {
    return &Camera{radius, polarAngle, azimuth, distanceToScreen}
}

func (c *Camera) RotatePolar (angle float64) {
    c.polarAngle += angle
}

func (c *Camera) ScaleLinear (distance float64) {
    c.distanceToScreen += distance
}

func (c *Camera) fromWorldToView (v geom.Vertex) geom.Vertex {
    x := -v.X() * math.Sin(c.polarAngle) + v.Y() * math.Cos(c.polarAngle)
    y := -v.X() * math.Cos (c.azimuth) * math.Cos (c.polarAngle) - v.Y() * math.Cos (c.azimuth) * math.Sin(c.polarAngle) +
        v.Z() * math.Sin (c.azimuth)
    z := -v.X() * math.Sin (c.azimuth) * math.Cos (c.polarAngle) - v.Y() * math.Sin (c.azimuth) * math.Sin(c.polarAngle) -
        v.Z() * math.Cos (c.azimuth) + c.radius
        
    return geom.NewVertex (x, y, z)
}

// from world to view coords
func (c *Camera) Transform (v geom.Vertex) geom.Vertex {
    view_v := c.fromWorldToView (v)
    return geom.NewVertex (view_v.X(), view_v.Y(), view_v.Z())
}
