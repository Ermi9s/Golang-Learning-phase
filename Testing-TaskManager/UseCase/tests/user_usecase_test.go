package tests

import (
	"errors"
	"testing"

	usecase "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/UseCase"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/infrastructure"
	mocks "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseSuite struct {
	suite.Suite
	usecase domain.User_Usecase_interface
	repositoy *mocks.User_Repository_interface
	service *mocks.Services
}

func (suite *UserUsecaseSuite)SetupTest() {
	repository := new(mocks.User_Repository_interface)
	usecase := usecase.New_User_Usecase(repository)

	suite.repositoy = repository
	suite.usecase = usecase
	suite.service = new(mocks.Services)
}

func (suite *UserUsecaseSuite)TestGetUserPositive() {
	user := domain.User{
		ID: primitive.NewObjectID(),
		UserName: "TestingUser",
		Email: "testing@gmail.com",
		Password: "1234",
		Is_admin: false,
	}

	id := "1"
	suite.repositoy.On("GetUserDocumentById" , id).Return(user , nil)
	result , err := suite.usecase.GetUser(id)

	suite.NoError(err , "Error found on GetUser")
	suite.Equal(result , user , "return result not equal with specified")
}

func (suite *UserUsecaseSuite)TestGetUserNegative() {
	id := "1"
	r_error := errors.New("Internal server error")
	suite.repositoy.On("GetUserDocumentById" , id).Return(domain.User{} , r_error)
	result , err := suite.usecase.GetUser(id)

	suite.Error(err , "Error not found on GetUser")
	suite.Equal(result , domain.User{} , "return result not equal with specified")
}

func (suite *UserUsecaseSuite)TestGetUsersPositive() {
	user1 := domain.User{
		ID: primitive.NewObjectID(),
		UserName: "TestingUser2",
		Email: "testing1@gmail.com",
		Password: "1234",
		Is_admin: false,
	}
	user2 := domain.User{
		ID: primitive.NewObjectID(),
		UserName: "TestingUser2",
		Email: "testing2@gmail.com",
		Password: "1234",
		Is_admin: false,
	}

	data := []domain.User{user1 , user2}
	filter := make(map[string]string)
	suite.repositoy.On("GetUserDocumentByFilter" , filter).Return(data , nil)
	result,err := suite.usecase.GetUsers()

	suite.NoError(err , "Error found in GetUsers")
	suite.Equal(result , data , "returned result not the same as repository return")
}

func (suite *UserUsecaseSuite)TestGetUsersNegative() {
	r_error := errors.New("Internal server error")
	filter := make(map[string]string)
	suite.repositoy.On("GetUserDocumentByFilter" , filter).Return(nil , r_error)
	result,err := suite.usecase.GetUsers()

	suite.Error(err , "No Error found in GetUsers")
	suite.Equal(result , []domain.User{} , "returned result not the same as repository return")
}


func (suite *UserUsecaseSuite)TestCreateUserPositive() {
	user := domain.User{
		ID: primitive.NewObjectID(),
		UserName: "TestingUser2",
		Email: "testing1@gmail.com",
		Password: "1234",
		Is_admin: false,
	}
	serv := infrastructure.KeyServices{}

	typeUser := mock.AnythingOfType("domain.User")
	suite.repositoy.On("InsertUserDocument",typeUser).Return(user.ID.Hex() , nil)
	
	s_user,token,err := suite.usecase.CreateUser(user)
	
	passerr := bcrypt.CompareHashAndPassword([]byte(s_user.Password) , []byte("1234"))
	s_token ,_:= serv.Encode(user.ID.Hex() , user.Email , user.Is_admin)
	
	suite.Nil(passerr , "Hash password has errors")
	suite.Equal(token , s_token , "token doesnt match")
	suite.NoError(err , "Error found in CreateUser")

}

func (suite *UserUsecaseSuite)TestCreateUserNegative() {
	user := domain.User{
		ID: primitive.NewObjectID(),
		UserName: "TestingUser2",
		Email: "testing1@gmail.com",
		Password: "1234",
		Is_admin: false,
	}
	serv := infrastructure.KeyServices{}

	typeUser := mock.AnythingOfType("domain.User")
	r_error := errors.New("Internal server error")
	suite.repositoy.On("InsertUserDocument",typeUser).Return("" , r_error)
	
	s_user,token,err := suite.usecase.CreateUser(user)
	
	passerr := bcrypt.CompareHashAndPassword([]byte(s_user.Password) , []byte("1234"))
	s_token ,_:= serv.Encode(user.ID.Hex() , user.Email , user.Is_admin)
	
	suite.NotNil(passerr , "Hash should password have errors")
	suite.NotEqual(token , s_token , "token should not be generated")
	suite.Error(err , "Error should be found in CreateUser")

}

func (suite *UserUsecaseSuite)TestDeleteUserPositive() {
	id := "1"
	suite.repositoy.On("DeleteUserDocument" , id).Return(nil)

	err := suite.usecase.DeleteUser(id)

	suite.Nil(err , "Error found in DeleteUser")

}

func (suite *UserUsecaseSuite)TestDeleteUserNegative() {
	id := "1"
	r_error := errors.New("Internal server error")
	suite.repositoy.On("DeleteUserDocument" , id).Return(r_error)

	err := suite.usecase.DeleteUser(id)

	suite.NotNil(err , "Error should be found in DeleteUser")
}

func (suite *UserUsecaseSuite)TestLogInPositive() {
	user := domain.User{
		ID: primitive.NewObjectID(),
		UserName: "TestingUser2",
		Email: "testing1@gmail.com",
		Password: "1234",
		Is_admin: false,
	}
	
	serv := infrastructure.KeyServices{}
	user.Password = serv.HashPassWord(user.Password)

	auser := domain.AuthUser{
		ID: user.ID.Hex(),
		UserName: "TestingUser2",
		Email: "testing1@gmail.com",
		Password: "1234",
		Is_admin: false,
	}

	output := []domain.User{}
	output = append(output, user)
	suite.repositoy.On("GetUserDocumentByFilter" , mock.Anything).Return(output, nil)
	result,err := suite.usecase.LogIn(auser)

	suite.Nil(err , "Error found in LogIn")
	suite.Equal(result , user , "Returned user not the same")

}

func (suite *UserUsecaseSuite)TestLogInNegative() {
	user := domain.User{
		ID: primitive.NewObjectID(),
		UserName: "TestingUser2",
		Email: "testing1@gmail.com",
		Password: "1234",
		Is_admin: false,
	}
	
	serv := infrastructure.KeyServices{}
	user.Password = serv.HashPassWord(user.Password)

	auser := domain.AuthUser{
		ID: user.ID.Hex(),
		UserName: "TestingUser2",
		Email: "testing1@gmail.com",
		Password: "1234",
		Is_admin: false,
	}

	r_error := errors.New("Internal server Error")
	suite.repositoy.On("GetUserDocumentByFilter" , mock.Anything).Return([]domain.User{}, r_error)
	result,err := suite.usecase.LogIn(auser)

	suite.NotNil(err , "Error should be found in LogIn")
	suite.NotEqual(result , user , "Returned user should not the same")

}

func (suite *UserUsecaseSuite)TestUpdateUserPositive(){
	id := "1"
	user := domain.User{
		ID: primitive.NewObjectID(),
		UserName: "TestingUser2",
		Email: "testing1@gmail.com",
		Password: "1234",
		Is_admin: false,
	}

	suite.repositoy.On("UpdateUserDocumentById" , id , user).Return(nil)
	result,err := suite.usecase.UpdateUser(id , user)

	suite.Equal(result , user , "Models should be the same")
	suite.Nil(err , "Error returned")
	
}

func (suite *UserUsecaseSuite)TestUpdateUserNegative(){
	id := "1"
	user := domain.User{
		ID: primitive.NewObjectID(),
		UserName: "TestingUser2",
		Email: "testing1@gmail.com",
		Password: "1234",
		Is_admin: false,
	}

	r_error := errors.New("Internal server error")
	suite.repositoy.On("UpdateUserDocumentById" , id , user).Return(r_error)
	result,err := suite.usecase.UpdateUser(id , user)

	suite.NotEqual(result , user , "Models should not be the same")
	suite.NotNil(err , "Error returned")
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t , new(UserUsecaseSuite))
}