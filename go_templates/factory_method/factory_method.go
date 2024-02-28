package factorymethod

type Food interface {
	GetFoodName() string
	PrintPeculiarities()
}

func GetFood(foodType string) Food {
	switch foodType {
	case "fruits":
		return NewFruits()
	case "streat":
		return NewStreat()
	case "lunch":
		return NewLunch()
	default:
		return NewDessert()
	}
}
