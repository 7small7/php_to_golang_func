package pstring

import "unicode/utf8"

// Strlen 统计字符串长度
func Strlen(s string) int {
	if s == "" {
		return 0
	}
	return utf8.RuneCountInString(s)
}
