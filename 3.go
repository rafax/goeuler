package main

import "fmt"
import "./utils"

func main() {
	num := 600851475143
	f := utils.PrimeFactors(num)
	fmt.Println(f[len(f)-1])
}
