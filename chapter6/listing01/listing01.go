// 这个示例程序演示如何创建 goroutine
// 以及调度器的行为。
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// main 是所有 Go 程序的入口点。
func main() {
	// 为调度器分配 1 个逻辑处理器使用。
	runtime.GOMAXPROCS(1)

	// wg 用于等待程序完成。
	// 添加计数 2，每个 goroutine 一个。
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数并创建一个 goroutine。
	go func() {
		// 安排调用 Done 以告诉 main 我们已完成。
		defer wg.Done()

		// 显示字母表三次
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 声明一个匿名函数并创建一个 goroutine。
	go func() {
		// 安排调用 Done 以告诉 main 我们已完成。
		defer wg.Done()

		// 显示字母表三次
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	// 等待 goroutine 完成。
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
