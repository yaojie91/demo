package main

import (
	"bytes"
	"fmt"
)

func main() {
	x := []byte("我爱你中国")
	fmt.Println("-------", x)
	r1 := bytes.NewReader(x)
	fmt.Println("============", r1)
	d1 := make([]byte, len(x))
	fmt.Println("_-------", d1)
	n, _ := r1.Read(d1)
	fmt.Println(n, string(d1))

	//var tmp = make([]byte, 30)
	//copy(tmp, x)
	//fmt.Println("---------", tmp)
	//r2 := bytes.Reader{}
	ttt := []byte("诶诶发额")
	r1.Reset(ttt)
	//d2 := make([]byte, len(ttt))
	n, _ = r1.Read(d1)
	fmt.Println(n, string(d1))
}
