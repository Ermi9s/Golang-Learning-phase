package controller

import (
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
)

type DataBaseManager struct {
	Usecase domain.Usecase_interface
}

func NewDatabaseManager(usecase domain.Usecase_interface) *DataBaseManager {
	return &DataBaseManager{
		Usecase: usecase,
	}
}