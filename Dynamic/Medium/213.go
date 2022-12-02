package Medium

// https://leetcode.cn/problems/house-robber-ii/
func rob1(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	result1 := robRange(nums, 0, length-2)
	result2 := robRange(nums, 1, length-1)
	return max(result1, result2)
}

func robRange(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	length := end - start + 1
	dp := make([]int, length)
	dp[0] = nums[start]
	dp[1] = max(nums[start], nums[start+1])
	for i := 2; i < length; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i+start])
	}
	return dp[length-1]
}
