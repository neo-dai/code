// 示例程序，展示为什么不能总是
// 获取值的地址。
package main

import "fmt"

// duration 是基本类型为 int 的类型。
type duration int

// format 美化打印 duration 值。
func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}

// main 是应用程序的入口点。
func main() {
	duration(42).pretty()

	// ./listing46.go:17: cannot call pointer method on duration(42)
	// ./listing46.go:17: cannot take the address of duration(42)
}
