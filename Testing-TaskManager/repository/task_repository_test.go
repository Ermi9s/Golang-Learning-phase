package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/database/databaseDomain/mocks"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository struct {
	suite.Suite
	main_repo *Repository

	task_repo domain.Task_Repository_interface

	single_result *mocks.SingleResult
	collection    *mocks.Collection
	database      *mocks.Database
	cursor        *mocks.Cursor
}

func (suite *TaskRepository) SetupTest() {
	client := new(mocks.Client)

	suite.collection = new(mocks.Collection)
	suite.database = new(mocks.Database)
	suite.single_result = new(mocks.SingleResult)

	suite.cursor = new(mocks.Cursor)
	suite.main_repo = NewRepository(client, suite.database)
	suite.task_repo = New_Task_Repository(*suite.main_repo, suite.collection)
}

func(suite *UserRepository)TestGetTaskDocumentById(){
	task := domain.Task{
		ID: primitive.NewObjectID(),
	}
	filter := bson.D{{Key : "_id" , Value: task.ID}}

	suite.collection.On("FindOne" , context.TODO() ,filter).Return(suite.single_result)
	suite.single_result.On("Decode" , mock.Anything).Return(nil)

	id := task.ID.Hex()
	_, err := suite.user_repo.GetUserDocumentById(id)

	suite.Nil(err , "Found error GetUserDocumentById")
}

func (suite *TaskRepository) TestGetTaskDocumentByFilter() {
	task := domain.Task{
		ID:        primitive.NewObjectID(),
		Creator: primitive.NewObjectID(),
		Title:     "Test Task",
		Description:   "This is a test task",
	}
	filter := map[string]string{
		"creator_id": task.Creator.Hex(),
		"title":      task.Title,
	}

	dbfilter := bson.D{
		{Key: "creator_id", Value: task.Creator},
		{Key: "title", Value: task.Title},
	}

	suite.cursor.On("Next", context.TODO()).Return(true).Once()
	suite.cursor.On("Next", context.TODO()).Return(false).Once()
	suite.cursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.Task)
		*arg = task
	}).Return(nil)

	suite.collection.On("Find", context.TODO(), dbfilter).Return(suite.cursor, nil)

	// Act
	result, err := suite.task_repo.GetTaskDocumentByFilter(filter)

	suite.Nil(err, "Expected no error")
	suite.NotNil(result, "Expected non-nil result")
	suite.Equal(1, len(result), "Expected one task in the result")
	suite.Equal(task, result[0], "Expected the returned task to match the mock task")
}

func (suite *TaskRepository)TestInsertTaskDocument() {
	task := domain.Task{
		ID:        primitive.NewObjectID(),
		Creator: primitive.NewObjectID(),
		Title:     "Test Task",
		Description:   "This is a test task",
	}

	var doc bson.D
	byteModel,_ := bson.Marshal(task)
	bson.Unmarshal(byteModel , &doc)

	suite.collection.On("InsertOne" , context.TODO() , doc).Return(task.ID, nil)

	id , err := suite.task_repo.InsertTaskDocument(task)
	suite.Nil(err , "error found in Insert")
	suite.Equal(id , task.ID.Hex() , "Id not the same")
}

func (suite *TaskRepository)TestDeleteUserDocument() {
	id := primitive.NewObjectID()
	filter := bson.D{{Key : "_id" , Value:id}}

	var r int64
	suite.collection.On("DeleteOne" , context.TODO() , filter).Return(r, nil)

	err := suite.task_repo.DeleteTaskDocument(id.Hex())

	suite.Nil(err , "Error found on delete")
}


//negative test cases
func (suite *TaskRepository)TestGetTaskDocumentByIdNegative(){
	id := primitive.NewObjectID()
	filter := bson.D{{Key : "_id" , Value:id}}

	suite.collection.On("FindOne" , context.TODO() ,filter).Return(suite.single_result)
	suite.single_result.On("Decode" , mock.Anything).Return(errors.New("Error decoding"))

	_, err := suite.task_repo.GetTaskDocumentById(id.Hex())

	suite.NotNil(err , "No error found GetUserDocumentById")
}

func (suite *TaskRepository) TestGetTaskDocumentByFilterNegative() {
	task := domain.Task{
		ID:        primitive.NewObjectID(),
		Creator: primitive.NewObjectID(),
		Title:     "Test Task",
		Description:   "This is a test task",
	}
	filter := map[string]string{
		"creator_id": task.Creator.Hex(),
		"title":      task.Title,
	}

	dbfilter := bson.D{
		{Key: "creator_id", Value: task.Creator},
		{Key: "title", Value: task.Title},
	}

	suite.cursor.On("Next", context.TODO()).Return(true).Once()
	suite.cursor.On("Next", context.TODO()).Return(false).Once()
	suite.cursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.Task)
		*arg = task
	}).Return(errors.New("Error decoding"))

	suite.collection.On("Find", context.TODO(), dbfilter).Return(suite.cursor, nil)

	// Act
	result, err := suite.task_repo.GetTaskDocumentByFilter(filter)

	suite.NotNil(err, "Expected error")
	suite.Nil(result, "Expected nil result")
}

func (suite *TaskRepository)TestInsertTaskDocumentNegative() {
	task := domain.Task{
		ID:        primitive.NewObjectID(),
		Creator: primitive.NewObjectID(),
		Title:     "Test Task",
		Description:   "This is a test task",
	}

	var doc bson.D
	byteModel,_ := bson.Marshal(task)
	bson.Unmarshal(byteModel , &doc)

	suite.collection.On("InsertOne" , context.TODO() , doc).Return(primitive.NilObjectID, errors.New("Error inserting"))

	_, err := suite.task_repo.InsertTaskDocument(task)
	suite.NotNil(err , "error not found in Insert")
}

func (suite *TaskRepository)TestDeleteUserDocumentNegative() {
	id := primitive.NewObjectID()
	filter := bson.D{{Key : "_id" , Value:id}}

	var r int64
	suite.collection.On("DeleteOne" , context.TODO() , filter).Return(r, errors.New("Error deleting"))

	err := suite.task_repo.DeleteTaskDocument(id.Hex())

	suite.NotNil(err , "Error not found on delete")
}



func TestTaskRepository(t *testing.T) {
	suite.Run(t , new(TaskRepository))
}
