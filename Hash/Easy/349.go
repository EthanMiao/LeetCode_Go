package Easy

// https://leetcode.cn/problems/intersection-of-two-arrays/
func intersection(nums1 []int, nums2 []int) []int {
	var set1 = map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	var set2 = map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	var result []int
	for v := range set1 {
		if _, has := set2[v]; has {
			result = append(result, v)
		}
	}
	return result
}
