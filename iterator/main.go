package main

import "fmt"

func main() {
	lib.IterateBooks(myBookCallback)

	lib.IterateBooks(func(b Book) error {
		fmt.Println("Book title:", b.Name)
		return nil
	})

	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Printf("Book %+v", book.Name)
	}
}

func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
