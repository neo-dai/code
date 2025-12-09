// 这个示例程序演示如何使用无缓冲通道
// 来模拟四个 goroutine 之间的接力赛。
package main

import (
	"fmt"
	"sync"
	"time"
)

// wg 用于等待程序完成。
var wg sync.WaitGroup

// main 是所有 Go 程序的入口点。
func main() {
	// 创建一个无缓冲通道。
	baton := make(chan int)

	// 为最后一个跑步者添加计数 1。
	wg.Add(1)

	// 第一个跑步者就位。
	go Runner(baton)

	// 开始比赛。
	baton <- 1

	// 等待比赛结束。
	wg.Wait()
}

// Runner 模拟一个人在接力赛中跑步。
func Runner(baton chan int) {
	var newRunner int

	// 等待接收接力棒。
	runner := <-baton

	// 开始绕跑道跑步。
	fmt.Printf("Runner %d Running With Baton\n", runner)

	// 新跑步者上线。
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// 绕跑道跑步。
	time.Sleep(100 * time.Millisecond)

	// 比赛结束了吗。
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// 为下一个跑步者交换接力棒。
	fmt.Printf("Runner %d Exchange With Runner %d\n",
		runner,
		newRunner)

	baton <- newRunner
}
