package Easy

// https://leetcode.cn/problems/happy-number/
func isHappy(n int) bool {
	var hashMap = map[int]bool{}
	for n != 1 && !hashMap[n] {
		hashMap[n] = true
		n = nextStep(n)
	}
	return n == 1
}

func nextStep(n int) int {
	sum := 0
	for n > 0 {
		num := n % 10
		sum += num * num
		n /= 10
	}
	return sum
}
