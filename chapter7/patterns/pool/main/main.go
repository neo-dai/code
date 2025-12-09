// 这个示例程序演示如何使用 pool 包
// 来共享模拟的一组数据库连接。
package main

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/goinaction/code/chapter7/patterns/pool"
)

const (
	maxGoroutines   = 25 // 要使用的例程数量。
	pooledResources = 2  // 池中的资源数量
)

// dbConnection 模拟要共享的资源。
type dbConnection struct {
	ID int32
}

// Close 实现 io.Closer 接口，以便 dbConnection
// 可以由池管理。Close 执行任何资源
// 释放管理。
func (dbConn *dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

// idCounter 为每个连接提供唯一 id 的支持。
var idCounter int32

// createConnection 是一个工厂方法，当需要
// 新连接时，池将调用它。
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id}, nil
}

// main 是所有 Go 程序的入口点。
func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建池来管理我们的连接。
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 使用池中的连接执行查询。
	for query := 0; query < maxGoroutines; query++ {
		// 每个 goroutine 需要自己的查询值副本，
		// 否则它们都将共享相同的查询
		// 变量。
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	// 等待 goroutine 完成。
	wg.Wait()

	// 关闭池。
	log.Println("Shutdown Program.")
	p.Close()
}

// performQueries 测试连接的资源池。
func performQueries(query int, p *pool.Pool) {
	// 从池中获取连接。
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	// 将连接释放回池。
	defer p.Release(conn)

	// 等待以模拟查询响应。
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
