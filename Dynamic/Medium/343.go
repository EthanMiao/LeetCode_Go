package Medium

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/integer-break/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		max := math.MinInt
		for j := 1; j < i; j++ {
			temp := maxTwo([]int{j * dp[i-j], j * (i - j)})
			if temp > max {
				max = temp
			}
			fmt.Printf("max is %v\n", max)
		}
		dp[i] = max
		fmt.Printf("%v\n", dp)
	}
	return dp[n]
}

func maxTwo(a []int) int {
	max := math.MinInt
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
