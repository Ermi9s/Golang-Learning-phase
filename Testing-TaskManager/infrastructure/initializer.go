package infrastructure

import (
	"context"

	database "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/database/connection"
	databasedomain "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/database/databaseDomain"
	"github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Root domain.User = domain.User{
	UserName: "root",
	Email: "root@gmail.com",
	Password: "1234",
	Is_admin: true,
}

var DataBase databasedomain.Database
var User_Collection databasedomain.Collection
var client database.ServerConnection

func Start() databasedomain.Client{
	//make connection 
	client.Connect_could()
	DataBase = client.Client.Database("TaskManager")
	User_Collection = client.Client.Database("TaskManager").Collection("Users")

	//make email unique
	indexModel := mongo.IndexModel{
		Keys: bson.D{{
			Key: "email",
			Value: 1,
		}},
		Options: options.Index().SetUnique(true),
	}

	//ensure creation by inserting a root user
	filter := bson.D{{Key : "username" , Value: Root.UserName} , {Key : "email" , Value: Root.Email}}
	User_Collection.Indexes().CreateOne(context.TODO(), indexModel)
	User_Collection.FindOne(context.TODO() ,filter)

	return client.Client
}