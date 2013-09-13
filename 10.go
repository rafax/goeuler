package main

import (
	"./utils"
	"fmt"
)

func main() {
	max := 2000000
	sum := 2
	for i := 3; i < max; i += 2 {
		if utils.IsPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
