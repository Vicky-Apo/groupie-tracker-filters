package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/data"
	"groupie-tracker/internal/utils"
)

type SearchResult struct {
	Value  string `json:"Value"`
	Type   string `json:"Type"`
	Artist string `json:"Artist"`
	ID     int    `json:"ID"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("query"))
	query = utils.NormalizeQuery(query)
	if query == "" {
		json.NewEncoder(w).Encode([]SearchResult{})
		return
	}

	var results []SearchResult
	seen := make(map[string]bool)

	// Match by artist, members, dates
	for _, artist := range data.AllArtists {
		// Artist name
		if strings.HasPrefix(strings.ToLower(artist.Name), query) {
			key := artist.Name + "_Artist"
			if !seen[key] {
				results = append(results, SearchResult{
					Value:  artist.Name + " — Artist/Band",
					Type:   "artist/band",
					Artist: artist.Name,
					ID:     artist.ID,
				})
				seen[key] = true
			}
		}

		// Members
		for _, member := range artist.Members {
			if strings.HasPrefix(strings.ToLower(member), query) {
				key := member + "_Member"
				if !seen[key] {
					results = append(results, SearchResult{
						Value:  member + " — Member",
						Type:   "member",
						Artist: artist.Name,
						ID:     artist.ID,
					})
					seen[key] = true
				}
				// Also add band
				bandKey := artist.Name + "_Artist"
				if !seen[bandKey] {
					results = append(results, SearchResult{
						Value:  artist.Name + " — Artist/Band",
						Type:   "artist/band",
						Artist: artist.Name,
						ID:     artist.ID,
					})
					seen[bandKey] = true
				}
			}
		}

		// Creation date
		creationStr := strconv.Itoa(artist.CreationDate)
		if strings.HasPrefix(creationStr, query) {
			key := creationStr + "_Creation"
			if !seen[key] {
				results = append(results, SearchResult{
					Value:  creationStr + " — Creation Date",
					Type:   "creation date",
					Artist: artist.Name,
					ID:     artist.ID,
				})
				seen[key] = true
			}
			artistKey := artist.Name + "_Artist_Creation"
			if !seen[artistKey] {
				results = append(results, SearchResult{
					Value:  artist.Name + " — Artist/Band",
					Type:   "artist/band",
					Artist: artist.Name,
					ID:     artist.ID,
				})
				seen[artistKey] = true
			}
		}

		// First album date (exact match, can be mid-string)
		cleanedAlbumDate := utils.CleanDate(artist.FirstAlbum)
		if strings.Contains(strings.ToLower(cleanedAlbumDate), query) {
			key := cleanedAlbumDate + "_First_Album"
			if !seen[key] {
				results = append(results, SearchResult{
					Value:  cleanedAlbumDate + " — First Album Date",
					Type:   "first album date",
					Artist: artist.Name,
					ID:     artist.ID,
				})
				seen[key] = true
			}
			artistKey := artist.Name + "_Artist_FirstAlbum"
			if !seen[artistKey] {
				results = append(results, SearchResult{
					Value:  artist.Name + " — Artist/Band",
					Type:   "artist/band",
					Artist: artist.Name,
					ID:     artist.ID,
				})
				seen[artistKey] = true
			}
		}
	}

	// Match by location (outside artist loop)
	for _, rel := range data.AllRelations.Index {
		for location := range rel.DatesLocations {
			if len(query) >= 3 && strings.Contains(strings.ToLower(location), query) {
				locKey := location + "_Location"
				artist := findArtistByID(rel.ID)
				if artist != nil {
					if !seen[locKey] {
						results = append(results, SearchResult{
							Value:  utils.FormatLocation(location) + " — Location",
							Type:   "location",
							Artist: artist.Name,
							ID:     artist.ID,
						})
						seen[locKey] = true
					}
					for _, date := range rel.DatesLocations[location] {
						cleanedDate := utils.CleanDate(date)
						dateKey := location + cleanedDate + "_ConcertDate"

						if strings.Contains(strings.ToLower(cleanedDate), query) && !seen[dateKey] {
							results = append(results, SearchResult{
								Value:  cleanedDate + " — Concert Date",
								Type:   "concert date",
								Artist: artist.Name,
								ID:     artist.ID,
							})
							seen[dateKey] = true
						}
					}

					artistKey := artist.Name + "_Artist_Location"
					if !seen[artistKey] {
						results = append(results, SearchResult{
							Value:  artist.Name + " — Artist/Band",
							Type:   "artist/band",
							Artist: artist.Name,
							ID:     artist.ID,
						})
						seen[artistKey] = true
					}
				}
			}
		}
	}

	json.NewEncoder(w).Encode(results)
}

func findArtistByID(id int) *data.Artist {
	for _, artist := range data.AllArtists {
		if artist.ID == id {
			return &artist
		}
	}
	return nil
}
