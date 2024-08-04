package main

import (
	"context"
	"github.com/Ermi9s/Golang-Learning-phase/DataBase-E-Task-manager/connection"
)
func main() {
	remote := connection.ServerConnection{}
	remote.Connect_could()
	local := connection.LServerCon{}
	local.Connect_local()
	
	remote.Client.Database("TaskManager").Collection("NEWTEST").Drop(context.TODO())
	local.Client.Database("TaskManager").Collection("NEWTEST").Drop(context.TODO())

	
}

