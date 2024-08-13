package domain

import "github.com/gin-gonic/gin"

//main repository interface 
type Repository_interface interface{}

type Task_Repository_interface interface {
	GetTaskDocumentById(id string) (Task, error)
	GetTaskDocumentByFilter(filter map[string]string) ([]Task, error)
	UpdateTaskDocumentById(id string, update Task) error
	InsertTaskDocument(object Task) (string, error)
	DeleteTaskDocument(id string) error
}

type User_Repository_interface interface {
	GetUserDocumentById(id string) (User, error)
	GetUserDocumentByFilter(filter map[string]string) ([]User, error)
	UpdateUserDocumentById(id string, update User) error
	InsertUserDocument(object User) (string, error)
	DeleteUserDocument(id string) error
}

type Task_Usecase_interface interface {
	GetTask(id string) (Task, error)
	GetTasks(filter map[string]string) ([]Task, error)
	CreateTask(model Task , user_id string) (string, error)
	UpdateTask(id string, model Task) (Task, error)
	DeleteTask(id string) error
}

type User_Usecase_interface interface {
	GetUser(id string) (User, error)
	GetUsers() ([]User, error)
	CreateUser(model User) (AuthUser , string, error)
	UpdateUser(id string, model User) (User, error)
	DeleteUser(id string) error
	LogIn(model AuthUser) (User, error)
	Promote(id string) (User, error)
}

type Task_Controler_interface interface {
	GetOneTask() func(context *gin.Context)
	GetTasks() func(context *gin.Context)
	DeleteTask() func(contest *gin.Context)
	UpdateTask() func(context *gin.Context)
	CreateTask() func(context *gin.Context)
}

type User_Controller_interface interface {
	GetOneTask() func(context *gin.Context)
	GetTasks() func(context *gin.Context)
	DeleteTask() func (contest *gin.Context)
	UpdateTask() func (context *gin.Context)
	CreateTask() func(context *gin.Context)
}