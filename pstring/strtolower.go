package pstring

import "strings"

// 将字符串转为小写
func Strtolower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s)
}
