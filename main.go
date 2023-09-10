package main

import (
	"net/http"

	"github.com/gorilla/mux"
	auth "github.com/jeypc/go-jwt-mux/controllers/auth"
	"github.com/jeypc/go-jwt-mux/controllers/product"
	"github.com/jeypc/go-jwt-mux/middleware"
	models "github.com/jeypc/go-jwt-mux/models"
	"github.com/rs/cors"
)

func main() {
	models.ConnectionDatabase()
	router := mux.NewRouter()

	router.HandleFunc("/api/ticket/login", auth.Login).Methods("POST")
	router.HandleFunc("/api/ticket/register", auth.Register).Methods("POST")
	router.HandleFunc("/api/ticket/logout", auth.Logout).Methods("DELETE")

	router.HandleFunc("/api/ticket/products", product.GetDataAll).Methods("GET")
	router.HandleFunc("/api/ticket/{id}/product", product.FindOne).Methods("GET")
	router.HandleFunc("/api/ticket/product", product.CreateProduct).Methods("POST")
	router.HandleFunc("/api/ticket/{id}/product", product.UpdateProduct).Methods("PATCH")
	router.HandleFunc("/api/ticket/{id}/product", product.DeleteProduct).Methods("DELETE")
	router.Use(middleware.JWTMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(":8080", handler)
}
