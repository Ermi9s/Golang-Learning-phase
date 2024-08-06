package router

import (
	"context"

	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/connection"
	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/models"
	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var Root models.User = models.User{
	UserName: "root",
	Email: "root@gmail.com",
	Password: "12345678",
	Is_admin: true,
}


var DataBase *mongo.Database
var Task_Collection *mongo.Collection
var User_Collection *mongo.Collection
var client connection.ServerConnection

func Start() services.DataBaseManager {

	client.Connect_could()
	DataBase = client.Client.Database("TaskManager")
	Task_Collection = client.Client.Database(DataBase.Name()).Collection("Tasks")
	User_Collection = client.Client.Database(DataBase.Name()).Collection("Users")

	var DataBaseManager services.DataBaseManager = services.DataBaseManager{
		Client: client.Client,
		Tasks: Task_Collection,
		Users: User_Collection,
	}

	indexModel := mongo.IndexModel{
		Keys: bson.D{{
			Key: "email",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	}

	filter := bson.D{{Key : "username" , Value: Root.UserName} , {Key : "email" , Value: Root.Email}}
	//ensure creation by inserting a root user
	DataBaseManager.Users.Indexes().CreateOne(context.TODO(), indexModel)
	root := DataBaseManager.Users.FindOne(context.TODO() ,filter)
	if root.Err() != nil {
		DataBaseManager.CreateUser(Root)
	}
	
	return DataBaseManager
}