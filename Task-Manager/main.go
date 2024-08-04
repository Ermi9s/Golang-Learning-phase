package main

import (
	"fmt"
	"github.com/Ermi9s/Golang-Learning-phase/tree/main/Task-Manager/models"
	"github.com/Ermi9s/Golang-Learning-phase/tree/main/Task-Manager/router"
	"github.com/Ermi9s/Golang-Learning-phase/tree/main/Task-Manager/services"
)



func main() {	
	var manager services.TaskManager = services.TaskManager{
		Tasks: make(map[string]*models.Task),
		NextId: 1,
	}
	fmt.Println("Server running on http://127.0.0.1:8080/")
	router.Run(&manager)
}