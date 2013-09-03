package main

import "fmt"

func main() {
	sumOfSquares := 0
	sqareOfSum := 0
	for i := 1; i <= 100; i++ {
		sumOfSquares += i * i
		sqareOfSum += i
	}
	sqareOfSum *= sqareOfSum
	fmt.Println(sqareOfSum - sumOfSquares)
}
