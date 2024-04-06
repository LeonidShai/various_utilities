package iterator

import (
	"fmt"
)

func IteratorWorker() {
	fmt.Println("---------- Iterator -----------")

	// создаем объекты, по которым будем итерироваться
	book1 := &Book{
		Name:   "War and Piece",
		Author: "Tolstoy",
	}

	book2 := &Book{
		Name:   "Dubrovskii",
		Author: "Pushkin",
	}

	// создаем коллекцию
	booksCollection := Collection{
		Books: []*Book{book1, book2},
	}

	// создаем итератор
	it := booksCollection.CreateIterator()

	// и можем перебирать объекты:
	for it.HasNext() {
		book := it.Next()
		fmt.Printf("Book: %s Author: %s\n", book.Name, book.Author)
	}

	fmt.Println("-------------------")
}
