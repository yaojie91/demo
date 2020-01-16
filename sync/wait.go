package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go fun1(&wg)
	go fun2(&wg)
	fmt.Println("开始等待")
	wg.Wait()
	fmt.Println("解除阻塞")
}

func fun1(wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		fmt.Println("[---func1----", i)
	}
	wg.Done()
}

func fun2(wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < 1000; j++ {
		fmt.Println("----func2-----", j)
	}
}
