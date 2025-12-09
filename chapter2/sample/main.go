package main

import (
	"log"
	"os"

	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
)

// init 在 main 之前被调用。
func init() {
	// 将日志输出设备更改为标准输出。
	log.SetOutput(os.Stdout)
}

// main 是程序的入口点。
func main() {
	// 对指定的术语执行搜索。
	search.Run("president")
}
