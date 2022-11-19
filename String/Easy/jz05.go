package Easy

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
func replaceSpace(s string) string {
	b := []byte(s)
	result := make([]byte, 0)
	for i := 0; i < len(b); i++ {
		if b[i] == ' ' {
			result = append(result, []byte("%20")...)
		} else {
			result = append(result, b[i])
		}
	}
	return string(result)
}

// 原地修改
func replaceSpace2(s string) string {
	b := []byte(s)
	length := len(b)
	spaceCount := 0
	// count number of space
	for _, v := range b {
		if v == ' ' {
			spaceCount++
		}
	}
	// extend slice
	tmp := make([]byte, 2*spaceCount)
	b = append(b, tmp...)
	i := length - 1
	j := len(b) - 1
	for i >= 0 {
		if b[i] != ' ' {
			b[j] = b[i]
			i--
			j--
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			i--
			j = j - 3
		}
	}
	return string(b)
}
