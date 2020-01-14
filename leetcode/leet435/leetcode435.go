package leet435

import (
	"fmt"
	"sort"
)

func main() {
	params := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	result := eraseOverlapIntervals(params)
	fmt.Println("------------------", result)
}

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	end := intervals[0][1]
	minCount := 0
	for i := 1; i < len(intervals); i++ {
		if end > intervals[i][0] {
			if end > intervals[i][1] {
				end = intervals[i][1]
			}
			minCount++
		} else {
			end = intervals[i][1]
		}
	}
	return minCount
}
