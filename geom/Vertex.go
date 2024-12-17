// Vertex.go
package geom

type Vertex struct {
    myCoords Coords3d
}

func NewVertex(x float64, y float64, z float64) Vertex {
    return Vertex{Coords3d{x, y, z}}
}

func (v *Vertex) Subtract (other Vertex) {
    v.myCoords.Subtract (other.myCoords)
}

func (v *Vertex) Subtracted (other Vertex) Vertex {
    res := Vertex{v.myCoords}
    res.Subtract (other)
    return res
}

func (v *Vertex) Add (other Vertex) {
    v.myCoords.Add (other.myCoords)
}

func (v *Vertex) Added (other Vertex) Vertex {
    res := Vertex{v.myCoords}
    res.Add (other)
    return res
}

func (v *Vertex) Distance(other Vertex) float64 {
    return v.myCoords.Distance (other.myCoords)
}

func (v *Vertex) X() float64 {
    return v.myCoords.X
}

func (v *Vertex) Y() float64 {
    return v.myCoords.Y
}

func (v *Vertex) Z() float64 {
    return v.myCoords.Z
}