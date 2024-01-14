package database

import "go.mongodb.org/mongo-driver/mongo"

var ConnectString = "mongodb+srv://duongquangdat172004:quangdat@go-gin-gonic.ra4g6s0.mongodb.net/?retryWrites=true&w=majority"
var DBName = "ToDoApp"
var ColName = "ToDoList"
var Collection *mongo.Collection
