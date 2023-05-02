package main

import (
	"log"
	"net/http"

	"boringstories/internal/routes"
)

func main() {
	mux := routes.SetupRoutes()

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
	// // Create a new handler for the home page
	// homeHandler := http.HandlerFunc(handlers.HomeHandler)

	// // Serve static files from the "static" directory
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// // Route requests to the home page to the homeHandler
	// http.HandleFunc("/", homeHandler)

	// // Start the server
	// log.Println("Server listening on port 8080...")
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
