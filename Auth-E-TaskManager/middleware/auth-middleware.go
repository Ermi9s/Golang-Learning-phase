package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func Validate(is_special bool , DBM *services.DataBaseManager) func(context gin.Context) {
	var err error = godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}

	var SecretKey = []byte(os.Getenv("SECRETKEY"))
	return func(context gin.Context) {
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

		token , err := jwt.Parse(authSplitted[1] , func(t *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Invalid Token"})
			context.Abort()
			return
		}

		if is_special {
			var id int
			if payload , ok := token.Claims.(*UserClaims); !ok {
				context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Invalid Token"})
				context.Abort()
				return
			}else{
				id  = payload.ID
			}
			user ,err:= DBM.GetUser(id)
			if err != nil {
				context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Invalid Token"})
				context.Abort()
				return
			}
			
			if !user.Is_admin {
				context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Unauthorized"})
				context.Abort()
				return
			}

		}
		context.Next()
	}
}