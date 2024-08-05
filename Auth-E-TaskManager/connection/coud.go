package connection

import (
	"context"
	"fmt"
	"log"

	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type ServerConnection struct {
	Client *mongo.Client
}


func (SC *ServerConnection)Connect_could(){

	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}

	password := os.Getenv("PASSWORD")
	url := fmt.Sprintf("mongodb+srv://ermias:%s@mngo101.bbjbuu3.mongodb.net/?retryWrites=true&w=majority&appName=Mngo101" , password)

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	options := options.Client().ApplyURI(url).SetServerAPIOptions(serverAPI)
	client,connetion_err := mongo.Connect(context.TODO() , options)

	if connetion_err != nil {
		log.Panic("Failed to connect to server" , connetion_err.Error())
	}


	if err := client.Database("TaskManager").RunCommand(context.TODO() , bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Panic("Ping failed" , err.Error())
	}

	SC.Client = client
	log.Println("Connected to server")
}