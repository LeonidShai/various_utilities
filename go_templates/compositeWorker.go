package main

import (
	"fmt"
	"go_templates/composite"
)

func compositeWorker() {
	fmt.Println("---------- Composite -----------")

	// Создадим структуру папок с файлами
	folder1 := composite.NewFolder("Documents")
	file1 := composite.NewFile("book.txt")
	file2 := composite.NewFile("doc.txt")
	folder1.AddComponents(file1, file2)

	folder2 := composite.NewFolder("home")
	file3 := composite.NewFile("test.txt")
	folder2.AddComponents(file3, folder1)

	// вызовем поиск по папкам и файлам
	folder2.Search("keyword")

	fmt.Println("-------------------")
}
