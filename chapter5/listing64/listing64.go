// 示例程序，展示程序如何无法访问
// 另一个包中的未导出标识符。
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing64/counters"
)

// main 是应用程序的入口点。
func main() {
	// 创建一个未导出类型的变量
	// 并将值初始化为 10。
	counter := counters.alertCounter(10)

	// ./listing64.go:15: cannot refer to unexported name
	//                                         counters.alertCounter
	// ./listing64.go:15: undefined: counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
}
