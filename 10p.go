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
	primes := make(chan int,max)
	sum := make(chan int)
	var done sync.WaitGroup
	go sumIt(primes, sum)
	for i := 3; i < max; i += 2 {
		done.Add(1)
		go func(num int) {
			if utils.IsPrime(num) {
				primes <- num
			}
			done.Done()
		}(i)
	}
	done.Wait()
	close(primes)
	fmt.Println(<-sum)
}

func sumIt(primes, send chan int){
	sum := 2
	for i:= range primes{
		sum+=i
	}
	send <- sum
}
