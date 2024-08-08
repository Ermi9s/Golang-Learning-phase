package Infrastructure

import (
	"context"
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



var Root domain.User = domain.User{
	UserName: "root",
	Email: "root@gmail.com",
	Password: "12345",
	Is_admin: true,
}


var DataBase *mongo.Database
var Task_Collection *mongo.Collection
var User_Collection *mongo.Collection
var client connection.ServerConnection

func Start() {

	client.Connect_could()
	DataBase = client.Client.Database("TaskManager")
	Task_Collection = client.Client.Database(DataBase.Name()).Collection("Tasks")
	User_Collection = client.Client.Database(DataBase.Name()).Collection("Users")

	var DataBaseManager domain.DataBaseManager = services.DataBaseManager{
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
	
}