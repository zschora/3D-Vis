// Vertex2d.go
package geom

type Vertex2d struct {
    myCoords Coords2d
}

func NewVertex2d(x float64, y float64) Vertex2d {
    return Vertex2d{Coords2d{x, y}}
}

func (v *Vertex2d) Subtract (other Vertex2d) {
    v.myCoords.Subtract (other.myCoords)
}

func (v *Vertex2d) Subtracted (other Vertex2d) Vertex2d {
    res := Vertex2d{v.myCoords}
    res.Subtract (other)
    return res
}

func (v *Vertex2d) Distance(other Vertex2d) float64 {
    return v.myCoords.Distance (other.myCoords)
}

func (v *Vertex2d) X() float64 {
    return v.myCoords.X
}

func (v *Vertex2d) Y() float64 {
    return v.myCoords.Y
}
