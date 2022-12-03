package Easy

// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/
func findLengthOfLCIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}

	result := 1
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

// 优化空间
func findLengthOfLCIS1(nums []int) int {
	a, result := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			a += 1
			if result < a {
				result = a
			}
		} else {
			a = 1
		}
	}
	return result
}
