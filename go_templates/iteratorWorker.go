package main

import (
	"fmt"
	"go_templates/iterator"
)

func iteratorWorker() {
	fmt.Println("---------- Iterator -----------")

	// создаем объекты, по которым будем итерироваться
	book1 := &iterator.Book{
		Name:   "War and Piece",
		Author: "Tolstoy",
	}

	book2 := &iterator.Book{
		Name:   "Dubrovskii",
		Author: "Pushkin",
	}

	// создаем коллекцию
	booksCollection := iterator.Collection{
		Books: []*iterator.Book{book1, book2},
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
