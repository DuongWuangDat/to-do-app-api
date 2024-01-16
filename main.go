package main

import (
	"context"

	"github.com/DuongWuangDat/to-do-app-api/database"
	"github.com/DuongWuangDat/to-do-app-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDB()
	defer func() {
		if err := database.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	r := gin.Default()
	r.GET("/tasks", routes.GetAllTask)
	r.GET("/tasks/:id", routes.GetOneTask)
	r.POST("/signup", routes.SignUp)
	r.POST("/login", routes.Login)
	r.POST("/tasks", routes.AddNewTask)
	r.PUT("/tasks/:id", routes.UpdateTask)
	r.DELETE("/tasks/:id", routes.DeleteTask)
	r.Run(":8080")
}
