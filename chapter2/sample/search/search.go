package search

import (
	"log"
	"sync"
)

// 用于搜索的已注册匹配器的映射。
var matchers = make(map[string]Matcher)

// Run 执行搜索逻辑。
func Run(searchTerm string) {
	// 检索要搜索的订阅源列表。
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个无缓冲通道来接收要显示的匹配结果。
	results := make(chan *Result)

	// 设置等待组，以便我们可以处理所有订阅源。
	var waitGroup sync.WaitGroup

	// 设置我们需要等待的 goroutine 数量，
	// 它们会处理各个订阅源。
	waitGroup.Add(len(feeds))

	// 为每个订阅源启动一个 goroutine 来查找结果。
	for _, feed := range feeds {
		// 检索用于搜索的匹配器。
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动 goroutine 执行搜索。
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个 goroutine 来监控所有工作何时完成。
	go func() {
		// 等待所有内容被处理。
		waitGroup.Wait()

		// 关闭通道以向 Display 函数发出信号，
		// 表示我们可以退出程序。
		close(results)
	}()

	// 在结果可用时开始显示结果，
	// 并在显示最终结果后返回。
	Display(results)
}

// Register 被调用以注册供程序使用的匹配器。
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
