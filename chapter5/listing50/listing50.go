// 示例程序，展示如何将类型嵌入到另一个类型中
// 以及内部类型和外部类型之间的关系。
package main

import (
	"fmt"
)

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
	user  // 嵌入类型
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

	// 我们可以直接访问内部类型的方法。
	ad.user.notify()

	// 内部类型的方法被提升。
	ad.notify()
}
