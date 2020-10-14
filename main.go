package main

import (
	"encoding/hex"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"library/repo"
	"library/view"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := repo.NewLibraryDatabase("mysql", "root:322453az@/library")
	handleError(err)

	router := gin.Default()
	apiRouter := router.Group("/api")
	apiUsersRouter := apiRouter.Group("/users")

	apiUsersRouter.POST("/", func(c *gin.Context) {
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
			Surname:        body.Surname,
			Group:          body.Group,
			Grade:          body.Grade,
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
	})

	apiUsersRouter.GET("/:id/books", func(c *gin.Context) {
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
	})

	apiUsersRouter.GET("/:id/books/returned", func(c *gin.Context) {
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
	})

	apiUsersRouter.GET("/:id/books/on-hold", func(c *gin.Context) {
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
	})

	apiUsersRouter.GET("/:id/rent-book",
		func(c *gin.Context) {
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
		})

	apiBooksRouter := apiRouter.Group("/books")

	apiBooksRouter.POST("/", func(c *gin.Context) {
		body := new(view.AddBookRequest)
		err := c.BindJSON(body)

		if err != nil {
			c.String(http.StatusBadRequest,
				"couldn't read the body")

			return
		}

		inventoryNumberBytes := make([]byte, 16)
		rand.Seed(time.Now().UTC().UnixNano())
		_, err = rand.Read(inventoryNumberBytes)

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't generate an inventory number for the book")

			return
		}

		inventoryNumber := hex.EncodeToString(inventoryNumberBytes)
		book := &repo.Book{
			Name:            body.Name,
			AuthorName:      body.AuthorName,
			AuthorSurname:   body.AuthorSurname,
			InventoryNumber: inventoryNumber,
		}

		id, err := db.AddBook(book)

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't add the book in the database")

			return
		}

		reply := &view.BookAddedReply{
			ID:              id,
			InventoryNumber: inventoryNumber,
		}
		c.JSON(http.StatusAccepted, reply)
	})

	apiBooksRouter.GET("/on-hold", func(c *gin.Context) {
		books, err := db.GetBooksOnHold()

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't get books being currently on hold")

			return
		}

		c.JSON(http.StatusOK, books)
	})

	apiBooksRouter.GET("/expired", func(c *gin.Context) {
		books, err := db.GetBooksExpired()

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't get expired books")

			return
		}

		c.JSON(http.StatusOK, books)
	})

	err = router.Run("127.0.0.1:8000")
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
