package routes

import (
	"boringstories/internal/handlers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Define your routes here
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/blog", handlers.BlogHandler)

	return mux
}
