package TwoPointer

// https://leetcode.cn/problems/remove-element/solutions/
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

// 优化
func removeElement2(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right]
			right--
		} else {
			left++
		}
	}
	return left
}
