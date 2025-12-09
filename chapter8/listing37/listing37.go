// 示例程序，展示标准库中的不同函数
// 如何使用 io.Writer 接口。
package main

import (
	"bytes"
	"fmt"
	"os"
)

// main 是应用程序的入口点。
func main() {
	// 创建一个 Buffer 值并将字符串写入缓冲区。
	// 使用实现 io.Writer 的 Write 方法。
	var b bytes.Buffer
	b.Write([]byte("Hello "))

	// 使用 Fprintf 将字符串连接到 Buffer。
	// 为 io.Writer 传递 bytes.Buffer 值的地址。
	fmt.Fprintf(&b, "World!")

	// 将 Buffer 的内容写入标准输出设备。
	// 为 io.Writer 传递 os.File 值的地址。
	b.WriteTo(os.Stdout)
}
