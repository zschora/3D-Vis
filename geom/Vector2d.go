// Vector2d.go
package geom

import "math"

type Vector2d struct {
    myCoords Coords2d
}

func NewVector2d(x float64, y float64) Vector2d {
    return Vector2d{Coords2d{x, y}}
}

func NewVector2dFromVertex(v Vertex2d) Vector2d {
    return Vector2d{v.myCoords}
}

func NewVector2dFromVertices(begin Vertex2d, end Vertex2d) Vector2d {
    return NewVector2dFromVertex(end.Subtracted (begin))
}

func (v *Vector2d) Subtract (other Vector2d) {
    v.myCoords.Subtract (other.myCoords);
}

func (v Vector2d) Equals(other Vector2d) bool {
    return v.myCoords.Equals(other.myCoords)
}

func (v Vector2d) Length() float64 {
    return v.myCoords.Length()
}

func (v *Vector2d) Scale(scalar float64) {
    v.myCoords.Scale(scalar)
}

func (v *Vector2d) Normalize() {
    v.myCoords.Normalize()
}

func (v Vector2d) Dot(other Vector2d) float64 {
    return v.myCoords.X * other.myCoords.X + 
        v.myCoords.Y*other.myCoords.Y
}

func (v Vector2d) Cross(other Vector2d) float64 {
    return v.myCoords.X * other.myCoords.Y - v.myCoords.Y * other.myCoords.X
}

func (v Vector2d) Angle(other Vector2d) float64 {
    sc := v.Dot(other)
    l1 := v.Length()
    l2 := other.Length()
    if l1 == 0 || l2 == 0 {
        return math.NaN()
    }
    
    cosine := sc / l1 / l2
    return math.Acos(cosine) 
}

func (v Vector2d) X() float64 {
    return v.myCoords.X
}

func (v Vector2d) Y() float64 {
    return v.myCoords.Y
}