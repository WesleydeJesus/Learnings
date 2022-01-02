package routes

import (
	"Learnings/go/firstApi/controllers"
	"Learnings/go/firstApi/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest(){
	router := mux.NewRouter()
	router.Use(middleware.ContentType)
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	router.HandleFunc("/api/personalities/{id}", controllers.ReturnPersonalityById).Methods("Get")
	router.HandleFunc("/api/personalities", controllers.PostNewPersonality).Methods("Post")
	router.HandleFunc("/api/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	router.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(router)))
}