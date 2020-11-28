package api

import (
	"encoding/hex"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"library/repo"
	"library/view"

	"github.com/gin-gonic/gin"
)

// addUserHandler returns an API function for addding a new user.
func addUserHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
		body := new(view.AddUserRequest)
		err := c.BindJSON(body)

		if err != nil {
			c.String(http.StatusBadRequest,
				"couldn't read the body", err)

			return
		}

		personalNumberBytes := make([]byte, 8)
		rand.Seed(time.Now().UTC().UnixNano())
		_, err = rand.Read(personalNumberBytes)

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't generate a personal number for the user")

			return
		}

		personalNumber := hex.EncodeToString(personalNumberBytes)
		user := &repo.User{
			PersonalNumber: personalNumber,
			Nickname:       body.Nickname,
			Name:           body.Name,
			Email:          body.Email,
			Surname:        body.Surname,
			Group:          body.Group,
			Grade:          body.Grade,
			Password:       body.Password,
		}

		id, err := db.AddUser(user)

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't add the user in the database", err)

			return
		}

		reply := &view.UserAddedReply{
			ID:             id,
			PersonalNumber: personalNumber,
		}

		c.JSON(http.StatusAccepted, reply)
	}
}

// getUserBooksHandler returns a web API function for
// obtaining user's books.
func getUserBooksHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.String(http.StatusBadRequest,
				"invalid value for parameter 'id'")

			return
		}

		books, err := db.GetUserBooks(id)

		if err != nil {
			c.String(http.StatusBadRequest,
				"couldn't extract books for the user", err)

			return
		}

		c.JSON(http.StatusOK, books)
	}
}

// getUserBooksReturnedHandler returns a web API function
// for obtaining all the books ever returned by the user.
func getUserBooksReturnedHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.String(http.StatusBadRequest,
				"invalid value for parameter 'id'")

			return
		}

		books, err := db.GetUserBooksReturned(id)

		if err != nil {
			c.String(http.StatusBadRequest,
				"couldn't extract books for the user", err)

			return
		}

		c.JSON(http.StatusOK, books)
	}
}

// getUserBooksOnHoldHandler returns a web API function
// to obtain all the books currently held by the user.
func getUserBooksOnHoldHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.String(http.StatusBadRequest,
				"invalid value for parameter 'id'")

			return
		}

		books, err := db.GetUserBooksOnHold(id)

		if err != nil {
			c.String(http.StatusBadRequest,
				"couldn't extract books for the user", err)

			return
		}

		c.JSON(http.StatusOK, books)
	}
}

// userRentBookHandler returns a web API method for
// renting a certain book for the user.
func userRentBookHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
		userIDStr := c.Param("id")
		bookIDStr := c.Query("book_id")

		userID, err := strconv.Atoi(userIDStr)

		if err != nil {
			c.String(http.StatusBadRequest,
				"invalid value for parameter 'user_id'")

			return
		}

		bookID, err := strconv.Atoi(bookIDStr)

		if err != nil {
			c.String(http.StatusBadRequest,
				"invalid value for query parameter 'book_id'")

			return
		}

		result, err := db.RentBookForUser(userID, bookID)

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't rent the book for the user", err)

			return
		}

		c.JSON(http.StatusAccepted, result)
	}
}

// UserRoutes sets up the routes for handling
// requests for library users.
func UserRoutes(rg *gin.RouterGroup, db *repo.LibraryDatabase) {
	usersRouter := rg.Group("/users")

	usersRouter.POST("/", addUserHandler(db))
	usersRouter.GET("/:id/books", getUserBooksHandler(db))
	usersRouter.GET("/:id/books/returned", getUserBooksReturnedHandler(db))
	usersRouter.GET("/:id/books/on-hold", getUserBooksOnHoldHandler(db))
	usersRouter.GET("/:id/rent-book", userRentBookHandler(db))
}
