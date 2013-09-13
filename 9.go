package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 1; a < 1000; a++ {
		for b := a; b < 1000-a; b++ {
			hass, c := hasSqrt(a*a + b*b)
			if hass && a+b+c == 1000 {
				fmt.Printf("%d^2 + %d^2 = %d^2\n", a*a, b*b, c*c)
				fmt.Printf("%d + %d + %d = 1000\n", a, b, c)

			}
		}
	}
}

func hasSqrt(n int) (bool, int) {
	sq := math.Sqrt(float64(n))
	return float64(int(sq)) == sq, int(sq)
}
