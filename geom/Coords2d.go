// Coords2d.go
package geom

import "math"

const (
	// DefaultTolerance is the default tolerance for floating-point comparisons
	DefaultTolerance = 1e-9
)

type Coords2d struct {
	X, Y float64
}

func (c Coords2d) Equals(other Coords2d) bool {
	return c.Distance(other) < DefaultTolerance
}

func (c Coords2d) Length() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y)
}

func (c *Coords2d) Subtract(other Coords2d) {
	c.X -= other.X
	c.Y -= other.Y
}

func (c *Coords2d) Subtracted(other Coords2d) Coords2d {
	res := Coords2d{c.X, c.Y}
	res.Subtract(other)
	return res
}

func (c *Coords2d) Scale(scalar float64) {
	c.X *= scalar
	c.Y *= scalar
}

func (c *Coords2d) Normalize() {
	length := c.Length()
	if length == 0 {
		return
	}

	c.Scale(1.0 / length)
}

func (c *Coords2d) Distance(other Coords2d) float64 {
	vec := c.Subtracted(other)
	return vec.Length()
}
