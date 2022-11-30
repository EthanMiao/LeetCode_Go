package Medium

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit(prices []int) int {
	profit := 0
	if len(prices) <= 1 {
		return profit
	}
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}
