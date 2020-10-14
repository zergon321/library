package view

// AddUserRequest represents a JSON
// request to add a new user.
type AddUserRequest struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Group    string `json:"group"`
	Grade    int16  `json:"grade"`
}

// AddBookRequest represents a JSON
// request to add a new book.
type AddBookRequest struct {
	Name          string `json:"name"`
	AuthorName    string `json:"author_name"`
	AuthorSurname string `json:"author_surname"`
}
