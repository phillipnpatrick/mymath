package algebra

import (
	"fmt"
	"mymath/basicmath"
	"sort"
	"strings"
)

type Monomial struct {
	coefficient *basicmath.Fraction
	variables []*Variable
}

// #region Monomial Constructors

func NewMonomial(coefficient *basicmath.Fraction, letter string) *Monomial {
	return &Monomial{
		coefficient: coefficient,
		variables: []*Variable{NewVariableWithDegree(letter, basicmath.NewInteger(1))},
	}
}

func NewMonomialConstant(coefficient *basicmath.Fraction) *Monomial {
	return &Monomial{
		coefficient: coefficient,
		variables: []*Variable{NewVariableWithDegree("x", basicmath.NewInteger(0))},
	}
}

func NewMonomialWithDegree(coefficient *basicmath.Fraction, letter string, degree *basicmath.Fraction) *Monomial {
	return &Monomial{
		coefficient: coefficient,
		variables: []*Variable{NewVariableWithDegree(letter, degree)},
	}
}

func NewMonomialWithVariables(coefficient *basicmath.Fraction, variables ...*Variable) *Monomial {
	return &Monomial{
		coefficient: coefficient,
		variables:   variables,
	}
}

// #endregion

// #region Comparable

func (m *Monomial) Equals(other *Monomial) bool {// Check if the number of monomials is the same
	if len(m.variables) != len(other.variables) {
		return false
	}

	if !m.coefficient.Equals(other.coefficient) {
		return false
	}
	
	// Create maps to track variables by their variable and degree
	v1 := make(map[string]*Variable)
	v2 := make(map[string]*Variable)

	for _, variable := range m.variables {
		key := fmt.Sprintf("%s^%s", string(variable.letter), variable.degree.String())
		v1[key] = variable
	}

	for _, variable := range other.variables {
		key := fmt.Sprintf("%s^%s", string(variable.letter), variable.degree.String())
		v2[key] = variable
	}

	// Compare both maps
	for key, variable1 := range v1 {
		variable2, exists := v2[key]
		if !exists || !variable1.Equals(variable2) {
			return false
		}
	}

	return true	
}

// #endregion

// #region Factorable
// #endregion

// #region LaTeXer

func (m Monomial) LaTeX() string {
	c := fmt.Sprintf("%v", m.coefficient.LaTeX())

	if m.coefficient.Equals(basicmath.NewInteger(1)) {
		c = ""
	}

	var sb strings.Builder
	sb.WriteString(c)
	for _, variable := range m.variables {
		sb.WriteString(variable.LaTeX())
	}

	return sb.String()
}

// #endregion

// #region Operable
func (m *Monomial) Add(others ...*Monomial) *Polynomial {
	p := &Polynomial{}
	p.AddMonomial(m)

	for _, other := range others {
		p.AddMonomial(other)
	}

	return p
}

func (m *Monomial) Subtract(others ...*Monomial) *Polynomial {
	p := &Polynomial{}
	p.AddMonomial(m)

	for _, other := range others {
		o := NewMonomialWithVariables(other.coefficient.Multiply(basicmath.NewInteger(-1)), other.variables...)

		p.AddMonomial(o)
	}

	return p
}

// func (m *Monomial) Multiply(others ...*Monomial) *Polynomial {
// 	temp := NewMonomialWithDegree(m.coefficient, m.letter, m.degree)

// 	for _, other := range others {
// 		if AreLikeTerms(temp, other) {
// 			temp.coefficient = temp.coefficient.Multiply(other.coefficient)
// 			temp.degree = temp.degree.Add(other.degree)
// 		} else {
// 			return nil
// 		}
// 	}

// 	return temp
// }

// func (m *Monomial) Divide(others ...*Monomial) *Polynomial {
// 	temp := NewMonomialWithDegree(m.coefficient, m.letter, m.degree)

// 	for _, other := range others {
// 		if AreLikeTerms(temp, other) {
// 			temp.coefficient = temp.coefficient.Divide(other.coefficient)
// 			temp.degree = temp.degree.Subtract(other.degree)
// 		} else {

// 		}
// 	}

// 	return temp
// }

// #endregion

// #region Stringer

func (m Monomial) String() string {
	c := fmt.Sprintf("%v", m.coefficient)

	if m.coefficient.Equals(basicmath.NewInteger(1)) {
		c = ""
	}

	var sb strings.Builder
	sb.WriteString(c)
	for _, variable := range m.variables {
		sb.WriteString(variable.String())
	}

	return sb.String()
}

// #endregion

// #region Public Methods

func AreLikeTerms(monomials ...*Monomial) bool {
	like := true

	for i := 1; i < len(monomials); i++ {
		like = like && areLike(monomials[i-1], monomials[i])
	}

	return like
}

// #endregion

// #region Private Methods

func areLike(a, b *Monomial) bool {
	if len(a.variables) != len(b.variables) {
		return false
	}

	// Create maps to track variables by their variable and degree
	v1 := make(map[string]*Variable)
	v2 := make(map[string]*Variable)

	for _, variable := range a.variables {
		key := fmt.Sprintf("%s^%s", string(variable.letter), variable.degree.String())
		v1[key] = variable
	}

	for _, variable := range b.variables {
		key := fmt.Sprintf("%s^%s", string(variable.letter), variable.degree.String())
		v2[key] = variable
	}

	// Compare both maps
	for key, variable1 := range v1 {
		variable2, exists := v2[key]
		if !exists || !variable1.IsLikeTerm(*variable2) {
			return false
		}
	}

	return true	
}

func sortString(s string) string {
	// Convert string to a slice of runes to handle Unicode characters
	runes := []rune(s)
	// Sort the slice of runes
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	// Convert the sorted slice back to a string
	return string(runes)
}

// #endregion
