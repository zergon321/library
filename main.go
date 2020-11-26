package main

import (
	"library/api"
	"library/repo"
	"library/www"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := repo.NewLibraryDatabase("mysql", "root:322453az@/library")
	handleError(err)

	router := gin.Default()

	router.Use(sessions.Sessions("lib-session",
		sessions.NewCookieStore([]byte(
			"ad0c973f-84d4-496b-85c1-d24ed49c3882"))))

	apiRouter := router.Group("/api")

	api.UserRoutes(apiRouter, db)
	api.BookRoutes(apiRouter, db)

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")
	www.Pages(router)

	err = router.Run("127.0.0.1:8000")
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
