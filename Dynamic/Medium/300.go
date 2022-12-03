package Medium

// https://leetcode.cn/problems/longest-increasing-subsequence/
func lengthOfLIS(nums []int) int {
	length := len(nums)
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}

	result := 1
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] > result {
			result = dp[i]
		}
	}
	return result
}
