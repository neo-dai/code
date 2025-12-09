// 这个示例程序演示如何使用 atomic
// 包的 Store 和 Load 函数来提供对
// 数值类型的安全访问。
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	// shutdown 是一个标志，用于通知正在运行的 goroutine 关闭。
	shutdown int64

	// wg 用于等待程序完成。
	wg sync.WaitGroup
)

// main 是所有 Go 程序的入口点。
func main() {
	// 添加计数 2，每个 goroutine 一个。
	wg.Add(2)

	// 创建两个 goroutine。
	go doWork("A")
	go doWork("B")

	// 给 goroutine 运行的时间。
	time.Sleep(1 * time.Second)

	// 安全地标记现在是关闭的时间。
	fmt.Println("Shutdown Now")
	atomic.StoreInt64(&shutdown, 1)

	// 等待 goroutine 完成。
	wg.Wait()
}

// doWork 模拟一个执行工作的 goroutine
// 并检查 Shutdown 标志以提前终止。
func doWork(name string) {
	// 安排调用 Done 以告诉 main 我们已完成。
	defer wg.Done()

	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond)

		// 我们需要关闭吗。
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}
