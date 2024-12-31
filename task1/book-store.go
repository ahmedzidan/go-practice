package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func updatePages(book *Book, pages int) {
	// your code goes here
	book.Pages = pages
}

func main() {

	/*
		Create 3 Book Structs with the following data:

		Book 1:
		Title: "The Great Gatsby"
		Author: "F. Scott Fitzgerald"
		Pages: 180

		Book 2
		Title: "To Kill a Mockingbird"
		Author: "Harper Lee"
		Pages: 281

		Book 3
		Title: "Pride and Prejudice"
		Author: "Jane Austen"
		Pages: 279
	*/
	book1 := Book{
		Title:  "The Great Gatsby",
		Author: "F. Scott Fitzgerald",
		Pages:  180,
	}

	book2 := Book{
		Title:  "To Kill a Mockingbird",
		Author: "Harper Lee",
		Pages:  281,
	}
	book3 := Book{
		Title:  "Pride and Prejudice",
		Author: "Jane Austen",
		Pages:  279,
	}
	updatePages(&book1, 210)
	updatePages(&book2, 250)
	updatePages(&book3, 295)

	// your code for function calls to updatePages goes here

	/*
		Print all the struct objects
		fmt.Println(book)
	*/
	fmt.Println(book1)
	fmt.Println(book2)
	fmt.Println(book3)

	// your code for printing objects goes here
}
