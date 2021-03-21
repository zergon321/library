package view

import "time"

// BookInfo holds the information on
// the book taken by the user.
type BookInfo struct {
	Name          string     `db:"name"             json:"name"`
	AuthorName    string     `db:"author_name"      json:"author_name"`
	AuthorSurname string     `db:"author_surname"   json:"author_surname"`
	VendorCode    string     `db:"vendor_code"      json:"vendor_code"`
	Price         float64    `db:"price"            json:"price"`
	Ordered       time.Time  `db:"ordered"          json:"ordered"`
	Delivered     *time.Time `db:"delivered"        json:"delivered"`
}
