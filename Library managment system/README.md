# Library Management System

This is a simple command-line based Library Management System written in Go. It allows users to sign up, log in, and manage books within the library. Users can add, remove, borrow, and return books. The application uses buffered input for user interactions.

## Features

- Sign Up and Log In
- Add books to the library
- Remove books from the library
- Borrow books from the library
- Return borrowed books
- View a list of available and borrowed books
- Log out

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/ermi9s/library-management-system.git
    cd library-management-system
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

## Usage

1. **Run the application:**

    ```sh
    go run main.go
    ```

2. **Follow the prompts in the command-line interface:**

    - **Sign Up:** Enter `0` to sign up as a new user.
    - **Log In:** Enter your user ID to log in.
    - **Exit:** Type `exit` to quit the application.

3. **Menu Options:**
    - **1. ADD BOOK:** Add a new book to the library.
    - **2. REMOVE BOOK:** Remove a book from the library.
    - **3. BORROW BOOK:** Borrow a book from the library.
    - **4. RETURN BOOK:** Return a borrowed book to the library.
    - **5. YOUR BORROW LIST:** View your borrowed books.
    - **6. LOGOUT:** Log out and return to the main menu.

## Project Structure

- **controllers:** Contains the logic for adding, removing, borrowing, and returning books.
- **models:** Contains the data models for `Book` and `Member`.
- **services:** Contains the `Library` service which manages the book and member data.


