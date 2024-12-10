package algebra

import "mymath/basicmath"

type Exponent struct {
	*basicmath.Fraction
}

// #region Constructors

func NewExponent(exponent *basicmath.Fraction) *Exponent {
	return &Exponent{Fraction: basicmath.NewFraction(exponent.Numerator(), exponent.Denominator())}
}

// #endregion