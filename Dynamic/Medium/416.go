package Medium

import "fmt"

// https://leetcode.cn/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)
	for _, v := range nums {
		for i := target; i >= v; i-- {
			dp[i] = max(dp[i], dp[i-v]+v)
		}
		fmt.Printf("%v\n", dp)
	}
	return dp[target] == target
}
