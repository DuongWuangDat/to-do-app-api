package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DuongWuangDat/to-do-app-api/database"
	"github.com/DuongWuangDat/to-do-app-api/routes"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(database.ConnectString))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database ping successfully")
	database.Collection = client.Database(database.DBName).Collection(database.ColName)
	fmt.Println("Connection instance is prepared successfully")
	r := gin.Default()
	r.GET("/tasks", routes.GetAllTask)
	r.GET("/tasks/:id", routes.GetOneTask)
	r.POST("/tasks", routes.AddNewTask)
	r.PUT("/tasks/:id", routes.UpdateTask)
	r.DELETE("/tasks/:id", routes.DeleteTask)
	r.Run(":8080")
}
