package ptime

import "time"

func date() string {
	return time.Now().Format("2006-01-01")
}