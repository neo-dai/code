// 这个示例代码实现一个简单的 web 服务。
package main

import (
	"log"
	"net/http"

	"github.com/goinaction/code/chapter9/listing17/handlers"
)

// main 是应用程序的入口点。
func main() {
	handlers.Routes()

	log.Println("listener : Started : Listening on :4000")
	http.ListenAndServe(":4000", nil)
}
