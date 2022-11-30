package Medium

// https://leetcode.cn/problems/wiggle-subsequence/
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0 {
		ans++
	}
	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		// 统计上升和下降趋势的变化次数，即峰与谷的次数
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans++
			prevDiff = diff
		}
	}
	return ans
}
