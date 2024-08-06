package controller

import (
	"bytes"
	"io"
	"net/http"
	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/middleware"
	models "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/models"
	service "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func GetOneUser(manager *service.DataBaseManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		objID,err := primitive.ObjectIDFromHex(id)
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : "invalid Id" , "error" : err.Error()})
			return
		}

		user, err := manager.GetUser(objID)
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : "User not found" , "errror" : err})
			return 

		}
		token,err := middleware.Encode(user.ID , user.Email)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
			return
		}

		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : map[string]interface{}{"token" : token,"user" : user}})

	}
}

func GetUsers(manager *service.DataBaseManager) func(context *gin.Context) {

	return func(context *gin.Context) {
		users , err :=  manager.GetUsers()
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"data" : users})

	}
}

func CreateUser(manager *service.DataBaseManager) func(context *gin.Context) {
	var user models.User
	return func(context *gin.Context) {
		err := context.Request.ParseForm()
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return 
		}

		//these two are so that the context.bind can read the body multiple times 
		// as form data can only be read once normaly.
		byteBody,_ := io.ReadAll(context.Request.Body)
		context.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))

		if err := context.BindJSON(&user); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return 
		}

		new_user,err := manager.CreateUser(user)
		if err != nil {
			context.IndentedJSON(http.StatusOK , gin.H{"error" : err.Error()})
			return
		}

		token,err := middleware.Encode(new_user.ID , new_user.Email)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
			return
		}

		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : map[string]interface{}{"token" : token,"user" : user}})

	}
}

func UpdateUser(manager *service.DataBaseManager) func(conetext *gin.Context) {
	var user models.User
	return func(context *gin.Context) {
		id := context.Param("id")
		objID,_ := primitive.ObjectIDFromHex(id)
		err := context.Request.ParseForm()
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return 
		}
		byteBody,_:= io.ReadAll(context.Request.Body)
		context.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))

		if err := context.BindJSON(&user); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return 
		}

		new_user , err := manager.UpdateUser(objID , user)
		if err != nil {
			context.IndentedJSON(http.StatusOK , gin.H{"error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : new_user})
	}
}

func DeleteUser(manager *service.DataBaseManager)func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		objID,_ := primitive.ObjectIDFromHex(id)
		err := manager.DeleteUser(objID)
		if err != nil {
			context.IndentedJSON(http.StatusOK , gin.H{"error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusAccepted , gin.H{"message" : "deleted successfully"})
	}
}

func PromoteUser(manager *service.DataBaseManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		obj,_ := primitive.ObjectIDFromHex(id)

		user,err := manager.MakeAdmin(obj)
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}

		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : user})

	}
}