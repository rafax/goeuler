package utils

import "math"

func IsPrime(num int) bool {
	floatNum := float64(num)
	for i := 2; i <= int(math.Floor(math.Sqrt(floatNum))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Gcd(a, b int) int {
	var min, max int
	if a < b {
		min, max = a, b
	} else {
		min, max = b, a
	}
	for min != 0 {
		max, min = min, max%min
	}
	return max
}

func PrimeFactors(n int) []int {
	factors := make([]int, 0)
	d := 2
	for n > 1 {
		for n%d == 0 {
			factors = append(factors, d)
			n /= d
		}
		d += 1
		if d*d > n {
			if n > 1 {
				factors = append(factors, n)
			}
			break
		}
	}
	return factors
}
