package repository

import (
	databasedomain "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/database/databaseDomain"
)

type Repository struct {
	Client databasedomain.Client
	Database databasedomain.Database
}

func NewRepository(client databasedomain.Client, data_base databasedomain.Database) *Repository {
	return &Repository {
		Client: client,
		Database: data_base,
	}
}