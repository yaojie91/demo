package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{3, 7, 6, 18, 10, 8}
	s := []int{3, 8, 15, 4, 6}
	r := findContentChildren(g, s)
	fmt.Println("==", r)
}

// 贪心算法
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gC, sC := 0, 0
	for gC < len(g) && sC < len(s) {
		if g[gC] <= s[sC] {
			gC++
		}
		sC++
	}
	return gC
}
