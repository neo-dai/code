// 示例程序，展示如何声明方法以及 Go
// 编译器如何支持它们。
package main

import (
	"fmt"
)

// user 在程序中定义一个用户。
type user struct {
	name  string
	email string
}

// notify 实现一个具有值接收者的方法。
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail 实现一个具有指针接收者的方法。
func (u *user) changeEmail(email string) {
	u.email = email
}

// main 是应用程序的入口点。
func main() {
	// user 类型的值可用于调用
	// 用值接收者声明的方法。
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// user 类型的指针也可用于调用
	// 用值接收者声明的方法。
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// user 类型的值可用于调用
	// 用指针接收者声明的方法。
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// user 类型的指针可用于调用
	// 用指针接收者声明的方法。
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
