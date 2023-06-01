package main

import (
	"github.com/MHafizAF/bookself-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()
	models.ConnectDatabase()

	routes.GET("/api/books", productcontroller.Index)
	routes.GET("/api/books/:id", productcontroller.Show)
	routes.POST("/api/books", productcontroller.Create)
	routes.PUT("/api/books/:id", productcontroller.Update)
	routes.DELETE("/api/books/:id", productcontroller.Delete)

	routes.Run()
}
