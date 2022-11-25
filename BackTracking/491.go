package BackTracking

// https://leetcode.cn/problems/increasing-subsequences/
var ans5 [][]int
var path5 []int

func findSubsequences(nums []int) [][]int {
	ans5 = make([][]int, 0)
	path5 = make([]int, 0)
	backTracking5(nums, 0)
	return ans5
}

func backTracking5(nums []int, index int) {
	if len(path5) >= 2 {
		temp := make([]int, len(path5))
		copy(temp, path5)
		ans5 = append(ans5, temp)
	}
	if len(nums) == index {
		return
	}
	history := make([]int, 201)
	for i := index; i < len(nums); i++ {
		if len(path5) > 0 && nums[i] < path5[len(path5)-1] {
			continue
		}
		if history[nums[i]+100] == 1 {
			continue
		}
		history[nums[i]+100] = 1
		path5 = append(path5, nums[i])
		backTracking5(nums, i+1)
		path5 = path5[:len(path5)-1]
	}
}
