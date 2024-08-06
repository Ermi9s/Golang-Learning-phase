package controller

import (
	"bytes"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	serv "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/services"
	model "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/models"
)


func GetOneTask(DBM *serv.DataBaseManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")

		objID,err := primitive.ObjectIDFromHex(id)
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : "Invalid ID"})
			return
		}

		task,err := DBM.GetTask(objID);
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task not found!"})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"data" : task})
	}
}

func GetTasks(DBM *serv.DataBaseManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		tasks,err := DBM.GetTasks();
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task not found!"})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"data" : tasks})
	}
}


func DeleteTask(DBM *serv.DataBaseManager) func (contest *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")

		objID,err := primitive.ObjectIDFromHex(id)
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : "Invalid ID"})
			return
		}

		err = DBM.DeleteTask(objID);
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task doesn't exist"})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"message" : "Deleted successfully"})
	}
}


func UpdateTask(DBM *serv.DataBaseManager) func (context *gin.Context) {
	var task model.Task
	return func(context *gin.Context) {
		id := context.Param("id")
	
		if err := context.BindJSON(&task); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : err.Error()})
			return 
		}

		objID,err := primitive.ObjectIDFromHex(id)
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : "Invalid ID"})
			return
		}
		updated_task , err := DBM.UpdateTask(objID , task)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"message" : "Internal server error", "error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"data" : updated_task})
	}
}


func CreateTask(DBM *serv.DataBaseManager) func(context *gin.Context) {
	var task model.Task
	return func(context *gin.Context) {
		context.Request.ParseForm()

		bodyBytes,_ := io.ReadAll(context.Request.Body)
		context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		 
		if err := context.BindJSON(&task); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : err.Error()})
			return 
		}
		value,_ := context.Get("user_id")
		task.Creater = value.(primitive.ObjectID)
		new_task , err := DBM.CreateTask(task)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"message" : "Internal server error", "error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusOK, gin.H{"data" : new_task})
	}
}

