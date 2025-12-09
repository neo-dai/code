// 示例程序，展示如何在 Go 中使用接口。
package main

import (
	"fmt"
)

// notifier 是定义通知
// 类型行为的接口。
type notifier interface {
	notify()
}

// user 在程序中定义一个用户。
type user struct {
	name  string
	email string
}

// notify 实现一个具有指针接收者的方法。
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// main 是应用程序的入口点。
func main() {
	// 创建一个 User 类型的值并发送通知。
	u := user{"Bill", "bill@email.com"}

	sendNotification(u)

	// ./listing36.go:32: cannot use u (type user) as type
	//                     notifier in argument to sendNotification:
	//   user does not implement notifier
	//                          (notify method has pointer receiver)
}

// sendNotification 接受实现 notifier
// 接口的值并发送通知。
func sendNotification(n notifier) {
	n.notify()
}
