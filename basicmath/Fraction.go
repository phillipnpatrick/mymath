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

// FractionOptions gives public access to the parts of a fraction
type FractionOptions struct {
	Numerator   int
	Denominator int
}

// FractionOption 
type FractionOption func(*FractionOptions)

func WithNumerator(n int) FractionOption {
	return func(fo *FractionOptions) {
		fo.Numerator = n
	}
}

func WithDenominator(d int) FractionOption {
	return func(fo *FractionOptions) {
		fo.Denominator = d
	}
}

func NewFraction(options ...FractionOption) *Fraction {
	opts := &FractionOptions{}
	for _, option := range options {
		option(opts)
	}

	if opts.Denominator == 0 {
		opts.Denominator = 1
	}

	fraction := &Fraction{
		n: opts.Numerator,
		d: opts.Denominator,
	}

	return fraction
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
	return fmt.Sprintf(`\dfrac{%d}{%d}`, f.n, f.d)
}

// #endregion

// #region Operable
func (f *Fraction) Add(T ...any) any {
	if f == nil {
		return nil // Return a default value or error to prevent nil dereference
	}
	for _, value := range T {
		if f2, ok := value.(*Fraction); ok {
			var temp *Fraction

			f.Simplify()
			f2.Simplify()

			if f.d == f2.d {
				temp = &Fraction{
					n: f.n + f2.n,
					d: f.d,
				}
			} else {
				lcm := LCM(f.d, f2.d)
				left := &Fraction{n: lcm/f.d, d: lcm/f.d}
				right := &Fraction{n: lcm/f2.d}
				
				temp = &Fraction{
					n: f.n*left.n + f2.n*right.n,
					d: f.d*left.d,
				}
			}

			temp.Simplify()
			
			return temp
		}
	}
	return nil
}

func (f *Fraction) Subtract(T ...any) any {
	if f == nil {
		return nil // Return a default value or error to prevent nil dereference
	}
	for _, value := range T {
		if f2, ok := value.(*Fraction); ok {
			f1 := f2.Multiply(NewFraction(WithNumerator(-1)))

			return f.Add(f1)
		}
	}
	return nil
}

func (f *Fraction) Multiply(T ...any) any {
	if f == nil {
		return nil // Return a default value or error to prevent nil dereference
	}
	for _, value := range T {
		if f2, ok := value.(*Fraction); ok {
			temp := &Fraction{
				n: f.n * f2.n,
				d: f.d * f2.d,
			}
			temp.Simplify()
			return temp
		}
	}
	return nil
}

func (f *Fraction) Divide(T ...any) any {
	if f == nil {
		return nil // Return a default value or error to prevent nil dereference
	}
	for _, value := range T {
		if f2, ok := value.(*Fraction); ok {
			f1 := NewFraction(WithNumerator(f2.d), WithDenominator(f2.n))
			
			return f.Multiply(f1)
		}
	}
	return nil
}

// #endregion

// #region Simplifiable

func (f *Fraction) Simplify() {
	factors_n := FactorInt(f.n)
	factors_d := FactorInt(f.d)

	for factor, count_n := range factors_n {
		count := Min(count_n, factors_d[factor])
		for i := 0; i < count; i++ {
			factors_n[factor]--
			factors_d[factor]--

			if factors_n[factor] == 0 {
				delete(factors_n, factor)
			}

			if factors_d[factor] == 0 {
				delete(factors_d, factor)
			}
		}
	}

	f.n = 1
	for factor, count := range factors_n {
		for i := 0; i < count; i++ {
			f.n *= factor
		}
	}

	f.d = 1
	for factor, count := range factors_d {
		for i := 0; i < count; i++ {
			f.d *= factor
		}
	}
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

// #region Private Methods

// #endregion