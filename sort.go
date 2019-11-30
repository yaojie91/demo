package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	array := make([]int, 0)
	for i := 0; i < 500; i++ {
		num := rand.Intn(1000)
		array = append(array, num)
	}

	//fmt.Println("---array----", array)
	// 冒泡法
	tBegin := time.Now().UnixNano()
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	//fmt.Println(array)
	tEnd := time.Now().UnixNano()
	fmt.Println("---冒泡------", tEnd-tBegin)

	// 库函数
	timeBegin := time.Now().UnixNano()
	sort.Ints(array)
	timeEnd := time.Now().UnixNano()
	fmt.Println("-------库函数--------", timeEnd-timeBegin)

	// 堆排序
	tHeapBegin := time.Now().UnixNano()
	duiSort(array)
	tHeapEnd := time.Now().UnixNano()
	//fmt.Println("----", array)
	fmt.Println("-----堆排序-----", tHeapEnd-tHeapBegin)
}

func duiSort(array []int) {
	n := len(array) - 1
	for root := n / 2; root >= 0; root-- {
		minHeap(root, n, array)
	}
	for end := n; end >= 0; end-- {
		if array[0] > array[end] {
			array[0], array[end] = array[end], array[0]
			minHeap(0, end, array)
		}
	}
}

func minHeap(root int, end int, arr []int) {
	child := 2*root + 1
	for {
		if child > end {
			break
		}
		if child+1 <= end && arr[child] > arr[child+1] {
			child++
		}
		if arr[root] > arr[child] {
			arr[root], arr[child] = arr[child], arr[root]
			root = child
		} else {
			break
		}
	}
}
