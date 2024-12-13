package algebra

import (
	"fmt"
	"mymath/basicmath"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Monomial struct {
	coefficient *basicmath.Fraction
	variables   []*Variable         // Example: (x^2)(y^3)
	degree      *basicmath.Fraction // Total degree of the term
}

// #region Monomial Constructors

func NewMonomial(coefficient *basicmath.Fraction, letter string) *Monomial {
	m := &Monomial{
		coefficient: coefficient,
		variables:   []*Variable{NewVariableWithExponent(letter, basicmath.NewInteger(1))},
	}

	m.degree = calculateDegree(m)

	return m
}

func NewMonomialConstant(coefficient *basicmath.Fraction) *Monomial {
	m := &Monomial{
		coefficient: coefficient,
	}

	m.degree = calculateDegree(m)

	return m
}

func NewMonomialWithExponent(coefficient *basicmath.Fraction, letter string, exponent *basicmath.Fraction) *Monomial {
	m := &Monomial{
		coefficient: coefficient,
		variables:   []*Variable{NewVariableWithExponent(letter, exponent)},
	}

	m.degree = calculateDegree(m)

	return m
}

func NewMonomialWithVariables(coefficient *basicmath.Fraction, variables ...*Variable) *Monomial {
	m := &Monomial{
		coefficient: coefficient,
		variables:   variables,
	}

	m.degree = calculateDegree(m)

	return m
}

// #endregion

// #region Properties

// The total degree of the monomial (a sum of the exponents)
func (m *Monomial) Degree() *basicmath.Fraction {
	m.degree = calculateDegree(m)
	return m.degree
}

// A string representation of the variables (without coefficient)
func (m *Monomial) Variables() string {
	var sb strings.Builder
	for _, variable := range m.variables {
		sb.WriteString(variable.String())
	}

	return sb.String()
}

// #endregion

// #region Comparable

func (m *Monomial) Equals(other *Monomial) bool { // Check if the number of monomials is the same
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
		key := fmt.Sprintf("%s^%s", string(variable.letter), variable.exponent.String())
		v1[key] = variable
	}

	for _, variable := range other.variables {
		key := fmt.Sprintf("%s^%s", string(variable.letter), variable.exponent.String())
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

	if len(m.variables) == 0 {
		return c
	}

	if m.coefficient.Equals(basicmath.NewInteger(1)) {
		c = ""
	} else if m.coefficient.Equals(basicmath.NewInteger(-1)) {
		c = "-"
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

func (m *Monomial) Multiply(others ...*Monomial) *Monomial {
	// temp := NewMonomialWithVariables(m.coefficient, m.variables...)
	temp := makeCopy(*m)

	for _, other := range others {
		temp = multiplyTwoMonomials(*temp, *other)
		// temp.coefficient = temp.coefficient.Multiply(other.coefficient)
		// if len(temp.variables) == 0 {
		// 	temp.variables = other.variables
		// } else {
		// 	for i, tempVar := range temp.variables {
		// 		for _, otherVar := range other.variables {
		// 			if tempVar.letter == otherVar.letter {
		// 				temp.variables[i].exponent = tempVar.exponent.Add(otherVar.exponent)
		// 			}
		// 		}
		// 	}
		// }
	}
	temp.degree = calculateDegree(temp)

	return temp
}

func (m *Monomial) Divide(others ...*Monomial) *Monomial {
	// temp := NewMonomialWithVariables(m.coefficient, m.variables...)
	temp := makeCopy(*m)

	for _, other := range others {
		temp.coefficient = temp.coefficient.Divide(other.coefficient)
		for i, tempVar := range temp.variables {
			for _, otherVar := range other.variables {
				if tempVar.letter == otherVar.letter {
					temp.variables[i].exponent = tempVar.exponent.Subtract(otherVar.exponent)
				}
			}
		}
	}
	temp.degree = calculateDegree(temp)

	return temp
}

// #endregion

// #region Stringer

func (m Monomial) String() string {
	c := fmt.Sprintf("%v", m.coefficient)

	if len(m.variables) == 0 {
		return c
	}

	if m.coefficient.Equals(basicmath.NewInteger(1)) {
		c = ""
	} else if m.coefficient.Equals(basicmath.NewInteger(-1)) {
		c = "-"
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

func ParseToVariables(variables string) []*Variable {
	if variables == "" {
		return nil
	}
	vars := []*Variable{}
	var sb strings.Builder

	for _, char := range variables {
		if unicode.IsLetter(char) && sb.Len() > 0 {
			vars = append(vars, parseToVariable(sb.String()))

			sb.Reset()
		}
		sb.WriteRune(char)
	}

	if sb.Len() > 0 {
		vars = append(vars, parseToVariable(sb.String()))
	}

	return vars
}

func (m *Monomial) StandardForm() *Monomial {
	// Sort variables alphabetically
	sort.Slice(m.variables, func(i int, j int) bool {
		a := m.variables[i]
		b := m.variables[j]
		
		return a.Letter() < b.Letter()
	})

	return m
}

// #endregion

// #region Private Methods

func (a *Monomial) multiplyVariable(v *Variable) {
	for i, variable := range a.variables {
		if AreLikeVariables(*v, *variable) {
			a.variables[i].exponent = a.variables[i].exponent.Add(v.exponent)
			return
		}
	}
	// If no similar variable is found, append it
	a.variables = append(a.variables, v)
}

func areLike(a, b *Monomial) bool {
	if len(a.variables) != len(b.variables) {
		return false
	}

	// Create maps to track variables by their variable and degree
	v1 := make(map[string]*Variable)
	v2 := make(map[string]*Variable)

	for _, variable := range a.variables {
		key := fmt.Sprintf("%s^%s", string(variable.letter), variable.exponent.String())
		v1[key] = variable
	}

	for _, variable := range b.variables {
		key := fmt.Sprintf("%s^%s", string(variable.letter), variable.exponent.String())
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

func calculateDegree(m *Monomial) *basicmath.Fraction {
	m.degree = basicmath.NewInteger(0)
	if len(m.variables) > 0 {
		if m.variables[0].exponent != nil {
			m.degree = basicmath.NewFraction(m.variables[0].exponent.Numerator(), m.variables[0].exponent.Denominator())
			for _, variable := range m.variables[1:] {
				m.degree = m.degree.Add(variable.exponent)
			}
		}
	}
	return m.degree
}

func makeCopy(m Monomial) *Monomial {
	copy := &Monomial{}

	copy.coefficient = basicmath.NewFraction(m.coefficient.Numerator(), m.coefficient.Denominator())

	var variables []*Variable
	for _, variable := range m.variables {
		variables = append(variables, NewVariableWithExponent(variable.Letter(),
			basicmath.NewFraction(variable.exponent.Numerator(), variable.exponent.Denominator())))
	}

	copy.variables = variables

	copy.degree = calculateDegree(copy)

	return copy
}

func multiplyTwoMonomials(a, b Monomial) *Monomial {
	m := &Monomial{}
	m.coefficient = a.coefficient.Multiply(b.coefficient)
	m.variables = a.variables

	for _, other := range b.variables {
		m.multiplyVariable(other)
	}

	return m.StandardForm()
}

func parseToVariable(part string) *Variable {
	var letter string
	numerator, denominator := 1, 1

	for _, item := range part {
		if unicode.IsLetter(item) {
			letter = string(item)
		} else if unicode.IsNumber(item) {
			if numerator == 1 {
				numerator, _ = strconv.Atoi(string(item))
			} else if denominator == 1 {
				denominator, _ = strconv.Atoi(string(item))
			}
		}
	}

	return NewVariableWithExponent(letter, basicmath.NewFraction(numerator, denominator))
}

// #endregion
