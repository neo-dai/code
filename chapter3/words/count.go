// Package words 提供计数单词的支持。
package words

import "strings"

// CountWords 计算指定字符串中
// 单词的数量并返回计数。
func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
