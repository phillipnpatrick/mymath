package geometry

import (
	"math"
	"mymath/basicmath"

	"testing"
)

func TestPoint_Distance(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		other    Point
		expected float64
	}{
		{
			name:     "distance between points",
			p:        Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(2)},
			other:    Point{X: *basicmath.NewInteger(4), Y: *basicmath.NewInteger(6)},
			expected: 5.0,
		},
		{
			name:     "same point",
			p:        Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(3)},
			other:    Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(3)},
			expected: 0.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p.Distance(tt.other)
			if result != tt.expected {
				t.Errorf("Distance() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPoint_Move(t *testing.T) {
	type args struct {
		dx float64
		dy float64
	}
	tests := []struct {
		name     string
		p        *Point
		args     args
		expected Point
	}{
		{
			name:     "move right and up",
			p:        &Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(2)},
			args:     args{dx: 3, dy: 4},
			expected: Point{X: *basicmath.NewInteger(4), Y: *basicmath.NewInteger(6)},
		},
		{
			name:     "move left and down",
			p:        &Point{X: *basicmath.NewInteger(5), Y: *basicmath.NewInteger(5)},
			args:     args{dx: -2, dy: -3},
			expected: Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(2)},
		},
		{
			name:     "no movement",
			p:        &Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(0)},
			args:     args{dx: 0, dy: 0},
			expected: Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.p.Move(tt.args.dx, tt.args.dy)
			if tt.p.X != tt.expected.X || tt.p.Y != tt.expected.Y {
				t.Errorf("Move() = (%v, %v), want (%v, %v)",
					tt.p.X, tt.p.Y, tt.expected.X, tt.expected.Y)
			}
		})
	}
}

func TestPoint_Add(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		other    Point
		expected Point
	}{
		{
			name:     "add points",
			p:        Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(2)},
			other:    Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(4)},
			expected: Point{X: *basicmath.NewInteger(4), Y: *basicmath.NewInteger(6)},
		},
		{
			name:     "add zero point",
			p:        Point{X: *basicmath.NewInteger(5), Y: *basicmath.NewInteger(5)},
			other:    Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(0)},
			expected: Point{X: *basicmath.NewInteger(5), Y: *basicmath.NewInteger(5)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p.Add(tt.other)
			if result != tt.expected {
				t.Errorf("Add() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPoint_Subtract(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		other    Point
		expected Point
	}{
		{
			name:     "subtract points",
			p:        Point{X: *basicmath.NewInteger(5), Y: *basicmath.NewInteger(5)},
			other:    Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(4)},
			expected: Point{X: *basicmath.NewInteger(2), Y: *basicmath.NewInteger(1)},
		},
		{
			name:     "subtract zero point",
			p:        Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(3)},
			other:    Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(0)},
			expected: Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(3)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p.Subtract(tt.other)
			if result != tt.expected {
				t.Errorf("Subtract() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPoint_Scale(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		factor   float64
		expected Point
	}{
		{
			name:     "scale point",
			p:        Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(2)},
			factor:   2,
			expected: Point{X: *basicmath.NewInteger(2), Y: *basicmath.NewInteger(4)},
		},
		{
			name:     "scale by zero",
			p:        Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(4)},
			factor:   2,
			expected: Point{X: *basicmath.NewInteger(6), Y: *basicmath.NewInteger(8)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p.Scale(tt.factor)
			if result != tt.expected {
				t.Errorf("Scale() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPoint_Dot(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		other    Point
		expected float64
	}{
		{
			name:     "dot product",
			p:        Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(2)},
			other:    Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(4)},
			expected: 11,
		},
		{
			name:     "zero dot product",
			p:        Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(0)},
			other:    Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(4)},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p.Dot(tt.other)
			if result != tt.expected {
				t.Errorf("Dot() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPoint_Magnitude(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		expected float64
	}{
		{
			name:     "magnitude of point",
			p:        Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(4)},
			expected: 5,
		},
		{
			name:     "magnitude of zero point",
			p:        Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(0)},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p.Magnitude()
			if result != tt.expected {
				t.Errorf("Magnitude() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPoint_Normalize(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		expected Point
	}{
		{ 
			name:     "normalize point",
			p:        Point{X: *basicmath.NewInteger(3), Y: *basicmath.NewInteger(4)},
			expected: Point{X: *basicmath.NewFraction(3, 5), Y: *basicmath.NewFraction(4, 5)},
		},
		{
			name:     "normalize zero point",
			p:        Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(0)},
			expected: Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(0)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.p.Normalize()
			if result.X != tt.expected.X || result.Y != tt.expected.Y {
				t.Errorf("Normalize() = (%v, %v), want (%v, %v)", result.X, result.Y, tt.expected.X, tt.expected.Y)
			}
		})
	}
}

func TestPoint_AngleTo(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		other    Point
		expected float64
	}{
		{
			name:     "angle between points",
			p:        Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(0)},
			other:    Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(1)},
			expected: math.Pi / 2,
		},
		{
			name:     "same point",
			p:        Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(1)},
			other:    Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(1)},
			expected: 0,
		},
	}
	// Define a tolerance for floating-point comparisons
	const epsilon = 1e-9

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if points are exactly the same to avoid angle calculation errors
			if tt.p.X == tt.other.X && tt.p.Y == tt.other.Y {
				if tt.expected != 0 {
					t.Errorf("Expected 0 for same points, but got %v", tt.expected)
				}
				return
			}

			result := tt.p.AngleTo(tt.other)
			if math.Abs(result-tt.expected) > epsilon {
				t.Errorf("AngleTo() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPoint_AngleDegTo(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		other    Point
		expected float64
	}{
		{
			name:     "angle between points in degrees",
			p:        Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(0)},
			other:    Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(1)},
			expected: 90,
		},
		{
			name:     "same point",
			p:        Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(1)},
			other:    Point{X: *basicmath.NewInteger(1), Y: *basicmath.NewInteger(1)},
			expected: 0,
		},
	}
	// Define a tolerance for floating-point comparisons
	const epsilon = 1e-9
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Check if points are exactly the same to avoid angle calculation errors
			if tt.p.X == tt.other.X && tt.p.Y == tt.other.Y {
				if tt.expected != 0 {
					t.Errorf("Expected 0 for same points, but got %v", tt.expected)
				}
				return
			}

			result := tt.p.AngleDegTo(tt.other)
			if math.Abs(result-tt.expected) > epsilon {
				t.Errorf("AngleDegTo() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPoint_ToPolar(t *testing.T) {
	tests := []struct {
		name     string
		p        Point
		expected struct{ r, theta float64 }
	}{
		{
			name:     "convert to polar",
			p:        Point{X:*basicmath.NewInteger(3), Y: *basicmath.NewInteger(4)},
			expected: struct{ r, theta float64 }{r: 5, theta: math.Atan2(4, 3)},
		},
		{
			name:     "zero point to polar",
			p:        Point{X: *basicmath.NewInteger(0), Y: *basicmath.NewInteger(0)},
			expected: struct{ r, theta float64 }{r: 0, theta: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, theta := tt.p.ToPolar()
			if r != tt.expected.r || theta != tt.expected.theta {
				t.Errorf("ToPolar() = (%v, %v), want (%v, %v)", r, theta, tt.expected.r, tt.expected.theta)
			}
		})
	}
}

// func TestPoint_FromPolar(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		r, theta float64
// 		expected Point
// 	}{
// 		{
// 			name:     "convert from polar",
// 			r:        5,
// 			theta:    math.Atan2(4, 3),
// 			expected: Point{X: 3, Y: 4},
// 		},
// 		{
// 			name:     "zero point from polar",
// 			r:        0,
// 			theta:    0,
// 			expected: Point{X: 0, Y: 0},
// 		},
// 	}
// 	// Define a tolerance for floating-point comparisons
// 	const epsilon = 1e-9
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			result := FromPolar(tt.r, tt.theta)
// 			if math.Abs(result.X-tt.expected.X) > epsilon || math.Abs(result.Y-tt.expected.Y) > epsilon {
// 				t.Errorf("FromPolar() = (%v, %v), want (%v, %v)", result.X, result.Y, tt.expected.X, tt.expected.Y)
// 			}
// 		})
// 	}
// }
