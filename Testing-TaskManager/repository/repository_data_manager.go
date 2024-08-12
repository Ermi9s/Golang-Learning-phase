package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Client *mongo.Client
	Database *mongo.Database
}

func NewRepository(client *mongo.Client , data_base *mongo.Database) *Repository {
	return &Repository {
		Client: client,
		Database: data_base,
	}
}