package routes

import (
	"log"
	controller "github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/Delivery/controllers"
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/Repository"
	usecase "github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/UseCase"
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/infrastructure"
)

var DataBaseManager controller.DataBaseManager

//runs even before main
func init() {
	client , rootExists,root := infrastructure.Start()
	dataBase := client.Database("TaskManager")

	//initialize everything
	Repository := repository.NewRepository(client , dataBase)
	UseCase := usecase.NewUsecase(Repository)
	DataBaseManager = *controller.NewDatabaseManager(UseCase)

	if !rootExists {
		_,err := DataBaseManager.Usecase.CreateUser(root)
		if err != nil {
			log.Panic("Root not initialized!")
		}
	}
}

func Dependancy()*controller.DataBaseManager {
	return &DataBaseManager
}