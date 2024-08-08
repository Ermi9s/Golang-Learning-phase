package routes

import (
	control "github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/Delivery/controllers"
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/infrastructure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(taskmanager *control.DataBaseManager) {
	route := gin.Default()
	
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	route.Use(cors.New(corsConfig))

	onlyAdminMiddleWare := infrastructure.Validate(true , true)
	onlyOwnerMiddleWare := infrastructure.Validate(true , false)
	publicMiddleWare := infrastructure.Validate(false , false)

	protected := route.Group("/protected")
	{
		protected.PUT("/promote/:id",onlyAdminMiddleWare , control.PromoteUser(taskmanager))
		protected.GET("/user" ,onlyAdminMiddleWare, control.GetUsers(taskmanager) )
	}

	protected_public := route.Group("/s-api") 
	{
		protected_public.GET("/task/:id",onlyOwnerMiddleWare ,control.GetOneTask(taskmanager))
		protected_public.GET("/task" ,onlyOwnerMiddleWare, control.GetTasks(taskmanager))
		protected_public.PUT("/task/:id",onlyOwnerMiddleWare , control.UpdateTask(taskmanager) )
		protected_public.DELETE("/task/:id",onlyOwnerMiddleWare , control.DeleteTask(taskmanager))
		protected_public.PUT("/user/:id",onlyOwnerMiddleWare ,control.UpdateUser(taskmanager) )
		protected_public.DELETE("/user/:id",onlyOwnerMiddleWare , control.DeleteUser(taskmanager))
		protected_public.GET("/user/:id",onlyOwnerMiddleWare , control.GetOneUser(taskmanager))
	}


	public := route.Group("/api")
	{
		public.POST("/task/" ,publicMiddleWare ,control.CreateTask(taskmanager))
		public.POST("/log-in/" , control.LogIN(taskmanager))
		public.POST("/register/" , control.CreateUser(taskmanager)) 
	}

	
	route.Run("localhost:8080")
}