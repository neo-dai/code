// 这个示例程序演示如何使用基本的 log 包。
package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 写入标准记录器。
	log.Println("message")

	// Fatalln 是 Println() 后跟对 os.Exit(1) 的调用。
	log.Fatalln("fatal message")

	// Panicln 是 Println() 后跟对 panic() 的调用。
	log.Panicln("panic message")
}
