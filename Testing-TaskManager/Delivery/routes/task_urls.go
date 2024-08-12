package routes

import (
	controller "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/Delivery/controllers"
	"github.com/gin-gonic/gin"
)

func Task_Routes_Run(route *gin.Engine , task_controller controller.Task_Controller , middleWaer MiddleWares) {
	protected_public := route.Group("/s-api") 
	{
		protected_public.GET("/task/:id",middleWaer.OnlyOwnerMiddleWare ,task_controller.GetOneTask())
		protected_public.GET("/task" ,middleWaer.OnlyOwnerMiddleWare, task_controller.GetTasks())
		protected_public.PUT("/task/:id",middleWaer.OnlyOwnerMiddleWare , task_controller.UpdateTask())
		protected_public.DELETE("/task/:id",middleWaer.OnlyOwnerMiddleWare, task_controller.DeleteTask())
	}

	public := route.Group("/api")
	{
		public.POST("/task/" ,middleWaer.PublicMiddleWare ,task_controller.CreateTask())
	}
}