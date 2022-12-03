package Medium

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfit3(prices []int) int {
	length := len(prices)
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 3)
	}
	dp[0][0] = -prices[0]
	// dp[i] 代表第i天的累计最大收益
	// dp[i][0] 代表买入状态
	// dp[i][1] 代表卖出状态，并处于冷冻期
	// dp[i][2] 代表卖出状态，不处于冷冻期
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[length-1][1], dp[length-1][2])
}
