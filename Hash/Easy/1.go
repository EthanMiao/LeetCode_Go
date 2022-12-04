package Easy

// https://leetcode.cn/problems/two-sum/
func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, x := range nums {
		if p, ok := hashMap[target-x]; ok {
			return []int{p, i}
		}
		hashMap[x] = i
	}
	return nil
}
