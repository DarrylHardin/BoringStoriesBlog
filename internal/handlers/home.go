package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type HomePageData struct {
	Title   string
	Content string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := HomePageData{
		Title:   "My Website!",
		Content: "Welcome to my website!",
	}
	// Build absolute file paths using the project root as a starting point
	baseFilePath := filepath.Join("..", "..", "internal", "templates", "base.html")
	homeFilePath := filepath.Join("..", "..", "internal", "templates", "home.html")

	// Parse templates using absolute file paths
	tmpl, err := template.ParseFiles(baseFilePath, homeFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
