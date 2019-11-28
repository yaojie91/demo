// 在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
// 示例：输入: [3,2,1,5,6,4] 和 k = 2，输出：5
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	t := findKthLargest(nums, 2)
	fmt.Println("-------", t)
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
