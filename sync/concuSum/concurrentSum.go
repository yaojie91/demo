package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 5

var mu sync.Mutex
var wg sync.WaitGroup

func main() {
	sl := getSl()
	sum := 0
	wg.Add(N)
	seg := len(sl) / N
	for k := 0; k < N; k++ {
		go func(kk int) {
			mu.Lock()
			defer mu.Unlock()
			sum += cal(sl[kk*seg : (kk+1)*seg])
			wg.Done()
		}(k)
	}
	wg.Wait()
	fmt.Println(sum)
}

func getSl() []int {
	rand.Seed(time.Now().UnixNano())
	intSl := make([]int, 100)
	for i := 0; i < 100; i++ {
		intSl[i] = i
	}
	return intSl
}

func cal(sl []int) int {
	sum := 0
	for _, v := range sl {
		sum += v
	}
	return sum
}
