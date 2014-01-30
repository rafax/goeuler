package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	biggie := new(big.Int)
	base := new(big.Int)
	pow := new(big.Int)
	biggie.Exp(base.SetInt64(2), pow.SetInt64(1000), nil)
	digits := strings.Split(biggie.String(), "")
	sum := 0
	for _, s := range digits {
		v, _ := strconv.Atoi(s)
		sum += v
	}
	fmt.Println(sum)
}
