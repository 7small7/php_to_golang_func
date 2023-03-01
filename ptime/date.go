package ptime

import "time"

func CurrentDate() string {
	return time.Now().Format("2006-01-01")
}