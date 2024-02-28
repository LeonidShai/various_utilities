package abstractfactory

import "fmt"

// интерфейс для одинаковых элементов (Фрукты) с реализацией
type Fruit interface {
	GetFruitName() string
	GetProperties() string
}

type Apple struct {
	Name  string
	Shape string
	Color string
}

func (a *Apple) GetFruitName() string {
	return a.Name
}

func (a *Apple) GetProperties() string {
	return fmt.Sprintf("Apple properties are %v and %v", a.Shape, a.Color)
}

type Banana struct {
	Name  string
	Color string
}

func (b *Banana) GetFruitName() string {
	return b.Name
}

func (b *Banana) GetProperties() string {
	return fmt.Sprintf("Banana property is %v", b.Color)
}

// интерфейс для одинаковых элементов с реализацией
type Item interface {
	GetItemName() string
	GetProperties() string
}

type Pencil struct {
	Name  string
	Color string
}

func (p *Pencil) GetItemName() string {
	return p.Name
}

func (p *Pencil) GetProperties() string {
	return fmt.Sprintf("Pencil property is %v", p.Color)
}

type Chair struct {
	Name     string
	Property string
	Color    string
}

func (c *Chair) GetItemName() string {
	return c.Name
}

func (c *Chair) GetProperties() string {
	return fmt.Sprintf("Chair properties are %v and %v", c.Property, c.Color)
}
