package BackTracking

// https://leetcode.cn/problems/subsets/
var result3 [][]int
var path3 []int

func subsets(nums []int) [][]int {
	result3 = make([][]int, 0)
	path3 = make([]int, 0)
	backTrackingSublist(nums, 0)
	return result3
}

func backTrackingSublist(nums []int, startIndex int) {
	temp := make([]int, len(path3))
	copy(temp, path3)
	result3 = append(result3, temp)
	if len(nums) == startIndex {
		return
	}
	for i := startIndex; i < len(nums); i++ {
		path3 = append(path3, nums[i])
		backTrackingSublist(nums, i+1)
		path3 = path3[:len(path3)-1]
	}
}
