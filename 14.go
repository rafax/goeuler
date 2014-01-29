package main

import "fmt"

func main() {
	limit := 1000000
	max := 1
	maxInd := 1
	clen := make(map[int]int, 2*limit)
	for i := 1; i < limit; i++ {
		len := chainLen(i, clen)
		if len >= max {
			max = len
			maxInd = i
		}
	}
	fmt.Println(maxInd)
}

func chainLen(n int, clen map[int]int) int {
	if n == 1 {
		return 1
	}
	len, ok := clen[n]
	if ok {
		return len
	}
	ret := -1
	if n%2 == 0 {
		ret = chainLen(n/2, clen) + 1
	} else {
		ret = chainLen(3*n+1, clen) + 1
	}
	clen[n] = ret
	return ret
}
