// Package search : seachers.go 包含
// 现有搜索器的所有不同实现。
package search

import (
	"log"
	"math/rand"
	"time"
)

// init 在 main 之前被调用。
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Google 提供对 Google 搜索的支持。
type google struct{}

// Search 实现 Searcher 接口。它对
// Google 执行搜索。
func (g google) Search(term string, results chan<- []Result) {
	log.Printf("Google : Search : Started : search term[%s]\n", term)

	// 结果切片。
	var r []Result

	// 模拟搜索的时间量。
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	// 模拟搜索结果。
	r = append(r, Result{
		Engine:      "Google",
		Title:       "The Go Programming Language",
		Description: "The Go Programming Language",
		Link:        "https://golang.org/",
	})

	log.Printf("Google : Search : Completed : Found[%d]\n", len(r))
	results <- r
}

// Bing 提供对 Bing 搜索的支持。
type bing struct{}

// Search 实现 Searcher 接口。它对
// Bing 执行搜索。
func (b bing) Search(term string, results chan<- []Result) {
	log.Printf("Bing : Search : Started : search term [%s]\n", term)

	// 结果切片。
	var r []Result

	// 模拟搜索的时间量。
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	// 模拟搜索结果。
	r = append(r, Result{
		Engine:      "Bing",
		Title:       "A Tour of Go",
		Description: "Welcome to a tour of the Go programming language.",
		Link:        "http://tour.golang.org/",
	})

	log.Printf("Bing : Search : Completed : Found[%d]\n", len(r))
	results <- r
}

// Yahoo 提供对 Yahoo 搜索的支持。
type yahoo struct{}

// Search 实现 Searcher 接口。它对
// Yahoo 执行搜索。
func (y yahoo) Search(term string, results chan<- []Result) {
	log.Printf("Yahoo : Search : Started : search term [%s]\n", term)

	// 结果切片。
	var r []Result

	// 模拟搜索的时间量。
	time.Sleep(time.Millisecond * time.Duration(rand.Int63n(900)))

	// 模拟搜索结果。
	r = append(r, Result{
		Engine:      "Yahoo",
		Title:       "Go Playground",
		Description: "The Go Playground is a web service that runs on golang.org's servers",
		Link:        "http://play.golang.org/",
	})

	log.Printf("Yahoo : Search : Completed : Found[%d]\n", len(r))
	results <- r
}
