package concurrency

import (
	"fmt"
	"sync"
)

type Book struct {
	Id       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "The Lord of the Rings", false},
	{2, "The Hobbit", false},
	{3, "The Catcher in the Rye", false},
	{4, "The Grapes of Wrath", false},
	{5, "The Great Gatsby", false},
	{6, "Black Mirror", false},
	{7, "The Hunger Games", false},
	{8, "The Da Vinci Code", false},
	{9, "The Lion, the Witch and the Wardrobe", false},
	{10, "The Witcher", false},
}

func FindBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var searchBook *Book

	m.RLock()
	for i, book := range books {
		fmt.Printf("Book %v\n", book)
		if book.Id == id {
			index = i
			searchBook = book
			break
		}
	}
	m.RUnlock()
	fmt.Printf("Book founded: %v\n", searchBook)
	return index, searchBook
}

func FinishBook(id int, m *sync.RWMutex) {
	index, book := FindBook(id, m)
	fmt.Printf("Index: %v\n", index)
	if index < 0 {
		return
	}

	m.Lock()
	book.Finished = true
	books[index] = book
	m.Unlock()
	fmt.Printf("Finished book %s\n", book.Title)
}
