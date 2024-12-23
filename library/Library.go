package library

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	books map[string]Book
}

func NewLibrary() *Library {
	return &Library{
		books: make(map[string]Book),
	}
}

func (l *Library) AddBook(book Book) {
	if _, exists := l.books[book.ID]; exists {
		fmt.Println("Book with this ID already exists.")
		return
	}
	l.books[book.ID] = book
	fmt.Println("Book added successfully.")
}

func (l *Library) BorrowBook(id string) {
	book, exists := l.books[id]
	if !exists {
		fmt.Println("Book not found.")
		return
	}
	if book.IsBorrowed {
		fmt.Println("Book is already borrowed.")
		return
	}
	book.IsBorrowed = true
	l.books[id] = book
	fmt.Println("Book borrowed successfully.")
}

func (l *Library) ReturnBook(id string) {
	book, exists := l.books[id]
	if !exists {
		fmt.Println("Book not found.")
		return
	}
	if !book.IsBorrowed {
		fmt.Println("Book is not currently borrowed.")
		return
	}
	book.IsBorrowed = false
	l.books[id] = book
	fmt.Println("Book returned successfully.")
}

func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for _, book := range l.books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

func main() {
	library := NewLibrary()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add")
		fmt.Println("2. Borrow")
		fmt.Println("3. Return")
		fmt.Println("4. List")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter book ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Enter book Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter book Author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			library.AddBook(Book{ID: id, Title: title, Author: author, IsBorrowed: false})

		case "2":
			fmt.Print("Enter book ID to borrow: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			library.BorrowBook(id)

		case "3":
			fmt.Print("Enter book ID to return: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)
			library.ReturnBook(id)

		case "4":
			library.ListBooks()

		case "5":
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
