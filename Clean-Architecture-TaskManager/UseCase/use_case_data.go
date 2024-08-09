package usecase

import (
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
)
type UseCaseData struct {
	Repo domain.Repository_interface
}

func NewUsecase(repository domain.Repository_interface) domain.Usecase_interface {
	return &UseCaseData{
		Repo: repository,
	}
}