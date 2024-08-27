package controller

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	domain "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserControllerSuite struct {
	suite.Suite
	controller *User_Controller
	usecase *mocks.User_Usecase_interface
	service *mocks.Services
}

func (suite *UserControllerSuite) SetupTest() {
	usecase := new(mocks.User_Usecase_interface)
	controller := New_User_Controller(usecase)
	suite.controller = controller
	suite.usecase = usecase
	suite.service = new(mocks.Services)
}

func (suite *UserControllerSuite) TestGetOneUserPositive() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		UserName: "TestUser",
		Email:    "test@example.com",
		Password: "password",
		Is_admin: false,
	}
	id := user.ID.Hex()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Params = gin.Params{gin.Param{Key: "id", Value: id}}

	suite.usecase.On("GetUser", id).Return(user, nil)

	handler := suite.controller.GetOneUser()
	handler(context)

	suite.Equal(http.StatusAccepted, context.Writer.Status())
}

func (suite *UserControllerSuite) TestGetOneUserNegative() {
	id := primitive.NewObjectID().Hex()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Params = gin.Params{gin.Param{Key: "id", Value: id}}

	r_error := errors.New("Internal server error")
	suite.usecase.On("GetUser", id).Return(domain.User{}, r_error)

	handler := suite.controller.GetOneUser()
	handler(context)

	suite.Equal(400, context.Writer.Status())
}

func (suite *UserControllerSuite) TestCreateUserPositive() {
	user := domain.AuthUser{
		UserName: "TestUser",
		Email:    "test@example.com",
		Password: "password",
		Is_admin: false,
	}

	suser := domain.User {
		UserName: "TestUser",
		Email:    "test@example.com",
		Password: "password",
		Is_admin: false,
	}

	expectedToken := "token"
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"userName": "TestUser",
		"email": "test@example.com",
		"password": "password",
		"is_admin": false
	}`)

	context.Request = httptest.NewRequest(http.MethodPost, "/user", body)
	context.Request.Header.Set("Content-Type", "application/json")

	suite.usecase.On("CreateUser", suser).Return(user,expectedToken, nil)

	handler := suite.controller.CreateUser()
	handler(context)

	suite.Equal(http.StatusAccepted, w.Code)
}


func (suite *UserControllerSuite) TestCreateUserNegative() {
	suser := domain.User {
		UserName: "TestUser",
		Email:    "test@example.com",
		Password: "password",
		Is_admin: false,
	}
	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"userName": "TestUser",
		"email": "test@example.com",
		"password": "password",
		"is_admin": false
	}`)
	context.Request = httptest.NewRequest(http.MethodPost, "/user", body)
	context.Request.Header.Set("Content-Type", "application/json")

	expectedError := errors.New("error creating user")
	suite.usecase.On("CreateUser", suser).Return(domain.AuthUser{}, "", expectedError)

	handler := suite.controller.CreateUser()
	handler(context)

	suite.Equal(200, w.Code)
}

func (suite *UserControllerSuite)TestUpdateUserPositive() {
    user := domain.User{
		ID: primitive.NewObjectID(),
        Email:    "test@example.com",
    }

    payload := &domain.UserClaims{
        ID:    user.ID.Hex(),
        Email: user.Email,
        Is_admin: true,
    }
	
	w := httptest.NewRecorder()
	context,_ := gin.CreateTestContext(w)
	body := bytes.NewBufferString(`{
		"userName": "TestUser",
		"email": "test@example.com",
		"password": "password",
		"is_admin": false
	}`)
	context.Request = httptest.NewRequest(http.MethodPut, "/user/:id", body)
	context.Params = gin.Params{gin.Param{Key: "id" , Value: user.ID.Hex()}}
	context.Set("payload" , payload)

	suite.usecase.On("UpdateUser" ,user.ID.Hex() , mock.Anything).Return(user , nil)

	handler := suite.controller.UpdateUser()
	handler(context)

	suite.Equal(202 , w.Code)
}

func (suite *UserControllerSuite)TestUpdateUserNegative() {
    user := domain.User{
		ID: primitive.NewObjectID(),
        Email:    "test@example.com",
    }

    payload := &domain.UserClaims{
        ID:    user.ID.Hex(),
        Email: user.Email,
        Is_admin: true,
    }
	
	w := httptest.NewRecorder()
	context,_ := gin.CreateTestContext(w)
	body := bytes.NewBufferString(`{
		"userName": "TestUser",
		"email": "test@example.com",
		"password": "password",
		"is_admin": false
	}`)
	context.Request = httptest.NewRequest(http.MethodPut, "/user/:id", body)
	context.Params = gin.Params{gin.Param{Key: "id" , Value: user.ID.Hex()}}
	context.Set("payload" , payload)

	r_error := errors.New("Internal server error")
	suite.usecase.On("UpdateUser" ,user.ID.Hex() , mock.Anything).Return(domain.User{} , r_error)

	handler := suite.controller.UpdateUser()
	handler(context)

	suite.Equal(400 , w.Code)
}

func (suite *UserControllerSuite) TestLogInPositive() {
	user := domain.User{
		UserName: "TestUser",
		Email:    "test@example.com",
		Password: "password",
		Is_admin: false,
		ID:       primitive.NewObjectID(),
	}

	luser := domain.AuthUser{
		UserName: "TestUser",
		Email:    "test@example.com",
		Password: "password",
		Is_admin: false,
	}

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"UserName": "TestUser",
		"email": "test@example.com",
		"password": "password"
	}`)
	context.Request = httptest.NewRequest(http.MethodPost, "/log-in", body)
	context.Request.Header.Set("Content-Type", "application/json")

	suite.usecase.On("LogIn", luser).Return(user, nil)

	suite.service.On("Encode", user.ID.Hex(), user.Email, user.Is_admin).Return("mocked-token", nil)

	handler := suite.controller.LogIN()
	handler(context)

	suite.Equal(http.StatusAccepted, w.Code)
}

func (suite *UserControllerSuite) TestLogInNegative() {
	user := domain.User{
		Email:    "test@example.com",
		Is_admin: false,
		ID:       primitive.NewObjectID(),
	}

	luser := domain.AuthUser{
		UserName: "TestUser",
		Email:    "test@example.com",
		Password: "password",
		Is_admin: false,
	}

	w := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"UserName": "TestUser",
		"email": "test@example.com",
		"password": "password"
	}`)
	context.Request = httptest.NewRequest(http.MethodPost, "/log-in", body)
	context.Request.Header.Set("Content-Type", "application/json")

	r_error := errors.New("Internal server error")
	suite.usecase.On("LogIn", luser).Return(domain.User{}, r_error)

	suite.service.On("Encode", user.ID.Hex(), user.Email, user.Is_admin).Return("mocked-token", nil)

	handler := suite.controller.LogIN()
	handler(context)

	suite.Equal(404, w.Code)
}

func (suite *UserControllerSuite)TestDeleteUserPositive() {
    user := domain.User{
        Email:    "test@example.com",
    }


    payload := &domain.UserClaims{
        ID:    user.ID.Hex(),
        Email: user.Email,
        Is_admin: true,
    }
	w := httptest.NewRecorder()
	context,_ := gin.CreateTestContext(w)

	id := primitive.NewObjectID().Hex()
	context.Params = gin.Params{gin.Param{Key: "id" , Value: id}}
	context.Set("payload" , payload)
	suite.usecase.On("DeleteUser" , id).Return(nil)

	handler := suite.controller.DeleteUser()
	handler(context)

	suite.Equal(202 , w.Code)
}

func (suite *UserControllerSuite)TestDeleteUserNegative() {
    user := domain.User{
        Email:    "test@example.com",
    }

    payload := &domain.UserClaims{
        ID:    primitive.NewObjectID().Hex(),
        Email: user.Email,
        Is_admin: false,
    }
	w := httptest.NewRecorder()
	context,_ := gin.CreateTestContext(w)

	id := primitive.NewObjectID().Hex()
	context.Params = gin.Params{gin.Param{Key: "id" , Value: id}}
	context.Set("payload" , payload)
	suite.usecase.On("DeleteUser" , id).Return(nil)

	handler := suite.controller.DeleteUser()
	handler(context)

	suite.Equal(406, w.Code)
}

func (suite *UserControllerSuite)TestPromoteUserPositive(){
	user := domain.User{
		ID:       primitive.NewObjectID(),
		UserName: "TestUser",
		Email:    "test@example.com",
		Password: "password",
		Is_admin: false,
	}
	id := user.ID.Hex()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Params = gin.Params{gin.Param{Key: "id", Value: id}}

	suite.usecase.On("Promote", id).Return(user, nil)

	handler := suite.controller.PromoteUser()
	handler(context)

	suite.Equal(http.StatusAccepted, context.Writer.Status())
}

func (suite *UserControllerSuite)TestPromoteUserNegative(){
	id := primitive.NewObjectID().Hex()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	context.Params = gin.Params{gin.Param{Key: "id", Value: id}}

	r_error := errors.New("Internal server error")
	suite.usecase.On("Promote", id).Return(domain.User{}, r_error)

	handler := suite.controller.PromoteUser()
	handler(context)

	suite.Equal(400, context.Writer.Status())
}


func TestUserController(t *testing.T) {
	suite.Run(t , new(UserControllerSuite))
}