// 示例由 Jason Waldrip 提供帮助。
// Package work 管理一个 goroutine 池来执行工作。
package work

import "sync"

// Worker 必须由想要使用
// 工作池的类型实现。
type Worker interface {
	Task()
}

// Pool 提供一个 goroutine 池，可以执行
// 提交的任何 Worker 任务。
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 创建一个新的工作池。
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run 向池提交工作。
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有 goroutine 关闭。
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
