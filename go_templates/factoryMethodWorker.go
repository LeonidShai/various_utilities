package main

import (
	"fmt"
	factorymethod "go_templates/factory_method"
)

func factoryMethodWorker() {
	fmt.Println("---------- Factory Method -----------")

	foodTypes := []string{"fruits", "streat", "lunch", "dessert"}

	for _, ft := range foodTypes {
		food := factorymethod.GetFood(ft)
		fmt.Println(food.GetFoodName())
		food.PrintPeculiarities()
		fmt.Println()
	}

	fmt.Println("-------------------")
}
