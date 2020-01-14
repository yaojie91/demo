package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "abc一二三🌚"
	strByte := []byte(str)
	for i := 0; i < len(str); {
		c, size := utf8.DecodeRune(strByte[i:])
		fmt.Printf("---%d--%c---\n", i, c)
		i += size
	}
}
