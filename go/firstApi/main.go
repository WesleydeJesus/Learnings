package main

import (
	"Learnings/go/firstApi/database"
	"Learnings/go/firstApi/routes"
	"fmt"
)


func main() {
	database.ConnectDB()
	fmt.Println("Initialize server")
	routes.HandleRequest()
}