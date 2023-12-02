package util

import "time"

//存放时间处理函数

// GetResTime 获取当前时间
func GetResTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetTodayString 获取时间字符串
func GetTodayString() string {
	return time.Now().Format("20060102150405")
}

// GetUnixTime 获取当前时间戳
func GetUnixTime() int64 {
	return time.Now().UnixNano() / 1e6
}
