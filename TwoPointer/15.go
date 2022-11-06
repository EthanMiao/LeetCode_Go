package TwoPointer

import "sort"

// https://leetcode.cn/problems/3sum/
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n-2; first++ {
		n1 := nums[first]
		if n1 > 0 {
			return ans
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := n - 1
		for second < third {
			n2, n3 := nums[second], nums[third]
			total := n1 + n2 + n3
			if total == 0 {
				ans = append(ans, []int{n1, n2, n3})
				for second < third && nums[second] == n2 {
					second++
				}
				for second < third && nums[third] == n3 {
					third--
				}
			} else if total < 0 {
				second++
			} else {
				third--
			}
		}
	}
	return ans
}
