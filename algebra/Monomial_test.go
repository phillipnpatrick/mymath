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
		// TODO: Add test cases.
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
			m: 	&Monomial{coefficient: basicmath.NewInteger(88)},
			want: basicmath.NewInteger(0),
		},
		{ // x
			name: "Monomial_Degree_Test08",
			m: 	&Monomial{variables: []*Variable{{letter: 'x'}},},
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
