package view

// UserAddedReply is sent back
// to the client when a new user
// has been added in the database.
type UserAddedReply struct {
	ID             int64  `json:"id"`
	PersonalNumber string `json:"personal_number"`
}

// BookAddedReply is sent back
// to the client when a new book
// has been added in the database.
type BookAddedReply struct {
	ID              int64  `json:"id"`
	InventoryNumber string `json:"inventory_number"`
}
