// 示例程序，展示如何将 bytes.Buffer
// 与 io.Copy 函数一起使用。
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// main 是应用程序的入口点。
func main() {
	var b bytes.Buffer

	// 将字符串写入缓冲区。
	b.Write([]byte("Hello"))

	// 使用 Fprintf 将字符串连接到 Buffer。
	fmt.Fprintf(&b, "World!")

	// 将 Buffer 的内容写入标准输出。
	io.Copy(os.Stdout, &b)
}
