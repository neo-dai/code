// 这个示例程序演示如何使用 atomic
// 包来提供对数值类型的安全访问。
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// counter 是一个被所有 goroutine 递增的变量。
	counter int64

	// wg 用于等待程序完成。
	wg sync.WaitGroup
)

// main 是所有 Go 程序的入口点。
func main() {
	// 添加计数 2，每个 goroutine 一个。
	wg.Add(2)

	// 创建两个 goroutine。
	go incCounter(1)
	go incCounter(2)

	// 等待 goroutine 完成。
	wg.Wait()

	// 显示最终值。
	fmt.Println("Final Counter:", counter)
}

// incCounter 递增包级别的 counter 变量。
func incCounter(id int) {
	// 安排调用 Done 以告诉 main 我们已完成。
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 安全地将 Counter 加 1。
		atomic.AddInt64(&counter, 1)

		// 让出线程并放回队列。
		runtime.Gosched()
	}
}
