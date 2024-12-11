package algebra

import (
	"fmt"
	"mymath/basicmath"
	"reflect"
	"testing"
)

func TestPolynomial_Equals(t *testing.T) {
	type args struct {
		other *Polynomial
	}
	tests := []struct {
		name string
		p    *Polynomial
		args args
		want bool
	}{
		{ // 9x^2 + 4x^3
			name: "Polynomial_Equals_Test01",
			p: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(3))),
			args: args{other: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(3)),
				NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2))),
			},
			want: true,
		},
		{ // 5x^2 - 12y^3
			name: "Polynomial_Equals_Test02",
			p: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(-12), "y", basicmath.NewInteger(3))),
			args: args{other: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(-12), "y", basicmath.NewInteger(3))),
			},
			want: true,
		},
		{ // 5x^2 - 12y^3
			name: "Polynomial_Equals_Test03",
			p: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(-12), "y", basicmath.NewInteger(3))),
			args: args{other: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(3)),
				NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2))),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Equals(tt.args.other); got != tt.want {
				t.Errorf("Polynomial.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolynomial_LaTeX(t *testing.T) {
	tests := []struct {
		name string
		p    Polynomial
		want string
	}{
		{ // 9x^2 + 4x^3
			name: "Polynomial_LaTeX_Test01",
			p: *NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(3))),
			want: "9x^2 + 4x^3",
		},
		{ // 5x^2 - 12y^3
			name: "Polynomial_LaTeX_Test02",
			p: *NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(-12), "y", basicmath.NewInteger(3))),
			want: "5x^2 - 12y^3",
		},
		{ // 14x^3
			name: "Polynomial_LaTeX_Test03",
			p: *NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(14), "x", basicmath.NewInteger(3))),
			want: "14x^3",
		},
		{ // x^2 + 3x^3 + 5
			name: "Polynomial_LaTeX_Test04",
			p: *NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(1), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(3), "x", basicmath.NewInteger(3)),
				NewMonomialConstant(basicmath.NewInteger(5)),
			),
			want: "x^2 + 3x^3 + 5",
		},
		{ // 14x^2 + 11x + 4y
			name: "Polynomial_LaTeX_Test05",
			p: *NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(14), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(11), "x"),
				NewMonomial(basicmath.NewInteger(4), "y"),
			),
			want: "14x^2 + 11x + 4y",
		},
		{ // 7xy + 20x^2 + 5y^2
			name: "Polynomial_LaTeX_Test06",
			p: *NewPolynomial(
				NewMonomialWithVariables(basicmath.NewInteger(7), NewVariable("x"), NewVariable("y")),
				NewMonomialWithExponent(basicmath.NewInteger(20), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(5), "y", basicmath.NewInteger(2)),
			),
			want: "7xy + 20x^2 + 5y^2",
		},
		{ // 23m – 11n
			name: "Polynomial_LaTeX_Test07",
			p: *NewPolynomial(
				NewMonomial(basicmath.NewInteger(23), "m"),
				NewMonomial(basicmath.NewInteger(-11), "n"),
			),
			want: "23m - 11n",
		},
		{ // x^2 + 4x + 5
			name: "Polynomial_LaTeX_Test08",
			p: *NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(1), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(4), "x"),
				NewMonomialConstant(basicmath.NewInteger(5)),
			),
			want: "x^2 + 4x + 5",
		},
		{ // 6x^3 – 14x^2 + 7x – 6
			name: "Polynomial_LaTeX_Test09",
			p: *NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(6), "x", basicmath.NewInteger(3)),
				NewMonomialWithExponent(basicmath.NewInteger(-14), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(7), "x"),
				NewMonomialConstant(basicmath.NewInteger(-6)),
			),
			want: "6x^3 - 14x^2 + 7x - 6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.LaTeX(); got != tt.want {
				t.Errorf("Polynomial.LaTeX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolynomial_StandardForm(t *testing.T) {
	tests := []struct {
		name string
		p    *Polynomial
		want *Polynomial
	}{
		{ // 3x + 5x^2 - 4 = 5x^2 + 3x - 4
			name: "Polynomial_StandardForm_Test01",
			p: NewPolynomial(
				NewMonomial(basicmath.NewInteger(3), "x"),
				NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(2)),
				NewMonomialConstant(basicmath.NewInteger(-4)),
			),
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(3), "x"),
				NewMonomialConstant(basicmath.NewInteger(-4)),
			),
		},
		{ // 2x^2 + 3x + x^2 - 4 = 3x^2 + 3x - 4
			name: "Polynomial_StandardForm_Test02",
			p: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(2), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(3), "x"),
				NewMonomialWithExponent(basicmath.NewInteger(1), "x", basicmath.NewInteger(2)),
				NewMonomialConstant(basicmath.NewInteger(-4)),
			),
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(3), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(3), "x"),
				NewMonomialConstant(basicmath.NewInteger(-4)),
			),
		},
		{ // (1/2)x^3 + (3/4)x^2 - (1/4)x^2 = (1/2)x^3 + (1/2)x^2
			name: "Polynomial_StandardForm_Test03",
			p: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewFraction(1, 2), "x", basicmath.NewInteger(3)),
				NewMonomialWithExponent(basicmath.NewFraction(3, 4), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewFraction(-1, 4), "x", basicmath.NewInteger(2)),
			),
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewFraction(1, 2), "x", basicmath.NewInteger(3)),
				NewMonomialWithExponent(basicmath.NewFraction(1, 2), "x", basicmath.NewInteger(2)),
			),
		},
		{ // x^2 - 2x + 1 - x^2 + x = -x + 1
			name: "Polynomial_StandardForm_Test04",
			p: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(1), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(-2), "x"),
				NewMonomialConstant(basicmath.NewInteger(1)),
				NewMonomialWithExponent(basicmath.NewInteger(-1), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(1), "x"),
			),
			want: NewPolynomial(
				NewMonomial(basicmath.NewInteger(-1), "x"),
				NewMonomialConstant(basicmath.NewInteger(1)),
			),
		},
		{ // 2x^2y + 3xy^2 - 5x^2y + 4 = 3xy^2 - 3x^2y + 4	(Ordered by degree: degree 3, then degree 2, and constant.)
			name: "Polynomial_StandardForm_Test05",
			p: NewPolynomial(
				NewMonomialWithVariables(basicmath.NewInteger(2), NewVariableWithExponent("x", basicmath.NewInteger(2)), NewVariable("y")),
				NewMonomialWithVariables(basicmath.NewInteger(3), NewVariable("x"), NewVariableWithExponent("y", basicmath.NewInteger(2))),
				NewMonomialWithVariables(basicmath.NewInteger(-5), NewVariableWithExponent("x", basicmath.NewInteger(2)), NewVariable("y")),
				NewMonomialConstant(basicmath.NewInteger(4)),
			),
			want: NewPolynomial(
				NewMonomialWithVariables(basicmath.NewInteger(-3), NewVariableWithExponent("x", basicmath.NewInteger(2)), NewVariable("y")),
				NewMonomialWithVariables(basicmath.NewInteger(3), NewVariable("x"), NewVariableWithExponent("y", basicmath.NewInteger(2))),				
				NewMonomialConstant(basicmath.NewInteger(4)),
			),
		},
		{ // x^2y + xyz^2 - 3xy^2z + 5xyz^2 - 2x^2y = 6xyz^2 - 3xy^2z -x^2y	(Ordered by degree: 5, then 5, and finally 3.)
			name: "Polynomial_StandardForm_Test06",
			p: NewPolynomial(
				NewMonomialWithVariables(basicmath.NewInteger(1),
					NewVariableWithExponent("x", basicmath.NewInteger(2)),
					NewVariable("y")),
				NewMonomialWithVariables(
					basicmath.NewInteger(1),
					NewVariable("x"),
					NewVariable("y"),
					NewVariableWithExponent("z", basicmath.NewInteger(2))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-3),
					NewVariable("x"),
					NewVariableWithExponent("y", basicmath.NewInteger(2)),
					NewVariable("z")),
				NewMonomialWithVariables(
					basicmath.NewInteger(5),
					NewVariable("x"),
					NewVariable("y"),
					NewVariableWithExponent("z", basicmath.NewInteger(2))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-2),
					NewVariableWithExponent("x", basicmath.NewInteger(2)),
					NewVariable("y")),
			),
			want: NewPolynomial(
				NewMonomialWithVariables(
					basicmath.NewInteger(6),
					NewVariable("x"),
					NewVariable("y"),
					NewVariableWithExponent("z", basicmath.NewInteger(2))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-3),
					NewVariable("x"),
					NewVariableWithExponent("y", basicmath.NewInteger(2)),
					NewVariable("z")),
				NewMonomialWithVariables(
					basicmath.NewInteger(-1),
					NewVariableWithExponent("x", basicmath.NewInteger(2)),
					NewVariable("y")),
			),
		},
		{ // x^3y^2 + x^2y^3 - xy^4 + 2x^3y^2 - y^5 + 4 = 3x^3y^2 + x^2y^3 - xy^4 - y^5 + 4	(Ordered by degree: 5 ,5, 4, 5, and constant)
			name: "Polynomial_StandardForm_Test07",
			p: NewPolynomial(
				NewMonomialWithVariables(
					basicmath.NewInteger(1),
					NewVariableWithExponent("x", basicmath.NewInteger(3)),
					NewVariableWithExponent("y", basicmath.NewInteger(2))),
				NewMonomialWithVariables(
					basicmath.NewInteger(1),
					NewVariableWithExponent("x", basicmath.NewInteger(2)),
					NewVariableWithExponent("y", basicmath.NewInteger(3))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-1),
					NewVariable("x"),
					NewVariableWithExponent("y", basicmath.NewInteger(4))),
				NewMonomialWithVariables(
					basicmath.NewInteger(2),
					NewVariableWithExponent("x", basicmath.NewInteger(3)),
					NewVariableWithExponent("y", basicmath.NewInteger(2))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-1),
					NewVariableWithExponent("y", basicmath.NewInteger(5))),
				NewMonomialConstant(basicmath.NewInteger(4)),
			),
			want: NewPolynomial(
				NewMonomialWithVariables(
					basicmath.NewInteger(3),
					NewVariableWithExponent("x", basicmath.NewInteger(3)),
					NewVariableWithExponent("y", basicmath.NewInteger(2))),
				NewMonomialWithVariables(
					basicmath.NewInteger(1),
					NewVariableWithExponent("x", basicmath.NewInteger(2)),
					NewVariableWithExponent("y", basicmath.NewInteger(3))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-1),
					NewVariable("x"),
					NewVariableWithExponent("y", basicmath.NewInteger(4))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-1),
					NewVariableWithExponent("y", basicmath.NewInteger(5))),
				NewMonomialConstant(basicmath.NewInteger(4)),
			),
		},
		{ // 4x^2z^3 - 2xyz + 3x^3y^2 - 5x^2z^3 + xy^2z^2 = 3x^3y^2 - x^2z^3 + xy^2z^2 - 2xyz	(Ordered by degree: 5, 5, 5, and 3)
			name: "Polynomial_StandardForm_Test08",
			p: NewPolynomial(
				NewMonomialWithVariables(
					basicmath.NewInteger(4),
					NewVariableWithExponent("x", basicmath.NewInteger(2)),
					NewVariableWithExponent("z", basicmath.NewInteger(3))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-2),
					NewVariable("x"),
					NewVariable("y"),
					NewVariable("z")),
				NewMonomialWithVariables(
					basicmath.NewInteger(3),
					NewVariableWithExponent("x", basicmath.NewInteger(3)),
					NewVariableWithExponent("y", basicmath.NewInteger(2))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-5),
					NewVariableWithExponent("x", basicmath.NewInteger(2)),
					NewVariableWithExponent("z", basicmath.NewInteger(3))),
				NewMonomialWithVariables(
					basicmath.NewInteger(1),
					NewVariable("x"),
					NewVariableWithExponent("y", basicmath.NewInteger(2)),
					NewVariableWithExponent("z", basicmath.NewInteger(2))),
			),
			want: NewPolynomial(
				NewMonomialWithVariables(
					basicmath.NewInteger(3),
					NewVariableWithExponent("x", basicmath.NewInteger(3)),
					NewVariableWithExponent("y", basicmath.NewInteger(2))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-1),
					NewVariableWithExponent("x", basicmath.NewInteger(2)),
					NewVariableWithExponent("z", basicmath.NewInteger(3))),
				NewMonomialWithVariables(
					basicmath.NewInteger(1),
					NewVariable("x"),
					NewVariableWithExponent("y", basicmath.NewInteger(2)),
					NewVariableWithExponent("z", basicmath.NewInteger(2))),
				NewMonomialWithVariables(
					basicmath.NewInteger(-2),
					NewVariable("x"),
					NewVariable("y"),
					NewVariable("z")),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.StandardForm()
			if !reflect.DeepEqual(got, tt.want) {
				compareValues("monomials", got.monomials, tt.want.monomials)
				if len(got.monomials) != len(tt.want.monomials) {
					fmt.Printf("monomial count not equal: got %d, tt.want %d\n", len(got.monomials), len(tt.want.monomials))
				} else {
					for i := 0; i < len(got.monomials); i++ {
						a := got.monomials[i]
						b := tt.want.monomials[i]
					
						compareValues(fmt.Sprintf("%d: coefficient", i), a.coefficient, b.coefficient)

						if !a.coefficient.Equals(b.coefficient) {
							fmt.Printf("got.monomials[%d].coefficient (%v) not equal to tt.want.monomials[%d].coefficient (%v)\n", 
							i, got.monomials[i].coefficient, i, tt.want.monomials[i].coefficient)
						}

						compareValues(fmt.Sprintf("%d: degree", i), a.degree, b.degree)

						if !a.degree.Equals(b.degree) {
							fmt.Printf("got.monomials[%d].degree (%v) not equal to tt.want.monomials[%d].degree (%v)\n", 
							i, got.monomials[i].degree, i, tt.want.monomials[i].degree)
						}

						if len(a.variables) != len(b.variables) {
							fmt.Printf("variable count not equal: got %d, tt.want %d\n", len(a.variables), len(b.variables))
						} else {
							compareValues("variables", a.variables, b.variables)
							for j := 0; j < len(a.variables); j++ {
								x := a.variables[j]
								y := a.variables[j]

								compareValues(fmt.Sprintf("%d, %d: letter", i, j), x.letter, y.letter)

								if x.letter != y.letter {
									fmt.Printf("got.monomials[%d].variables[%d].letter (%s) not equal to tt.want.monomials[%d].variables[%d].letter (%s) \n", 
									i, j, string(got.monomials[i].variables[j].letter),
									i, j, string(tt.want.monomials[i].variables[j].letter))
								}

								compareValues(fmt.Sprintf("%d, %d: exponent", i, j), x.exponent, y.exponent)

								if !x.exponent.Equals(y.exponent) {
									fmt.Printf("got.monomials[%d].variables[%d].exponent (%v) not equal to tt.want.monomials[%d].variables[%d].exponent (%v) \n", 
									i, j, got.monomials[i].variables[j].exponent,
									i, j, tt.want.monomials[i].variables[j].exponent)
								}
							}
						}
					}
				}
				t.Errorf("Polynomial.StandardForm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareValues(name string, v1, v2 interface{}) {
	if !reflect.DeepEqual(v1, v2) {
		fmt.Printf("Mismatch in %s: %v != %v\n", name, v1, v2)
	}
}
