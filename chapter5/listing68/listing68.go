// 示例程序，展示程序如何访问
// 另一个包中未导出标识符的值。
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing68/counters"
)

// main 是应用程序的入口点。
func main() {
	// 使用 counters 包中导出的 New 函数
	// 创建一个未导出类型的变量。
	counter := counters.New(10)

	fmt.Printf("Counter: %d\n", counter)
}
