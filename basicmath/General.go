package basicmath

import (
	"math"
	"mymath/interfaces"
)

// #region Public Methods

// AddTwo adds two Addable types
func AddTwo[T interfaces.Addable[T]](a, b T) T {
	return a.Add(b)
}

// DivideTwo adds two Dividable types
func DivideTwo[T interfaces.Dividable[T]](a, b T) T {
	return a.Divide(b)
}

// MultiplyTwo adds two Multipliable types
func MultiplyTwo[T interfaces.Multipliable[T]](a, b T) T {
	return a.Multiply(b)
}

// SubtractTwo adds two Subtractable types
func SubtractTwo[T interfaces.Subtractable[T]](a, b T) T {
	return a.Subtract(b)
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

func Min(numbers ...int) int {
	min := math.MaxInt
	for _, number := range numbers {
		min = getMin(min, number)
	}
	return min
}

// #endregion

// #region Private Methods

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getGCF(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return abs(a)
}

func getLCM(a,b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return abs(a*b) / getGCF(a,b)
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// #endregion