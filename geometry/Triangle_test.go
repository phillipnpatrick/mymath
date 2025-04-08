package geometry

import (
	"math"
	"testing"
)

func floatEquals(a, b, epsilon float64) bool {
	return math.Abs(a-b) <= epsilon
}

func pointsEqual(p1, p2 Point, epsilon float64) bool {
	return floatEquals(p1.X, p2.X, epsilon) && floatEquals(p1.Y, p2.Y, epsilon)
}

func TestTriangle_Area(t *testing.T) {
	tests := []struct {
		name     string
		triangle Triangle
		want     float64
		eps      float64
	}{
		{
			name: "right triangle",
			triangle: Triangle{
				A: Point{0, 0},
				B: Point{4, 0},
				C: Point{0, 3},
			},
			want: 6,
			eps:  1e-9,
		},
		{
			name: "equilateral triangle side 2",
			triangle: Triangle{
				A: Point{0, 0},
				B: Point{2, 0},
				C: Point{1, math.Sqrt(3)},
			},
			want: math.Sqrt(3),
			eps:  1e-9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.triangle.Area()
			if !floatEquals(got, tt.want, tt.eps) {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangle_Perimeter(t *testing.T) {
	tests := []struct {
		name     string
		triangle Triangle
		want     float64
		eps      float64
	}{
		{
			name: "right triangle 3-4-5",
			triangle: Triangle{
				A: Point{0, 0},
				B: Point{4, 0},
				C: Point{0, 3},
			},
			want: 12,
			eps:  1e-9,
		},
		{
			name: "equilateral triangle side 2",
			triangle: Triangle{
				A: Point{0, 0},
				B: Point{2, 0},
				C: Point{1, math.Sqrt(3)},
			},
			want: 2 + 2 + 2,
			eps:  1e-9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.triangle.Perimeter()
			if !floatEquals(got, tt.want, tt.eps) {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangle_RotateAround(t *testing.T) {
	tests := []struct {
		name     string
		triangle Triangle
		pivot    Point
		angle    float64 // in radians
		want     Triangle
		eps      float64
	}{
		{
			name: "rotate 90 degrees around origin",
			triangle: Triangle{
				A: Point{1, 0},
				B: Point{2, 0},
				C: Point{1, 1},
			},
			pivot: Point{0, 0},
			angle: math.Pi / 2, // 90 degrees
			want: Triangle{
				A: Point{0, 1},
				B: Point{0, 2},
				C: Point{-1, 1},
			},
			eps: 1e-9,
		},
		{
			name: "rotate 180 degrees around (1,1)",
			triangle: Triangle{
				A: Point{0, 0},
				B: Point{2, 0},
				C: Point{1, 2},
			},
			pivot: Point{1, 1},
			angle: math.Pi, // 180 degrees
			want: Triangle{
				A: Point{2, 2},
				B: Point{0, 2},
				C: Point{1, 0},
			},
			eps: 1e-9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tri := tt.triangle
			tri.RotateAround(tt.pivot, tt.angle)
			if !pointsEqual(tri.A, tt.want.A, tt.eps) {
				t.Errorf("A = %+v, want %+v", tri.A, tt.want.A)
			}
			if !pointsEqual(tri.B, tt.want.B, tt.eps) {
				t.Errorf("B = %+v, want %+v", tri.B, tt.want.B)
			}
			if !pointsEqual(tri.C, tt.want.C, tt.eps) {
				t.Errorf("C = %+v, want %+v", tri.C, tt.want.C)
			}
		})
	}
}

func TestNewEquilateralTriangle(t *testing.T) {
	type args struct {
		start      Point
		sideLength float64
	}
	tests := []struct {
		name string
		args args
		want EquilateralTriangle
		eps   float64
	}{
		{
			name: "origin centered triangle",
			args: args{
				start:      Point{0, 0},
				sideLength: 2,
			},
			want: EquilateralTriangle{
				Triangle: Triangle{
					A: Point{0, 0},
					B: Point{2, 0},
					C: Point{1, math.Sqrt(3)},
				},
				SideLength: 2,
			},
			eps: 1e-9,
		},
		{
			name: "offset triangle",
			args: args{
				start:      Point{1, 1},
				sideLength: 4,
			},
			want: EquilateralTriangle{
				Triangle: Triangle{
					A: Point{1, 1},
					B: Point{5, 1},
					C: Point{3, 1 + (math.Sqrt(3) * 2)},
				},
				SideLength: 4,
			},
			eps: 1e-9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewEquilateralTriangle(tt.args.start, tt.args.sideLength)

			if !pointsEqual(got.A, tt.want.A, tt.eps) {
				t.Errorf("A = %+v, want %+v", got.A, tt.want.A)
			}
			if !pointsEqual(got.B, tt.want.B, tt.eps) {
				t.Errorf("B = %+v, want %+v", got.B, tt.want.B)
			}
			if !pointsEqual(got.C, tt.want.C, tt.eps) {
				t.Errorf("C = %+v, want %+v", got.C, tt.want.C)
			}
			if !floatEquals(got.SideLength, tt.want.SideLength, tt.eps) {
				t.Errorf("SideLength = %v, want %v", got.SideLength, tt.want.SideLength)
			}
		})
	}
}
