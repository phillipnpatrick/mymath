package basicmath

import (
	"fmt"
	"strings"
)

// #region Constructor

// Fraction represents a fraction with numerator and denominator.
type Fraction struct {
	n, d int
}

func NewFraction(numerator int, denominator int) *Fraction {
	if denominator < 0 {
		numerator *= -1
		denominator *= -1
	}

	return &Fraction{n: numerator, d: denominator}
}

func NewInteger(value int) *Fraction {
	return &Fraction{n: value, d: 1}
}

// #endregion

// #region Properties

// Numerator returns the numerator of the fraction
func (f *Fraction) Numerator() int {
	return f.n
}

// Denominator returns the denominator of the fraction
func (f *Fraction) Denominator() int {
	return f.d
}

// #endregion

// #region Comparable

func (f *Fraction) Compare(other *Fraction) int {
	if f.n*other.d < other.n*f.d {
		return -1
	} else if f.n*other.d > other.n*f.d {
		return 1
	}

	return 0
}

func (f *Fraction) Equals(other *Fraction) bool {
	// Cross-multiply to compare without floating-point operations
	return f.n*other.d == other.n*f.d
}

func (f *Fraction) GreaterThan(other *Fraction) bool {
	return f.n*other.d > other.n*f.d
}

func (f *Fraction) GreaterThanOrEqualTo(other *Fraction) bool {
	return f.n*other.d >= other.n*f.d
}

func (f *Fraction) LessThan(other *Fraction) bool {
	return f.n*other.d < other.n*f.d
}

func (f *Fraction) LessThanOrEqualTo(other *Fraction) bool {
	return f.n*other.d <= other.n*f.d
}

// #endregion

// #region Factorable

func (f *Fraction) Factor() string {
	factors_n := FactorInt(f.n)
	factors_d := FactorInt(f.d)

	var num strings.Builder
	var denom strings.Builder

	for factor, count := range factors_n {
		for i := 0; i < count; i++ {
			if num.Len() > 0 {
				num.WriteString(" * ")
			}
			num.WriteString(fmt.Sprintf("%d", factor))
		}
	}

	for factor, count := range factors_d {
		for i := 0; i < count; i++ {
			if denom.Len() > 0 {
				denom.WriteString(" * ")
			}
			denom.WriteString(fmt.Sprintf("%d", factor))
		}
	}

	var result string
	if len(factors_d) > 0 {
		result = fmt.Sprintf("(%s)/(%s)", num.String(), denom.String())
	} else {
		result = num.String()
	}

	return result
}

// #endregion

// #region LaTeXer

func (f *Fraction) LaTeX() string {
	if f.d == 1 {
		return fmt.Sprintf("%d", f.n)
	}

	if f.n < 0 || f.d < 0 {
		return fmt.Sprintf(`-\dfrac{%d}{%d}`, Abs(f.n), Abs(f.d))
	}

	return fmt.Sprintf(`\dfrac{%d}{%d}`, f.n, f.d)
}

// #endregion

// #region Operable

func (f *Fraction) Add(others ...*Fraction) *Fraction {
	f.Simplify()
	temp := NewFraction(f.n, f.d)

	for _, other := range others {
		other.Simplify()

		if temp.d == other.d {
			temp.n = temp.n + other.n
			temp.d = other.d
		} else {
			lcm := LCM(temp.d, other.d)
			left := &Fraction{n: lcm / temp.d, d: lcm / temp.d}
			right := &Fraction{n: lcm / other.d, d: lcm / other.d}

			temp.n = temp.n*left.n + other.n*right.n
			temp.d = temp.d * left.d
		}

		temp.Simplify()
	}

	return temp
}

func (f *Fraction) Subtract(others ...*Fraction) *Fraction {
	temp := NewFraction(f.n, f.d)

	for _, other := range others {
		f1 := other.Multiply(NewInteger(-1))

		temp = temp.Add(f1)
	}

	return temp
}

func (f *Fraction) Multiply(others ...*Fraction) *Fraction {
	temp := NewFraction(f.n, f.d)

	for _, other := range others {
		temp.n = temp.n * other.n
		temp.d = temp.d * other.d
		temp.Simplify()
	}

	return temp
}

func (f *Fraction) Divide(others ...*Fraction) *Fraction {
	temp := NewFraction(f.n, f.d)

	for _, other := range others {
		f1 := NewFraction(other.d, other.n)
		temp = temp.Multiply(f1)
	}

	return temp
}

// #endregion

// #region Simplifiable

func (f *Fraction) Simplify() {
	g := GCF(f.n, f.d)
	f.n /= g
	f.d /= g
}

// #endregion

// #region Stringer

func (f *Fraction) String() string {
	if f.d == 1 {
		return fmt.Sprintf("%d", f.n)
	}

	return fmt.Sprintf("%d/%d", f.n, f.d)
}

// #endregion

// #region Public Methods

func (f *Fraction) IsInteger() bool {
	return f.d == 1
}

// #endregion

// #region Private Methods

// #endregion
