package controllers

import (
	"Learnings/go/firstApi/database"
	"Learnings/go/firstApi/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Server Ok")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request){
	var listPersonalities []models.Personality
	database.DB.Find(&listPersonalities)
	json.NewEncoder(w).Encode(listPersonalities)
}

func ReturnPersonalityById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func PostNewPersonality(w http.ResponseWriter, r *http.Request){
	var newPersonality models.Personality
	json.NewDecoder(r.Body).Decode(&newPersonality)
	database.DB.Create(&newPersonality)
	json.NewEncoder(w).Encode(newPersonality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)
}