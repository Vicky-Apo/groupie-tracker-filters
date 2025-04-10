package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"groupie-tracker/internal/data"
)

// DetailHandler handles the "/artist/{name}" route
func DetailHandler(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Extract artist name from URL and replace "-" back to spaces
		artistName := strings.TrimPrefix(r.URL.Path, "/artist/")
		artistName = strings.ReplaceAll(artistName, "-", " ")

		if artistName == "" {
			http.NotFound(w, r)
			return
		}

		// Find the artist by name (case-insensitive)
		var foundArtist *data.Artist
		for i := range data.AllArtists {
			if strings.EqualFold(data.AllArtists[i].Name, artistName) {
				foundArtist = &data.AllArtists[i]
				break
			}
		}
		if foundArtist == nil {
			handler404(tpl, w)
			return
		}

		// Find locations for the artist
		var locations []string
		for _, loc := range data.AllLocations.Index {
			if loc.ID == foundArtist.ID {
				locations = loc.Locations
				break
			}
		}

		// Find dates for the artist
		var dates []string
		for _, date := range data.AllDates.Index {
			if date.ID == foundArtist.ID {
				dates = date.Dates
				break
			}
		}

		// Find relation data (dates & locations)
		var relationMap map[string][]string
		for _, relation := range data.AllRelations.Index {
			if relation.ID == foundArtist.ID {
				relationMap = relation.DatesLocations
				break
			}
		}

		// Prepare data to send to the template
		dataToSend := data.CombinedData{
			Artist:      *foundArtist,
			Locations:   locations,
			Dates:       dates,
			DatesLocMap: relationMap,
		}

		// Render template
		err := tpl.ExecuteTemplate(w, "artist.html", dataToSend)
		if err != nil {
			fmt.Println("Error rendering template:", err)
			http.Error(w, "Internal Server Error while rendering detail", http.StatusInternalServerError)
			return
		}
	}
}
