package models

import (
	"context"
	"log"
	"time"

	"github.com/DuongWuangDat/to-do-app-api/database"
	"github.com/DuongWuangDat/to-do-app-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Title     string             `bson:"title" json:"title"`
	IsDone    bool               `bson:"isdone" json:"isdone"`
	CreatedAt int                `bson:"createdat" json:"createdat"`
}

func GetAll(tokenstring string) ([]Task, error) {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	var tasks []Task
	claim, err := utils.ParseToken(tokenstring)
	if err != nil {
		return tasks, err
	}
	filter := bson.M{
		"user_id": claim.ID,
	}
	cur, err := database.Collection.Find(context.Background(), filter)
	if err != nil {
		return tasks, err
	}
	err = cur.All(context.Background(), &tasks)
	return tasks, err
}

func GetOne(taskId string, tokenstring string) (Task, error) {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	claim, err := utils.ParseToken(tokenstring)
	if err != nil {
		log.Fatal(err)
	}
	id, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{
		"_id":     id,
		"user_id": claim.ID,
	}
	var task Task
	err = database.Collection.FindOne(context.Background(), filter).Decode(&task)
	return task, err

}

func (d *Task) CreateTask(tokenstring string) (string, error) {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	claim, err := utils.ParseToken(tokenstring)
	if err != nil {
		log.Fatal(err)
	}
	d.UserID = claim.ID
	rs, err := database.Collection.InsertOne(context.TODO(), d)
	return rs.InsertedID.(primitive.ObjectID).String(), err
}

func DeleteTask(taskId string, tokenstring string) error {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	claim, err := utils.ParseToken(tokenstring)
	if err != nil {
		log.Fatal(err)
	}
	id, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{
		"_id":     id,
		"user_id": claim.ID,
	}
	_, err = database.Collection.DeleteOne(context.Background(), filter)
	return err

}

func (d *Task) UpdateTask(taskID string, tokenstring string) error {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	claim, err := utils.ParseToken(tokenstring)
	if err != nil {
		log.Fatal(err)
	}
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{
		"_id":     id,
		"user_id": claim.ID,
	}
	update := bson.M{
		"$set": bson.M{
			"title":     d.Title,
			"isdone":    d.IsDone,
			"createdat": time.Now().Unix(),
		},
	}
	_, err = database.Collection.UpdateOne(context.Background(), filter, update)
	return err

}
