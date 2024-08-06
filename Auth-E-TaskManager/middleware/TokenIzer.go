package middleware

import (
	"log"
	"os"

	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func Encode( id primitive.ObjectID , email string) (string , error) {
	var err error = godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}
	var SecretKey = []byte(os.Getenv("SECRETKEY"))

	itoken := jwt.NewWithClaims(jwt.SigningMethodES256 , models.UserClaims{})
	token, err := itoken.SignedString(SecretKey)
	if err != nil {
		return "" , nil
	}
	return token , nil
}