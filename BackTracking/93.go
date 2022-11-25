package BackTracking

import "strconv"

// https://leetcode.cn/problems/restore-ip-addresses/
var result []string
var ip []string

func restoreIpAddresses(s string) []string {
	result = make([]string, 0)
	ip = make([]string, 0)
	backTrackIp(s, 0)
	return result
}

func backTrackIp(s string, index int) {
	if len(ip) == 4 && index == len(s) {
		result = append(result, convertToIp(ip))
		return
	}
	if len(ip) == 4 {
		return
	}
	if index == len(s) {
		return
	}
	for i := index; i < len(s); i++ {
		subStr := s[index : i+1]
		if string(subStr[0]) == "0" && len(subStr) != 1 {
			return
		}
		num, _ := strconv.Atoi(subStr)
		if num >= 0 && num <= 255 {
			ip = append(ip, subStr)
			backTrackIp(s, i+1)
			ip = ip[:len(ip)-1]
		} else {
			return
		}
	}
}

func convertToIp(path []string) string {
	result := ""
	for _, v := range path {
		result += v
		result += "."
	}
	return result[:len(result)-1]
}
