// 这个示例程序演示如何在程序中创建竞态条件。
// 我们不希望这样做。
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	// counter 是一个被所有 goroutine 递增的变量。
	counter int

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
	fmt.Println("Final Counter:", counter)
}

// incCounter 递增包级别的 counter 变量。
func incCounter(id int) {
	// 安排调用 Done 以告诉 main 我们已完成。
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 捕获 Counter 的值。
		value := counter

		// 让出线程并放回队列。
		runtime.Gosched()

		// 递增 Counter 的本地值。
		value++

		// 将值存回 Counter。
		counter = value
	}
}
