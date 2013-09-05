package main

import (
	"fmt"
	"strconv"
)

func main() {
	max := -1
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			mult := i * j
			if isPalindrome(mult) {
				if mult > max {
					max = mult
				}
			}
		}
	}
	fmt.Println(max)
}

func isPalindrome(num int) bool {
	chars := strconv.Itoa(num)
	count := len(chars)
	for ind := range chars {
		if chars[ind] != chars[count-ind-1] {
			return false
		}
	}
	return true
}
