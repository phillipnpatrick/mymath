package algebra

import (
	"fmt"
	"mymath/basicmath"
	"mymath/latex"
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

// #region Factorable

func (p *Polynomial) Factor() []*Polynomial {
	var sb strings.Builder
	factors := []*Polynomial{}

	if isQuadraticTrinomial(p) {
		sb.WriteString(latex.WriteMathLine(fmt.Sprintf("%v", p.LaTeX())))
		trinomial := &Polynomial{}

		a, b, c := getQuadraticTrinomialTerms(p)

		// gcf
		gcf := basicmath.GetFractionGCF(a.coefficient, b.coefficient, c.coefficient)

		if !a.coefficient.IsInteger() || !b.coefficient.IsInteger() || !c.coefficient.IsInteger() {
			denominators := []int{a.coefficient.Denominator(), b.coefficient.Denominator(), c.coefficient.Denominator()}
			gcf = basicmath.NewFraction(1, basicmath.Max(denominators...))
		}

		if !gcf.Equals(basicmath.NewInteger(1)) {
			factors = append(factors, NewPolynomial(NewMonomialConstant(gcf)))
			a.coefficient = a.coefficient.Divide(gcf)
			b.coefficient = b.coefficient.Divide(gcf)
			c.coefficient = c.coefficient.Divide(gcf)
		}
		trinomial.monomials = append(trinomial.monomials, a)
		trinomial.monomials = append(trinomial.monomials, b)
		trinomial.monomials = append(trinomial.monomials, c)

		if !gcf.Equals(basicmath.NewInteger(1)) {
			sb.WriteString(latex.WriteMathLine(fmt.Sprintf("%v%s", gcf.LaTeX(), latex.WrapInParentheses(fmt.Sprintf("%v", trinomial)))))
		}

		// factor of a*c with sum b
		f1, f2, isPrime := basicmath.FactorsWithSum(b.coefficient, a.coefficient.Multiply(c.coefficient))

		if isPrime {
			factors = append(factors, trinomial)
		} else {
			// rewrite middle term
			left := NewPolynomial(makeCopyOfMonomial(*a), NewMonomial(f1, a.getVariableLetter()))
			right := NewPolynomial(NewMonomial(f2, a.getVariableLetter()), makeCopyOfMonomial(*c))

			if gcf.Equals(basicmath.NewInteger(1)) {
				sb.WriteString(latex.WriteMathLine(fmt.Sprintf("%s", latex.ConnectWithPlusSign(left.LaTeX(), right.LaTeX()))))
			} else {
				sb.WriteString(latex.WriteMathLine(fmt.Sprintf("%v%s", gcf.LaTeX(), latex.WrapInParentheses(latex.ConnectWithPlusSign(left.LaTeX(), right.LaTeX())))))
			}

			// group terms and factor
			leftGCF := GetMonomialGCF(left.monomials...)
			rightGCF := GetMonomialGCF(right.monomials...)

			leftFactored := NewPolynomial(left.monomials[0].Divide(leftGCF), left.monomials[1].Divide(leftGCF))
			rightFactored := NewPolynomial(right.monomials[0].Divide(rightGCF), right.monomials[1].Divide(rightGCF))

			if right.monomials[0].coefficient.LessThan(basicmath.NewInteger(0)) {
				rightGCF = rightGCF.Multiply(NewMonomialConstant(basicmath.NewInteger(-1)))
				rightFactored = rightFactored.Multiply(NewPolynomial(NewMonomialConstant(basicmath.NewInteger(-1))))
			}

			if gcf.Equals(basicmath.NewInteger(1)) {
				sb.WriteString(latex.WriteMathLine(
					latex.ConnectWithPlusSign(
						fmt.Sprintf("%v%s", leftGCF, latex.WrapInParentheses(fmt.Sprintf("%v", leftFactored))),
						fmt.Sprintf("%v%s", rightGCF, latex.WrapInParentheses(fmt.Sprintf("%v", rightFactored))),
					)))
			} else {
				sb.WriteString(latex.WriteMathLine(fmt.Sprintf("%v%s", gcf,
					latex.WrapInBrackets(
						latex.ConnectWithPlusSign(
							fmt.Sprintf("%v%s", leftGCF, latex.WrapInParentheses(fmt.Sprintf("%v", leftFactored))),
							fmt.Sprintf("%v%s", rightGCF, latex.WrapInParentheses(fmt.Sprintf("%v", rightFactored))),
						)))))
			}

			// factor common binomial
			factors = append(factors, NewPolynomial(leftGCF, rightGCF))

			if leftFactored.Equals(rightFactored) {
				factors = append(factors, leftFactored)
			}

			sb.WriteString("$")
			for _, factor := range factors {
				if len(factor.monomials) == 1 {
					sb.WriteString(gcf.LaTeX())
				} else {
					sb.WriteString(latex.WrapInParentheses(fmt.Sprintf("%v", factor)))
				}
			}
			sb.WriteString(`$\newline`)
		}
	}

	// fmt.Println(sb.String())

	return factors
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
			sb.WriteString(temp.LaTeX())
		} else {
			sb.WriteString(monomial.LaTeX())
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

func (p *Polynomial) Add(others ...*Polynomial) *Polynomial {
	temp := NewPolynomial(p.monomials...)

	for _, other := range others {
		for _, mono := range other.monomials {
			temp.AddMonomial(mono)
		}
	}

	return temp.StandardForm()
}

func (p *Polynomial) Subtract(others ...*Polynomial) *Polynomial {
	temp := NewPolynomial(p.monomials...)

	for _, other := range others {
		for _, mono := range other.monomials {
			temp.AddMonomial(mono.Multiply(NewMonomialConstant(basicmath.NewInteger(-1))))
		}
	}

	return temp.StandardForm()
}

func (p *Polynomial) Multiply(others ...*Polynomial) *Polynomial {
	if len(others) == 0 {
		return p
	}
	var monomials []*Monomial

	product := multiplyTwoPolynomials(p, others[0])
	monomials = append(monomials, product...)
	temp := NewPolynomial(monomials...)

	if len(others) > 1 {
		monomials = monomials[:0]

		for _, other := range others[1:] {
			product = multiplyTwoPolynomials(temp, other)
			monomials = append(monomials, product...)
			temp.monomials = append(temp.monomials, monomials...)
		}
	}

	return NewPolynomial(monomials...).StandardForm()
}

// Synthetic division
func (p *Polynomial) DividedBy(linearBinomial *Polynomial) *Polynomial {
	var coefficients []*basicmath.Fraction

	for _, monomial := range p.monomials {
		coefficients = append(coefficients, monomial.coefficient)
	}

	a := GetMonomialByDegree(basicmath.NewInteger(1), linearBinomial.monomials...)
	b := GetMonomialByDegree(basicmath.NewInteger(0), linearBinomial.monomials...)
	c := b.coefficient.Divide(a.coefficient)

	newCoefficients := []*basicmath.Fraction{coefficients[0]}
	j := 0
	for _, f := range coefficients[1:] {
		t := c.Multiply(newCoefficients[j])
		t = t.Add(f)
		newCoefficients = append(newCoefficients, t)
		j++
	}

	originalDegree := getHighestDegreeTerm(p)
	degree := originalDegree.degree.Subtract(basicmath.NewInteger(1))
	variable := originalDegree.variables[0].Letter()
	var terms []*Monomial

	for _, coeff := range newCoefficients {
		if degree.GreaterThan(basicmath.NewInteger(1)) {
			terms = append(terms, NewMonomialWithVariables(coeff, NewVariableWithExponent(variable, degree)))
		} else if degree.Equals(basicmath.NewInteger(0)) {
			terms = append(terms, NewMonomialConstant(coeff))
		}

		degree = degree.Subtract(basicmath.NewInteger(1))
	}

	return NewPolynomial(terms...)
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
	sort.Slice(p.monomials, func(i int, j int) bool {
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

// #region Private Methods

func getHighestDegreeTerm(p *Polynomial) *Monomial {
	term := p.monomials[0]
	for _, mono := range p.monomials[1:] {
		if term.degree.LessThan(mono.degree) {
			term = mono
		}
	}
	return term
}

func getQuadraticTrinomialTerms(p *Polynomial) (a *Monomial, b *Monomial, c *Monomial) {
	a = GetMonomialByDegree(basicmath.NewInteger(2), p.monomials...)
	b = GetMonomialByDegree(basicmath.NewInteger(1), p.monomials...)
	c = GetMonomialByDegree(basicmath.NewInteger(0), p.monomials...)

	return a, b, c
}

func isQuadraticTrinomial(p *Polynomial) bool {
	return len(p.monomials) == 3 &&
		len(p.monomials[0].variables) == 1 &&
		p.monomials[0].variables[0].exponent.Equals(basicmath.NewInteger(2))
}

func makeCopyOfPolynomial(p *Polynomial) *Polynomial {
	copy := &Polynomial{}

	for _, mono := range p.monomials {
		copy.monomials = append(copy.monomials, makeCopyOfMonomial(*mono))
	}

	return copy
}

func multiplyTwoPolynomials(a, b *Polynomial) []*Monomial {
	var monomials []*Monomial

	for _, am := range a.monomials {
		for _, bm := range b.monomials {
			temp := am.Multiply(bm)
			monomials = append(monomials, temp)
		}
	}

	return monomials
}

// #endregion
