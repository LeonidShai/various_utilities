package decorator

import "fmt"

// интерфейс для декорирования
type DecorIf interface {
	Print()
	GetName() string
}

// Обычная структурка
type simpleItem struct {
	name string
}

func NewSimpleItem(name string) simpleItem {
	return simpleItem{
		name: name,
	}
}

func (s *simpleItem) GetName() string {
	return s.name
}

func (s *simpleItem) Print() {
	fmt.Println("This is Simple Item")
}

// Можем задекорировать следующей структурой с методами
type decorItem struct {
	decor DecorIf
}

func NewDecorItem(d DecorIf) decorItem {
	return decorItem{
		decor: d,
	}
}

func (d *decorItem) Print() {
	fmt.Printf("Decor the ")
	d.decor.Print()
}

func (d *decorItem) GetName() string {
	return d.decor.GetName() + "_decor"
}

// decoration function
// simple func:
func Double(x int) int {
	return x * 2
}

type Fn func(int) int

// decor for func:
func Wrapper(f Fn) Fn {
	return func(x int) int {
		fmt.Printf("Calculate double for %d\n", x)

		result := f(x)

		fmt.Printf("Result is %d\n", result)

		return result
	}
}
