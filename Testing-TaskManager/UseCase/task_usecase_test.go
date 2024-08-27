package usecase

import (
	"errors"
	"testing"

	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	mocks "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskUsecaseSuite struct {
	suite.Suite
	usecase domain.Task_Usecase_interface
	repository *mocks.Task_Repository_interface
}

func (suite *TaskUsecaseSuite)SetupTest() {
	repository := new(mocks.Task_Repository_interface)
	usecase := New_Task_Usecase(repository)

	suite.usecase = usecase
	suite.repository = repository
}


func (suite *TaskUsecaseSuite)TestGetTasksPositive(){
	task := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "Test Task",
		Description: "Testing task get usecase",
		Staus: "DONE",
		Date: primitive.DateTime(primitive.NilObjectID.Timestamp().Hour()),
		DueDate: primitive.DateTime(primitive.NewObjectID().Timestamp().Hour()),
	}
	task2 := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "Test Task",
		Description: "Testing task get usecase",
		Staus: "DONE",
		Date: primitive.DateTime(primitive.NilObjectID.Timestamp().Hour()),
		DueDate: primitive.DateTime(primitive.NewObjectID().Timestamp().Hour()),
	}

	data := []domain.Task{task , task2}
	filter := make(map[string]string)

	suite.repository.On("GetTaskDocumentByFilter" , filter).Return(data , nil)
	result,err := suite.usecase.GetTasks(filter)

	suite.NoError(err , "There shouldn't be an error fetching all tasks")
	suite.Equal(len(result) , len(data) , "The data returned from repository should be equal length to the one returned by the usecase")
	suite.Equal(result , data ,  "The data returned from repository should be the same as the one returned by the usecase")
}

func (suite *TaskUsecaseSuite)TestGetTasksNegative(){
	task := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "Test Task",
		Description: "Testing task get usecase",
		Staus: "DONE",
		Date: primitive.DateTime(primitive.NilObjectID.Timestamp().Hour()),
		DueDate: primitive.DateTime(primitive.NewObjectID().Timestamp().Hour()),
	}
	task2 := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "Test Task",
		Description: "Testing task get usecase",
		Staus: "DONE",
		Date: primitive.DateTime(primitive.NilObjectID.Timestamp().Hour()),
		DueDate: primitive.DateTime(primitive.NewObjectID().Timestamp().Hour()),
	}

	data := []domain.Task{task , task2}
	filter := make(map[string]string)

	r_error := errors.New("Failed to find by filter")

	suite.repository.On("GetTaskDocumentByFilter" , filter).Return([]domain.Task{} , r_error)
	result,err := suite.usecase.GetTasks(filter)

	suite.Error(err , "No error returned")
	suite.NotEqual(len(result) , len(data) , "empty should be returned")
}


func (suite *TaskUsecaseSuite)TestGetTaskPositive(){
	task := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "Test Task",
		Description: "Testing task get usecase",
		Staus: "DONE",
		Date: primitive.DateTime(primitive.NilObjectID.Timestamp().Hour()),
		DueDate: primitive.DateTime(primitive.NewObjectID().Timestamp().Hour()),
	}

	id := "1"
	suite.repository.On("GetTaskDocumentById" , id).Return(task , nil)
	result , err := suite.usecase.GetTask(id)

	suite.NoError(err , "There shouldn't be an error fetching a task")
	suite.Equal(result , task , "The task returned from repository should be the same as the usecase")
}

func (suite *TaskUsecaseSuite)TestGetTaskNegative(){

	id := "1"
	r_error := errors.New("Couldnt find document")
	suite.repository.On("GetTaskDocumentById" , id).Return(domain.Task{} , r_error)
	result , err := suite.usecase.GetTask(id)

	suite.Error(err , "There should be an error return")
	suite.Equal(result , domain.Task{} , "returned task should be empty")
}

func (suite *TaskUsecaseSuite)TestCreatTaskPositive() {
	task := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "Test Task",
		Description: "Testing task get usecase",
		Staus: "DONE",
		Date: primitive.DateTime(primitive.NilObjectID.Timestamp().Hour()),
		DueDate: primitive.DateTime(primitive.NewObjectID().Timestamp().Hour()),
	}
	id := primitive.NewObjectID().Hex()
	suite.repository.On("InsertTaskDocument", task).Return(id , nil)
	result,err := suite.usecase.CreateTask(task , "2")

	suite.NoError(err , "There shouldn't be an error creating tasks")
	suite.Equal(result , id , "The id returned from repository should be the same as the usecase")
}

func (suite *TaskUsecaseSuite)TestCreatTaskNegative() {
	task := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "Test Task",
		Description: "Testing task get usecase",
		Staus: "DONE",
		Date: primitive.DateTime(primitive.NilObjectID.Timestamp().Hour()),
		DueDate: primitive.DateTime(primitive.NewObjectID().Timestamp().Hour()),
	}

	r_error := errors.New("Internal server error")
	suite.repository.On("InsertTaskDocument", task).Return("" , r_error)
	result,err := suite.usecase.CreateTask(task , "2")

	suite.Error(err , "There should be an error creating tasks")
	suite.Equal(result , "" , "The id returned from repository should be empty")
}

func (suite *TaskUsecaseSuite)TestUpdateTaskPositive(){
	task := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "Test Task",
		Description: "Testing task get usecase",
		Staus: "DONE",
		Date: primitive.DateTime(primitive.NilObjectID.Timestamp().Hour()),
		DueDate: primitive.DateTime(primitive.NewObjectID().Timestamp().Hour()),
	}
	id := primitive.NewObjectID().Hex()

	suite.repository.On("UpdateTaskDocumentById" , id , task).Return(nil)
	result , err := suite.usecase.UpdateTask(id , task)
	suite.NoError(err , "There shouldn't be an error updating a task")
	suite.Equal(result , task , "The task returned from repository should be the same as the usecase")
}

func (suite *TaskUsecaseSuite)TestUpdateTaskNegative(){
	task := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "Test Task",
		Description: "Testing task get usecase",
		Staus: "DONE",
		Date: primitive.DateTime(primitive.NilObjectID.Timestamp().Hour()),
		DueDate: primitive.DateTime(primitive.NewObjectID().Timestamp().Hour()),
	}

	id := primitive.NewObjectID().Hex()
	r_error := errors.New("Internal server error")
	suite.repository.On("UpdateTaskDocumentById" , id , task).Return(r_error)
	result , err := suite.usecase.UpdateTask(id , task)
	suite.Error(err , "There should be an error updating a task")
	suite.Equal(result , domain.Task{} , "The task returned from repository should be the same as the usecase")
}

func (suite *TaskUsecaseSuite)TestDeleteTaskPositive() {
	id := primitive.NewObjectID().Hex()
	suite.repository.On("DeleteTaskDocument" , id).Return(nil)
	
	err := suite.usecase.DeleteTask(id)
	suite.NoError(err , "Error found on deleting task")
}

func (suite *TaskUsecaseSuite)TestDeleteTaskNegative() {
	id := primitive.NewObjectID().Hex()
	r_error := errors.New("Internal server error")
	suite.repository.On("DeleteTaskDocument" , id).Return(r_error)
	
	err := suite.usecase.DeleteTask(id)
	suite.Error(err , "Error not found on deleting task")
}

//negative testcases for user repository

func (suite *TaskUsecaseSuite)TestGetTaskDocumentByIdNegative(){
	id := primitive.NewObjectID()

	suite.repository.On("GetTaskDocumentById" , id.Hex()).Return(domain.Task{} , errors.New("Error decoding"))

	_, err := suite.usecase.GetTask(id.Hex())

	suite.NotNil(err , "No error found GetUserDocumentById")
}

func (suite *TaskUsecaseSuite)TestGetTaskDocumentByFilterNegative() {
	task := domain.Task{
		Title:     "Test Task",
	}
	filter := map[string]string{
		"title":      task.Title,
	}

	suite.repository.On("GetTaskDocumentByFilter" , filter).Return([]domain.Task{} , errors.New("Error decoding"))

	_, err := suite.usecase.GetTasks(filter)

	suite.NotNil(err , "No error found GetUserDocumentById")
}

func (suite *TaskUsecaseSuite)TestInsertTaskDocumentNegative() {
	task := domain.Task{
		ID:        primitive.NewObjectID(),
		Title:     "Test Task",
		Description:   "This is a test task",
	}

	suite.repository.On("InsertTaskDocument" , task).Return("" , errors.New("Error decoding"))

	_, err := suite.usecase.CreateTask(task , "2")

	suite.NotNil(err , "No error found GetUserDocumentById")
}

func (suite *TaskUsecaseSuite)TestUpdateTaskDocumentByIdNegative() {
	task := domain.Task{
		ID:        primitive.NewObjectID(),
		Title:     "Test Task",
		Description:   "This is a test task",
	}

	suite.repository.On("UpdateTaskDocumentById" , task.ID.Hex() , task).Return(errors.New("Error decoding"))

	_, err := suite.usecase.UpdateTask(task.ID.Hex() , task)

	suite.NotNil(err , "No error found GetUserDocumentById")
}

func (suite *TaskUsecaseSuite)TestDeleteTaskDocumentNegative() {
	id := primitive.NewObjectID().Hex()

	suite.repository.On("DeleteTaskDocument" , id).Return(errors.New("Error decoding"))

	err := suite.usecase.DeleteTask(id)

	suite.NotNil(err , "No error found GetUserDocumentById")
}

func TestTaskUsecase(t *testing.T) {
	suite.Run(t , new(TaskUsecaseSuite))
}