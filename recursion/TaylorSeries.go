package recursion

import (
	"math"
)

func factorial(n uint64) uint64 {
	if n > 0 {
		return n * factorial(n-1)
	}

	return 1
}

func calc(n, x float64) float64 {
	return math.Pow(x, n) / float64(factorial(uint64(n)))
}

func calcNaturalExponential(sum float64, n, x float64) float64 {
	if n > 0 {
		sum += calc(n , x)
		return calcNaturalExponential(sum, n-1, x)
	}
	return sum
}

func NaturalExponential(n, x float64) float64 {
	//e^x = 1 + x/1 + x^2/2! + x^3/3! + x^4/4! + ... n terms
	return calcNaturalExponential(1, n, x)
}

// Horner's Rule
func calcHorner(sum float64, n, x float64) float64 {
	if n > 0 {
		sum = 1 + x * sum / n

		return calcHorner(sum, n - 1, x)
	}

	return sum
}

func NaturalExponentialHorners(n,x float64) float64 {
	return calcHorner(1, n, x)
}
