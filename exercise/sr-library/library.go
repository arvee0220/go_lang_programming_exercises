//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type book struct {
	title string
	checkedOut bool
	checkOutTime string
	returnTime string
}

type member struct {
	name string
	id string
}

type library struct {
	books []book
	members []member
}

func findMemberByID(members []member, id string) *member {
	for i := range members {
		if members[i].id == id {
			return &members[i]
		}
	}
	return nil
}

func findBookByTitle(books []book, bookTitle string) *book {
	for i := range books {
		if books[i].title == bookTitle {
			return &books[i]
		}
	}

	return nil
}

func borrowBook(lib *library, bookTitle string, memberID string) {

	member := findMemberByID(lib.members, memberID)
	if member == nil {
		fmt.Println("Checkout failed: member not found")
		return
	}

	book := findBookByTitle(lib.books, bookTitle)
	if book == nil || book.checkedOut {
		fmt.Println("Checkout failed: book not found or already checked out")
		return
	}

	book.checkedOut = true
	book.checkOutTime = time.Now().Format(time.RFC3339)
	fmt.Printf("The book titled \"%s\" was checked out by %s at %s\n", bookTitle, member.name, book.checkOutTime)
}

func returnBook(lib *library, bookTitle string, memberID string) {
	member := findMemberByID(lib.members, memberID)
	book := findBookByTitle(lib.books, bookTitle)

	if book == nil || !book.checkedOut {
		fmt.Println("Return failed: book not found or not checked out")
		return
	}

	book.checkedOut = false
	book.returnTime = time.Now().Format(time.RFC3339)

	if member != nil {
		fmt.Printf("The book titled \"%s\" was returned by %s at %s\n", bookTitle, member.name, book.returnTime)
	} else {
		fmt.Printf("The book titled \"%s\" was returned by a non-member at %s\n", bookTitle, book.returnTime)
	}
}

func main() {
	lib := library{
		books: []book{
			{title: "Learning Go - An Idiomatic Approach to Real-world Go Programming, 2nd Edition", checkedOut: false},
			{title: "Learning React - Modern Patterns for Developing React Apps, 2nd Edition", checkedOut: false},
			{title: "Learning Test-Driven Development - A Polyglot Guide to Writing Uncluttered Code", checkedOut: false},
			{title: "Learning TypeScript - Enhance Your Web Development Skills Using Type-Safe JavaScript", checkedOut: false},
		},

		members: []member{
			{name: "Alice", id: "M001"},
			{name: "Bob", id: "M002"},
			{name: "Charlie", id: "M003"},
		},
	}

	fmt.Println("Library Information:", lib.books)

	borrowBook(&lib, "Learning Go - An Idiomatic Approach to Real-world Go Programming, 2nd Edition", "M001")
	fmt.Println("After first checkout:", lib.books)

	borrowBook(&lib, "Learning React - Modern Patterns for Developing React Apps, 2nd Edition", "M002")
	fmt.Println("After second checkout:", lib.books)

	returnBook(&lib, "Learning React - Modern Patterns for Developing React Apps, 2nd Edition", "M002")
	fmt.Println("After return:", lib.books)

}