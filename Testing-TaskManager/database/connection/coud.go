package database

import (
	"fmt"
	"log"
	"os"
	databasedomain "github.com/Ermi9s.Golang-Learning-phase/Testing-TaskManager/database/databaseDomain"
	"github.com/joho/godotenv"
)

type ServerConnection struct {
	Client databasedomain.Client
}

func (SC *ServerConnection)Connect_could(){
	err := godotenv.Load()
	if err != nil {
		log.Panic("Failed to load .env" , err.Error())
	}

	password := os.Getenv("PASSWORD")
	url := fmt.Sprintf("mongodb+srv://ermias:%s@mngo101.bbjbuu3.mongodb.net/?retryWrites=true&w=majority&appName=Mngo101" , password)

	client,connetion_err := databasedomain.NewClient(url)

	if connetion_err != nil {
		log.Panic("Failed to connect to server\n" , connetion_err.Error())
		return
	}

	SC.Client = client
	log.Println("Connected to server")
}