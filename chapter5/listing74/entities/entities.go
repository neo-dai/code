// Package entities 包含系统中
// 人员类型的支持。
package entities

// user 在程序中定义一个用户。
type user struct {
	Name  string
	Email string
}

// Admin 在程序中定义一个管理员。
type Admin struct {
	user   // 嵌入类型是未导出的。
	Rights int
}
