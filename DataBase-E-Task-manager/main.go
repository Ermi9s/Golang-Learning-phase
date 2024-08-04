package main

import (
	"log"

	"github.com/Ermi9s/Golang-Learning-phase/DataBase-E-Task-manager/connection"
	router "github.com/Ermi9s/Golang-Learning-phase/DataBase-E-Task-manager/router"
	service "github.com/Ermi9s/Golang-Learning-phase/DataBase-E-Task-manager/services"
)


var _DataBaseMmanager = service.DataBaseManager{}

func main() {
	remote := connection.ServerConnection{}
	remote.Connect_could()

	// local := connection.LServerCon{}
	// local.Connect_local()

	//local is local instance of the server
	//remote is remote instance of the server
	_DataBaseMmanager.Client = remote.Client
	_DataBaseMmanager.Tasks = remote.Client.Database("TaskManager").Collection("Tasks")
	router.Run(&_DataBaseMmanager)
	log.Println("Server is running on localhost:8080")
}

