// 示例程序，展示如何无法直接访问
// 导出结构类型中的未导出字段。
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing74/entities"
)

// main 是应用程序的入口点。
func main() {
	// 从 entities 包创建一个 Admin 类型的值。
	a := entities.Admin{
		Rights: 10,
	}

	// 设置来自未导出
	// 内部类型的导出字段。
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}
