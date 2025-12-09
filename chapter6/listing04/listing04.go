// 这个示例程序演示 goroutine 调度器
// 如何在单个线程上对 goroutine 进行时间切片。
package main

import (
	"fmt"
	"runtime"
	"sync"
)

// wg 用于等待程序完成。
var wg sync.WaitGroup

// main 是所有 Go 程序的入口点。
func main() {
	// 为调度器分配 1 个逻辑处理器使用。
	runtime.GOMAXPROCS(1)

	// 添加计数 2，每个 goroutine 一个。
	wg.Add(2)

	// 创建两个 goroutine。
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// 等待 goroutine 完成。
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime 显示前 5000 个数字的质数。
func printPrime(prefix string) {
	// 安排调用 Done 以告诉 main 我们已完成。
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
