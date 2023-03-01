package ptime

import "time"

// 获取当前日期
func date() string {
	return time.Now().Format("2006-01-01")
}
