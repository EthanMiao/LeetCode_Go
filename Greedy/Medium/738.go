package Medium

import (
	"strconv"
)

// https://leetcode.cn/problems/monotone-increasing-digits/
func monotoneIncreasingDigits(n int) int {
	s := strconv.Itoa(n)
	ss := []byte(s)
	length := len(ss)
	if length <= 1 {
		return n
	}
	for i := length - 1; i > 0; i-- {
		if ss[i-1] > ss[i] {
			ss[i-1] -= 1
			for j := i; j < length; j++ {
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}
