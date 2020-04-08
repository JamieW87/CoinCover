package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"techTask/models"
	"techTask/service"
)

//CreateUser handles the response for creating a single user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	service.InsertUser(user)
	json.NewEncoder(w).Encode(user)
	fmt.Println("New person entered...")

}

//GetAllUsers gets all entries from the DB
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	payload := service.GetAllUsers()
	json.NewEncoder(w).Encode(payload)

}
