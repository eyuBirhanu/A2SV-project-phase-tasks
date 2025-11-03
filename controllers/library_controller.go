package controllers

import (
	"bufio"
	"fmt"
	"library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	service services.LibraryManager
}

func NewLibraryController(service services.LibraryManager) *LibraryController {
	return &LibraryController{service: service}
}

func (lc *LibraryController) Run() {
	reader := bufio.NewReader(os.Stdin)

	lc.service.AddBook(models.Book{ID: 1, Title: "The Go Programming Language", Author: "Alan Donovan", Status: "Available"})
	lc.service.AddBook(models.Book{ID: 2, Title: "Structure and Interpretation of Computer Programs", Author: "Harold Abelson", Status: "Available"})

	for {
		fmt.Println("\n--- Library Management System ---")
		fmt.Println("1. Add a new book")
		fmt.Println("2. Remove a book")
		fmt.Println("3. Borrow a book")
		fmt.Println("4. Return a book")
		fmt.Println("5. List all available books")
		fmt.Println("6. List borrowed books for a member")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
		if err != nil {
			fmt.Println("Invalid choice. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			lc.addBook(reader)
		case 2:
			lc.removeBook(reader)
		case 3:
			fmt.Println("NOTE: For this demo, we will assume you are Member ID 1.")
			lc.borrowBook(reader, 1)
		case 4:
			fmt.Println("NOTE: For this demo, we will assume you are Member ID 1.")
			lc.returnBook(reader, 1)
		case 5:
			lc.listAvailableBooks()
		case 6:
			fmt.Println("NOTE: For this demo, we will assume you are Member ID 1.")
			lc.listBorrowedBooks(1)
		case 7:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (lc *LibraryController) addBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')

	book := models.Book{
		ID:     id,
		Title:  strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
		Status: "Available",
	}
	lc.service.AddBook(book)
}

func (lc *LibraryController) removeBook(reader *bufio.Reader) {
	fmt.Print("Enter book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	if err := lc.service.RemoveBook(id); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (lc *LibraryController) listAvailableBooks() {
	books := lc.service.ListAvailableBooks()
	if len(books) == 0 {
		fmt.Println("No available books.")
		return
	}
	fmt.Println("\n--- Available Books ---")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) borrowBook(reader *bufio.Reader, memberID int) {
	fmt.Print("Enter book ID to borrow: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	if err := lc.service.BorrowBook(bookID, memberID); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (lc *LibraryController) returnBook(reader *bufio.Reader, memberID int) {
	fmt.Print("Enter book ID to return: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	if err := lc.service.ReturnBook(bookID, memberID); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func (lc *LibraryController) listBorrowedBooks(memberID int) {
	books := lc.service.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Printf("Member with ID %d has no borrowed books.\n", memberID)
		return
	}
	fmt.Printf("\n--- Books Borrowed by Member %d ---\n", memberID)
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}