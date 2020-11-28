package view

import "time"

// BookInfo holds the information on
// the book taken by the user.
type BookInfo struct {
	Name            string     `db:"name"             json:"name"`
	AuthorName      string     `db:"author_name"      json:"author_name"`
	AuthorSurname   string     `db:"author_surname"   json:"author_surname"`
	InventoryNumber string     `db:"inventory_number" json:"inventory_number"`
	Taken           time.Time  `db:"taken"            json:"taken"`
	Expires         time.Time  `db:"expires"          json:"expires"`
	Returned        *time.Time `db:"returned"         json:"returned"`
}
