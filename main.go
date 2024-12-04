package main

import (
	"go_crud/database"
	"go_crud/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/entries", handlers.GetEntries).Methods("GET")
	r.HandleFunc("/entries", handlers.CreateEntry).Methods("POST")
	r.HandleFunc("/entries/{id}", handlers.UpdateEntry).Methods("PUT")
	r.HandleFunc("/entries/{id}", handlers.DeleteEntry).Methods("DELETE")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
