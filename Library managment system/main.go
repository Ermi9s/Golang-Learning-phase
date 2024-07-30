package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	io "fmt"
	controller "github.com/ermi9s/go/controllers"
	model "github.com/ermi9s/go/models"
	service "github.com/ermi9s/go/services"
)

func main() {
	var library = service.Library{
		Books : make(map[int]*model.Book),
		Members: make(map[int]*model.Member),
		NextMemberId: 1,
		NextBookId: 1,
	}

	session := true
	for session{
		HOME:
		var current_user model.Member
		var found bool;
		reader := bufio.NewReader(os.Stdin);


		io.Println("\n\t\t*** WELCOME TO THE LIBRARY ***");io.Print("\n");
		
		var memberID int;
		
		{
		InvalidInput := true;
		for InvalidInput {
			io.Println("Enter 0 to SignUp or Type 'exit' to EXIT")
			io.Print("Enter Your ID: ")
			input , err := reader.ReadString('\n');

			if err == nil {
				if strings.TrimSpace((input)) == "exit" {
					return;
				}
				id ,er:= strconv.Atoi(strings.TrimSpace((input)))
				if er != nil {
					io.Println("Invalid Id!")
					continue
				}else{
					memberID = id;
					break;
				}

			}else{
				io.Println("Invalid Input!")
				continue
			}

		}}

		if memberID == 0 {
			var name string;
			{
			InvalidInput := true;
			for InvalidInput {
				io.Print("Enter Your Name: ");
				input,err := reader.ReadString('\n');

				if err == nil {
					name = input;
					break;
				}else {
					io.Println("Invalid Input");
				}
			}}

			newID := library.NextMemberId;
			var borrowedBooks []model.Book
			var newMember  = model.Member{
				Id : newID,
				Name: name,
				BorrowedBooks: borrowedBooks,
			};
			
			library.Members[newID] = &newMember;
			io.Println(newMember.Name ," You have registered with id: ",newID);
			memberID = newID;
			library.NextMemberId ++;
		}

		if user,ok := library.Members[memberID]; ok {
			current_user = *user
			found = true;
		}
	
		ENRTY:
		if found {
			io.Println("\n\t\t*****Logged-In*****")
			io.Println()
			io.Println("User: ", current_user.Name);
			io.Println("User ID: ", current_user.Id);
			io.Println()
			io.Printf("%-30s %-30s %-30s %-30s\n" , "ID","Title","Author","Status");
			io.Println(strings.Repeat("-" , 120))
			Abooks := library.ListAvailableBooks();
			for i := 0; i < len(Abooks); i++{
				io.Printf("%-30d %-30s %-30s %-30s\n" , Abooks[i].Id , Abooks[i].Title , Abooks[i].Author , Abooks[i].Status);
			}

			Bbooks := library.ListBorrowedBooks()
			for i := 0; i < len(Bbooks); i++ {
				io.Printf("%-30d %-30s %-30s %-30s\n" , Bbooks[i].Id , Bbooks[i].Title , Bbooks[i].Author , Bbooks[i].Status);
			}
		
			{
				
				io.Println("1.ADD BOOK")
				io.Println("2.REMOVE BOOK")
				io.Println("3.BORROW BOOK")
				io.Println("4.RETURN BOOK")
				io.Println("5.YOUR BORROW LIST")
				io.Println("6.LOGOUT")
				io.Print("#ENTER YOUR CHOICE: ")
				icoice,_ := reader.ReadString('\n')
				choice,err := strconv.Atoi(strings.TrimSpace(icoice))
				if err != nil {
					io.Println("Invalid Input!")
					goto ENRTY
				}
				switch choice {
				case 1:
					controller.AddBook(library);
					library.NextBookId ++;
					reader.Reset(os.Stdin);
					goto ENRTY;
				case 2:
					controller.RemoveBook(library);
					reader.Reset(os.Stdin);
					goto ENRTY;
				case 3:
					controller.BorrowBook(library , current_user.Id);
					reader.Reset(os.Stdin);
					goto ENRTY;
				case 4:
					controller.ReturnBook(library , current_user.Id);
					reader.Reset(os.Stdin);
					goto ENRTY;
				case 5:
					controller.UserBorrowList(library , current_user.Id);
					reader.Reset(os.Stdin);
					goto ENRTY;
				case 6:
					goto HOME;
				default:
					goto ENRTY;
				}
			}
			
		}else{
			io.Println("User NOt Found!");
		}
	}
	
}