package basicmath

import "math"

// #region Public Methods

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