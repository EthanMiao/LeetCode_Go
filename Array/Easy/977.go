package Easy

// https://leetcode.cn/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	len := len(nums)
	result := make([]int, len)
	left, right := 0, len-1

	for pos := len - 1; pos >= 0; pos-- {
		numLeft := nums[left] * nums[left]
		numRight := nums[right] * nums[right]
		if numLeft <= numRight {
			result[pos] = numRight
			right--
		} else {
			result[pos] = numLeft
			left++
		}
	}
	return result
}
