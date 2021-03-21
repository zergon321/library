package api

import (
	"encoding/hex"
	"math/rand"
	"net/http"
	"time"

	"library/repo"
	"library/view"

	"github.com/gin-gonic/gin"
)

func addBookHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
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
			Name:          body.Name,
			AuthorName:    body.AuthorName,
			AuthorSurname: body.AuthorSurname,
			VendorCode:    inventoryNumber,
			Price:         body.Price,
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
	}
}

func getBooksOnHoldHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
		books, err := db.GetBooksOnHold()

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't get books being currently on hold")

			return
		}

		c.JSON(http.StatusOK, books)
	}
}

// BookRoutes sets up the routes for web functions
// for working with books.
func BookRoutes(rg *gin.RouterGroup, db *repo.LibraryDatabase) {
	booksRouter := rg.Group("/books")

	booksRouter.POST("/", addBookHandler(db))
	booksRouter.GET("/on-hold", getBooksOnHoldHandler(db))
}
