package infrastructure

import (
	"log"
	"os"

	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type KeyServices struct {}

func (KeyServices)Encode(id string , email string , is_admin bool) (string , error) {
	var err error = godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env\n" , err.Error())
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

func (KeyServices)HashPassWord(spass string) string {
	hasshedPasskey,_ := bcrypt.GenerateFromPassword([]byte(spass) , bcrypt.DefaultCost);
	return string(hasshedPasskey)
}