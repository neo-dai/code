// 这个示例程序演示如何使用有缓冲通道
// 使用预定义数量的 goroutine
// 处理多个任务。
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberGoroutines = 4  // 要使用的 goroutine 数量。
	taskLoad         = 10 // 要处理的工作量。
)

// wg 用于等待程序完成。
var wg sync.WaitGroup

// init 在执行任何其他代码之前由
// Go 运行时调用以初始化包。
func init() {
	// 为随机数生成器设置种子。
	rand.Seed(time.Now().Unix())
}

// main 是所有 Go 程序的入口点。
func main() {
	// 创建一个有缓冲通道来管理任务负载。
	tasks := make(chan string, taskLoad)

	// 启动 goroutine 来处理工作。
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 添加一堆要完成的工作。
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 关闭通道，以便在所有工作完成时
	// goroutine 会退出。
	close(tasks)

	// 等待所有工作完成。
	wg.Wait()
}

// worker 作为 goroutine 启动，以从
// 有缓冲通道处理工作。
func worker(tasks chan string, worker int) {
	// 报告我们刚刚返回。
	defer wg.Done()

	for {
		// 等待分配工作。
		task, ok := <-tasks
		if !ok {
			// 这意味着通道为空且已关闭。
			fmt.Printf("Worker: %d : Shutting Down\n", worker)
			return
		}

		// 显示我们正在开始工作。
		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		// 随机等待以模拟工作时间。
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 显示我们完成了工作。
		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
