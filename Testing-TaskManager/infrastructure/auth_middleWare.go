package infrastructure

import (
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func Validate(is_special bool , onlyadmin bool) func(context *gin.Context) {
	var err error = godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
		return nil
	}
	var SecretKey = []byte(os.Getenv("SECRETKEY"))

	return func(context *gin.Context) {
		authToken := context.GetHeader("Authorization")
		if authToken == "" {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"Message":"Authorization Header needed"})
			context.Abort()
			return
		}

		authSplitted := strings.Split(authToken , " ")
		if len(authSplitted) != 2 || strings.ToLower(authSplitted[0]) != "bearer" {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Token-Header invalid"})
			context.Abort()
			return
		}
		
		token , err := jwt.ParseWithClaims(authSplitted[1] ,&domain.UserClaims{} ,func(t *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Invalid Token a"})
			context.Abort()
			return
		}

		payload,ok := token.Claims.(*domain.UserClaims); 

		if !ok {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Invalid Token Claims"})
			context.Abort()
			return
		}

		if is_special && onlyadmin && !payload.Is_admin {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Not Authorized"})
			context.Abort()
			return
		}

		// log.Println(payload)
		context.Set("payload" , payload)
		context.Next()
	}
}