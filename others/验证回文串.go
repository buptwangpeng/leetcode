package others

import "strings"

// leetcode 125. 验证回文串 简单

func IsPalindrome(s string) bool {
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}
	sgood = strings.ToLower(sgood)
	left := 0
	right := len(sgood) - 1
	for left < right {
		if sgood[left] != sgood[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
