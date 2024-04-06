package composite

import (
	"fmt"
)

func CompositeWorker() {
	fmt.Println("---------- Composite -----------")

	// Создадим структуру папок с файлами
	folder1 := NewFolder("Documents")
	file1 := NewFile("book.txt")
	file2 := NewFile("doc.txt")
	folder1.AddComponents(file1, file2)

	folder2 := NewFolder("home")
	file3 := NewFile("test.txt")
	folder2.AddComponents(file3, folder1)

	// вызовем поиск по папкам и файлам
	folder2.Search("keyword")

	fmt.Println("-------------------")
}
