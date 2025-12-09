// 这个示例程序演示如何实现一个模式，
// 用于从不同系统并发请求结果，然后
// 等待所有结果返回或仅等待第一个。
package main

import (
	"log"

	"github.com/goinaction/code/chapter7/patterns/search"
)

// main 是所有 Go 程序的入口点。
func main() {
	// 提交搜索并显示结果。
	results := search.Submit(
		"golang",
		search.OnlyFirst,
		search.Google,
		search.Bing,
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("main : Results : Info : %+v\n", result)
	}

	// 这次我们要等待所有结果。
	results = search.Submit(
		"golang",
		search.Google,
		search.Bing,
		search.Yahoo,
	)

	for _, result := range results {
		log.Printf("main : Results : Info : %+v\n", result)
	}
}
