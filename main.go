package main

import (
	"log"
	"net/http"

	"github.com/jdrada/go-auth-v1/api/router"
	"github.com/jdrada/go-auth-v1/db"
)

func main() {
	db.Connect() // Init database

	r := router.Router() // Init router

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
