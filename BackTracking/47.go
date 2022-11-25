package BackTracking

import "sort"

// https://leetcode.cn/problems/permutations-ii/
var ans7 [][]int
var path7 []int

func permuteUnique(nums []int) [][]int {
	ans7 = make([][]int, 0)
	path7 = make([]int, 0)
	sort.Ints(nums)
	backTracking7(nums, len(nums))
	return ans7
}

func backTracking7(nums []int, numsLen int) {
	if len(path7) == numsLen {
		temp := make([]int, numsLen)
		copy(temp, path7)
		ans7 = append(ans7, temp)
	}
	visited := [21]bool{}
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if visited[cur+10] {
			continue
		}
		visited[cur+10] = true
		path7 = append(path7, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backTracking7(nums, numsLen)
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		path7 = path7[:len(path7)-1]
	}
}
