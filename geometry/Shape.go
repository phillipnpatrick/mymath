package geometry

type Shape interface {
	Area() float64
	Perimeter() float64
	Rotate(theta float64)
	RotateAround(pivot Point, theta float64)
}