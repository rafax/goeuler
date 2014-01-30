package main

import "fmt"

func main() {
	limit := 21
	arr := make([]int, limit*limit, limit*limit)
	arr[0] = 1
	for i := 0; i < limit; i++ {
		for j := 0; j < limit; j++ {
			if i > 0 {
				arr[i*limit+j] += arr[(i-1)*limit+j]
			}
			if j > 0 {
				arr[i*limit+j] += arr[i*limit+j-1]
			}
		}
	}
	fmt.Println(arr[limit*limit-1])
}
