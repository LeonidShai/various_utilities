package flyweight

import "fmt"

// заведем интерфейс для простых структур
type Dress interface {
	WhichDress()
	Color() string
	DressType() string
}

// заведем, структуры удовлетворяющие интерфейсу
type homeDress struct {
	color     string
	dressType string
}

func NewHomeDress(color string) homeDress {
	return homeDress{
		color:     color,
		dressType: "home",
	}
}

func (d *homeDress) Color() string {
	fmt.Printf("Home dress main color is %s\n", d.color)

	return d.color
}

func (d *homeDress) WhichDress() {
	fmt.Printf("Home dress trousers and t-shirt with %s color\n", d.color)
}

func (d *homeDress) DressType() string {
	return d.dressType
}

type streetDress struct {
	color     string
	dressType string
}

func NewStreetDress(color string) streetDress {
	return streetDress{
		color:     color,
		dressType: "street",
	}
}

func (d *streetDress) Color() string {
	fmt.Printf("Street dress main color is %s\n", d.color)

	return d.color
}

func (d *streetDress) WhichDress() {
	fmt.Printf("Street dress trousers and jacket with %s color\n", d.color)
}

func (d *streetDress) DressType() string {
	return d.dressType
}
