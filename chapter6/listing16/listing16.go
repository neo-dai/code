// 这个示例程序演示如何使用互斥锁
// 来定义需要同步访问的
// 关键代码段。
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

	// mutex 用于定义关键代码段。
	mutex sync.Mutex
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
	fmt.Printf("Final Counter: %d\n", counter)
}

// incCounter 使用 Mutex 递增包级别的 Counter 变量
// 以同步并提供安全访问。
func incCounter(id int) {
	// 安排调用 Done 以告诉 main 我们已完成。
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// 一次只允许一个 goroutine 通过
		// 这个关键代码段。
		mutex.Lock()
		{
			// 捕获 counter 的值。
			value := counter

			// 让出线程并放回队列。
			runtime.Gosched()

			// 递增 counter 的本地值。
			value++

			// 将值存回 counter。
			counter = value
		}
		mutex.Unlock()
		// 释放锁并允许任何
		// 等待的 goroutine 通过。
	}
}
