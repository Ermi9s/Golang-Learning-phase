package controller

import (
	"bytes"
	"io"
	"net/http"

	domain "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"github.com/gin-gonic/gin"
)

type Task_Controller struct {
	Task_Usecase domain.Task_Usecase_interface
}

func New_Task_Controller(taskusecase domain.Task_Usecase_interface) *Task_Controller {
	return &Task_Controller{
		Task_Usecase: taskusecase,
	}
}

func (DBM *Task_Controller)GetOneTask() func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")

		task,err := DBM.Task_Usecase.GetTask(id);
	
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)

		if task.Creator.Hex() != payload.ID && !payload.Is_admin {
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

func (DBM *Task_Controller) GetTasks() func(context *gin.Context) {
	return func(context *gin.Context) {
		filter := make(map[string]string)
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)
		filter["creator_id"] = payload.ID
		if payload.Is_admin {
			delete(filter , "creator_id")
		}
		// log.Println(filter["creator_id"] , "ke",payload.ID)
		tasks,err := DBM.Task_Usecase.GetTasks(filter);
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task not found!"})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"data" : tasks})
	}
}


func (DBM *Task_Controller)DeleteTask() func (contest *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		task,err := DBM.Task_Usecase.GetTask(id)
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

		err = DBM.Task_Usecase.DeleteTask(id)
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "error deleting task"})
			return
		}
		context.IndentedJSON(http.StatusOK , gin.H{"message" : "Deleted successfully"})
	}
}


func (DBM *Task_Controller)UpdateTask() func (context *gin.Context) {
	var task domain.Task
	return func(context *gin.Context) {
		id := context.Param("id")
		ipayload,_ := context.Get("payload")
		payload := ipayload.(*domain.UserClaims)

		if err := context.BindJSON(&task); err != nil {
			context.IndentedJSON(http.StatusBadRequest , gin.H{"message" : err.Error()})
			return 
		}
		itask,err := DBM.Task_Usecase.GetTask(id)
		if err != nil {
			context.IndentedJSON(http.StatusNotFound , gin.H{"message" : "task not found!"})
			return
		}

		if itask.Creator.Hex() != payload.ID &&  !payload.Is_admin{
			context.IndentedJSON(http.StatusMethodNotAllowed , gin.H{"message" : "task belongs to other user"})
			return
		}

		updated_task , err := DBM.Task_Usecase.UpdateTask(id, task)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"message" : "Internal server error", "error" : err.Error()})
			return
		}
		updated_task.ID = itask.ID
		updated_task.Creator = itask.Creator
		context.IndentedJSON(http.StatusOK , gin.H{"data" : updated_task})
	}
}


func (DBM *Task_Controller)CreateTask() func(context *gin.Context) {
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
		new_task , err := DBM.Task_Usecase.CreateTask(task , payload.ID)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError , gin.H{"message" : "Internal server error", "error" : err.Error()})
			return
		}
		context.IndentedJSON(http.StatusOK, gin.H{"ID" : new_task})
	}
}

