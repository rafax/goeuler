package main

import (
	"./utils"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	max := 2000000
	primes := findPrimes(max)
	sum := 2
	for i := range primes {
		sum += i
	}
	fmt.Println(sum)
}

func findPrimes(max int) <-chan int {
	primes := make(chan int, max)
	var done sync.WaitGroup
	for i := 3; i < max; i += 2 {
		done.Add(1)
		go func(num int) {
			if utils.IsPrime(num) {
				primes <- num
			}
			done.Done()
		}(i)
	}
	go func() {
		done.Wait()
		close(primes)
	}()
	return primes
}
