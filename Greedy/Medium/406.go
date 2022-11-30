package Medium

import (
	"sort"
)

// https://leetcode.cn/problems/queue-reconstruction-by-height/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(a, b int) bool {
		if people[a][0] != people[b][0] {
			// descending
			return people[a][0] > people[b][0]
		} else {
			// Ascending
			return people[a][1] < people[b][1]
		}
	})

	result := make([][]int, 0)
	for _, person := range people {
		idx := person[1]
		result = append(result[:idx], append([][]int{person}, result[idx:]...)...)
	}
	return result
}
