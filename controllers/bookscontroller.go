package bookscontroller

import (
	"net/http"

	"github.com/MHafizAF/bookself-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var books []models.Book

	models.DB.Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"message": "Success get all books",
		"data":    books,
	})

}

func Show(c *gin.Context) {

	var book models.Book

	id := c.Param("id")
	errors := models.DB.First(&book, id).Error

	if errors != nil {
		switch errors {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Book not found",
			})

			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": errors.Error(),
			})

			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book Found",
		"data":    book,
	})

}

func Create(c *gin.Context) {

	var book models.Book

	errors := c.ShouldBindJSON(&book)

	if errors != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errors.Error(),
		})

		return
	}

	models.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{
		"message": "New book has been created",
		"data":    book,
	})

}

func Update(c *gin.Context) {
	var book models.Book

	id := c.Param("id")
	errors := c.ShouldBindJSON(&book)

	if errors != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": errors.Error(),
		})

		return
	}

	if models.DB.Model(&book).Where("id = ?", id).Updates(&book).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot update book",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book updated successfully",
	})

}

func Delete(c *gin.Context) {

	var book models.Book

	id := c.Param("id")

	if models.DB.Delete(&book, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Cannot delete book",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book deleted successfully",
	})

}
