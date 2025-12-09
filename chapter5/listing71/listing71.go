// 示例程序，展示如何无法直接访问
// 导出结构类型中的未导出字段。
package main

import (
	"fmt"

	"github.com/goinaction/code/chapter5/listing71/entities"
)

// main 是应用程序的入口点。
func main() {
	// 从 entities 包创建一个 User 类型的值。
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com",
	}

	// ./example71.go:16: unknown entities.User field 'email' in
	//                    struct literal

	fmt.Printf("User: %v\n", u)
}
