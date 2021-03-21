package ajax

import (
	"encoding/hex"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"library/repo"
	"library/view"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// rentBookHandler rents book for the
// current authorized user.
func rentBookHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userIDData := session.Get("user-id")

		if userIDData == nil {
			c.String(http.StatusUnauthorized,
				"authorization required")
		}

		userID := userIDData.(int)
		bookIDParam := c.Query("book_id")
		bookID, err := strconv.Atoi(bookIDParam)

		if err != nil {
			c.String(http.StatusBadRequest,
				"invalid format: %v", bookIDParam)
		}

		_, err = db.RentBookForUser(userID, bookID)

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't rent the book for the user: %s", err.Error())
		}

		c.Status(http.StatusOK)
	}
}

// authHandler returns a web API method for user authentication.
func authHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		body := new(view.UserAuthRequest)
		err := c.BindJSON(body)

		if err != nil {
			c.String(http.StatusBadRequest,
				"couldn't read the body", err)

			return
		}

		user, err := db.GetUserByNickname(body.Login)

		if err != nil {
			c.String(http.StatusUnauthorized,
				"couldn't extract any data for the login", err)

			return
		}

		if user.Password != body.Password {
			c.String(http.StatusUnauthorized,
				"password incorrect")
		}

		session.Set("user-id", user.ID)
		session.Set("username", user.Nickname)
		err = session.Save()

		if err != nil {
			c.String(http.StatusInternalServerError,
				"couldn't save the session", err)
		}

		c.Status(http.StatusOK)
	}
}

// signUpHandler returns an API function for addding a new user.
func signUpHandler(db *repo.LibraryDatabase) func(c *gin.Context) {
	return func(c *gin.Context) {
		body := new(view.SignUpRequest)
		err := c.BindJSON(body)

		if err != nil {
			c.String(http.StatusBadRequest,
				"couldn't read the body", err)

			return
		}

		if body.Password != body.ConfirmPassword {
			c.String(http.StatusBadRequest,
				"password and confirmation do not match")
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
			Nickname: body.Nickname,
			Name:     body.Name,
			Email:    body.Email,
			Surname:  body.Surname,
			Patronim: body.Patronim,
			Password: body.Password,
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

func UserAJAX(rg *gin.RouterGroup, db *repo.LibraryDatabase) {
	usersRouter := rg.Group("/user")

	usersRouter.POST("/sign-up", signUpHandler(db))
	usersRouter.POST("/auth", authHandler(db))
	usersRouter.GET("/rent-book", rentBookHandler(db))
}
