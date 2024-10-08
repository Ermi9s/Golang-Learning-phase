package services

import (
	"context"
	"errors"
	models "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseTask interface {
	GetTask(id primitive.ObjectID) (models.Task , error)
	GetTasks() ([]models.Task , error)
	CreateTask(model models.Task) (models.Task ,error)
	UpdateTask(id primitive.ObjectID, model models.Task) (models.Task, error)
	DeleteTask(id primitive.ObjectID) error
}

type DataBaseUser interface {
	GetUser(id primitive.ObjectID) (models.User , error)
	GetUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User , error)
	UpdateUser(id primitive.ObjectID , user models.User)(models.User , error)
	DeleteUser(id primitive.ObjectID) error
	MakeAdmin(id primitive.ObjectID)(models.User , error)
	LogIn(user models.LogIN)(models.LoggedInUser , error)
}

type DataBaseManager struct {
	Client *mongo.Client
	Tasks *mongo.Collection
	Users *mongo.Collection
	Filter bson.D
}

var DataBase string = "TaskManager"
var Collection string = "Tasks"

func (DBM *DataBaseManager)GetTask(id primitive.ObjectID) (models.Task , error) {
	var decoded models.Task
	filter := bson.D{{Key: "_id" , Value: id}}
	err := DBM.Tasks.FindOne(context.TODO() , filter).Decode(&decoded)

	if err != nil {
		return models.Task{},err
	}

	return decoded , nil
}

func (DBM *DataBaseManager)GetTasks() ([]models.Task , error) {
	filter := bson.D{}
	var decoded []models.Task
	if DBM.Filter != nil {
		mp := DBM.Filter.Map()
		filter = bson.D{{Key: "creator_id" , Value: mp["_id"]}}
	}

	Curr,err := DBM.Tasks.Find(context.TODO() , filter)
	for  Curr.Next(context.TODO()) {
		var task models.Task
		err := Curr.Decode(&task)
		if err != nil {
			return nil,nil
		}
		decoded = append(decoded, task)
	}
	if err != nil {
		return []models.Task{},err
	}
	return decoded , nil
}

func (DBM *DataBaseManager)CreateTask(model models.Task) (models.Task ,error) {
	var doc bson.M
	model.ID = primitive.NewObjectID()
	bsonModel,err := bson.Marshal(model)
	
	if err != nil {
		return models.Task{} , err
	}
	err = bson.Unmarshal(bsonModel , &doc)
	if err != nil {
		return models.Task{} , err
	}
	_ , err =DBM.Tasks.InsertOne(context.TODO() , doc)
	if err != nil {
		return models.Task{} , err
	}

	return model , nil
}

func (DBM *DataBaseManager) UpdateTask(id primitive.ObjectID, model models.Task) (models.Task, error) {
	mp := DBM.Filter.Map()
	if model.Creater != mp["_id"] {
		return models.Task{},errors.New("not authorized")
	}

	bsonModel, err := bson.Marshal(model)
	if err != nil {
		return models.Task{}, err
	}

	var doc bson.M
	err = bson.Unmarshal(bsonModel, &doc)
	if err != nil {
		return models.Task{}, err
	}

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: doc}}

	_,err = DBM.Tasks.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return models.Task{}, err
	}

	return model, nil
}

func (DBM *DataBaseManager)DeleteTask(id primitive.ObjectID) error {
	filter := bson.D{{Key: "_id" , Value: id}}
	_,err := DBM.Tasks.DeleteOne(context.TODO() , filter)
	return err
}

