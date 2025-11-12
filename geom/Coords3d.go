// Coords3d.go
package geom

import "math"

type Coords3d struct {
	X, Y, Z float64
}

func (c Coords3d) Equals(other Coords3d) bool {
	return math.Abs(c.X-other.X) < DefaultTolerance &&
		math.Abs(c.Y-other.Y) < DefaultTolerance &&
		math.Abs(c.Z-other.Z) < DefaultTolerance
}

func (c Coords3d) Length() float64 {
	return math.Sqrt(c.X*c.X + c.Y*c.Y + c.Z*c.Z)
}

func (c *Coords3d) Subtract(other Coords3d) {
	c.X -= other.X
	c.Y -= other.Y
	c.Z -= other.Z
}

func (c *Coords3d) Add(other Coords3d) {
	c.X += other.X
	c.Y += other.Y
	c.Z += other.Z
}

func (c *Coords3d) Subtracted(other Coords3d) Coords3d {
	res := Coords3d{c.X, c.Y, c.Z}
	res.Subtract(other)
	return res
}

func (c *Coords3d) Scale(scalar float64) {
	c.X *= scalar
	c.Y *= scalar
	c.Z *= scalar
}

func (c *Coords3d) Normalize() {
	length := c.Length()
	if length == 0 {
		return
	}

	c.Scale(1.0 / length)
}

func (c *Coords3d) Distance(other Coords3d) float64 {
	vec := c.Subtracted(other)
	return vec.Length()
}
