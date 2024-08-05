package middleware

import (
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type UserClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.Claims
}


func Encode( id int , email string) (string , error) {
	var err error = godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}
	var SecretKey = []byte(os.Getenv("SECRETKEY"))

	itoken := jwt.NewWithClaims(jwt.SigningMethodES256 , UserClaims{})
	token, err := itoken.SignedString(SecretKey)
	if err != nil {
		return "" , nil
	}
	return token , nil
}