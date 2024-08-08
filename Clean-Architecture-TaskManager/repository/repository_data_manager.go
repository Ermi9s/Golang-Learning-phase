package repository

import (
	// domain "github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/Domain"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Client *mongo.Client
	Database *mongo.Database
}

// type DataBaseUser interface {
// 	GetUser(id primitive.ObjectID) (domain.User , error)
// 	GetUsers() ([]domain.User, error)
// 	CreateUser(user domain.User) (domain.User , error)
// 	UpdateUser(id primitive.ObjectID , user domain.User)(domain.User , error)
// 	DeleteUser(id primitive.ObjectID) error
// 	MakeAdmin(id primitive.ObjectID)(domain.User , error)
// 	LogIn(user domain.AuthUser)(domain.LoggedInUser , error)
// }

// type DatabaseTask interface {
// 	GetTask(id primitive.ObjectID) (domain.Task , error)
// 	GetTasks() ([]domain.Task , error)
// 	CreateTask(model domain.Task) (domain.Task ,error)
// 	UpdateTask(id primitive.ObjectID, model domain.Task) (domain.Task, error)
// 	DeleteTask(id primitive.ObjectID) error
// }