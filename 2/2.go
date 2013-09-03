package main

import "fmt"

func main() {
	sum := 0
	curr := 1
	prev := 1
	for curr <= 4000000 {
		tmp := curr
		curr = curr + prev
		prev = tmp
		if curr%2 == 0 {
			sum += curr
		}
	}
	fmt.Println(sum)
}
