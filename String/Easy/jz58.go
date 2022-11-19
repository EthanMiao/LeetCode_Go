package Easy

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/description/
func reverseLeftWords(s string, n int) string {
	l := len(s)
	if l <= n {
		return s
	}
	b := []byte(s)
	return string(b[n:]) + string(b[:n])
}

// 原地反转
func reverseLeftWords2(s string, n int) string {
	b := []byte(s)
	reverse(b, 0, n-1)
	reverse(b, n, len(b)-1)
	reverse(b, 0, len(b)-1)
	return string(b)
}

func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
