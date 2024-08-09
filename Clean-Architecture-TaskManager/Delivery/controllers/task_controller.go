package controller

import (
	"bytes"
	"io"
	"net/http"

	domain "github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
	"github.com/gin-gonic/gin"
)

func GetOneTask(DBM *DataBaseManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")

		task,err := DBM.Usecase.GetTask(id);
	
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)

		if task.Creator.Hex() != payload.ID.Hex() && !payload.Is_admin {
			context.IndentedJSON(http.StatusMethodNotAllowed , gin.H{"message" : "task belongs to other user"})
			return
		}
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task not found!"})
			return
		}

		context.IndentedJSON(http.StatusOK , gin.H{"data" : task})
	}
}

func GetTasks(DBM *DataBaseManager) func(context *gin.Context) {
	return func(context *gin.Context) {
		filter := make(map[string]string)
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)
		filter["creator_id"] = payload.ID.Hex()
		if payload.Is_admin {
			delete(filter , "creator_id")
		}
		// log.Println(filter["creator_id"] , "ke",payload.ID)
		tasks,err := DBM.Usecase.GetTasks(filter);
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task not found!"})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"data" : tasks})
	}
}


func DeleteTask(DBM *DataBaseManager) func (contest *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		task,err := DBM.Usecase.GetTask(id)
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task doesn't exist"})
			return
		}
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)

		if task.Creator.Hex() != id &&  !payload.Is_admin{
			context.IndentedJSON(http.StatusMethodNotAllowed , gin.H{"message" : "task belongs to other user"})
			return
		}

		err = DBM.Usecase.DeleteTask(id)
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "error deleting task"})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"message" : "Deleted successfully"})
	}
}


func UpdateTask(DBM *DataBaseManager) func (context *gin.Context) {
	var task domain.Task
	return func(context *gin.Context) {
		id := context.Param("id")
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)

		if err := context.BindJSON(&task); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : err.Error()})
			return 
		}
		itask,err := DBM.Usecase.GetTask(id)
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task not found!"})
			return
		}

		if itask.Creator != payload.ID &&  !payload.Is_admin{
			context.IndentedJSON(http.StatusMethodNotAllowed , gin.H{"message" : "task belongs to other user"})
			return
		}

		updated_task , err := DBM.Usecase.UpdateTask(id, task)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"message" : "Internal server error", "error" : err.Error()})
			return
		}
		updated_task.ID = itask.ID
		updated_task.Creator = itask.Creator
		context.IndentedJSON(http.StatusOK , gin.H{"data" : updated_task})
	}
}


func CreateTask(DBM *DataBaseManager) func(context *gin.Context) {
	var task domain.Task
	return func(context *gin.Context) {
		context.Request.ParseForm()

		bodyBytes,_ := io.ReadAll(context.Request.Body)
		context.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		 
		if err := context.BindJSON(&task); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : err.Error()})
			return 
		}
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)
		task.Creator = payload.ID
		new_task , err := DBM.Usecase.CreateTask(task)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"message" : "Internal server error", "error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusOK, gin.H{"data" : new_task})
	}
}

