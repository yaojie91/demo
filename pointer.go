package main

import "fmt"

func main() {
	ss := []string{"a", "b", "c"}
	pss := &ss
	//fmt.Println("------", ss, &ss)
	//小张=你
	fmt.Printf("数据地址=%p，指针地址=%p\n", ss, pss)

	ss1 := ss
	pss1 := &ss1
	//fmt.Println("-------------", ss, ss1)
	//张三=你
	fmt.Printf("数据地址=%p，指针地址=%p\n", ss1, pss1)

	ss = []string{"d"}
	//小张=另外一个同事
	fmt.Printf("数据地址=%p，指针地址=%p\n", ss, &ss)
}
