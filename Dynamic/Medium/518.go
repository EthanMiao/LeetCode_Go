package Medium

import "fmt"

// https://leetcode.cn/problems/coin-change-ii/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, v := range coins {
		for j := v; j <= amount; j++ {
			dp[j] += dp[j-v]
			fmt.Printf("%v\n", dp)
		}
		fmt.Println("##############")
	}
	return dp[amount]
}
