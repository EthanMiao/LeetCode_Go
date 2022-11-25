package BackTracking

import "sort"

// https://leetcode.cn/problems/subsets-ii/
var ans4 [][]int
var path4 []int

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ans4 = make([][]int, 0)
	path4 = make([]int, 0)
	backTracking4(nums, 0)
	return ans4
}

func backTracking4(nums []int, index int) {
	temp := make([]int, len(path4))
	copy(temp, path4)
	ans4 = append(ans4, temp)
	if index == len(nums) {
		return
	}
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		path4 = append(path4, nums[i])
		backTracking4(nums, i+1)
		path4 = path4[:len(path4)-1]
	}
}
