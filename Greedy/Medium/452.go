package Medium

import (
	"sort"
)

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(a, b int) bool {
		return points[a][1] < points[b][1]
	})
	maxRight := points[0][1]
	count := 1
	for _, point := range points {
		if point[0] > maxRight {
			count++
			maxRight = point[1]
		}
	}
	return count
}
