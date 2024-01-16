package utils

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	ID primitive.ObjectID `bson: "id" json:"id"`
	jwt.RegisteredClaims
}

func GenerateTokenString(claim Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenstring, err := token.SignedString([]byte("SECRETKEY"))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenstring, err
}
