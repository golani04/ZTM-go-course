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

type Name string
type Title string

type Audits struct {
	checkedIn  time.Time
	checkedOut time.Time
}

type Member struct {
	name  Name
	books map[Title]Audits
}

type Book struct {
	title  Title
	total  int
	lended int
}

type Library struct {
	books   map[Title]Book
	members map[Name]Member
}

func checkedOutBook(title Title, member *Member, library *Library) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println(title, "is absent")
		return false
	}

	if book.lended == book.total {
		fmt.Println("no more books are available")
		return false
	}

	book.lended += 1
	// copy back in order to save changes
	library.books[title] = book

	member.books[title] = Audits{checkedOut: time.Now()}
	return true
}

func checkedInBook(title Title, member *Member, library *Library) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println(title, "is absent")
		return false
	}

	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.lended -= 1
	// copy back in order to save changes
	library.books[title] = book

	audit.checkedIn = time.Now()
	member.books[title] = audit
	return true
}

func printLibInfo(library *Library) {
	fmt.Println("Total books:")
	for _, book := range library.books {
		fmt.Print(book.title, " total: ", book.total, ", landed: ", book.lended, "\n")
		if book.lended == 0 {
			continue
		}

		for _, member := range library.members {
			audit, found := member.books[book.title]
			if !found {
				continue
			}
			if audit.checkedIn.IsZero() {
				fmt.Println("- Member", member.name, "still has this book")
			}
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("Starting...")
	library := Library{}
	library.books = make(map[Title]Book)
	library.books["The Adventures of Tom Sawyer"] = Book{"The Adventures of Tom Sawyer", 3, 0}
	library.books["Adventures of Huckleberry Finn"] = Book{"Adventures of Huckleberry Finn", 2, 0}
	library.books["Gulliver's Travels"] = Book{"Gulliver's Travels", 1, 0}
	library.books["The Treasure Island"] = Book{"The Treasure Island", 4, 0}

	library.members = make(map[Name]Member)
	library.members["Anthony"] = Member{"Anthony", make(map[Title]Audits)}
	library.members["Jennifer"] = Member{"Jennifer", make(map[Title]Audits)}
	library.members["John"] = Member{"John", make(map[Title]Audits)}

	printLibInfo(&library)

	member, found := library.members["John"]
	if !found {
		fmt.Println("no member with this name exists")
		return
	}

	checkedOutBook("The Adventures of Tom Sawyer", &member, &library)
	printLibInfo(&library)
	checkedInBook("The Adventures of Tom Sawyer", &member, &library)
	printLibInfo(&library)
}
