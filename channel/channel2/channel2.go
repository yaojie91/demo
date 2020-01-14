package main

import (
	"fmt"
	"time"
)

// 两种循环遍历通道
// for 如果没有判断是否关闭，channel关闭之后可以继续取值，值为对应类型的nil，比如int类型，0
// for range 如果channel关闭，可以在取完值之后跳出循环
func main() {
	ch := make(chan int)
	go genNum(ch)
	//for {
	//	if v, ok := <-ch; ok {
	//		fmt.Println("-------", v)
	//	} else {
	//		fmt.Println("---111----", v)
	//	}
	//	//fmt.Println("--------", <-ch)
	//	time.Sleep(1e9)
	//}
	for v := range ch {
		fmt.Println("[-", v)
		time.Sleep(1e9)
	}
}

func genNum(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
