package main

import (
	bookscontroller "github.com/MHafizAF/bookself-api/controllers"
	"github.com/MHafizAF/bookself-api/controllers/authentication"
	"github.com/MHafizAF/bookself-api/middlewares"
	"github.com/MHafizAF/bookself-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()
	models.ConnectDatabase()

	// authentication routes
	routes.POST("/api/register", authentication.Register)
	routes.POST("/api/sign-in", authentication.SignIn)

	routes.Use(middlewares.JwtAuthMiddleware())

	// current user
	routes.GET("/api/current-user", authentication.CurrentUser)

	// books routes
	routes.GET("/api/books", bookscontroller.Index)
	routes.GET("/api/books/:id", bookscontroller.Show)
	routes.POST("/api/books", bookscontroller.Create)
	routes.PUT("/api/books/:id", bookscontroller.Update)
	routes.DELETE("/api/books/:id", bookscontroller.Delete)

	routes.Run()
}
