package BackTracking

// https://leetcode.cn/problems/permutations/
var ans6 [][]int
var path6 []int

func permute(nums []int) [][]int {
	ans6 = make([][]int, 0)
	path6 = make([]int, 0)
	backTracking6(nums, len(nums))
	return ans6
}

func backTracking6(nums []int, numLen int) {
	if len(path6) == numLen {
		temp := make([]int, len(path6))
		copy(temp, path6)
		ans6 = append(ans6, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		path6 = append(path6, cur)
		nums = append(nums[:i], nums[i+1:]...)
		backTracking6(nums, numLen)
		// 回溯的时候，切片也需要复原
		nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
		path6 = path6[:len(path6)-1]
	}
}
