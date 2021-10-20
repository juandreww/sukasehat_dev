package geometry
import "math"
type Point struct{ X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
return math.Hypot(q.Xp.
X, q.Yp.
Y)
}

// same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
return math.Hypot(q.Xp.
X, q.Yp.
Y)
}