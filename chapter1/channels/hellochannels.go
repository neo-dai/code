package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printer(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d ", i)
	}
	wg.Done()
}

// main 是程序的入口点。
func main() {
	c := make(chan int)
	go printer(c)
	wg.Add(1)

	// 在通道上发送 10 个整数。
	for i := 1; i <= 10; i++ {
		c <- i
	}

	close(c)
	wg.Wait()
}
