// Package counters 提供警报计数器支持。
package counters

// alertCounter 是一个未导出类型，
// 包含用于警报的整数计数器。
type alertCounter int

// New 创建并返回未导出类型
// alertCounter 的值。
func New(value int) alertCounter {
	return alertCounter(value)
}
