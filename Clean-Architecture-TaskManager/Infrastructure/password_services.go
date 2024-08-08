package infrastructure

import (
	"log"
	"os"

	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func Encode(id primitive.ObjectID , email string , is_admin bool) (string , error) {
	var err error = godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}
	var SecretKey = []byte(os.Getenv("SECRETKEY"))
	
	itoken := jwt.NewWithClaims(jwt.SigningMethodHS256 , domain.UserClaims{
		ID: id,
		Email: email,
		Is_admin: is_admin,
	})
	token, err := itoken.SignedString(SecretKey)
	if err != nil {
		return "" , err
	}
	return token , nil
}