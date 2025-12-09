// 这个示例程序演示如何使用通道
// 监控程序运行的时间量，并在
// 运行时间过长时终止程序。
package main

import (
	"log"
	"os"
	"time"

	"github.com/goinaction/code/chapter7/patterns/runner"
)

// timeout 是程序完成所需的秒数。
const timeout = 3 * time.Second

// main 是程序的入口点。
func main() {
	log.Println("Starting work.")

	// 为此次运行创建新的计时器值。
	r := runner.New(timeout)

	// 添加要运行的任务。
	r.Add(createTask(), createTask(), createTask())

	// 运行任务并处理结果。
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

// createTask 返回一个示例任务，该任务根据
// id 休眠指定的秒数。
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
