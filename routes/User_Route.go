package routes

import (
	"net/http"

	"github.com/DuongWuangDat/to-do-app-api/models"
	"github.com/DuongWuangDat/to-do-app-api/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind body",
		})
		return
	}
	user.PassWord, err = utils.HashPassword(user.PassWord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	err = models.SignUp(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to bind body",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Sign up new user successfully",
		"data":    user,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	tokenstring, err := models.Login(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": tokenstring,
	})
}
