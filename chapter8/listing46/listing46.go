// 示例程序，展示如何使用 io.Reader 和
// io.Writer 接口支持编写 curl 的简单版本。
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// main 是应用程序的入口点。
func main() {
	// 这里的 r 是一个响应，r.Body 是一个 io.Reader。
	r, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// 创建一个文件来持久化响应。
	file, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// 使用 MultiWriter，以便我们可以在同一写入操作中
	// 写入标准输出和文件。
	dest := io.MultiWriter(os.Stdout, file)

	// 读取响应并写入两个目标。
	io.Copy(dest, r.Body)
	if err := r.Body.Close(); err != nil {
		log.Println(err)
	}
}
