package Medium

import "math"

// https://leetcode.cn/problems/maximum-subarray/
func maxSubArray(nums []int) int {
	sum := 0
	max := math.MinInt
	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += v
		if max < sum {
			max = sum
		}
	}
	return max
}
