package main

import (
	"log"
	"net/http"

	"github.com/agripaa/ticket-app/controllers/authcontroller"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", authcontroller.Login).Methods("POST")
	router.HandleFunc("/register", authcontroller.Register).Methods("POST")
	router.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
