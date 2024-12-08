package basicmath

func FactorInt(a int) map[int]int {
	factors := make(map[int]int)
	if a < 0 {
		factors[-1] = 1
		a *= -1
	}

	if a == 0 {
		factors[0]++
		return factors
	}

	if a == 1 {
		factors[1]++
		return factors
	}

	product := 1

	for a%2 == 0 {
		factors[2]++
		product *= 2
		a /= 2
	}

	for i := 3; i <= a || product == a; i += 2 {
		for a%i == 0 {
			factors[i]++
			product *= i
			a /= i
		}
	}
	return factors
}