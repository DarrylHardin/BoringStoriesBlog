package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Post struct {
	Title   string
	Content string
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	data := Post{
		Title:   "fasfas",
		Content: "dfdfdf!",
	}
	// Build absolute file paths using the project root as a starting point
	baseFilePath := filepath.Join("..", "..", "internal", "templates", "base.html")
	blogFilePath := filepath.Join("..", "..", "internal", "templates", "blog.html")

	// Parse templates using absolute file paths
	tmpl, err := template.ParseFiles(baseFilePath, blogFilePath)
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
