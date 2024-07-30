package controllers

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	io "fmt"
	model "github.com/ermi9s/go/models"
	main "github.com/ermi9s/go/services"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin);

func AddBook(library main.Library) {
	var new_Title string
	var new_Author string
	new_Id := library.NextBookId;

	io.Print("Enter Book-Name: ")
	input_title,err := reader.ReadString('\n');
	if err != nil {
		io.Println("Invalid Input!")
		return
	}
	new_Title = strings.TrimSpace(input_title);

	io.Print("Enter Book-Author: ")
	input_author,err := reader.ReadString('\n');
	if err != nil {
		io.Println("Invalid Input!")
		return
	}
	new_Author = strings.TrimSpace(input_author);

	newBook :=  model.Book{
		Id : new_Id,
		Title: new_Title,
		Author: new_Author,
		Status: "Available",
	}

	library.AddBook(newBook);
	io.Println("Book Added Successfully.")
}

func RemoveBook(library main.Library) {
	var BookId int;

	io.Print("Enter Book-ID: ")
	input_id,err := reader.ReadString('\n') 
	if err != nil {
		io.Println("Invalid Input!")
		return
	}
	BookId,_  = strconv.Atoi(strings.TrimSpace(input_id));
	library.RemoveBooks(BookId);
	io.Println("Book Removed Successfully.")
}

func BorrowBook(library main.Library , memberID int) {
	var bookId int;

	io.Print("Enter Book-ID: ")
	input_id,err := reader.ReadString('\n') 
	if err != nil {
		io.Println("Invalid Input!")
		return
	}
	bookId,_  = strconv.Atoi(strings.TrimSpace(input_id));

	if message := library.BorrowBook(bookId , memberID); message != nil {
		io.Println(message.Error());
	}else {
		io.Println("Book Borrowed Successfully")
	}
}

func ReturnBook(library main.Library ,memberId int) {
	var BookId int;

	io.Print("Enter Book-ID: ")
	input_id,err := reader.ReadString('\n') 
	if err != nil {
		io.Println("Invalid Input!")
		return
	}
	BookId,_  = strconv.Atoi(strings.TrimSpace(input_id));

	if message := library.ReturnBook(BookId , memberId); message != nil {
		io.Println(message.Error());
	}else {
		io.Println("Book Returned Successfully")
	}
}

func UserBorrowList(library main.Library , useId int) {
	Blist := library.Members[useId].BorrowedBooks;
	const (
		Reset  = "\033[0m"
		Yellow = "\033[33m"
	)

	io.Println(Yellow , "Borrowed Books Of " ,library.Members[useId].Name , Reset)
	io.Printf("%-30s %-30s %-30s %-30s\n" , "ID","Title","Author","Status");
	io.Println(strings.Repeat("-" , 120))
	for _,book := range Blist {
		io.Printf("%-30d %-30s %-30s %-30s\n" , book.Id , book.Title , book.Author , book.Status);
	}
	io.Println()
}