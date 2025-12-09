// 这个示例程序演示如何编组 JSON 字符串。
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// 创建一个键/值对的映射。
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	// 将映射编组为 JSON 字符串。
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))
}
