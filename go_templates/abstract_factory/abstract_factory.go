package abstractfactory

// интерфейс фабрики и 2 разные фабрики HomeThingsFactory и OfficeThingsFactory
type Factory interface {
	GetFruit(fruitType string) Fruit
	GetItem(itemType string) Item
}

type HomeThingsFactory struct{}

func (htf HomeThingsFactory) GetFruit(ft string) Fruit {
	if ft == "apple" {
		return &Apple{
			Name:  "apple",
			Shape: "round",
			Color: "red",
		}
	}

	return &Banana{
		Name:  "banana",
		Color: "yellow",
	}
}

func (htf HomeThingsFactory) GetItem(it string) Item {
	if it == "chair" {
		return &Chair{
			Name:     "chair",
			Property: "soft",
			Color:    "grey",
		}
	}

	return &Pencil{
		Name:  "pencil",
		Color: "black",
	}
}

type OfficeThingsFactory struct{}

func (otf OfficeThingsFactory) GetFruit(ft string) Fruit {
	if ft == "apple" {
		return &Apple{
			Name:  "apple",
			Shape: "round",
			Color: "green",
		}
	}

	return &Banana{
		Name:  "banana",
		Color: "yellow",
	}
}

func (tf OfficeThingsFactory) GetItem(it string) Item {
	if it == "chair" {
		return &Chair{
			Name:     "chair",
			Property: "soft",
			Color:    "red",
		}
	}

	return &Pencil{
		Name:  "pencil",
		Color: "blue",
	}
}

func GetFactory(factoryType string) Factory {
	if factoryType == "office" {
		return &OfficeThingsFactory{}
	}

	return &HomeThingsFactory{}
}
