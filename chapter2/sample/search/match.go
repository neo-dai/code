package search

import (
	"log"
)

// Result 包含搜索的结果。
type Result struct {
	Field   string
	Content string
}

// Matcher 定义想要实现新搜索类型的类型
// 所需的行为。
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 为每个单独的订阅源启动为 goroutine，
// 以并发地运行搜索。
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 针对指定的匹配器执行搜索。
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将结果写入通道。
	for _, result := range searchResults {
		results <- result
	}
}

// Display 在各个 goroutine 接收到结果时
// 将结果写入控制台窗口。
func Display(results chan *Result) {
	// 通道会阻塞，直到有结果被写入通道。
	// 一旦通道被关闭，for 循环就会终止。
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
