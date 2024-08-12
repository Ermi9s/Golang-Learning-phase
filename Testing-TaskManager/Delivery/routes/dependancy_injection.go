package routes

import (
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/infrastructure"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/repository"
)

var Repository *repository.Repository

//runs even before main
func init() {
	client := infrastructure.Start()
	dataBase := client.Database("TaskManager")
	
	//initialize main repository
	Repository = repository.NewRepository(client , dataBase)
}

func Dependancy() *repository.Repository {
	return Repository
}