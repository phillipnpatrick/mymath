package geometry

import (
	"math"
)

type Point struct {
	X float64
	Y float64
}

// Distance returns the Euclidean distance between two points.
func (p Point) Distance(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	sum := dx*dx + dy*dy

	return math.Sqrt(sum)
}

// Move shifts the point by dx and dy.
func (p *Point) Move(dx, dy float64) {
	p.X += dx
	p.Y += dy
}

// Add returns the vector addition of two points.
func (p Point) Add(other Point) Point {
	return Point{X: p.X+other.X, Y: p.Y+other.Y}
}

// Subtract returns the vector subtraction of two points.
func (p Point) Subtract(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

// Scale scales the vector by a given factor.
func (p Point) Scale(factor float64) Point {
	return Point{X: p.X*factor, Y: p.Y * factor}
}

// Dot returns the dot product of two vectors.
func (p Point) Dot(other Point) float64 {
	return p.X * other.X + p.Y * other.Y
}

// Magnitude returns the length of the vector.
func (p Point) Magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Normalize returns a unit vector in the same direction.
func (p Point) Normalize() Point {
	mag := p.Magnitude()
	if mag == 0 {
		return Point{X: 0, Y: 0}
	}
	return Point{p.X / mag, p.Y / mag}
}

// AngleTo returns the angle in radians between this vector and another.
func (p Point) AngleTo(other Point) float64 {
	dot := p.Dot(other)
	magProduct := p.Magnitude() * other.Magnitude()
	if magProduct == 0 {
		return 0
	}
	return math.Acos(dot / magProduct)
}

// AngleDegTo returns the angle in degrees between this vector and another.
func (p Point) AngleDegTo(other Point) float64 {
	return p.AngleTo(other) * 180 / math.Pi
}

// ToPolar converts the point to polar coordinates (r, θ in radians).
func (p Point) ToPolar() (r float64, theta float64) {
	r = p.Magnitude()
	theta = math.Atan2(p.Y, p.X)
	return
}

// FromPolar creates a Point from polar coordinates (r, θ in radians).
func FromPolar(r, theta float64) Point {
	return Point{
		X: r * math.Cos(theta),
		Y: r * math.Sin(theta),
	}
}