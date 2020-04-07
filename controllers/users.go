package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"techTask/models"
	"techTask/service"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	service.InsertUser(user)
	json.NewEncoder(w).Encode(user)
	fmt.Println("New person entered...")

}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	payload := service.GetAllUsers()
	json.NewEncoder(w).Encode(payload)

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}
