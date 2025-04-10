package geometry

import (
	"fmt"
	"mymath/basicmath"
)

type Line struct {
	Slope      basicmath.Fraction
	YIntercept float64
}

func NewLine(a Point, b Point) Line {
	return Line{
		Slope:      getSlope(a, b),
		YIntercept: getYIntercept(a, b),
	}
}

// #region LaTeXer

func (l *Line) LaTeX() string {
	if l.YIntercept < 0 {
		return fmt.Sprintf(`y = %vx - %v`, l.Slope, l.YIntercept)
	}
	return fmt.Sprintf(`y = %vx + %v`, l.Slope, l.YIntercept)
}

// #endregion

// --- Helper functions ---
func getSlope(a Point, b Point) basicmath.Fraction {
	dy := int(a.Y - b.Y)
	dx := int(a.X - b.X)

	return *basicmath.NewFraction(dx, dy)
}

func getYIntercept(a Point, b Point) float64 {
	m := getSlope(a, b)

	temp := m.MultiplyByFactor(a.X)

	return temp.PlusFloat(a.Y).ToFloat64()
}
