package main

import "fmt"

func main() {
	rangeEnd := 20
	for i := rangeEnd; ; i += rangeEnd {
		if divisibleByAll(i, rangeEnd) {
			fmt.Println(i)
			return
		}
	}
}

func divisibleByAll(num, rangeEnd int) bool {
	//for i := 1; i <= rangeEnd; i++ {
	for i := rangeEnd; i >= 1; i-- {
		if num%i != 0 {
			return false
		}
	}
	return true
}
