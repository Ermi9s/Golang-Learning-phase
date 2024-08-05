package services

import (
	"context"
	models "github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseTask interface {
	GetTask(id int) (models.Task , error)
	CreateTask(model models.Task) error
	GetTasks() (models.Task , error)
	UpdateTask(id primitive.ObjectID, model models.Task) (models.Task, error)
	DeleteTask(id primitive.ObjectID) error
	GetUser(id int) (models.User , error)
}

type DataBaseManager struct {
	Client *mongo.Client
	Tasks *mongo.Collection
	Users *mongo.Collection
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

func (DBM *DataBaseManager)GetTasks() (models.Task , error) {
	var decoded models.Task
	filter := bson.D{}
	err := DBM.Tasks.FindOne(context.TODO() , filter).Decode(&decoded)

	if err != nil {
		return models.Task{},err
	}

	return decoded , nil
}

func (DBM *DataBaseManager)CreateTask(model models.Task) (models.Task ,error) {
	var doc bson.M
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

	_, err = DBM.Tasks.UpdateOne(context.TODO(), filter, update)
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


func (DBM *DataBaseManager)GetUser(id int) (models.User , error) {
	filter := bson.D{{Key : "_id" , Value: id}}
	var user models.User

	err := DBM.Users.FindOne(context.TODO() , filter).Decode(&user)
	if err != nil {
		return models.User{} ,err
	}

	return user , nil
}