// 这个示例程序演示如何使用 work 包
// 来使用 goroutine 池来完成工作。
package main

import (
	"log"
	"sync"
	"time"

	"github.com/goinaction/code/chapter7/patterns/work"
)

// names 提供要显示的一组名称。
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// namePrinter 为打印名称提供特殊支持。
type namePrinter struct {
	name string
}

// Task 实现 Worker 接口。
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

// main 是所有 Go 程序的入口点。
func main() {
	// 创建一个有 2 个 goroutine 的工作池。
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// 迭代名称切片。
		for _, name := range names {
			// 创建一个 namePrinter 并提供
			// 特定名称。
			np := namePrinter{
				name: name,
			}

			go func() {
				// 提交要处理的任务。当 RunTask
				// 返回时，我们知道它正在被处理。
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	// 关闭工作池并等待所有现有工作
	// 完成。
	p.Shutdown()
}
