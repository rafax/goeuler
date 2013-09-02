package main

import "fmt"

func main() {
	sum := 0
	prev := 1
	prev2 := 1
	for prev <= 4000000 {
		tmp := prev
		prev = prev + prev2
		prev2 = tmp
		if prev%2 == 0 {
			sum += prev
		}
	}
	fmt.Println(sum)
}
