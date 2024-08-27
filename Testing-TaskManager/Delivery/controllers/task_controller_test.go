package controller

import (
	"net/http/httptest"
	"testing"

	domain "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerSuite struct {
	suite.Suite
	controller *Task_Controller
	usecase    *mocks.Task_Usecase_interface
	service    *mocks.Services
}

func (suite *TaskControllerSuite) SetupTest() {
	usecase := new(mocks.Task_Usecase_interface)
	controller := New_Task_Controller(usecase)
	suite.controller = controller
	suite.usecase = usecase
	suite.service = new(mocks.Services)
}


func (suite *TaskControllerSuite)TestGetOneTask() {
    user := domain.User{
		ID: primitive.NewObjectID(),
        Email:    "test@example.com",
    }

    payload := &domain.UserClaims{
        ID:    user.ID.Hex(),
        Email: user.Email,
        Is_admin: true,
    }

	task := domain.Task{
		ID: primitive.NewObjectID(),
		Title: "test",
		Creator: user.ID,
	}
	
	w := httptest.NewRecorder()
	context,_ := gin.CreateTestContext(w)
	context.Params = gin.Params{gin.Param{Key: "id" , Value: task.ID.Hex()}}
	context.Set("payload" , payload)

	suite.usecase.On("GetTask" , task.ID.Hex()).Return(task, nil)

	handler := suite.controller.GetOneTask()
	handler(context)

	suite.Equal(200 , w.Code)
}

func TestTaskController(t *testing.T) {
	suite.Run(t , new(TaskControllerSuite))
}