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
