package geometry

type Rectangle struct {
	A Point
	B Point
	C Point
	D Point
}

// Rectangle computes the perimeter of the rectangle
func (r Rectangle) Area() float64 {
	return distance(r.A, r.B) * distance(r.B, r.C)
}

// Perimeter computes the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (distance(r.A, r.B) + distance(r.B, r.C))
}

// Rotate rotates the triangle around the origin (0,0) by angle θ (in radians)
func (r *Rectangle) Rotate(theta float64) {
	r.A = rotatePoint(r.A, theta)
	r.B = rotatePoint(r.B, theta)
	r.C = rotatePoint(r.C, theta)
	r.D = rotatePoint(r.D, theta)
}

// RotateAround rotates the triangle around a given pivot point by angle θ (in radians)
func (r *Rectangle) RotateAround(pivot Point, theta float64) {
	r.A = rotateAround(r.A, pivot, theta)
	r.B = rotateAround(r.B, pivot, theta)
	r.C = rotateAround(r.C, pivot, theta)
	r.D = rotateAround(r.D, pivot, theta)
}

type Square struct {
	Rectangle
	SideLength float64
}

func NewSquare(start Point, sideLength float64) Square {
	// Create the 3 points of the triangle, assuming horizontal base
	a := start
	b := Point{X: start.X + sideLength, Y: start.Y}
	c := Point{X: b.X, Y: b.Y + sideLength}
	d := Point{X: start.X, Y: start.Y + sideLength}

	return Square{
		Rectangle:   Rectangle{A: a, B: b, C: c, D: d},
		SideLength: sideLength,
	}
}