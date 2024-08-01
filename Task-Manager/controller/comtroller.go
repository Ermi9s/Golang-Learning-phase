package controller

import (
	"bytes"
	"io"
	"net/http"
	"strconv"

	model "github.com/ermi9s/taskmanager/models"
	service "github.com/ermi9s/taskmanager/services"
	"github.com/gin-gonic/gin"
)


func GetTasks(tasks *service.TaskManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		var tasklist []model.Task

		for key := range tasks.Tasks {
			tasklist =append(tasklist,*tasks.Tasks[key])
		}

		context.IndentedJSON(http.StatusOK , gin.H{"data" : tasklist})
	}
}

func GetTask(tasks *service.TaskManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		id,err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "id conversion failed"})
		}
		task,err := tasks.GetTask(id);
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task not found!"})
			return
		}
		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : task})
	}
}

func DeleteTask(tasks *service.TaskManager) func (contest *gin.Context) {
	return func(context *gin.Context) {
		id,err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "id conversion failed"})
			return
		}

		err = tasks.DeleteTask(id);
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task doesn't exist"})
			return
		}
		context.IndentedJSON(http.StatusAccepted , gin.H{"message" : "Deleted successfully"})
	}
}


func UpdateTask(tasks *service.TaskManager) func (context *gin.Context) {
	var task model.Task
	return func(context *gin.Context) {
		id,err := strconv.Atoi(context.Param("id"))
		if err != nil {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "id conversion failed"})
			return
		}
	
		if err := context.BindJSON(&task); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : err.Error()})
			return 
		}
		oid := id
		tasks.Tasks[strconv.Itoa(id)] = &task
		tasks.Tasks[strconv.Itoa(id)].ID = strconv.Itoa(oid)

		context.IndentedJSON(http.StatusAccepted , gin.H{"data" : *tasks.Tasks[strconv.Itoa(id)]})
	}
}


func CreateTask(tasks *service.TaskManager) func(context *gin.Context) {
	var task model.Task
	return func(context *gin.Context) {
		context.Request.ParseForm()

		bodyBytes,_ := io.ReadAll(context.Request.Body)
		context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		 
		if err := context.BindJSON(&task); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : err.Error()})
			return 
		}

		new_task := tasks.CreateTask(task)
		context.IndentedJSON(http.StatusCreated, gin.H{"data" : new_task})
		tasks.NextId++;
	}
}