package algebra

import (
	"fmt"
	"mymath/basicmath"
)

type Variable struct {
	letter   rune
	exponent *basicmath.Fraction
}

// #region Constructors

func NewVariable(letter string) *Variable {
	l := rune(letter[0])
	return &Variable{letter: l, exponent: basicmath.NewInteger(1)}
}

func NewVariableWithExponent(letter string, exponent *basicmath.Fraction) *Variable {
	l := rune(letter[0])
	return &Variable{letter: l, exponent: exponent}
}

// #endregion

// #region Properties

func (v *Variable) Letter() string {
	return string(v.letter)
}

func (v *Variable) Exponent() *basicmath.Fraction {
	return v.exponent
}

// #endregion

// #region Comparable

func (v *Variable) Equals(other *Variable) bool {
	return v.letter == other.letter &&
		v.exponent.Equals(other.exponent)
}

// #endregion

// #region LaTeXer

func (v Variable) LaTeX() string {
	if v.exponent.Equals(basicmath.NewInteger(0)) {
		return ""
	}

	if v.exponent.Equals(basicmath.NewInteger(1)) {
		return string(v.letter)
	}

	if v.exponent.IsInteger() {
		return fmt.Sprintf("%s^%s", string(v.letter), v.exponent.LaTeX())
	}

	return fmt.Sprintf(`%s^\left(%s\right)`, string(v.letter), v.exponent.LaTeX())
}

// #endregion

// #region Stringer

func (v Variable) String() string {
	if v.exponent.Equals(basicmath.NewInteger(0)) {
		return ""
	}

	if v.exponent.Equals(basicmath.NewInteger(1)) {
		return string(v.letter)
	}

	if v.exponent.IsInteger() {
		return fmt.Sprintf("%s^%s", string(v.letter), v.exponent)
	}

	return fmt.Sprintf("%s^(%s)", string(v.letter), v.exponent)
}

// #endregion

// #region Public Methods

func (v Variable) IsLikeTerm(other Variable) bool {
	return v.letter == other.letter && v.exponent.Equals(other.exponent)
}

// #endregion
