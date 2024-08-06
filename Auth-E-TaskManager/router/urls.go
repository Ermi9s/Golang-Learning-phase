package router

import (
	control "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/controller"
	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/middleware"
	service "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(taskmanager *service.DataBaseManager) {
	route := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	route.Use(cors.New(corsConfig))

	route.POST("/task/" ,middleware.Validate(false , taskmanager) , control.CreateTask(taskmanager))
	route.GET("/task/:id" , middleware.Validate(false , taskmanager) , control.GetOneTask(taskmanager))
	route.GET("/task" , middleware.Validate(true , taskmanager , &taskmanager.Filter) , control.GetTasks(taskmanager))
	route.PUT("/task/:id" , middleware.Validate(false , taskmanager) , control.UpdateTask(taskmanager) )
	route.DELETE("/task/:id" , middleware.Validate(false , taskmanager) , control.DeleteTask(taskmanager))

	route.POST("/log-in/" , control.LogIN(taskmanager))//no need for token protection
	route.GET("/user/:id" , middleware.Validate(true , taskmanager) , control.GetOneUser(taskmanager) )
	route.GET("/user" ,middleware.Validate(true , taskmanager) , control.GetUsers(taskmanager) )
	route.POST("/register/" , control.CreateUser(taskmanager)) // no token needed for sign-up
	route.PUT("/user/:id" , middleware.Validate(false , taskmanager) ,control.UpdateUser(taskmanager) )
	route.DELETE("/user/:id" ,middleware.Validate(false , taskmanager), control.DeleteUser(taskmanager))
	
	route.PUT("/promote/:id" , middleware.Validate(true , taskmanager) , control.PromoteUser(taskmanager))
	route.Run("localhost:8080")
}