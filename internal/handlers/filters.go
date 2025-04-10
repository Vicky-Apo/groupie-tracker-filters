package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"groupie-tracker/internal/data"
	"groupie-tracker/internal/utils"
)

// FiltersResultHandler handles the AJAX/HTTP request for applying filters and returning the filtered list.
func FiltersResultHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We expect query parameters from the frontend (filters.js),
		// e.g. ?minCreation=1990&maxCreation=2005&minAlbum=1990&maxAlbum=2005&members=1&members=3&location=Texas, USA
		// Parse them below:

		// Parse creation year range
		minCreationStr := r.URL.Query().Get("minCreation")
		maxCreationStr := r.URL.Query().Get("maxCreation")
		minCreation, _ := strconv.Atoi(minCreationStr)
		maxCreation, _ := strconv.Atoi(maxCreationStr)

		// Parse first album year range
		minAlbumStr := r.URL.Query().Get("minAlbum")
		maxAlbumStr := r.URL.Query().Get("maxAlbum")
		minAlbum, _ := strconv.Atoi(minAlbumStr)
		maxAlbum, _ := strconv.Atoi(maxAlbumStr)

		// Parse "members" checkboxes (can be multiple)
		// e.g. members=1&members=4&members=6
		membersSelected := r.URL.Query()["members"] // returns []string

		// Parse "location" multi-select
		// e.g. location=Washington, USA&location=Texas, USA
		locationsSelected := r.URL.Query()["location"] // returns []string

		// Now do the actual filtering logic
		filtered := filterArtists(minCreation, maxCreation, minAlbum, maxAlbum, membersSelected, locationsSelected)

		sortedFiltered := utils.SortingArtists(filtered)

		// Return JSON list of filtered artists
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(sortedFiltered)
	}
}

// GetAllLocations returns a list of unique concert locations.
func GetAllLocations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Use AllLocations instead of AllRelations
	uniqueSet := make(map[string]bool)

	for _, artist := range data.AllLocations.Index {
		for _, loc := range artist.Locations {
			formatedLoc := utils.FormatLocation(loc)
			uniqueSet[formatedLoc] = true
		}
	}

	// Convert set to slice
	var locations []string
	for loc := range uniqueSet {
		locations = append(locations, loc)
	}

	// Sort the list alphabetically
	sort.Strings(locations)

	// Return as JSON
	json.NewEncoder(w).Encode(locations)
}

// filterArtists returns only those artists that match ALL of the given criteria.
func filterArtists(
	minCreation, maxCreation int,
	minAlbum, maxAlbum int,
	membersSelected []string,
	locationsSelected []string,
) []data.Artist {
	var results []data.Artist

	for _, artist := range data.AllArtists {
		// 1) Check creation date
		if !matchesCreationYear(artist.CreationDate, minCreation, maxCreation) {
			continue
		}

		// 2) Check first album year
		if !matchesAlbumYear(artist.FirstAlbum, minAlbum, maxAlbum) {
			continue
		}

		// 3) Check number of members
		if len(membersSelected) > 0 && !matchesMembers(len(artist.Members), membersSelected) {
			continue
		}

		// 4) Check location
		if len(locationsSelected) > 0 && !matchesLocation(artist.ID, locationsSelected) {
			continue
		}

		// If all checks pass, include this artist
		results = append(results, artist)
	}

	return results
}

// matchesCreationYear verifies if creationDate is within [min, max], ignoring if either bound is 0.
func matchesCreationYear(creationDate, min, max int) bool {

	if min > 0 && creationDate < min {
		return false
	}
	if max > 0 && creationDate > max {
		return false
	}
	return true
}

// matchesAlbumYear extracts the year from the artist's FirstAlbum string.
func matchesAlbumYear(firstAlbum string, min, max int) bool {

	parts := strings.Split(firstAlbum, "-")

	yearStr := parts[2]
	year, err := strconv.Atoi(yearStr)

	if err != nil {
		// If parse fails, treat it as no restriction
		return false
	}
	if min > 0 && year < min {
		return false
	}
	if max > 0 && year > max {
		return false
	}

	return true
}

// matchesMembers returns true if the band's member count matches ANY of the selected checkboxes.
func matchesMembers(memberCount int, membersSelected []string) bool {
	countStr := strconv.Itoa(memberCount)
	for _, sel := range membersSelected {
		if sel == countStr {
			return true
		}
	}
	return false
}

// matchesLocation checks if the band has a location that CONTAINS the user-chosen location(s).
// Some user-locations might be "Washington, USA" and the actual location might be "Seattle, Washington, USA".
func matchesLocation(artistID int, locationsSelected []string) bool {
	// Find this artist's dateLocations map from data.AllRelations
	var datesLocations map[string][]string
	for _, rel := range data.AllRelations.Index {
		if rel.ID == artistID {
			datesLocations = rel.DatesLocations
			break
		}
	}
	if datesLocations == nil {
		return false
	}

	// Normalize all selected user values
	for i, loc := range locationsSelected {
		locationsSelected[i] = strings.ToLower(utils.FormatLocation(loc))
	}

	// Compare formatted/normalized backend location to each user selection
	for loc := range datesLocations {
		formatted := strings.ToLower(utils.FormatLocation(loc))
		for _, selected := range locationsSelected {
			if strings.Contains(formatted, selected) {
				return true
			}
		}
	}

	return false
}
