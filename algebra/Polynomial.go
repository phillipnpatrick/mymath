package algebra

import (
	"mymath/basicmath"
	"strings"
)

type Polynomial struct {
	monomials []*Monomial
}

// #region Constructors

func NewPolynomial(monomials ...*Monomial) *Polynomial {
	p := &Polynomial{}

	p.monomials = append(p.monomials, monomials...)

	return p
}

// #endregion

// #region Comparable

func (p *Polynomial) Equals(other *Polynomial) bool {
	// Check if the number of monomials is the same
	if len(p.monomials) != len(other.monomials) {
		return false
	}

	// Create maps to track monomials by their variable and degree
	m1 := make(map[string]*Monomial)
	m2 := make(map[string]*Monomial)

	for _, mono := range p.monomials {
		key := mono.String()
		m1[key] = mono
	}

	for _, mono := range other.monomials {
		key := mono.String()
		m2[key] = mono
	}

	// Compare both maps
	for key, monomial1 := range m1 {
		monomial2, exists := m2[key]
		if !exists || !monomial1.Equals(monomial2) {
			return false
		}
	}

	return true
}

// #endregion

// #region Public Methods

func (p *Polynomial) AddMonomial(m *Monomial) {
	for i, mono := range p.monomials {
		if AreLikeTerms(m, mono) {
			p.monomials[i].coefficient = p.monomials[i].coefficient.Add(m.coefficient)
			return
		}
		// if mono.letter == m.letter && mono.degree.Equals(m.degree) {
		// 	// Combine coefficients if monomials are similar
		// 	p.monomials[i].coefficient = p.monomials[i].coefficient.Add(m.coefficient)
		// 	return
		// }
	}
	// If no similar monomial is found, append it
	p.monomials = append(p.monomials, m)
}

// #endregion

// #region Stringer

func (p *Polynomial) String() string {
	var sb strings.Builder

	for _, monomial := range p.monomials {
		if sb.Len() > 0 {
			var temp *Monomial
			if monomial.coefficient.LessThan(basicmath.NewInteger(0)) {
				sb.WriteString(" - ")
				temp = NewMonomialWithVariables(monomial.coefficient.Multiply(basicmath.NewInteger(-1)), monomial.variables...)
			} else {
				sb.WriteString(" + ")
				temp = NewMonomialWithVariables(monomial.coefficient, monomial.variables...)
			}
			sb.WriteString(temp.String())
		} else {
			sb.WriteString(monomial.String())
		}
	}

	return sb.String()
}

// #endregion
