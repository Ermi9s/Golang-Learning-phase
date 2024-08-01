package main

import (
	"fmt"
	"github.com/ermi9s/taskmanager/models"
	"github.com/ermi9s/taskmanager/router"
	"github.com/ermi9s/taskmanager/services"
)



func main() {	
	var manager services.TaskManager = services.TaskManager{
		Tasks: make(map[string]*models.Task),
		NextId: 1,
	}
	fmt.Println("Server running on http://127.0.0.1:8080/")
	router.Run(&manager)
}