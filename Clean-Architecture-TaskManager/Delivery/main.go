package main

import "github.com/Ermi9s.Golang-Learning-phase/Clean-Architecture-TaskManager/Delivery/routes"

func main(){
	databasemanger := routes.Dependancy()
	routes.Run(databasemanger)
}