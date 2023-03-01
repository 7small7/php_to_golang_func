package pstring

import "strings"

// 将字符串转为大写
func Strtoupper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s)
}
