package routes

import (
	"html/template"
	"net/http"

	"groupie-tracker/internal/handlers"
)

// NewRouter sets up the mux with all the necessary routes.
func NewRouter(tpl *template.Template) *http.ServeMux {
	mux := http.NewServeMux()

	// First welcome page route
	mux.HandleFunc("/", handlers.IntroHandler(tpl))

	// HomePage route
	mux.HandleFunc("/home", handlers.HomeHandler(tpl))

	// Artist detail route.
	mux.HandleFunc("/artist/", handlers.DetailHandler(tpl))

	// API Route for fetching artists
	mux.HandleFunc("/api/artists", handlers.GetArtists)

	// API Route for fetching all locations for filters
	mux.HandleFunc("/api/all-locations", handlers.GetAllLocations)

	// API Route for fetching filtered results
	mux.HandleFunc("/api/filters", handlers.FiltersResultHandler())

	//Search handler
	mux.HandleFunc("/search", handlers.SearchHandler)

	// About page section
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "about.html", nil)

	})

	// Serve static files.
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	return mux
}
