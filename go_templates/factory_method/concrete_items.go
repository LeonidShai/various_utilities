package factorymethod

import "fmt"

type Fruits struct {
	Name        string
	Ingredients string
}

func NewFruits() Food {
	return &Fruits{
		Name:        "fruits",
		Ingredients: "bananas, apples, oranges",
	}
}

func (f *Fruits) GetFoodName() string {
	return f.Name
}

func (f *Fruits) PrintPeculiarities() {
	fmt.Println(f.Ingredients)
}

// ---------------------
type Streat struct {
	Name        string
	Ingredients string
}

func NewStreat() Food {
	return &Streat{
		Name:        "streat food",
		Ingredients: "hamburger, pizza",
	}
}

func (s *Streat) GetFoodName() string {
	return s.Name
}

func (s *Streat) PrintPeculiarities() {
	fmt.Println(s.Ingredients)
}

// ---------------------
type Lunch struct {
	Name     string
	Drink    string
	Sandwich string
}

func NewLunch() Food {
	return &Lunch{
		Name:     "lunch",
		Drink:    "coffee",
		Sandwich: "sandwich",
	}
}

func (l *Lunch) GetFoodName() string {
	return l.Name
}

func (l *Lunch) PrintPeculiarities() {
	fmt.Printf("%v with %v\n", l.Drink, l.Sandwich)
}

// ---------------------
type Dessert struct {
	Name string
	Cake string
}

func NewDessert() Food {
	return &Dessert{
		Name: "lunch",
		Cake: "cake",
	}
}

func (d *Dessert) GetFoodName() string {
	return d.Name
}

func (d *Dessert) PrintPeculiarities() {
	fmt.Println(d.Cake)
}
