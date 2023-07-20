package router

import (
	"github.com/gorilla/mux"
	"github.com/jdrada/go-auth-v1/api/handler"
	"github.com/jdrada/go-auth-v1/api/middleware"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user/register", handler.Register).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/user/login", handler.Login).Methods("POST", "OPTIONS")

	protected := router.PathPrefix("/api/user/protected").Subrouter()
	protected.Use(middleware.JwtVerify)
	protected.HandleFunc("", handler.ProtectedEndpoint).Methods("GET", "OPTIONS")

	return router
}
