package database

import (
	"context"
	"log"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LServerCon struct {
	Client *mongo.Client
}

func (LS *LServerCon)Connect_local(){
	env_err := godotenv.Load()
	if env_err != nil {
		log.Panic("Failed to load env" ,env_err.Error())
	}

	url := "mongodb://localhost:27017/"
	options := options.Client().ApplyURI(url)

	client , err := mongo.Connect(context.TODO() , options)
	if err != nil {
		log.Panic("Faild to connect to local server" , err.Error())
	}

	LS.Client = client
}