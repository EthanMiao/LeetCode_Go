package Medium

import "regexp"

// https://leetcode.cn/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
	b := []byte(s)

	var reverse func(start, end int)
	reverse = func(start, end int) {
		for start < end {
			b[start], b[end] = b[end], b[start]
			start++
			end--
		}
	}

	reverse(0, len(b)-1)

	start, end, i := 0, 0, 0
	result := ""
	for i < len(b) {
		if b[i] != ' ' {
			start = i
			for i < len(b) && b[i] != ' ' {
				i++
			}
			end = i - 1
			reverse(start, end)
			result = result + " " + string(b[start:end+1])
		}
		i++
	}
	return result[1:]
}

func reverseWords2(s string) string {
	reg := regexp.MustCompile("\\s+")
	split := reg.Split(s, -1)
	result := ""
	for i := len(split) - 1; i >= 0; i-- {
		if split[i] == "" {
			continue
		}
		result += split[i] + " "
	}
	return result[:len(result)-1]
}
