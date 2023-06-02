package authentication

import (
	"net/http"

	"github.com/MHafizAF/bookself-api/models"
	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if errors := c.ShouldBindJSON(&input); errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.Error(),
		})

		return
	}

	user := models.User{}
	user.Username = input.Username
	user.Password = input.Password

	_, errors := user.SaveUser()

	if errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registered successfully",
	})

}
