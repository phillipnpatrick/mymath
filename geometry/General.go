package geometry

import "math"

// --- Helper functions ---

func distance(p1, p2 Point) float64 {
	return math.Hypot(p2.X-p1.X, p2.Y-p1.Y)
}

func rotatePoint(p Point, theta float64) Point {
	cosθ := math.Cos(theta)
	sinθ := math.Sin(theta)
	return Point{
		X: p.X*cosθ - p.Y*sinθ,
		Y: p.X*sinθ + p.Y*cosθ,
	}
}

func rotateAround(p, pivot Point, theta float64) Point {
	// Translate to origin
	translated := Point{X: p.X - pivot.X, Y: p.Y - pivot.Y}
	// Rotate
	rotated := rotatePoint(translated, theta)
	// Translate back
	return Point{X: rotated.X + pivot.X, Y: rotated.Y + pivot.Y}
}
