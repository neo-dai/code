// 示例由 Fatih Arslan 和 Gabriel Aszalos 提供帮助。
// Package pool 管理用户定义的一组资源。
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 管理一组可以被多个 goroutine
// 安全共享的资源。被管理的资源必须
// 实现 io.Closer 接口。
type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 在已关闭的池上
// 调用 Acquire 时返回。
var ErrPoolClosed = errors.New("Pool has been closed.")

// New 创建一个管理资源的池。池需要一个
// 可以分配新资源的函数以及
// 池的大小。
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// Acquire 从池中检索资源。
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有空闲资源。
	case r, ok := <-p.resources:
		log.Println("Acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	// 由于没有可用资源，提供新资源。
	default:
		log.Println("Acquire:", "New Resource")
		return p.factory()
	}
}

// Release 将新资源放入池中。
func (p *Pool) Release(r io.Closer) {
	// 用 Close 操作保护此操作。
	p.m.Lock()
	defer p.m.Unlock()

	// 如果池已关闭，丢弃资源。
	if p.closed {
		r.Close()
		return
	}

	select {
	// 尝试将新资源放入队列。
	case p.resources <- r:
		log.Println("Release:", "In Queue")

	// 如果队列已满，我们关闭资源。
	default:
		log.Println("Release:", "Closing")
		r.Close()
	}
}

// Close 将关闭池并关闭所有现有资源。
func (p *Pool) Close() {
	// 用 Release 操作保护此操作。
	p.m.Lock()
	defer p.m.Unlock()

	// 如果池已经关闭，什么都不做。
	if p.closed {
		return
	}

	// 将池设置为已关闭。
	p.closed = true

	// 在耗尽通道中的资源之前
	// 关闭通道。如果我们不这样做，将会出现死锁。
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}
