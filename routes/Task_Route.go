package routes

import (
	"net/http"

	"github.com/DuongWuangDat/to-do-app-api/models"
	"github.com/DuongWuangDat/to-do-app-api/utils"
	"github.com/gin-gonic/gin"
)

func GetAllTask(c *gin.Context) {
	tokenstring := utils.GetTokenStringFromHeader(c)
	if tokenstring == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Login first",
		})
		return
	}
	tasks, err := models.GetAll(tokenstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})

}

func GetOneTask(c *gin.Context) {
	tokenstring := utils.GetTokenStringFromHeader(c)
	if tokenstring == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Login first",
		})
		return
	}
	id := c.Param("id")
	task, err := models.GetOne(id, tokenstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}
func DeleteTask(c *gin.Context) {
	tokenstring := utils.GetTokenStringFromHeader(c)
	if tokenstring == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Login first",
		})
		return
	}
	id := c.Param("id")
	err := models.DeleteTask(id, tokenstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete successful",
	})
}
func AddNewTask(c *gin.Context) {
	tokenstring := utils.GetTokenStringFromHeader(c)
	if tokenstring == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Login first",
		})
		return
	}
	var body models.Task

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	id, err := body.CreateTask(tokenstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Add successful",
		"id":      id,
	})
}
func UpdateTask(c *gin.Context) {
	tokenstring := utils.GetTokenStringFromHeader(c)
	if tokenstring == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Login first",
		})
		return
	}
	id := c.Param("id")
	var body models.Task
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	task, err := models.GetOne(id, tokenstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if body.Title == "" {
		body.Title = task.Title
	}
	err = body.UpdateTask(id, tokenstring)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update successful",
	})
}
