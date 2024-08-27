package repository

import (
	"context"
	"testing"

	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/database/databaseDomain/mocks"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	suite.Suite
	main_repo *Repository

	user_repo domain.User_Repository_interface

	single_result *mocks.SingleResult
	collection *mocks.Collection
	database *mocks.Database
	cursor *mocks.Cursor
}

func (suite *UserRepository)SetupTest(){
	client := new(mocks.Client)

	suite.collection = new(mocks.Collection)
	suite.database = new(mocks.Database)
	suite.single_result = new(mocks.SingleResult)

	suite.cursor = new(mocks.Cursor)
	suite.main_repo = NewRepository(client , suite.database)
	suite.user_repo = New_User_Repository(*suite.main_repo , suite.collection)
}

func(suite *UserRepository)TestGetUserDocumentById(){
	user := domain.User{
		ID: primitive.NewObjectID(),
	}
	filter := bson.D{{Key : "_id" , Value: user.ID}}

	suite.collection.On("FindOne" , context.TODO() ,filter).Return(suite.single_result)
	suite.single_result.On("Decode" , mock.Anything).Return(nil)

	id := user.ID.Hex()
	_, err := suite.user_repo.GetUserDocumentById(id)

	suite.Nil(err , "Found error GetUserDocumentById")
}

func (suite *UserRepository) TestGetUserDocumentByFilter() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		UserName: "TestingUser",
		Email:    "testing@gmail.com",
		Password: "1234",
		Is_admin: false,
	}
	filter := map[string]string{
		"_id": user.ID.Hex(),
	}

	suite.cursor.On("Next", context.TODO()).Return(true).Once() 
	suite.cursor.On("Next", context.TODO()).Return(false).Once()
	suite.cursor.On("Decode", mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.User)
		*arg = user
	}).Return(nil)

	suite.collection.On("Find", context.TODO(), mock.Anything).Return(suite.cursor, nil)

	result, err := suite.user_repo.GetUserDocumentByFilter(filter)

	suite.Nil(err, "Expected no error")
	suite.NotNil(result, "Expected non-nil result")
	suite.Equal(1, len(result), "Expected one user in the result")
	suite.Equal(user, result[0], "Expected the returned user to match the mock user")
	}

	func (suite *UserRepository)TestUpdateUserDocumentById(){
		user := domain.User{
			ID:       primitive.NewObjectID(),
			UserName: "TestingUser",
			Email:    "testing@gmail.com",
			Password: "1234",
			Is_admin: false,
		}

		var doc bson.D
		byteModel,_ := bson.Marshal(user)
		bson.Unmarshal(byteModel , &doc)
		filter := bson.D{{Key : "_id" , Value: user.ID}}
		updater := bson.D{{Key: "$set" , Value: doc}}

		suite.collection.On("UpdateOne" , context.TODO() , filter , updater).Return(nil,nil)

		err := suite.user_repo.UpdateUserDocumentById(user.ID.Hex() , user)

		suite.Nil(err, "Error found on updateuser")
	}

	func (suite *UserRepository)TestInsertUserDocument() {
		user := domain.User{
			ID:       primitive.NewObjectID(),
			UserName: "TestingUser",
			Email:    "testing@gmail.com",
			Password: "1234",
			Is_admin: false,
		}

		var doc bson.D
		byteModel,_ := bson.Marshal(user)
		bson.Unmarshal(byteModel , &doc)
	
		suite.collection.On("InsertOne" , context.TODO() , doc).Return(user.ID, nil)

		id , err := suite.user_repo.InsertUserDocument(user)
		suite.Nil(err , "error found in Insert")
		suite.Equal(id , user.ID.Hex() , "Id not the same")
	}

	func (suite *UserRepository)TestDeleteUserDocument() {
		id := primitive.NewObjectID()
		filter := bson.D{{Key : "_id" , Value:id}}

		var r int64
		suite.collection.On("DeleteOne" , context.TODO() , filter).Return(r, nil)

		err := suite.user_repo.DeleteUserDocument(id.Hex())

		suite.Nil(err , "Error found on delete")
	}

func TestUserRepository(t *testing.T) {
	suite.Run(t , new(UserRepository))
}
