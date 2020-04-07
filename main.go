package main

import (
	"fmt"
	"net/http"
	"techTask/controllers"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("Starting the application...")
	r := mux.NewRouter()

	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/user/delete/{id}", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/user/update/{id}", controllers.UpdateUser).Methods("PUT")

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)

}
