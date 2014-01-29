package main

import (
	"fmt"
	"math"
)

func main() {
	target := 500
	triangle := 0
	for nth := 1; ; nth++ {
		triangle += nth
		numFactors := numOfDivisors(triangle)
		if numFactors > target {
			fmt.Println(triangle)
			return
		}
		if numFactors%100 == 0 {
			fmt.Println(numFactors)
		}
	}
}

func numOfDivisors(num int) int {
	bound := int(math.Ceil(math.Sqrt(float64(num)))) + 1
	factors := make(map[int]bool)
	cnt := 2
	for i := 2; i <= bound; i++ {
		if num%i == 0 && !factors[i] {
			factors[i] = true
			cnt++
		}
		res := num / i
		if res > 1 && num%res == 0 && !factors[res] {
			factors[res] = true
			cnt++
		}
	}
	return cnt
}
