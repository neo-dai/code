// 示例程序，演示用接口进行解耦。
package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// =============================================================================

// Data 是我们正在复制的数据的结构。
type Data struct {
	Line string
}

// =============================================================================

// Puller 声明提取数据的行为。
type Puller interface {
	Pull(d *Data) error
}

// Storer 声明存储数据的行为。
type Storer interface {
	Store(d Data) error
}

// =============================================================================

// Xenia 是我们需要从中提取数据的系统。
type Xenia struct{}

// Pull 知道如何从 Xenia 中提取数据。
func (Xenia) Pull(d *Data) error {
	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("Error reading data from Xenia")

	default:
		d.Line = "Data"
		fmt.Println("In:", d.Line)
		return nil
	}
}

// Pillar 是我们需要将数据存储到的系统。
type Pillar struct{}

// Store 知道如何将数据存储到 Pillar 中。
func (Pillar) Store(d Data) error {
	fmt.Println("Out:", d.Line)
	return nil
}

// =============================================================================

// System 将 Xenia 和 Pillar 包装到一个系统中。
type System struct {
	Xenia
	Pillar
}

// =============================================================================

// pull 知道如何从任何 Puller 批量提取数据。
func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// store 知道如何从任何 Storer 批量存储数据。
func store(s Storer, data []Data) (int, error) {
	for i, d := range data {
		if err := s.Store(d); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

// Copy 知道如何从 System 中提取和存储数据。
func Copy(sys *System, batch int) error {
	data := make([]Data, batch)

	for {
		i, err := pull(&sys.Xenia, data)
		if i > 0 {
			if _, err := store(&sys.Pillar, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {
			return err
		}
	}
}

// =============================================================================

func main() {

	// 初始化系统以供使用。
	sys := System{
		Xenia:  Xenia{},
		Pillar: Pillar{},
	}

	if err := Copy(&sys, 3); err != io.EOF {
		fmt.Println(err)
	}
}
