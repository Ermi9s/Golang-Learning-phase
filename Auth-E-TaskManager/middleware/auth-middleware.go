package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/models"
	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func Validate(is_special bool , DBM *services.DataBaseManager , filter ... *bson.D) func(context *gin.Context) {
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
		
		token , err := jwt.ParseWithClaims(authSplitted[1] ,&models.UserClaims{} ,func(t *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})
		
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Invalid Token a"})
			context.Abort()
			return
		}


		var id primitive.ObjectID
		payload,ok := token.Claims.(*models.UserClaims); 

		if !ok {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"Message" : "Invalid Token Claims"})
			context.Abort()
			return
		}

		id  = payload.ID
		context.Set("user_id" , id)
		context.Set("is_admin" , false)
		if is_special {
			user,_:= DBM.GetUser(id)
			if !user.Is_admin {
				if len(filter) != 0{
					*(filter)[0] = bson.D{{Key: "_id" , Value: id}}
				}else{
					context.IndentedJSON(http.StatusUnauthorized ,gin.H{"error" : "not authorized"})
					context.Abort()
					return 
				}	
			}else{
				if len(filter) > 0{
					*(filter)[0] = nil
				}
				context.Set("is_admin" , true)
			}
		}
		
		context.Next()
	}
}