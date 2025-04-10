package geometry

import (
	"fmt"
	"mymath/basicmath"
	"strings"
)

type Line struct {
	Slope      	basicmath.Fraction
	YIntercept 	basicmath.Fraction
}

func NewLine(a Point, b Point) Line {
	return Line{
		Slope:      getSlope(a, b),
		YIntercept: getYIntercept(a, b),
	}
}

// #region LaTeXer

func (l Line) LaTeX() string {
	if l.YIntercept.LessThan(basicmath.NewInteger(0)) {
		b := strings.ReplaceAll(l.YIntercept.LaTeX(), `-\dfrac`, `- \dfrac`)
		return fmt.Sprintf(`y = %sx %s`, l.Slope.LaTeX(), b)
	}

	return fmt.Sprintf(`y = %sx + %s`, l.Slope.LaTeX(), l.YIntercept.LaTeX())
}

// #endregion

// --- Helper functions ---
func getSlope(a Point, b Point) basicmath.Fraction {
	dy := int(a.Y - b.Y)
	dx := int(a.X - b.X)

	return *basicmath.NewFraction(dy, dx)
}

func getYIntercept(a Point, b Point) basicmath.Fraction {
	m := getSlope(a, b)

	mx1 := m.MultiplyByFactor(-a.X)
	y1 := mx1.PlusFloat(a.Y)

	return *y1
}
