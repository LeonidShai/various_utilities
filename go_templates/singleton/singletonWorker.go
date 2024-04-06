package singleton

import (
	"fmt"
)

func SingletonWorker() {
	fmt.Println("---------- Singleton -----------")
	name := "First"
	fmt.Printf("%v singleton creation\n", name)
	instance1 := GetSingletonInstance(name)
	fmt.Println(instance1.Title())

	name = "Second"
	fmt.Printf("\n%v singleton creation\n", name)
	instance2 := GetSingletonInstance(name)
	fmt.Println(instance2.Title())

	if instance1 == instance2 {
		fmt.Println("instance1 == instance2")
		fmt.Printf("addr instance1: %v addr instance2: %v\n", &instance1, &instance2)
		fmt.Printf("instance1 contain: %p instance2 contain: %p\n", instance1, instance2)
	}

	name = "Third"
	fmt.Printf("\n%v singleton creation\n", name)
	instance3 := *instance1
	fmt.Println(instance3.Title())

	if instance1 == &instance3 {
		fmt.Println("instance1 == instance3")
		fmt.Printf("addr instance1: %v addr instance3: %v\n", &instance1, &instance3)
		fmt.Printf("instance1 is: %T instance3 is: %T\n", instance1, instance3)
	} else {
		fmt.Printf("addr instance1: %v addr instance3: %p\n", &instance1, &instance3)
		fmt.Printf("instance1 is: %T instance3 is: %T\n", instance1, instance3)
	}

	name = "Fourth"
	fmt.Printf("\n%v singleton creation\n", name)
	instance4 := instance1
	fmt.Println(instance4.Title())

	if instance1 == instance4 {
		fmt.Println("instance1 == instance4")
		fmt.Printf("addr instance1: %v addr instance4: %v\n", &instance1, &instance4)
		fmt.Printf("instance1 contain: %p instance4 contain: %p\n", instance1, instance4)
	} else {
		fmt.Printf("addr instance1: %v addr instance4: %v\n", &instance1, &instance4)
		fmt.Printf("instance1 contain: %p instance4 contain: %p\n", instance1, instance4)
	}

	fmt.Println("\nTrying to change titles for instances")
	instance1.SetNewTitle("instance_1")
	instance2.SetNewTitle("instance_2")
	instance3.SetNewTitle("instance_3")
	instance4.SetNewTitle("instance_4")
	title1, _ := instance1.Title()
	title2, _ := instance2.Title()
	title3, _ := instance3.Title()
	title4, _ := instance4.Title()
	fmt.Printf("Title instance1: %v\n", title1)
	fmt.Printf("Title instance2: %v\n", title2)
	fmt.Printf("Title instance3: %v\n", title3)
	fmt.Printf("Title instance4: %v\n", title4)

	fmt.Println("\n-------------------")
}
