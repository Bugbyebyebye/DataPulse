package util

import "time"

//存放时间处理函数

// GetResTime 获取当前时间
func GetResTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
