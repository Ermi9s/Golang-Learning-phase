package controller

import (
	"bytes"
	"io"
	"net/http"

	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/infrastructure"
	"github.com/gin-gonic/gin"
)


func GetOneUser(DBM *DataBaseManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		user, err := DBM.Usecase.GetUser(id)
	
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : "User not found" , "errror" : err})
			return 
		}
	
		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : user})

	}
}

func GetUsers(DBM *DataBaseManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		users , err :=  DBM.Usecase.GetUsers()
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"data" : users})

	}
}

func CreateUser(DBM *DataBaseManager) func(context *gin.Context) {
	var user domain.User
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

		new_user,err := DBM.Usecase.CreateUser(user)
		if err != nil {
			context.IndentedJSON(http.StatusOK , gin.H{"error" : err.Error()})
			return
		}
		
		token,err :=  infrastructure.Encode(new_user.ID , new_user.Email , false)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"token-error" : err.Error()})
			return
		}

		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : map[string]interface{}{"token" : token,"user" : new_user}})

	}
}

func UpdateUser(DBM *DataBaseManager) func(conetext *gin.Context) {
	var user domain.User
	return func(context *gin.Context) {
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)

		id := context.Param("id")
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
		if id != payload.ID.Hex() {
			context.IndentedJSON(http.StatusNotAcceptable , gin.H{"message" : "can not update other users accounts"})
			return
		}

		new_user , err := DBM.Usecase.UpdateUser(id , user)
		if err != nil {
			context.IndentedJSON(http.StatusOK , gin.H{"error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : new_user})
	}
}

func LogIN(DBM *DataBaseManager)func(context *gin.Context){
	var loginForm domain.AuthUser
	return func(context *gin.Context) {
		
		err := context.Request.ParseForm()
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return 
		}

		byteBody,_ := io.ReadAll(context.Request.Body)
		context.Request.Body = io.NopCloser(bytes.NewBuffer(byteBody))

		if err := context.BindJSON(&loginForm); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return 
		}

		user,err := DBM.Usecase.LogIn(loginForm)
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"error" : err.Error()})
			return 
		}
		token,err :=infrastructure.Encode(user.ID , user.Email , user.Is_admin)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : map[string]interface{}{"token" : token,"user" : user}})

	}
}

func DeleteUser(DBM *DataBaseManager)func(context *gin.Context) {
	return func(context *gin.Context) {
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)

		id := context.Param("id")

		if id != payload.Id && !payload.Is_admin {
			context.IndentedJSON(http.StatusNotAcceptable , gin.H{"message" : "can not delete other users accounts"})
			return
		}

		err := DBM.Usecase.DeleteUser(id)
		if err != nil {
			context.IndentedJSON(http.StatusOK , gin.H{"error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusAccepted , gin.H{"message" : "deleted successfully"})
	}
}

func PromoteUser(DBM *DataBaseManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		user,err := DBM.Usecase.Promote(id)
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}

		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : user})
	}
}