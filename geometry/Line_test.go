package geometry

import "testing"

func TestLine_LaTeX(t *testing.T) {
	tests := []struct {
		name string
		l    Line
		want string
	}{
		{
			name: "Line_LaTeX",
			l: NewLine(Point{X: 2, Y:1}, Point{X:8, Y:6}),
			want: `y = \dfrac{5}{6}x - \dfrac{2}{3}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.LaTeX(); got != tt.want {
				t.Errorf("Line.LaTeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
