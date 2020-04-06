package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcdkfgehabc"
	s1 := "ebadcgfhcbak"
	fmt.Println(isEqual(s, s1))
}

func isEqual(s, s1 string) bool {
	if len(s) != len(s1) {
		return false
	}
	for _, v := range s {
		n := strings.LastIndex(s1, string(v))
		if n >= 0 {
			s1 = s1[:n] + s1[n+1:]
		} else {
			return false
		}
	}
	return true
}
