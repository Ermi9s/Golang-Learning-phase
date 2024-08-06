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
	
	log.Println(id , email)
	itoken := jwt.NewWithClaims(jwt.SigningMethodHS256 , models.UserClaims{
		ID: id,
		Email: email,
	})
	token, err := itoken.SignedString(SecretKey)
	if err != nil {
		return "" , err
	}
	return token , nil
}