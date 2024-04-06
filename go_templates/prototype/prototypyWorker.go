package prototype

import (
	"fmt"
)

func PrototypeWorker() {
	fmt.Println("---------- Prototype -----------")
	item1 := NewItem("buratino")
	item2 := item1.Clone()
	fmt.Printf("Item_1 name: %v and Item_2 name: %v\n", item1.GetName(), item2.GetName())
	fmt.Println("-------------------")
}
