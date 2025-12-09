// Package search : search.go 管理针对
// Google、Yahoo 和 Bing 的结果搜索。
package search

import "log"

// Result 表示找到的搜索结果。
type Result struct {
	Engine      string
	Title       string
	Description string
	Link        string
}

// Searcher 声明一个接口，用于利用不同的
// 搜索引擎来查找结果。
type Searcher interface {
	Search(searchTerm string, searchResults chan<- []Result)
}

// searchSession 保存有关当前搜索提交的信息。
// 它包含选项、搜索器和一个用于接收
// 结果的通道。
type searchSession struct {
	searchers  map[string]Searcher
	first      bool
	resultChan chan []Result
}

// Google 如果提供此选项，
// Google 搜索将添加到搜索会话中。
func Google(s *searchSession) {
	log.Println("search : Submit : Info : Adding Google")
	s.searchers["google"] = google{}
}

// Bing 如果提供此选项，
// Bing 搜索将添加到此搜索会话中。
func Bing(s *searchSession) {
	log.Println("search : Submit : Info : Adding Bing")
	s.searchers["bing"] = bing{}
}

// Yahoo 如果将此选项作为参数
// 提供给 Submit，则将启用 Yahoo 搜索。
func Yahoo(s *searchSession) {
	log.Println("search : Submit : Info : Adding Yahoo")
	s.searchers["yahoo"] = yahoo{}
}

// OnlyFirst 是一个选项，它将搜索会话
// 限制为仅第一个结果。
func OnlyFirst(s *searchSession) { s.first = true }

// Submit 使用 goroutine 和通道来并发地
// 对三个主要搜索引擎执行搜索。
func Submit(query string, options ...func(*searchSession)) []Result {
	var session searchSession
	session.searchers = make(map[string]Searcher)
	session.resultChan = make(chan []Result)

	for _, opt := range options {
		opt(&session)
	}

	// 并发执行搜索。使用映射是因为
	// 它每次都以随机顺序返回搜索器。
	for _, s := range session.searchers {
		go s.Search(query, session.resultChan)
	}

	var results []Result

	// 等待结果返回。
	for search := 0; search < len(session.searchers); search++ {
		// 如果我们只想要第一个结果，不要再等待，
		// 通过并发丢弃剩余的 searchResults。
		// 否则会导致 Searchers 永远阻塞。
		if session.first && search > 0 {
			go func() {
				r := <-session.resultChan
				log.Printf("search : Submit : Info : Results Discarded : Results[%d]\n", len(r))
			}()
			continue
		}

		// 等待接收结果。
		log.Println("search : Submit : Info : Waiting For Results...")
		result := <-session.resultChan

		// 将结果保存到最终切片。
		log.Printf("search : Submit : Info : Results Used : Results[%d]\n", len(result))
		results = append(results, result...)
	}

	log.Printf("search : Submit : Completed : Found [%d] Results\n", len(results))
	return results
}
