package models

import (
	"context"
	"log"
	"time"

	"github.com/DuongWuangDat/to-do-app-api/database"
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

func GetAll() ([]Task, error) {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	var tasks []Task
	cur, err := database.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return tasks, err
	}
	err = cur.All(context.Background(), &tasks)
	return tasks, err
}

func GetOne(taskId string) (Task, error) {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	id, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{
		"_id": id,
	}
	var task Task
	err = database.Collection.FindOne(context.Background(), filter).Decode(&task)
	return task, err

}

func (d *Task) CreateTask() (string, error) {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	rs, err := database.Collection.InsertOne(context.TODO(), d)
	return rs.InsertedID.(primitive.ObjectID).String(), err
}

func DeleteTask(taskId string) error {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	id, err := primitive.ObjectIDFromHex(taskId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{
		"_id": id,
	}
	_, err = database.Collection.DeleteOne(context.Background(), filter)
	return err

}

func (d *Task) UpdateTask(taskID string) error {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColToDoName)
	id, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{
		"_id": id,
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
