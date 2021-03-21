package repo

import (
	"time"
)

// User is a human who
// takes books from the library.
type User struct {
	ID       int    `db:"id"              json:"id"`
	Nickname string `db:"nickname"        json:"nickname"`
	Name     string `db:"name"            json:"name"`
	Email    string `db:"email"           json:"email"`
	Surname  string `db:"surname"         json:"surname"`
	Patronim string `db:"patronim"        json:"patronim"`
	Password string `db:"password"        json:"password"`
}

// Book represents a book
// stored in the library.
type Book struct {
	ID            int     `db:"id"               json:"id"`
	Name          string  `db:"name"             json:"name"`
	AuthorName    string  `db:"author_name"      json:"author_name"`
	AuthorSurname string  `db:"author_surname"   json:"author_surname"`
	VendorCode    string  `db:"vendor_code"      json:"vendor_code"`
	Price         float64 `db:"price"            json:"price"`
}

// UserToBook represents a
// relation between the book
// and the user who took it.
type UserToBook struct {
	UserID    int        `db:"user_id"   json:"user_id"`
	BookID    int        `db:"book_id"   json:"book_id"`
	Ordered   time.Time  `db:"ordered"   json:"ordered"`
	Delivered *time.Time `db:"delivered" json:"delivered"`
}
