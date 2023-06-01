package main

import (
	"github.com/MHafizAF/bookself-api/controllers/productscontroller"
	"github.com/MHafizAF/bookself-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	routes := gin.Default()
	models.ConnectDatabase()

	routes.GET("/api/books", productscontroller.Index)
	routes.GET("/api/books/:id", productscontroller.Show)
	routes.POST("/api/books", productscontroller.Create)
	routes.PUT("/api/books/:id", productscontroller.Update)
	routes.DELETE("/api/books/:id", productscontroller.Delete)

	routes.Run()
}
