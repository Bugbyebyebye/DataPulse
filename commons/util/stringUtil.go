package util

// In 判断 字符串是否存在于这个列表
func In(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}
