// 这个示例程序演示如何使用无缓冲通道
// 来模拟两个 goroutine 之间的网球比赛。
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg 用于等待程序完成。
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// main 是所有 Go 程序的入口点。
func main() {
	// 创建一个无缓冲通道。
	court := make(chan int)

	// 添加计数 2，每个 goroutine 一个。
	wg.Add(2)

	// 启动两名球员。
	go player("Nadal", court)
	go player("Djokovic", court)

	// 开始比赛。
	court <- 1

	// 等待比赛结束。
	wg.Wait()
}

// player 模拟一个人打网球。
func player(name string, court chan int) {
	// 安排调用 Done 以告诉 main 我们已完成。
	defer wg.Done()

	for {
		// 等待球被击回给我们。
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了。
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选择一个随机数，看看我们是否错过了球。
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道以表示我们输了。
			close(court)
			return
		}

		// 显示然后将击球次数加 1。
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球击回对手。
		court <- ball
	}
}
