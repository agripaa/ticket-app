package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/jeypc/go-jwt-mux/controllers/auth"
	"github.com/jeypc/go-jwt-mux/controllers/product"
	"github.com/jeypc/go-jwt-mux/middleware"
	models "github.com/jeypc/go-jwt-mux/models"
)

func main() {
	models.ConnectionDatabase()
	router := mux.NewRouter()

	router.HandleFunc("/login", auth.Login).Methods("POST")
	router.HandleFunc("/register", auth.Register).Methods("POST")
	router.HandleFunc("/logout", auth.Logout).Methods("DELETE")

	api := router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/products", product.GetDataAll).Methods("GET")
	api.HandleFunc("/:id/product", product.FindOne).Methods("GET")
	api.HandleFunc("/product", product.CreateProduct).Methods("CREATE")
	api.HandleFunc("/:id/product", product.UpdateProduct).Methods("PATCH")
	api.HandleFunc("/product", product.DeleteProduct).Methods("DELETE")
	api.Use(middleware.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", router))
}
