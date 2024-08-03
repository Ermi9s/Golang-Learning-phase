package main

import (
	"context"
	"github.com/ermi9s/db-taskmanager/connection"

)
func main() {
	remote := connection.ServerConnection{}
	remote.Connect_could()
	local := connection.LServerCon{}
	local.Connect_local()
	
	remote.Client.Database("TaskManager").CreateCollection(context.TODO() , "NEWTEST")
	local.Client.Database("TaskManager").CreateCollection(context.TODO() , "NEWTEST")
}