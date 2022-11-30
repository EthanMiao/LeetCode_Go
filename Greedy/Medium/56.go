package Medium

import (
	"sort"
)

// https://leetcode.cn/problems/merge-intervals/description/
func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	if len(intervals) == 0 {
		return ans
	}
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	start, end := intervals[0][0], intervals[0][1]
	for _, interval := range intervals {
		if interval[0] > end {
			ans = append(ans, []int{start, end})
			start = interval[0]
			end = interval[1]
		} else {
			end = maxEnd(end, interval[1])
		}
	}
	ans = append(ans, []int{start, end})
	return ans
}

func maxEnd(a, b int) int {
	if a > b {
		return a
	}
	return b
}
