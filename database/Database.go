package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var ConnectString = "mongodb+srv://duongquangdat172004:quangdat@go-gin-gonic.ra4g6s0.mongodb.net/?retryWrites=true&w=majority"
var DBName = "ToDoApp"
var ColToDoName = "ToDoList"
var ColUserName = "User"
var Collection *mongo.Collection
var Client *mongo.Client

func ConnectDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(ConnectString))
	Client = client
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	fmt.Println("Ping succesfully")
}
