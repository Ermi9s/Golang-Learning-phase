package usecase

import(
	repository "github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/Repository"
)
type UseCaseData struct {
	Repo *repository.Repository
}