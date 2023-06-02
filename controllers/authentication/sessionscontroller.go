package authentication

import (
	"net/http"

	"github.com/MHafizAF/bookself-api/models"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignIn(c *gin.Context) {
	var input LoginInput

	if errors := c.ShouldBindJSON(&input); errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errors.Error(),
		})

		return
	}

	user := models.User{}
	user.Username = input.Username
	user.Password = input.Password

	token, errors := models.SignInCheck(user.Username, user.Password)

	if errors != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "incorrect username or password",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sign in successfully",
		"token":   token,
	})

}
