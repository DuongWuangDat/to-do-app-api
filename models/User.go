package models

import (
	"context"
	"errors"

	"github.com/DuongWuangDat/to-do-app-api/database"
	"github.com/DuongWuangDat/to-do-app-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	UserName string             `bson:"username" json:"username"`
	PassWord string             `bson:"password" json:"password"`
}

func SignUp(user User) error {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColUserName)
	filter := bson.M{
		"username": user.UserName,
	}
	err := database.Collection.FindOne(context.Background(), filter).Err()
	if err == nil {
		return errors.New("username existed")
	}
	_, err = database.Collection.InsertOne(context.Background(), &user)
	return err
}

func Login(username string, password string) (string, error) {
	database.Collection = database.Client.Database(database.DBName).Collection(database.ColUserName)
	filter := bson.M{
		"username": username,
	}
	var user User
	err := database.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return "", errors.New("invalid username")
	}
	err = utils.ValidatePassword(user.PassWord, password)
	if err != nil {
		return "", errors.New("invalid username and password")
	}
	claim := utils.Claims{
		ID: user.ID,
	}
	token, err := utils.GenerateTokenString(claim)
	return token, err
}
