package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"groupie-tracker/internal/data"
	"groupie-tracker/internal/routes"
	"groupie-tracker/internal/utils"
)

func main() {
	// Load data from API.
	if err := data.LoadData(); err != nil {
		log.Fatalf("Error loading data: %v", err)
	}

	// Define custom template functions.
	funcMap := template.FuncMap{
		"replaceSpaces":  utils.ReplaceSpaces,
		"cleanDate":      utils.CleanDate,
		"formatLocation": utils.FormatLocation,
	}

	// Parse all templates.
	tmplPattern := filepath.Join("templates", "*.html")
	tpl := template.New("").Funcs(funcMap)
	tpl, err := tpl.ParseGlob(tmplPattern)
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	// Initialize the router from the routes package.
	mux := routes.NewRouter(tpl)

	// Start the server.
	addr := ":8080"
	fmt.Printf("Server is running at http://localhost%s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
