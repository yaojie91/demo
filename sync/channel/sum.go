package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const limit = 1e6

func main() {
	t1 := time.Now()
	sum1 := loopSum()
	fmt.Println("sum1 is: ", sum1)
	fmt.Printf("cost time %d ns\n", time.Now().Sub(t1))

	t2 := time.Now()
	sum2 := ConcurrentSum()
	fmt.Println("sum2 is: ", sum2)
	fmt.Printf("cost time %d ns\n", time.Now().Sub(t2))

	t3 := time.Now()
	sum3 := ChannelSum()
	fmt.Println("sum3 is: ", sum3)
	fmt.Printf("cost time %d ns\n", time.Now().Sub(t3))
}

func ConcurrentSum() int {
	n := runtime.GOMAXPROCS(0)
	sums := make([]int, n)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			start := (limit / n) * i
			end := start + limit/n
			for j := start; j < end; j++ {
				sums[i] += j
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	sum := 0
	for i := 0; i < n; i++ {
		sum += sums[i]
	}
	return sum
}

func loopSum() int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += i
	}
	return sum
}

func ChannelSum() int {
	n := runtime.GOMAXPROCS(0)
	ch := make(chan int)

	for i := 0; i < n; i++ {
		go func(i int, r chan int) {
			sum := 0
			start := (limit / n) * i
			end := start + limit/n
			for j := start; j < end; j++ {
				sum += j
			}
			r <- sum
		}(i, ch)
	}

	rs := 0
	for i := 0; i < n; i++ {
		rs += <-ch
	}
	return rs
}
