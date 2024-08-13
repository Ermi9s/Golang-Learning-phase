package usecase

import (
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task_usecase struct{
	Task_repo domain.Task_Repository_interface
}

func New_Task_Usecase(task_repo domain.Task_Repository_interface) domain.Task_Usecase_interface {
	return &Task_usecase{
		Task_repo: task_repo,
	}
}

func (taskusecase *Task_usecase)GetTask(id string) (domain.Task, error) {
	doc , err := taskusecase.Task_repo.GetTaskDocumentById(id)
	if err != nil {
		return domain.Task{},err
	}
	return doc, nil
}

func (taskusecase *Task_usecase)GetTasks(filter map[string]string) ([]domain.Task, error) {
	tasks, err := taskusecase.Task_repo.GetTaskDocumentByFilter(filter)
	if err != nil {
		return []domain.Task{} , err
	}
	return tasks,nil
}

func (taskusecase *Task_usecase)CreateTask(model domain.Task , user_id string) (string , error) {
	model.Creator,_= primitive.ObjectIDFromHex(user_id)
	id , err := taskusecase.Task_repo.InsertTaskDocument(model)
	if err != nil {
		return "", err
	}

	return id,nil
}

func (taskusecase *Task_usecase)UpdateTask(id string , model domain.Task) (domain.Task, error) {
	err := taskusecase.Task_repo.UpdateTaskDocumentById(id , model)
	if err != nil {
		return domain.Task{},err
	}
	return model , nil
}

func (taskusecase *Task_usecase)DeleteTask(id string ) error {
	err := taskusecase.Task_repo.DeleteTaskDocument(id)
	if err != nil {
		return err
	}
	return nil
}
