package Easy

// https://leetcode.cn/problems/reverse-string-ii/
func reverseStr(s string, k int) string {
	t := []byte(s)
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		sub := t[i:min(i+k, n)]
		for j, k := 0, len(sub); j < k/2; j++ {
			sub[j], sub[k-1-j] = sub[k-1-j], sub[j]
		}
	}
	return string(t)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
