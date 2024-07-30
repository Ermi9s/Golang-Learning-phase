# Library Management API

## `models` Package

### Types

- **`Book`**
  - `Id` (int): Unique book ID.
  - `Title` (string): Book title.
  - `Author` (string): Book author.
  - `Status` (string): "Available" or "Borrowed".

- **`Member`**
  - `Id` (int): Member ID.
  - `Name` (string): Member name.
  - `BorrowedBooks` ([]Book): List of borrowed books.

## `controllers` Package

### Functions

- **`AddBook(library main.Library)`**
  - Adds a book to the library.
  - Prompts for book title and author.

- **`RemoveBook(library main.Library)`**
  - Removes a book by its ID.

- **`BorrowBook(library main.Library, memberID int)`**
  - Allows a member to borrow a book.
  - Prompts for book ID.

- **`ReturnBook(library main.Library, memberId int)`**
  - Allows a member to return a borrowed book.
  - Prompts for book ID.

- **`UserBorrowList(library main.Library, useId int)`**
  - Displays borrowed books of a member.

