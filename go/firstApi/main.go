package main

import (
	"Learnings/go/firstApi/database"
	"Learnings/go/firstApi/models"
	"Learnings/go/firstApi/routes"
	"fmt"
)


func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Wesley", Description: "Opa1"},
		{Id: 2, Name: "Gabi", Description: "opa2" },
	} 
	database.ConnectDB()
	fmt.Println("Inicializando server")
	routes.HandleRequest()
}