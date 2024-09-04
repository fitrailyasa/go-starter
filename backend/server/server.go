package server

import (
	"go-starter/backend/db"
	"go-starter/backend/routes"
	"log"
	"net/http"
)

func Start() {
	db.InitDB()
	// db.SeedData()

	router := routes.SetupRoutes()

	fs := http.FileServer(http.Dir("./frontend/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
