package routes

import (
	controller "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/Delivery/controllers"
	"github.com/gin-gonic/gin"
)

func User_Routes_Run(route *gin.Engine , user_controller controller.User_Controller , middleWaer MiddleWares) {
	protected := route.Group("/protected")
	{
		protected.PUT("/promote/:id",middleWaer.OnlyAdminMiddleWare, user_controller.PromoteUser())
		protected.GET("/user" ,middleWaer.OnlyAdminMiddleWare, user_controller.GetUsers() )
	}

	protected_public := route.Group("/s-api") 
	{
		protected_public.PUT("/user/:id",middleWaer.OnlyOwnerMiddleWare ,user_controller.UpdateUser() )
		protected_public.DELETE("/user/:id",middleWaer.OnlyOwnerMiddleWare , user_controller.DeleteUser())
		protected_public.GET("/user/:id",middleWaer.OnlyOwnerMiddleWare , user_controller.GetOneUser())
	}

	public := route.Group("/api")
	{
		public.POST("/log-in/" , user_controller.LogIN())
		public.POST("/register/" , user_controller.CreateUser()) 
	}

}