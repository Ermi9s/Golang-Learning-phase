package usecase

import (
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
)

func (taskusecase *UseCaseData)GetTask(id string) (domain.Model, error) {
	doc , err := taskusecase.Repo.GetDocumentById("Tasks" , id)
	if err != nil {
		return &domain.Task{},err
	}
	return doc, nil
}

func (taskusecase *UseCaseData)GetTasks(filter map[string]string) ([]domain.Task, error) {
	decoded, err := taskusecase.Repo.GetDocumentByFilter("Tasks" , filter)
	if err != nil {
		return []domain.Task{} , err
	}
	result := []domain.Task{}
	for _,val := range decoded {
		new := val.(*domain.Task)
		result = append(result, *new)
	}
	return result,nil
}

func (taskusecase *UseCaseData)CreateTask(model domain.Model) (domain.Model, error) {
	err := taskusecase.Repo.InsertDocument("Tasks" , model)
	if err != nil {
		return &domain.Task{} , err
	}
	return model,nil
}

func (taskusecase *UseCaseData)UpdateTask(id string , model domain.Model) (domain.Model, error) {
	err := taskusecase.Repo.UpdateDocumentById("Tasks", id , model)
	if err != nil {
		return &domain.Task{},err
	}
	return model , nil
}

func (taskusecase *UseCaseData)DeleteTask(id string ) error {

	err := taskusecase.Repo.DeleteDocument("Tasks" , id)
	if err != nil {
		return err
	}
	return nil
}
