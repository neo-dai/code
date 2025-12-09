package search

import (
	"encoding/json"
	"os"
)

const dataFile = "data/data.json"

// Feed 包含处理订阅源所需的信息。
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并解组订阅源数据文件。
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件。
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 安排在函数返回后
	// 关闭文件。
	defer file.Close()

	// 将文件解码为指向
	// Feed 值的指针切片。
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// 我们不需要检查错误，调用者可以这样做。
	return feeds, err
}
