// 示例程序，展示如何使用 io.Reader 和
// io.Writer 接口支持编写 curl 的简单版本。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// init 在 main 之前被调用。
func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./example2 <url>")
		os.Exit(-1)
	}
}

// main 是应用程序的入口点。
func main() {
	// 从 web 服务器获取响应。
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 从 Body 复制到标准输出。
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
