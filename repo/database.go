package repo

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// LibraryDatabase is the database
// to store the information about
// books and users.
type LibraryDatabase struct {
	client *sqlx.DB
}

// GetUserByNickname returns the user if it exists,
// or error if doesn't.
func (db *LibraryDatabase) GetUserByNickname(nickname string) (*User, error) {
	user := new(User)
	query := "SELECT * FROM library.users WHERE users.nickname = ?;"
	err := db.client.Get(user, query, nickname)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByID returns the user if it exists,
// or error if doesn't.
func (db *LibraryDatabase) GetUserByID(id int) (*User, error) {
	user := new(User)
	query := "SELECT * FROM library.users WHERE users.id = ?;"
	err := db.client.Get(user, query, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// AddUser adss a new user in the database.
func (db *LibraryDatabase) AddUser(user *User) (int64, error) {
	query := `INSERT INTO library.users
			  (personal_number, nickname,
			  users.name, surname, email, users.group, grade, password)
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
	res, err := db.client.Exec(query, user.PersonalNumber, user.Nickname,
		user.Name, user.Surname, user.Email, user.Group, user.Grade, user.Password)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

// AddBook adds a new book in the database.
func (db *LibraryDatabase) AddBook(book *Book) (int64, error) {
	query := `INSERT INTO library.books
			  (books.name, author_name, author_surname,
			  inventory_number) VALUES (?, ?, ?, ?);`
	res, err := db.client.Exec(query, book.Name, book.AuthorName,
		book.AuthorSurname, book.InventoryNumber)

	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}

// RentBookForUser rents the book with the specified ID
// for the user with the specified ID.
func (db *LibraryDatabase) RentBookForUser(userID, bookID int) (*UserToBook, error) {
	onHold, err := db.IsBookOnHold(bookID)

	if err != nil {
		return nil, err
	}

	if onHold {
		return nil,
			fmt.Errorf("the book with ID = %d is currently on hold", bookID)
	}

	query := `INSERT INTO library.users_to_books
			  (user_id, book_id, taken, expires, returned)
			  VALUES (?, ?, ?, ?, NULL);`
	userToBook := &UserToBook{
		UserID:   userID,
		BookID:   bookID,
		Taken:    time.Now(),
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Returned: nil,
	}
	_, err = db.client.Exec(query, userToBook.UserID,
		userToBook.BookID, userToBook.Taken, userToBook.Expires)

	if err != nil {
		return nil, err
	}

	return userToBook, nil
}

// GetUserBooks returns all the books ever taken or being taken
// by the user with the specified ID.
func (db *LibraryDatabase) GetUserBooks(userID int) ([]*Book, error) {
	query := `SELECT books.id, books.name, books.author_name,
			  books.author_surname, books.inventory_number
			  FROM library.users
			  INNER JOIN library.users_to_books
			  ON users.id = users_to_books.user_id
			  INNER JOIN library.books
			  ON users_to_books.book_id = books.id
			  WHERE users.id = ?;`
	books := []*Book{}
	err := db.client.Select(&books, query, userID)

	if err != nil {
		return nil, err
	}

	return books, nil
}

// GetUserBooksReturned returns all the books returned by
// the user in the library.
func (db *LibraryDatabase) GetUserBooksReturned(userID int) ([]*Book, error) {
	query := `SELECT books.id, books.name, books.author_name,
			  books.author_surname, books.inventory_number
			  FROM library.users
			  INNER JOIN library.users_to_books
			  ON users.id = users_to_books.user_id
			  INNER JOIN library.books
			  ON users_to_books.book_id = books.id
			  WHERE users.id = ? AND returned IS NOT NULL;`
	books := []*Book{}
	err := db.client.Select(&books, query, userID)

	if err != nil {
		return nil, err
	}

	return books, nil
}

// GetUserBooksOnHold returns all the books being currently on hold
// by the user.
func (db *LibraryDatabase) GetUserBooksOnHold(userID int) ([]*Book, error) {
	query := `SELECT books.id, books.name, books.author_name,
			  books.author_surname, books.inventory_number
			  FROM library.users
			  INNER JOIN library.users_to_books
			  ON users.id = users_to_books.user_id
			  INNER JOIN library.books
			  ON users_to_books.book_id = books.id
			  WHERE users.id = ? AND returned IS NULL;`
	books := []*Book{}
	err := db.client.Select(&books, query, userID)

	if err != nil {
		return nil, err
	}

	return books, nil
}

// IsBookOnHold returns true if the book is on hold,
// and false otherwise.
func (db *LibraryDatabase) IsBookOnHold(bookID int) (bool, error) {
	query := `SELECT ? IN
			  (SELECT books.id
				FROM library.books
				INNER JOIN library.users_to_books
				ON books.id = users_to_books.book_id
				WHERE users_to_books.returned IS NULL) AS lol;`
	result := false
	err := db.client.Get(&result, query, bookID)

	if err != nil {
		return false, err
	}

	return result, nil
}

// GetBooksOnHold returns all the books currently being on hold.
func (db *LibraryDatabase) GetBooksOnHold() ([]*Book, error) {
	query := `SELECT books.id, books.name, books.author_name,
			  books.author_surname, books.inventory_number
			  FROM library.books
			  INNER JOIN library.users_to_books
			  ON books.id = users_to_books.book_id
			  WHERE users_to_books.returned IS NULL;`
	books := []*Book{}
	err := db.client.Select(&books, query)

	if err != nil {
		return nil, err
	}

	return books, nil
}

// GetBooksExpired returns all the expired books.
func (db *LibraryDatabase) GetBooksExpired() ([]*Book, error) {
	query := `SELECT books.id, books.name, books.author_name,
			  books.author_surname, books.inventory_number
			  FROM library.books
			  INNER JOIN library.users_to_books
			  ON books.id = users_to_books.book_id
			  WHERE users_to_books.returned IS NULL
				  AND NOW() > users_to_books.expires;`
	books := []*Book{}
	err := db.client.Select(&books, query)

	if err != nil {
		return nil, err
	}

	return books, nil
}

// NewLibraryDatabase creates a new library database
// with specified driver and connection string.
func NewLibraryDatabase(db, connStr string) (*LibraryDatabase, error) {
	instance, err := sqlx.Open(db, connStr)

	if err != nil {
		return nil, err
	}

	return &LibraryDatabase{
		client: instance,
	}, nil
}
