package algebra

import (
	"mymath/basicmath"
	"sort"
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

// #region LaTeXer

func (p Polynomial) LaTeX() string {
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

// #region Public Methods

func (p *Polynomial) AddMonomial(m *Monomial) {
	for i, mono := range p.monomials {
		if AreLikeTerms(m, mono) {
			p.monomials[i].coefficient = p.monomials[i].coefficient.Add(m.coefficient)
			return
		}
	}
	// If no similar monomial is found, append it
	p.monomials = append(p.monomials, m)
}

func (p *Polynomial) StandardForm() *Polynomial {
	// Rules:
	// Degree of term: a(x^m)(y^n), degree is m+n
	// Order by total degree: descending order; terms with same degree, order alphabetically
	// Combine like terms
	// Simplify coefficients
	//
	// Key Points:
	// Total Degree first: always prioritize terms with higher total degrees
	// Alphabetical order
	// Combine like terms

	monomialMap := make(map[string]*basicmath.Fraction)

	// Combine like terms by summing coefficients
	for _, monomial := range p.monomials {
		if value, exists := monomialMap[monomial.Variables()]; exists {
			c := value.Add(monomial.coefficient)
			monomialMap[monomial.Variables()] = c
		} else {
			monomialMap[monomial.Variables()] = monomial.coefficient
		}
	}

	// Create a simplified list of terms
	p.monomials = []*Monomial{}
	for vars, coefficient := range monomialMap {
		if !coefficient.Equals(basicmath.NewInteger(0)) { // skip zero coefficients
			v := ParseToVariables(vars)
			p.monomials = append(p.monomials, NewMonomialWithVariables(coefficient, v...))
		}
	}

	// Sort monomials by degree (descending) and then alphabetically
	sort.Slice(p.monomials, func(i int, j int) bool{
		a := p.monomials[i]
		b := p.monomials[j]
		if a.Degree().Equals(b.Degree()) {
			aVars := a.Variables()
			bVars := b.Variables()

			if aVars[0] == bVars[0] {
				return a.variables[0].exponent.GreaterThan(b.variables[0].exponent)
			}
			return aVars[0] < bVars[0]
		}
		return a.Degree().GreaterThan(b.Degree())
	})

	return p
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
