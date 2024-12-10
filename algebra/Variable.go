package algebra

import (
	"fmt"
	"mymath/basicmath"
)

type Variable struct {
	letter rune
	degree *basicmath.Fraction
}

// #region Constructors

func NewVariable(letter string) *Variable {
	l := rune(letter[0])
	return &Variable{letter: l, degree: basicmath.NewInteger(1)}
}

func NewVariableWithDegree(letter string, degree *basicmath.Fraction) *Variable {
	l := rune(letter[0])
	return &Variable{letter: l, degree: degree}
}

// #endregion

// #region Comparable

func (v *Variable) Equals(other *Variable) bool {
	return v.letter == other.letter &&
		v.degree.Equals(other.degree)
}

// #endregion

// #region LaTeXer

func (v Variable) LaTeX() string {
	if v.degree.Equals(basicmath.NewInteger(0)) {
		return ""
	}

	if v.degree.Equals(basicmath.NewInteger(1)) {
		return string(v.letter)
	}

	if v.degree.IsInteger() {
		return fmt.Sprintf("%s^%s", string(v.letter), v.degree.LaTeX())
	}

	return fmt.Sprintf(`%s^\left(%s\right)`, string(v.letter), v.degree.LaTeX())
}

// #endregion

// #region Stringer

func (v Variable) String() string {
	if v.degree.Equals(basicmath.NewInteger(0)) {
		return ""
	}

	if v.degree.Equals(basicmath.NewInteger(1)) {
		return string(v.letter)
	}

	if v.degree.IsInteger() {
		return fmt.Sprintf("%s^%s", string(v.letter), v.degree)
	}

	return fmt.Sprintf("%s^(%s)", string(v.letter), v.degree)
}

// #endregion

// #region Public Methods

func (v Variable) IsLikeTerm(other Variable) bool {
	return v.letter == other.letter && v.degree.Equals(other.degree)
}

// #endregion