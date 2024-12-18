package algebra

import (
	"mymath/basicmath"
	"reflect"
	"testing"
)

func TestMonomial_Add(t *testing.T) {
	type args struct {
		others []*Monomial
	}
	tests := []struct {
		name string
		m    *Monomial
		args args
		want *Polynomial
	}{
		{ // 9x^2 + 4x^3
			name: "Monomial_Add_Test01",
			m:    NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2)),
			args: args{others: []*Monomial{
				NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(3)),
			}},
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(3))),
		},
		{ // 5x^2 - 12y^3
			name: "Monomial_Add_Test02",
			m:    NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(2)),
			args: args{others: []*Monomial{
				NewMonomialWithExponent(basicmath.NewInteger(-12), "y", basicmath.NewInteger(3)),
			}},
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(-12), "y", basicmath.NewInteger(3))),
		},
		{ // 5x^3 + 9x^3 = 14x^3
			name: "Monomial_Add_Test03",
			m:    NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(3)),
			args: args{others: []*Monomial{
				NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(3)),
			}},
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(14), "x", basicmath.NewInteger(3))),
		},
		{ // 2x^2 + 5 - 4x^3 + 7x^3 - x^2 = x^2 + 3x^3 + 5
			name: "Monomial_Add_Test04",
			m:    NewMonomialWithExponent(basicmath.NewInteger(2), "x", basicmath.NewInteger(2)),
			args: args{others: []*Monomial{
				NewMonomialConstant(basicmath.NewInteger(5)),
				NewMonomialWithExponent(basicmath.NewInteger(-4), "x", basicmath.NewInteger(3)),
				NewMonomialWithExponent(basicmath.NewInteger(7), "x", basicmath.NewInteger(3)),
				NewMonomialWithExponent(basicmath.NewInteger(-1), "x", basicmath.NewInteger(2)),
			}},
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(1), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(3), "x", basicmath.NewInteger(3)),
				NewMonomialConstant(basicmath.NewInteger(5)),
			),
		},
		{ // 4x^2 + 3x + 4y + 8x + 10x^2 = 14x^2 + 11x + 4y
			name: "Monomial_Add_Test05",
			m:    NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(2)),
			args: args{others: []*Monomial{
				NewMonomial(basicmath.NewInteger(3), "x"),
				NewMonomial(basicmath.NewInteger(4), "y"),
				NewMonomial(basicmath.NewInteger(8), "x"),
				NewMonomialWithExponent(basicmath.NewInteger(10), "x", basicmath.NewInteger(2)),
			}},
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(14), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(11), "x"),
				NewMonomial(basicmath.NewInteger(4), "y"),
			),
		},
		{ // 2xy + 4x^2 + 5yx + 5y^2 + 16x^2 = 7xy + 20x^2 + 5y^2
			name: "Monomial_Add_Test06",
			m:    NewMonomialWithVariables(basicmath.NewInteger(2), NewVariable("x"), NewVariable("y")),
			args: args{others: []*Monomial{
				NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(2)),
				NewMonomialWithVariables(basicmath.NewInteger(5), NewVariable("y"), NewVariable("x")),
				NewMonomialWithExponent(basicmath.NewInteger(5), "y", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(16), "x", basicmath.NewInteger(2)),
			}},
			want: NewPolynomial(
				NewMonomialWithVariables(basicmath.NewInteger(7), NewVariable("x"), NewVariable("y")),
				NewMonomialWithExponent(basicmath.NewInteger(20), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(5), "y", basicmath.NewInteger(2)),
			),
		},
		{ // 7m + 14m – 6n – 5n + 2m = 23m – 11n
			name: "Monomial_Add_Test07",
			m:    NewMonomial(basicmath.NewInteger(7), "m"),
			args: args{others: []*Monomial{
				NewMonomial(basicmath.NewInteger(14), "m"),
				NewMonomial(basicmath.NewInteger(-6), "n"),
				NewMonomial(basicmath.NewInteger(-5), "n"),
				NewMonomial(basicmath.NewInteger(2), "m"),
			}},
			want: NewPolynomial(
				NewMonomial(basicmath.NewInteger(23), "m"),
				NewMonomial(basicmath.NewInteger(-11), "n"),
			),
		},
		{ // 2x^2 + 3x – 4 – x^2 + x + 9 = x^2 + 4x + 5
			name: "Monomial_Add_Test08",
			m:    NewMonomialWithExponent(basicmath.NewInteger(2), "x", basicmath.NewInteger(2)),
			args: args{others: []*Monomial{
				NewMonomial(basicmath.NewInteger(3), "x"),
				NewMonomialConstant(basicmath.NewInteger(-4)),
				NewMonomialWithExponent(basicmath.NewInteger(-1), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(1), "x"),
				NewMonomialConstant(basicmath.NewInteger(9)),
			}},
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(1), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(4), "x"),
				NewMonomialConstant(basicmath.NewInteger(5)),
			),
		},
		{ // 10x^3 – 14x^2 + 3x – 4x^3 + 4x – 6 = 6x^3 – 14x^2 + 7x – 6
			name: "Monomial_Add_Test09",
			m:    NewMonomialWithExponent(basicmath.NewInteger(10), "x", basicmath.NewInteger(3)),
			args: args{others: []*Monomial{
				NewMonomialWithExponent(basicmath.NewInteger(-14), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(3), "x"),
				NewMonomialWithExponent(basicmath.NewInteger(-4), "x", basicmath.NewInteger(3)),
				NewMonomial(basicmath.NewInteger(4), "x"),
				NewMonomialConstant(basicmath.NewInteger(-6)),
			}},
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(6), "x", basicmath.NewInteger(3)),
				NewMonomialWithExponent(basicmath.NewInteger(-14), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(7), "x"),
				NewMonomialConstant(basicmath.NewInteger(-6)),
			),
		},
		// [(6x – 8) – 2x] – [(12x – 7) – (4x – 5)] = –4x – 6
		// –4y – [3x + (3y – 2x + {2y – 7}) – 4x + 5] = 3x – 9y + 2
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.Add(tt.args.others...)

			// fmt.Printf("got: %v \t want: %v \n", got, tt.want)
			if !got.Equals(tt.want) {
				t.Errorf("Monomial.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonomial_Subtract(t *testing.T) {
	type args struct {
		others []*Monomial
	}
	tests := []struct {
		name string
		m    *Monomial
		args args
		want *Polynomial
	}{
		{ // 3x^2 - 5x - 2x^2 - 4x
			name: "Monomial_Subtract_Test01",
			m:    NewMonomialWithExponent(basicmath.NewInteger(3), "x", basicmath.NewInteger(2)),
			args: args{others: []*Monomial{
				NewMonomial(basicmath.NewInteger(5), "x"),
				NewMonomialWithExponent(basicmath.NewInteger(2), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(4), "x"),
			}},
			want: NewPolynomial(
				NewMonomialWithExponent(basicmath.NewInteger(1), "x", basicmath.NewInteger(2)),
				NewMonomial(basicmath.NewInteger(-9), "x")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Subtract(tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Monomial.Subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonomial_Equals(t *testing.T) {
	type args struct {
		other *Monomial
	}
	tests := []struct {
		name string
		m    *Monomial
		args args
		want bool
	}{
		{
			name: "Monomial_Equals_Test01",
			m:    NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2)),
			args: args{other: NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2))},
			want: true,
		},
		{
			name: "Monomial_Equals_Test02",
			m:    NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2)),
			args: args{other: NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(3))},
			want: false,
		},
		{
			name: "Monomial_Equals_Test03",
			m: NewMonomialWithVariables(basicmath.NewInteger(9),
				NewVariableWithExponent("a", basicmath.NewInteger(2)),
				NewVariableWithExponent("b", basicmath.NewInteger(3)),
				NewVariableWithExponent("c", basicmath.NewInteger(4))),
			args: args{other: NewMonomialWithVariables(basicmath.NewInteger(9),
				NewVariableWithExponent("c", basicmath.NewInteger(4)),
				NewVariableWithExponent("a", basicmath.NewInteger(2)),
				NewVariableWithExponent("b", basicmath.NewInteger(3)))},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Equals(tt.args.other); got != tt.want {
				t.Errorf("Monomial.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonomial_LaTeX(t *testing.T) {
	tests := []struct {
		name string
		m    Monomial
		want string
	}{
		{
			name: "Monomial_LaTeX_Test01",
			m:    *NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2)),
			want: "9x^2",
		},
		{
			name: "Monomial_LaTeX_Test02",
			m: *NewMonomialWithVariables(basicmath.NewInteger(9),
				NewVariableWithExponent("a", basicmath.NewInteger(2)),
				NewVariableWithExponent("b", basicmath.NewInteger(3)),
				NewVariableWithExponent("c", basicmath.NewInteger(4))),
			want: "9a^2b^3c^4",
		},
		{
			name: "Monomial_LaTeX_Test03",
			m: *NewMonomialConstant(basicmath.NewInteger(-1)),
			want: "-1",
		},
		{
			name: "Monomial_LaTeX_Test04",
			m: *NewMonomialConstant(basicmath.NewInteger(0)),
			want: "0",
		},
		{
			name: "Monomial_LaTeX_Test05",
			m: *NewMonomialConstant(basicmath.NewInteger(1)),
			want: "1",
		},
		{
			name: "Monomial_LaTeX_Test06",
			m: *NewMonomialWithExponent(basicmath.NewInteger(-1), "m", basicmath.NewInteger(2)),
			want: "-m^2",
		},
		{
			name: "Monomial_LaTeX_Test07",
			m: *NewMonomialWithExponent(basicmath.NewInteger(1), "m", basicmath.NewInteger(2)),
			want: "m^2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.LaTeX(); got != tt.want {
				t.Errorf("Monomial.LaTeX() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAreLikeTerms(t *testing.T) {
	type args struct {
		monomials []*Monomial
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{ // 9x^2 + 4x^3
			name: "Monomial_AreLikeTerms_Test01",
			args: args{monomials: []*Monomial{
				NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(2)),
				NewMonomialWithExponent(basicmath.NewInteger(4), "x", basicmath.NewInteger(3)),
			}},
			want: false,
		},
		{ // 5x^3 + 9x^3 = 14x^3
			name: "Monomial_AreLikeTerms_Test02",
			args: args{monomials: []*Monomial{
				NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(3)),
				NewMonomialWithExponent(basicmath.NewInteger(9), "x", basicmath.NewInteger(3)),
			}},
			want: true,
		},
		{ // 2xy + 5yx = 7xy
			name: "Monomial_AreLikeTerms_Test03",
			args: args{monomials: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(2), NewVariable("x"), NewVariable("y")),
				NewMonomialWithVariables(basicmath.NewInteger(5), NewVariable("y"), NewVariable("x")),
			}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreLikeTerms(tt.args.monomials...); got != tt.want {
				t.Errorf("AreLikeTerms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonomial_Degree(t *testing.T) {
	tests := []struct {
		name string
		m    *Monomial
		want *basicmath.Fraction
	}{
		{ // x^2
			name: "Monomial_Degree_Test01",
			m:    NewMonomialWithExponent(basicmath.NewInteger(2), "x", basicmath.NewInteger(2)),
			want: basicmath.NewInteger(2),
		},
		{ // -2x
			name: "Monomial_Degree_Test02",
			m:    NewMonomial(basicmath.NewInteger(-2), "x"),
			want: basicmath.NewInteger(1),
		},
		{ // 1
			name: "Monomial_Degree_Test03",
			m:    NewMonomialConstant(basicmath.NewInteger(1)),
			want: basicmath.NewInteger(0),
		},
		{ // 2x^2y
			name: "Monomial_Degree_Test04",
			m:    NewMonomialWithVariables(basicmath.NewInteger(2), NewVariableWithExponent("x", basicmath.NewInteger(2)), NewVariable("y")),
			want: basicmath.NewInteger(3),
		},
		{ // 3xy^2
			name: "Monomial_Degree_Test05",
			m:    NewMonomialWithVariables(basicmath.NewInteger(3), NewVariable("x"), NewVariableWithExponent("y", basicmath.NewInteger(2))),
			want: basicmath.NewInteger(3),
		},
		{ // -3xy^2z
			name: "Monomial_Degree_Test06",
			m: NewMonomialWithVariables(
				basicmath.NewInteger(-3),
				NewVariable("x"),
				NewVariableWithExponent("y", basicmath.NewInteger(2)),
				NewVariable("z")),
			want: basicmath.NewInteger(4),
		},
		{ // 88
			name: "Monomial_Degree_Test07",
			m:    &Monomial{coefficient: basicmath.NewInteger(88)},
			want: basicmath.NewInteger(0),
		},
		{ // x
			name: "Monomial_Degree_Test08",
			m:    &Monomial{variables: []*Variable{{letter: 'x'}}},
			want: basicmath.NewInteger(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Degree(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Monomial.Degree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseToVariables(t *testing.T) {
	type args struct {
		variables string
	}
	tests := []struct {
		name string
		args args
		want []*Variable
	}{
		{
			name: "Monomial_ParseToVariables_Test01",
			args: args{variables: "xy^2z"},
			want: []*Variable{NewVariable("x"), NewVariableWithExponent("y", basicmath.NewInteger(2)), NewVariable("z")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseToVariables(tt.args.variables); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseToVariables() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonomial_Multiply(t *testing.T) {
	type args struct {
		others []*Monomial
	}
	tests := []struct {
		name string
		m    *Monomial
		args args
		want *Monomial
	}{
		{ // 3x^2 -* 5x^3 = 15x^5
			name: "Monomial_Multiply_Test01",
			m:    NewMonomialWithExponent(basicmath.NewInteger(3), "x", basicmath.NewInteger(2)),
			args: args{others: []*Monomial{
				NewMonomialWithExponent(basicmath.NewInteger(5), "x", basicmath.NewInteger(3)),
			}},
			want: NewMonomialWithExponent(basicmath.NewInteger(15), "x", basicmath.NewInteger(5)),
		},
		{ // -2a^3 * 4a^2 = -8a^5
			name: "Monomial_Multiply_Test02",
			m:    NewMonomialWithExponent(basicmath.NewInteger(-2), "a", basicmath.NewInteger(3)),
			args: args{others: []*Monomial{
				NewMonomialWithExponent(basicmath.NewInteger(4), "a", basicmath.NewInteger(2)),
			}},
			want: NewMonomialWithExponent(basicmath.NewInteger(-8), "a", basicmath.NewInteger(5)),
		},
		{ // 6m^2n * 2m^3n^4 = 12m^5n^5
			name: "Monomial_Multiply_Test03",
			m: NewMonomialWithVariables(basicmath.NewInteger(6),
				NewVariableWithExponent("m", basicmath.NewInteger(2)),
				NewVariable("n")),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(2),
					NewVariableWithExponent("m", basicmath.NewInteger(3)),
					NewVariableWithExponent("n", basicmath.NewInteger(4))),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(12),
				NewVariableWithExponent("m", basicmath.NewInteger(5)),
				NewVariableWithExponent("n", basicmath.NewInteger(5))),
		},
		{ // 4p^3q^2 * -5p^4q^4 = -20p^7q^6
			name: "Monomial_Multiply_Test04",
			m: NewMonomialWithVariables(basicmath.NewInteger(4),
				NewVariableWithExponent("p", basicmath.NewInteger(3)),
				NewVariableWithExponent("q", basicmath.NewInteger(2))),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(-5),
					NewVariableWithExponent("p", basicmath.NewInteger(4)),
					NewVariableWithExponent("q", basicmath.NewInteger(4))),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(-20),
				NewVariableWithExponent("p", basicmath.NewInteger(7)),
				NewVariableWithExponent("q", basicmath.NewInteger(6))),
		},
		{ // -7x^2y^3 * 3xy^2 = -21x^3y^5
			name: "Monomial_Multiply_Test05",
			m: NewMonomialWithVariables(basicmath.NewInteger(-7),
				NewVariableWithExponent("x", basicmath.NewInteger(2)),
				NewVariableWithExponent("y", basicmath.NewInteger(3))),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(3),
					NewVariable("x"),
					NewVariableWithExponent("y", basicmath.NewInteger(2))),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(-21),
				NewVariableWithExponent("x", basicmath.NewInteger(3)),
				NewVariableWithExponent("y", basicmath.NewInteger(5))),
		},
		{ // a * b = ab
			name: "Monomial_Multiply_Test06",
			m:    NewMonomial(basicmath.NewInteger(1), "a"),
			args: args{others: []*Monomial{
				NewMonomial(basicmath.NewInteger(1), "b"),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(1), NewVariable("a"), NewVariable("b")),
		},
		{ // ab * b = ab^2
			name: "Monomial_Multiply_Test07",
			m:    NewMonomialWithVariables(basicmath.NewInteger(1), NewVariable("a"), NewVariable("b")),
			args: args{others: []*Monomial{
				NewMonomial(basicmath.NewInteger(1), "b"),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(1), NewVariable("a"), NewVariableWithExponent("b", basicmath.NewInteger(2))),
		},
		{ // ab * bc = ab^2c
			name: "Monomial_Multiply_Test08",
			m:    NewMonomialWithVariables(basicmath.NewInteger(1), NewVariable("a"), NewVariable("b")),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(1), NewVariable("b"), NewVariable("c")),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(1), 
			NewVariable("a"), 
			NewVariableWithExponent("b", basicmath.NewInteger(2)),
			NewVariable("c"), ),
		},
		{ // c * a * b = abc
			name: "Monomial_Multiply_Test09",
			m:    NewMonomial(basicmath.NewInteger(1), "c"),
			args: args{others: []*Monomial{
				NewMonomial(basicmath.NewInteger(1), "a"),
				NewMonomial(basicmath.NewInteger(1), "b"),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(1), NewVariable("a"), NewVariable("b"), NewVariable("c")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.Multiply(tt.args.others...)
			if !reflect.DeepEqual(got, tt.want) {
				compareValues("coefficient", got.coefficient, tt.want.coefficient)
				compareValues("degree", got.degree, tt.want.degree)
				compareValues("variables", got.variables, tt.want.variables)
				t.Errorf("Monomial.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMonomial_Divide(t *testing.T) {
	type args struct {
		others []*Monomial
	}
	tests := []struct {
		name string
		m    *Monomial
		args args
		want *Monomial
	}{
		{ // 12x^5 / 4x^2 = 3x^3
			name: "Monomial_Divide_Test01",
			m: NewMonomialWithVariables(basicmath.NewInteger(12),
				NewVariableWithExponent("x", basicmath.NewInteger(5))),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(4),
					NewVariableWithExponent("x", basicmath.NewInteger(2))),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(3),
				NewVariableWithExponent("x", basicmath.NewInteger(3))),
		},
		{ // -15a^7 / 5a^3 = -3a^4
			name: "Monomial_Divide_Test02",
			m: NewMonomialWithVariables(basicmath.NewInteger(-15),
				NewVariableWithExponent("a", basicmath.NewInteger(7))),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(5),
					NewVariableWithExponent("a", basicmath.NewInteger(3))),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(-3),
				NewVariableWithExponent("a", basicmath.NewInteger(4))),
		},
		{ // 18m^6n^4 / 6m^2n^3 = 3m^2n^3
			name: "Monomial_Divide_Test03",
			m: NewMonomialWithVariables(basicmath.NewInteger(18),
				NewVariableWithExponent("m", basicmath.NewInteger(6)),
				NewVariableWithExponent("n", basicmath.NewInteger(4))),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(6),
					NewVariableWithExponent("m", basicmath.NewInteger(2)),
					NewVariableWithExponent("n", basicmath.NewInteger(3))),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(3),
				NewVariableWithExponent("m", basicmath.NewInteger(4)),
				NewVariable("n")),
		},
		{ // -24x^8y^5 / 8x^3y^2 = -3x^5y^3
			name: "Monomial_Divide_Test04",
			m: NewMonomialWithVariables(basicmath.NewInteger(-24),
				NewVariableWithExponent("x", basicmath.NewInteger(8)),
				NewVariableWithExponent("y", basicmath.NewInteger(5))),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(8),
					NewVariableWithExponent("x", basicmath.NewInteger(3)),
					NewVariableWithExponent("y", basicmath.NewInteger(2))),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(-3),
				NewVariableWithExponent("x", basicmath.NewInteger(5)),
				NewVariableWithExponent("y", basicmath.NewInteger(3))),
		},
		{ // 40p^9q^7 / 10p^4q^3 = 4p^5q^4
			name: "Monomial_Divide_Test05",
			m: NewMonomialWithVariables(basicmath.NewInteger(40),
				NewVariableWithExponent("p", basicmath.NewInteger(9)),
				NewVariableWithExponent("q", basicmath.NewInteger(7))),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(10),
					NewVariableWithExponent("p", basicmath.NewInteger(4)),
					NewVariableWithExponent("q", basicmath.NewInteger(3))),
			}},
			want: NewMonomialWithVariables(basicmath.NewInteger(4),
				NewVariableWithExponent("p", basicmath.NewInteger(5)),
				NewVariableWithExponent("q", basicmath.NewInteger(4))),
		},
		{ // ab / abc = 1/c ???
			name: "Monomial_Divide_Test06",
			m: NewMonomialWithVariables(basicmath.NewInteger(1), NewVariable("a"), NewVariable("b")),
			args: args{others: []*Monomial{
				NewMonomialWithVariables(basicmath.NewInteger(1), NewVariable("a"), NewVariable("b"), NewVariable("c")),
			}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Divide(tt.args.others...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Monomial.Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
