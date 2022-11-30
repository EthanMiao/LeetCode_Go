package Medium

// https://leetcode.cn/problems/jump-game-ii/
func jump(nums []int) int {
	n := len(nums)
	step, curDistance, nextDistance := 0, 0, 0
	for i := 0; i < n-1; i++ {
		nextDistance = max(nextDistance, nums[i]+i)
		if i == curDistance {
			step++
			curDistance = nextDistance
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
