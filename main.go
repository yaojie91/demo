package main

import (
	"fmt"
	"sync"
)

func process(id int) {
	fmt.Println(id)
	for {
	}
}
func main() {
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go process(i)
	}
	wg.Wait()
}
