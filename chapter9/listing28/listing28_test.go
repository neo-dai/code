// 示例基准测试，以测试哪个函数更适合将
// 整数转换为字符串。首先使用 fmt.Sprintf 函数，
// 然后是 strconv.FormatInt 函数，然后是 strconv.Itoa。
package listing05_test

import (
	"fmt"
	"strconv"
	"testing"
)

// BenchmarkSprintf 为 fmt.Sprintf 函数
// 提供性能数据。
func BenchmarkSprintf(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat 为 strconv.FormatInt 函数
// 提供性能数据。
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa 为 strconv.Itoa 函数
// 提供性能数据。
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
