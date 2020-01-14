package main

import "fmt"

func main() {
	ch := make(chan int)
	done := make(chan int)
	go prod(ch)
	go print(ch, done)
	<-done
}

func prod(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i * 10
	}
	close(c)
}

func print(ch chan int, done chan int) {
	for v := range ch {
		fmt.Println("---------", v)
	}
	done <- 1
}
