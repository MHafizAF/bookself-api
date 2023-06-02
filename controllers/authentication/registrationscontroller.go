package authentication

import (
	"net/http"

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

	c.JSON(http.StatusOK, gin.H{
		"message": "Validated!",
	})

}
