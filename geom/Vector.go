// Vector.go
package geom

import "math"

type Vector struct {
    myCoords Coords3d
}

func NewVector(x float64, y float64, z float64) Vector {
    return Vector{Coords3d{x, y, z}}
}

func NewVectorFromVertex(v Vertex) Vector {
    return Vector{v.myCoords}
}

func NewVectorFromVertices(begin Vertex, end Vertex) Vector {
    return NewVectorFromVertex(end.Subtracted (begin))
}

func (v *Vector) Subtract (other Vector) {
    v.myCoords.Subtract (other.myCoords);
}

func (v *Vector) Subtracted (other Vector) Vector {
    res := Vector{v.myCoords}
    res.Subtract (other)
    return res
}

func (v *Vector) Add (other Vector) {
    v.myCoords.Add (other.myCoords)
}

func (v *Vector) Added (other Vector) Vector {
    res := Vector{v.myCoords}
    res.Add (other)
    return res
}

func (v Vector) Equals(other Vector) bool {
    return v.myCoords.Equals(other.myCoords)
}

func (v Vector) Length() float64 {
    return v.myCoords.Length()
}

func (v *Vector) Scale(scalar float64) {
    v.myCoords.Scale(scalar)
}

func (v *Vector) Normalize() {
    v.myCoords.Normalize()
}

func (v Vector) Dot(other Vector) float64 {
    return v.myCoords.X * other.myCoords.X + 
        v.myCoords.Y*other.myCoords.Y +
        v.myCoords.Z*other.myCoords.Z
}

func (v Vector) Cross(other Vector) Vector {
    return Vector{
        Coords3d{
            v.myCoords.Y * other.myCoords.Z - v.myCoords.Z * other.myCoords.Y,
            v.myCoords.Z * other.myCoords.X - v.myCoords.X * other.myCoords.Z,
            v.myCoords.X * other.myCoords.Y - v.myCoords.Y * other.myCoords.X,
        },
    }
}

func (v Vector) Angle(other Vector) float64 {
    sc := v.Dot(other)
    l1 := v.Length()
    l2 := other.Length()
    if l1 == 0 || l2 == 0 {
        return math.NaN()
    }
    
    cosine := sc / l1 / l2
    return math.Acos(cosine) 
}

func (v Vector) X() float64 {
    return v.myCoords.X
}

func (v Vector) Y() float64 {
    return v.myCoords.Y
}

func (v Vector) Z() float64 {
    return v.myCoords.Z
}