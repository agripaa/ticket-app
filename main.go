package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/jeypc/go-jwt-mux/controllers/auth"
	"github.com/jeypc/go-jwt-mux/controllers/product"
	models "github.com/jeypc/go-jwt-mux/models"
)

func main() {
	models.ConnectionDatabase()
	router := mux.NewRouter()

	router.HandleFunc("/login", auth.Login).Methods("POST")
	router.HandleFunc("/register", auth.Register).Methods("POST")
	router.HandleFunc("/logout", auth.Logout).Methods("DELETE")
	router.HandleFunc("/api/products", product.GetDataAll).Methods("GET")
	router.HandleFunc("/api/:id/product", product.FindOne).Methods("GET")
	router.HandleFunc("/api/product", product.CreateProduct).Methods("CREATE")
	router.HandleFunc("/api/:id/product", product.UpdateProduct).Methods("PATCH")
	router.HandleFunc("/api/product", product.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
