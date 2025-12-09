// 示例程序，演示编译器
// 何时提供隐式接口转换。
package main

import "fmt"

// =============================================================================

// Mover 为移动事物提供支持。
type Mover interface {
	Move()
}

// Locker 为锁定和解锁事物提供支持。
type Locker interface {
	Lock()
	Unlock()
}

// MoveLocker 为移动和锁定事物提供支持。
type MoveLocker interface {
	Mover
	Locker
}

// =============================================================================

// bike 表示示例的具体类型。
type bike struct{}

// Move 可以改变自行车的位置。
func (bike) Move() {
	fmt.Println("Moving the bike")
}

// Lock 防止自行车移动。
func (bike) Lock() {
	fmt.Println("Locking the bike")
}

// Unlock 允许自行车移动。
func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

// =============================================================================

func main() {

	// 声明 MoveLocker 和 Mover 接口的变量
	// 并将其设置为零值。
	var ml MoveLocker
	var m Mover

	// 创建一个 bike 类型的值并将该值分配给
	// MoveLocker 接口值。
	ml = bike{}

	// MoveLocker 类型的接口值可以隐式转换为
	// Mover 类型的值。它们都声明了一个名为 move 的方法。
	m = ml

	// prog.go:65: cannot use m (type Mover) as type MoveLocker in assignment:
	//	   Mover does not implement MoveLocker (missing Lock method)
	ml = m

	// 接口类型 Mover 没有声明名为 lock 和 unlock 的方法。
	// 因此，编译器无法执行隐式转换来将
	// Mover 接口类型的值分配给 MoveLocker 类型的接口值。
	// 存储在 Mover 接口值内部的 bike 类型的具体类型值
	// 实现了 MoveLocker 接口这一事实是无关紧要的。

	// 我们可以在运行时执行类型断言来支持赋值。

	// 对 Mover 接口值执行类型断言以访问
	// 存储在其内部的 bike 类型的具体类型值的副本。
	// 然后将具体类型的副本分配给
	// MoveLocker 接口。
	b := m.(bike)
	ml = b
}
