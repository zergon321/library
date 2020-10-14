package repo

import (
	"time"
)

// User is a human who
// takes books from the library.
type User struct {
	ID             int    `db:"id"              json:"id"`
	PersonalNumber string `db:"personal_number" json:"personal_number"`
	Nickname       string `db:"nickname"        json:"nickname"`
	Name           string `db:"name"            json:"name"`
	Surname        string `db:"surname"         json:"surname"`
	Group          string `db:"group"           json:"group"`
	Grade          int16  `db:"grade"           json:"grade"`
}

// Book represents a book
// stored in the library.
type Book struct {
	ID              int    `db:"id"               json:"id"`
	Name            string `db:"name"             json:"name"`
	AuthorName      string `db:"author_name"      json:"author_name"`
	AuthorSurname   string `db:"author_surname"   json:"author_surname"`
	InventoryNumber string `db:"inventory_number" json:"inventory_number"`
}

// UserToBook represents a
// relation between the book
// and the user who took it.
type UserToBook struct {
	UserID   int        `db:"user_id"  json:"user_id"`
	BookID   int        `db:"book_id"  json:"book_id"`
	Taken    time.Time  `db:"taken"    json:"taken"`
	Expires  time.Time  `db:"expires"  json:"expires"`
	Returned *time.Time `db:"returned" json:"returned"`
}
