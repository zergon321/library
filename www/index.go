package www

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	}
}

// IndexRoutes sets up the routes
// for the site index.
func IndexRoutes(r *gin.Engine) {
	r.GET("/index", indexHandler())
}
