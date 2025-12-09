// 示例程序，展示嵌入类型如何与接口一起工作。
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

// notify 实现一个可以通过
// user 类型的值调用的方法。
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin 表示具有权限的管理员用户。
type admin struct {
	user
	level string
}

// main 是应用程序的入口点。
func main() {
	// 创建一个管理员用户。
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 向管理员用户发送通知。
	// 嵌入的内部类型对接口的实现
	// 被"提升"到外部类型。
	sendNotification(&ad)
}

// sendNotification 接受实现 notifier
// 接口的值并发送通知。
func sendNotification(n notifier) {
	n.notify()
}
