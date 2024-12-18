package basicmath

import (
	"math"
	"mymath/interfaces"
)

// #region Public Methods

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// AddTwo adds two Addable types
func AddTwo[T interfaces.Addable[T]](a, b T) T {
	return a.Add(b)
}

// DivideTwo adds two Dividable types
func DivideTwo[T interfaces.Dividable[T]](a, b T) T {
	return a.Divide(b)
}

// Find factors of number that add to sum; returns flag to indicate if isPrime
func FactorsWithSum(sum *Fraction, number *Fraction) (*Fraction, *Fraction, bool) {
	var a, b *Fraction
	factors := FactorsOf(number)

	for key, value := range factors {
		posKey := NewFraction(Abs(key.n), Abs(key.d))
		negKey := NewFraction(Neg(key.n), Abs(key.d))
		posValue := NewFraction(Abs(value.n), Abs(value.d))
		negValue := NewFraction(Neg(value.n), Abs(value.d))

		if posKey.Multiply(posValue).Equals(number) && posKey.Add(posValue).Equals(sum) {
			a = NewFraction(posKey.n, posKey.d)
			b = NewFraction(posValue.n, posValue.d)
		} else if posKey.Multiply(negValue).Equals(number) && posKey.Add(negValue).Equals(sum) {
			a = NewFraction(posKey.n, posKey.d)
			b = NewFraction(negValue.n, negValue.d)
		} else if negKey.Multiply(posValue).Equals(number) && negKey.Add(posValue).Equals(sum) {
			a = NewFraction(negKey.n, negKey.d)
			b = NewFraction(posValue.n, posValue.d)
		} else if negKey.Multiply(negValue).Equals(number) && negKey.Add(negValue).Equals(sum) {
			a = NewFraction(negKey.n, negKey.d)
			b = NewFraction(negValue.n, negValue.d)
		}
	}

	return a, b, a == nil && b == nil
}

func GCF(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	gcf := numbers[0]
	for _, number := range numbers[1:] {
		gcf = getGCF(gcf, number)
	}
	return gcf
}

func LCM(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	lcm := numbers[0]
	for _, number := range numbers[1:] {
		lcm = getLCM(lcm, number)
	}
	return lcm
}

func Max(numbers ...int) int {
	max := math.MinInt
	for _, number := range numbers {
		max = getMax(max, number)
	}
	return max
}

func Min(numbers ...int) int {
	min := math.MaxInt
	for _, number := range numbers {
		min = getMin(min, number)
	}
	return min
}

// MultiplyTwo adds two Multipliable types
func MultiplyTwo[T interfaces.Multipliable[T]](a, b T) T {
	return a.Multiply(b)
}

func Neg(x int) int {
	if x < 0 {
		return x
	}
	return -x
}

// SubtractTwo adds two Subtractable types
func SubtractTwo[T interfaces.Subtractable[T]](a, b T) T {
	return a.Subtract(b)
}

// #endregion

// #region Private Methods

func getFactorsOf(n int) map[*Fraction]*Fraction {
	n = Abs(n)
	factors := make(map[*Fraction]*Fraction)

	for i := 1; i <= (int(math.Sqrt(float64(n)))); i++ {
		if n%i == 0 {
			factors[NewInteger(i)] = NewInteger(n / i)
		}
	}

	return factors
}

func getGCF(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return Abs(a)
}

func getLCM(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return Abs(a*b) / getGCF(a, b)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// #endregion
