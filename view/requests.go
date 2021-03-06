package view

// SignUpRequest represents a JSON
// request to add a new user.
type SignUpRequest struct {
	Nickname        string `json:"nickname"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Patronim        string `json:"patronim"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Agree           bool   `json:"agree"`
}

// AddUserRequest represents a JSON
// request to add a new user.
type AddUserRequest struct {
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Patronim string `json:"patronim"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AddBookRequest represents a JSON
// request to add a new book.
type AddBookRequest struct {
	Name          string  `json:"name"`
	AuthorName    string  `json:"author_name"`
	AuthorSurname string  `json:"author_surname"`
	Price         float64 `json:"price"`
}

// UserAuthRequest represents a JSON
// request to log in to the site.
type UserAuthRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
