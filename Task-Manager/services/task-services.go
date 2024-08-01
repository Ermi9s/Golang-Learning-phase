package services

import (
	"errors"
	"strconv"
	"time"
	model "github.com/ermi9s/taskmanager/models"
)

type TaskManagement interface {
	CreateTask(task *model.Task) model.Task
	UpdateTask(task *model.Task) (model.Task , error)
	DeleteTask(id int) error
	GetTask(id int) (model.Task , error)
}

type TaskManager struct {
	Tasks map[string]*model.Task
	NextId int
}

func (taskm *TaskManager) CreateTask(task model.Task) model.Task {
		currtime := time.Now()
		new_task := task
		new_task.Date = currtime.Format("2006/01/02")
		new_task.ID = strconv.Itoa(taskm.NextId) 
		taskm.Tasks[strconv.Itoa(taskm.NextId)] = &new_task

		return new_task
}

func (taskm *TaskManager) DeleteTask(id int) error {
	if _,ok := taskm.Tasks[strconv.Itoa(id)]; ok {
		delete(taskm.Tasks , strconv.Itoa(id))
		return nil
	}

	return errors.New("not found")
}
func (taskm *TaskManager) GetTask(id int) (model.Task , error) {
	if _,ok := taskm.Tasks[strconv.Itoa(id)]; ok {
		return *taskm.Tasks[strconv.Itoa(id)] , nil
	}
	return model.Task{}, errors.New("not found")	
}


