// 示例程序，展示接口的多态行为。
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

// notify 用指针接收者实现 notifier 接口。
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin 在程序中定义一个管理员。
type admin struct {
	name  string
	email string
}

// notify 用指针接收者实现 notifier 接口。
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// main 是应用程序的入口点。
func main() {
	// 创建一个 user 值并将其传递给 sendNotification。
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// 创建一个 admin 值并将其传递给 sendNotification。
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

// sendNotification 接受实现 notifier
// 接口的值并发送通知。
func sendNotification(n notifier) {
	n.notify()
}
