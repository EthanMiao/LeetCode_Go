package Medium

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/target-sum/
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if (sum+target)%2 == 1 {
		return 0
	}

	if abs(target) > sum {
		return 0
	}

	bag := (sum + target) / 2
	dp := make([]int, bag+1)
	dp[0] = 1
	for _, v := range nums {
		for j := bag; j >= v; j-- {
			dp[j] += dp[j-v]
		}
		fmt.Printf("%v\n", dp)
	}
	return dp[bag]
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
