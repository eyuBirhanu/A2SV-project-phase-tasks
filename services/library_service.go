package services

import (
	"errors"
	"fmt"
	"library_management/models"
	"sync"
	"time"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int) error
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ReserveBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
	mu      sync.Mutex
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.Books[book.ID] = book
	fmt.Printf("Book '%s' added successfully.\n", book.Title)
}

func (l *Library) RemoveBook(bookID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	_, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	delete(l.Books, bookID)
	fmt.Printf("Book with ID %d removed successfully.\n", bookID)
	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	var availableBooks []models.Book
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}

	if book.Status != "Available" && book.Status != "Reserved" {
		return errors.New("book is not available for borrowing")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	fmt.Printf("Book '%s' borrowed by '%s'.\n", book.Title, member.Name)
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	book, bookExists := l.Books[bookID]
	if !bookExists {
		return errors.New("book not found")
	}

	member, memberExists := l.Members[memberID]
	if !memberExists {
		return errors.New("member not found")
	}

	bookFound := false
	var bookIndex int
	for i, borrowedBook := range member.BorrowedBooks {
		if borrowedBook.ID == bookID {
			bookFound = true
			bookIndex = i
			break
		}
	}

	if !bookFound {
		return errors.New("this member did not borrow this book")
	}

	book.Status = "Available"
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks[:bookIndex], member.BorrowedBooks[bookIndex+1:]...)
	l.Members[memberID] = member

	fmt.Printf("Book '%s' returned by '%s'.\n", book.Title, member.Name)
	return nil
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	l.mu.Lock()
	defer l.mu.Unlock()

	member, exists := l.Members[memberID]
	if !exists {
		return nil
	}
	return member.BorrowedBooks
}

func (l *Library) ReserveBook(bookID int, memberID int) error {
	l.mu.Lock()
	book, exists := l.Books[bookID]
	if !exists {
		l.mu.Unlock()
		return errors.New("book not found")
	}

	if book.Status != "Available" {
		l.mu.Unlock()
		return fmt.Errorf("book %d is currently %s", bookID, book.Status)
	}

	book.Status = "Reserved"
	l.Books[bookID] = book
	l.mu.Unlock()

	fmt.Printf("Success: Book '%s' Reserved for Member %d. You have 5 seconds to borrow it.\n", book.Title, memberID)

	go func(bID int) {
		time.Sleep(5 * time.Second)

		l.mu.Lock()
		defer l.mu.Unlock()

		currentBook, ok := l.Books[bID]
		if ok && currentBook.Status == "Reserved" {
			currentBook.Status = "Available"
			l.Books[bID] = currentBook
			fmt.Printf("\n[Timeout] Reservation for Book ID %d expired. It is now Available.\n", bID)
		}
	}(bookID)

	return nil
}
