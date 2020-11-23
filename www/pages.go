package www

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Welcome to the public library",
		})
	}
}

func signUpHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign-up.html", nil)
	}
}

func signedUpHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		nickname := c.Query("nickname")
		personalNumber := c.Query("personal_number")

		c.HTML(http.StatusCreated, "signed-up.html", gin.H{
			"nickname":       nickname,
			"personalNumber": personalNumber,
		})
	}
}

func errorHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		status := c.Query("status")
		message := c.Query("message")

		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"status":  status,
			"message": message,
		})
	}
}

// Pages sets up the routes
// for the site pages.
func Pages(r *gin.Engine) {
	r.GET("/index", indexHandler())
	r.GET("/sign-up", signUpHandler())
	r.GET("/signed-up", signedUpHandler())
	r.GET("/error", errorHandler())
}
