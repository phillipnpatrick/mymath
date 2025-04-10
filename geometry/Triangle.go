package geometry

import (
	"math"
)

type Triangle struct {
	A Point
	B Point
	C Point
}

// Area computes the area using the shoelace formula
func (t Triangle) Area() float64 {
	return 0.5 * math.Abs(
		t.A.X*(t.B.Y-t.C.Y) +
			t.B.X*(t.C.Y-t.A.Y) +
			t.C.X*(t.A.Y-t.B.Y),
	)
}

// Perimeter computes the perimeter of the triangle
func (t Triangle) Perimeter() float64 {
	return distance(t.A, t.B) + distance(t.B, t.C) + distance(t.C, t.A)
}

// Rotate rotates the triangle around the origin (0,0) by angle θ (in radians)
func (t *Triangle) Rotate(theta float64) {
	t.A = rotatePoint(t.A, theta)
	t.B = rotatePoint(t.B, theta)
	t.C = rotatePoint(t.C, theta)
}

// RotateAround rotates the triangle around a given pivot point by angle θ (in radians)
func (t *Triangle) RotateAround(pivot Point, theta float64) {
	t.A = rotateAround(t.A, pivot, theta)
	t.B = rotateAround(t.B, pivot, theta)
	t.C = rotateAround(t.C, pivot, theta)
}

type EquilateralTriangle struct {
	Triangle
	SideLength float64
}

func NewEquilateralTriangle(start Point, sideLength float64) EquilateralTriangle {
	// Create the 3 points of the triangle, assuming horizontal base
	a := start
	b := Point{X: start.X + sideLength, Y: start.Y}
	height := (math.Sqrt(3) / 2) * sideLength
	c := Point{X: start.X + sideLength/2, Y: start.Y + height}

	return EquilateralTriangle{
		Triangle:   Triangle{A: a, B: b, C: c},
		SideLength: sideLength,
	}
}