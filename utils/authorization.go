package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Claims struct {
	ID primitive.ObjectID `bson: "id" json:"id"`
	jwt.RegisteredClaims
}

func GetTokenStringFromHeader(c *gin.Context) string {
	tokenstring := ""
	header := c.Request.Header.Get("Authorization")
	parse := strings.Split(header, " ")
	if parse[0] == "Bearer" {
		tokenstring = parse[1]
	}
	fmt.Println(tokenstring)
	return tokenstring
}
func ParseToken(tokenstring string) (Claims, error) {
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenstring, claim, func(t *jwt.Token) (interface{}, error) { return []byte("SECRETKEY"), nil })
	if err != nil {
		return Claims{}, err
	}
	if !token.Valid {
		return Claims{}, err
	}
	return *claim, err
}

func GenerateTokenString(claim Claims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenstring, err := token.SignedString([]byte("SECRETKEY"))
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return tokenstring, err
}
