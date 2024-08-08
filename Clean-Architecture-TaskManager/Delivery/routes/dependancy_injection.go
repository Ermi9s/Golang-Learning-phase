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

	var Repository repository.Repository = repository.Repository{
		Client:   client,
		Database: dataBase,
	}

	var UseCase usecase.UseCaseData = usecase.UseCaseData{
		Repo:&Repository,
	}

	//initialize everything
	DataBaseManager = controller.DataBaseManager{
		Usecase: &UseCase,
	}

	if !rootExists {
		_,err := DataBaseManager.Usecase.CreateUser(&root)
		if err != nil {
			log.Panic("Root not initialized!")
		}
	}
}

func Dependancy()*controller.DataBaseManager {
	return &DataBaseManager
}