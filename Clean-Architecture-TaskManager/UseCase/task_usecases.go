package usecase

import (
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
)

func (taskusecase *UseCaseData)GetTask(id string) (domain.Task, error) {
	doc , err := taskusecase.Repo.GetTaskDocumentById(id)
	if err != nil {
		return domain.Task{},err
	}
	return doc, nil
}

func (taskusecase *UseCaseData)GetTasks(filter map[string]string) ([]domain.Task, error) {
	tasks, err := taskusecase.Repo.GetTaskDocumentByFilter(filter)
	if err != nil {
		return []domain.Task{} , err
	}
	return tasks,nil
}

func (taskusecase *UseCaseData)CreateTask(model domain.Task) (domain.Task, error) {
	id , err := taskusecase.Repo.InsertTaskDocument(model)
	if err != nil {
		return domain.Task{} , err
	}
	new_task,err := taskusecase.GetTask(id)
	if err != nil {
		return domain.Task{} , err
	}
	return new_task,nil
}

func (taskusecase *UseCaseData)UpdateTask(id string , model domain.Task) (domain.Task, error) {
	err := taskusecase.Repo.UpdateTaskDocumentById(id , model)
	if err != nil {
		return domain.Task{},err
	}
	return model , nil
}

func (taskusecase *UseCaseData)DeleteTask(id string ) error {

	err := taskusecase.Repo.DeleteUserDocument(id)
	if err != nil {
		return err
	}
	return nil
}
