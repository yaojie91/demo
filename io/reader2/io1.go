package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("acb我爱你中国")
	fmt.Println("-----------", *reader)
	p := make([]byte, 20)
	n, err := reader.ReadAt(p, 1)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("--%c-------", p)
	fmt.Println("-------", p)
	fmt.Printf("------------%c----------", bytes.Runes(p))
	fmt.Println("-----string-----------", string(p))
	fmt.Println("---111-----", []rune(string(p)))
	t := []rune(string(p))
	fmt.Printf("---%d--%c----", n, t)
}
