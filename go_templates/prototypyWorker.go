package main

import (
	"fmt"
	"go_templates/prototype"
)

func prototypeWorker() {
	fmt.Println("---------- Prototype -----------")
	item1 := prototype.NewItem("buratino")
	item2 := item1.Clone()
	fmt.Printf("Item_1 name: %v and Item_2 name: %v\n", item1.GetName(), item2.GetName())
	fmt.Println("-------------------")
}
