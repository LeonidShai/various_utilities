package abstractfactory

import (
	"fmt"
)

func AbstractFactoryWorker() {
	fmt.Println("---------- Abstract Factory -----------")

	officeFactory := GetFactory("office")
	homeFactory := GetFactory("home")

	officeApple := officeFactory.GetFruit("apple")
	officeChair := officeFactory.GetItem("chair")

	homeBanana := homeFactory.GetFruit("banana")
	homePencil := homeFactory.GetItem("pencil")

	fmt.Printf("Office: fruit %v and item %v\n", officeApple.GetFruitName(), officeChair.GetItemName())
	fmt.Printf("fruit %v and item %v\n", officeApple.GetProperties(), officeChair.GetProperties())

	fmt.Printf("Home: fruit %v and item %v\n", homeBanana.GetFruitName(), homePencil.GetItemName())
	fmt.Printf("fruit %v and item %v\n", homeBanana.GetProperties(), homePencil.GetProperties())

	fmt.Println("-------------------")
}
