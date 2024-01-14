package main

import (
	"context"
	"fmt"
	"log"

	"github.com/DuongWuangDat/to-do-app-api/database"
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
}
