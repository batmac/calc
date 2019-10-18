package calc

import "math"

// Point is our  basic  type
type Point struct {
	x, y float64
}

// Hypot returns Sqrt(p*p + q*q), taking care to avoid
// unnecessary overflow and underflow.
func (p Point) Hypot() float64 {
	return math.Hypot(p.x, p.y)
}
