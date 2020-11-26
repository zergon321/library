package www

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
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
		session := sessions.Default(c)
		userID := session.Get("user-id")

		if userID == nil {
			c.HTML(http.StatusOK, "sign-up.html", nil)
		} else {
			c.HTML(http.StatusForbidden, "error.html", gin.H{
				"status":  http.StatusForbidden,
				"message": "alredy authorized",
			})
		}
	}
}

func loginHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "log-in.html", nil)
	}
}

func userHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		usernameData := session.Get("username")

		if usernameData == nil {
			c.Redirect(http.StatusPermanentRedirect, "/log-in")
			return
		}

		username := usernameData.(string)

		c.HTML(http.StatusOK, "user.html", gin.H{
			"nickname": username,
		})
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
	r.GET("/log-in", loginHandler())
	r.GET("/signed-up", signedUpHandler())
	r.GET("/error", errorHandler())
	r.GET("/user", userHandler())
}
