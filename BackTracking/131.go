package BackTracking

// https://leetcode.cn/problems/palindrome-partitioning/
var ans2 [][]string
var path2 []string

func partition(s string) [][]string {
	ans2 = make([][]string, 0)
	path2 = make([]string, 0)
	backTrackPalindrome(s, 0)
	return ans2
}

func backTrackPalindrome(s string, startIndex int) {
	if startIndex == len(s) {
		temp := make([]string, len(path2))
		copy(temp, path2)
		ans2 = append(ans2, temp)
	}
	for i := startIndex; i < len(s); i++ {
		if isPalindrome(s, startIndex, i) {
			path2 = append(path2, s[startIndex:i+1])
		} else {
			continue
		}
		backTrackPalindrome(s, i+1)
		path2 = path2[:len(path2)-1]
	}
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
