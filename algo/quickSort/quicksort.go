package main

import (
	"fmt"
)

func quickSort(array []int, l, r int) {
	if l < r {
		pos := partition(array, l, r)
		// 递归的时候去掉pos位置的元素
		quickSort(array, l, pos-1)
		quickSort(array, pos+1, r)
	}
}

func partition(array []int, l, r int) int {
	key := array[l]
	for l < r {
		// 从右边开始比较，从左边的话会丢失最后一个数字，可以画一下
		// 需要判断相等的情况
		for array[r] >= key && l < r {
			r--
		}
		array[l] = array[r]
		for array[l] <= key && l < r {
			l++
		}
		array[r] = array[l]
	}
	array[l] = key
	return l
}

func main() {
	array := []int{5, 2, 7, 7, 4, 6, 11, 8, 1}
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}
