package Medium

// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/description/
func findLength(nums1 []int, nums2 []int) int {
	length1 := len(nums1)
	length2 := len(nums2)
	dp := make([][]int, length1)
	for i := 0; i <= length1; i++ {
		dp[i] = make([]int, length2)
	}

	result := 0
	for i := 1; i <= length1; i++ {
		for j := 1; j <= length2; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}
	return result
}

// 滑动窗口
func findLength1(nums1 []int, nums2 []int) int {
	// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/solutions/310099/zui-chang-zhong-fu-zi-shu-zu-by-leetcode-solution/
	return 0
}
