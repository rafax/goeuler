package main

import "fmt"
import "./utils"

func nthPrime(n int) int {
	ind := 1
	for i := 3; ; i++ {
		if utils.IsPrime(i) {
			ind += 1
		}
		if ind == n {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(nthPrime(10001))
}
