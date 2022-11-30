package Easy

import (
	"sort"
)

// https://leetcode.cn/problems/assign-cookies/
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	count := 0
	for gIndex, sIndex := 0, 0; gIndex < len(g) && sIndex < len(s); {
		if s[sIndex] >= g[gIndex] {
			gIndex++
			sIndex++
			count++
		} else {
			sIndex++
		}
	}
	return count
}
