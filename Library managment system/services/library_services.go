package services

import (
	"errors"
	model "github.com/ermi9s/go/models"
)

type LibraryManager interface {
	AddBook(book model.Book)
	RemoveBooks(bookId int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []model.Book
	ListBorrowedBooks() []model.Book
}

type Library struct {
	Books   map[int] *model.Book
	NextBookId int; 
	Members map[int] *model.Member
	NextMemberId int;
}

func (library *Library) AddBook(book model.Book) {
	if _, ok := library.Books[book.Id]; !ok {
		library.Books[book.Id] = &book
	}

}


func (library *Library) RemoveBooks(bookId int) {
	delete(library.Books, bookId)
}

func (library *Library) BorrowBook(bookID int, memberID int) error {
	if book, ok := library.Books[bookID]; ok {
		if book.Status == "Available" {
			book.Status = "Borrowed"
			library.Members[memberID].BorrowedBooks = append(library.Members[memberID].BorrowedBooks, *book)
		} else {
			return errors.New("400!, Book not available")
		}
	}else {
		return errors.New("404!, Book not found")
	}
	return nil;
}

func (library *Library) ReturnBook(bookID int, memberID int) error {
	
	book_list := library.Members[memberID].BorrowedBooks;
	for idx := 0; idx < len(book_list); idx++ {
		if book_list[idx].Id == bookID {
			library.Books[bookID].Status = "Available";
			book_list = append(book_list[:idx] , book_list[idx+1:]...)
			library.Members[memberID].BorrowedBooks = book_list;
			return nil
		}
	}
	return errors.New("404!, Book not found")
}

func (library *Library) ListAvailableBooks() []model.Book {
	var available_books []model.Book;
	
	for _,book := range library.Books {
		if book.Status == "Available" {
			available_books = append(available_books, *book);
		}
	}

	return available_books;
}

func (library *Library)ListBorrowedBooks() []model.Book {
	var borrowed_books []model.Book;
	
	for _,book := range library.Books {
		if book.Status == "Borrowed" {
			borrowed_books = append(borrowed_books, *book);
		}
	}

	return borrowed_books;
}
