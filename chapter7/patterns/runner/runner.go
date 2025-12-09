// 示例由 Gabriel Aszalos 提供帮助。
// Package runner 管理进程的运行和生命周期。
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定超时时间内运行一组任务，并且可以
// 在操作系统中断时关闭。
type Runner struct {
	// interrupt 通道报告来自
	// 操作系统的信号。
	interrupt chan os.Signal

	// complete 通道报告处理已完成。
	complete chan error

	// timeout 报告时间已用完。
	timeout <-chan time.Time

	// tasks 保存一组按索引顺序
	// 同步执行的函数。
	tasks []func(int)
}

// ErrTimeout 在 timeout 通道上接收到值时返回。
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 在接收到来自操作系统的事件时返回。
var ErrInterrupt = errors.New("received interrupt")

// New 返回一个新的可用 Runner。
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 将任务附加到 Runner。任务是一个
// 接受 int ID 的函数。
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 运行所有任务并监控通道事件。
func (r *Runner) Start() error {
	// 我们想要接收所有基于中断的信号。
	signal.Notify(r.interrupt, os.Interrupt)

	// 在不同的 goroutine 上运行不同的任务。
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 处理完成时发出信号。
	case err := <-r.complete:
		return err

	// 时间用完时发出信号。
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 执行每个已注册的任务。
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检查来自操作系统的中断信号。
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务。
		task(id)
	}

	return nil
}

// gotInterrupt 验证是否已发出中断信号。
func (r *Runner) gotInterrupt() bool {
	select {
	// 发送中断事件时发出信号。
	case <-r.interrupt:
		// 停止接收任何进一步的信号。
		signal.Stop(r.interrupt)
		return true

	// 继续正常运行。
	default:
		return false
	}
}
