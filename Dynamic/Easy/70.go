package Easy

// https://leetcode.cn/problems/climbing-stairs/
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 1, 2, 0
	for i := 3; i <= n; i++ {
		c = a + b
		a, b = b, c
	}
	return c
}
