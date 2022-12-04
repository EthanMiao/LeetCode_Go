package Medium

import "math"

// https://leetcode.cn/problems/minimum-size-subarray-sum/
func minSubArrayLen(target int, nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}

	ans := math.MaxInt
	sum := 0
	start, end := 0, 0
	for end < len {
		sum += nums[end]
		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
