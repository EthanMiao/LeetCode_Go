package Medium

import "sort"

// https://leetcode.cn/problems/non-overlapping-intervals/
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	n := len(intervals)
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][1] < intervals[b][1]
	})
	maxRight := intervals[0][1]
	count := 1
	for _, interval := range intervals {
		if interval[0] >= maxRight {
			count++
			maxRight = interval[1]
		}
	}
	return n - count
}
