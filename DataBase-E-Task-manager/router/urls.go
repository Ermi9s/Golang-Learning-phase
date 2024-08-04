package router

import (
	control "github.com/Ermi9s/Golang-Learning-phase/DataBase-E-Task-manager/controller"
	service "github.com/Ermi9s/Golang-Learning-phase/DataBase-E-Task-manager/services"
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

	route.POST("/task/" , control.CreateTask(taskmanager))
	route.GET("/task/:id" , control.GetOneTask(taskmanager))
	route.GET("/task" , control.GetTasks(taskmanager))
	route.PUT("/task/:id" , control.UpdateTask(taskmanager))
	route.DELETE("/task/:id" , control.DeleteTask(taskmanager))

	route.Run("localhost:8080")
}