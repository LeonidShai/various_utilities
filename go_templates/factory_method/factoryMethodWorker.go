package factorymethod

import (
	"fmt"
)

func FactoryMethodWorker() {
	fmt.Println("---------- Factory Method -----------")

	foodTypes := []string{"fruits", "streat", "lunch", "dessert"}

	for _, ft := range foodTypes {
		food := GetFood(ft)
		fmt.Println(food.GetFoodName())
		food.PrintPeculiarities()
		fmt.Println()
	}

	fmt.Println("-------------------")
}
