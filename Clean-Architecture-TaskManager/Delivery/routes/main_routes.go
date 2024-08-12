package routes

import (
	controller "github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/Delivery/controllers"
	usecase "github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/UseCase"
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/infrastructure"
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/repository"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type MiddleWares struct {
	OnlyAdminMiddleWare func(c *gin.Context)
	OnlyOwnerMiddleWare func(c *gin.Context)
	PublicMiddleWare func(c *gin.Context)
}

func Run() {
	route := gin.Default()
	
	//cross origin resource sharing middleware
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	route.Use(cors.New(corsConfig))

	//custom middleware setup
	var middlewares MiddleWares
	middlewares.OnlyAdminMiddleWare = infrastructure.Validate(true , true)
	middlewares.OnlyOwnerMiddleWare = infrastructure.Validate(true , false)
	middlewares.PublicMiddleWare = infrastructure.Validate(false , false)

	_repository := Dependancy()
	//initialize repository for each route
	_task_repository := repository.New_Task_Repository(*_repository)
	_user_repository := repository.New_User_Repository(*_repository)

	//initialize usecase for each route
	_task_usecase := usecase.New_Task_Usecase(_task_repository)
	_user_usecase := usecase.New_User_Usecase(_user_repository)

	//initialize controllers
	_task_controller := controller.New_Task_Controller(_task_usecase)
	_user_controller := controller.New_User_Controller(_user_usecase)

	//run routes for each
	User_Routes_Run(route , *_user_controller , middlewares)
	Task_Routes_Run(route , *_task_controller , middlewares)
	
	
	route.Run("localhost:8080")
}