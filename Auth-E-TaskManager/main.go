package main

import (
	"log"
	"github.com/Ermi9s/Golang-Learning-phase/Auth-E-TaskManager/router"
)

func main() {
	//database manager
	DBM := router.Start()
	//run the routes
	router.Run(&DBM)
	log.Println("Server Running on 127.0.0.1:8080")
	
}