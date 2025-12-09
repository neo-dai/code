// 示例程序，展示当你需要为自己的包或测试
// 模拟具体类型时如何操作。
package main

import (
	"github.com/goinaction/code/chapter10/listing06/pubsub"
)

// publisher 是一个接口，允许此包
// 模拟 pubsub 包的支持。
type publisher interface {
	Publish(key string, v interface{}) error
	Subscribe(key string) error
}

// mock 是一个具体类型，用于帮助支持
// pubsub 包的模拟。
type mock struct{}

// Publish 为模拟实现 publisher 接口。
func (m *mock) Publish(key string, v interface{}) error {

	// 在此添加您对 PUBLISH 调用的模拟。
	return nil
}

// Subscribe 为模拟实现 publisher 接口。
func (m *mock) Subscribe(key string) error {

	// 在此添加您对 SUBSCRIBE 调用的模拟。
	return nil
}

func main() {

	// 创建一个 publisher 接口值的切片。分配
	// pubsub.PubSub 值的地址和
	// mock 值的地址。
	pubs := []publisher{
		pubsub.New("localhost"),
		&mock{},
	}

	// 遍历接口值以查看 publisher
	// 接口如何提供用户所需的解耦级别。
	// pubsub 包不需要提供接口类型。
	for _, p := range pubs {
		p.Publish("key", "value")
		p.Subscribe("key")
	}
}
