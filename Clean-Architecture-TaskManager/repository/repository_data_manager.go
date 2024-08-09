package repository

import (
	"github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/domain"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Client *mongo.Client
	Database *mongo.Database
}

func NewRepository(client *mongo.Client , data_base *mongo.Database) domain.Repository_interface {
	return &Repository {
		Client: client,
		Database: data_base,
	}
}