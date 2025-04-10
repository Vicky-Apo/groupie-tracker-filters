package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

// Global variables to hold the data after fetching
var (
	AllArtists   []Artist
	AllLocations LocationsIndex
	AllDates     DatesIndex
	AllRelations RelationIndex
)

// fetchJSON is a helper to fetch JSON from a URL and decode into target
func fetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(target); err != nil {
		return err
	}
	return nil
}

// LoadData fetches all data from the provided API endpoints concurrently
func LoadData() error {
	var wg sync.WaitGroup
	var errArtists, errLocations, errDates, errRelations error

	wg.Add(4)

	// /api/artists
	go func() {
		defer wg.Done()
		errArtists = fetchJSON("https://groupietrackers.herokuapp.com/api/artists", &AllArtists)
	}()

	// /api/locations
	go func() {
		defer wg.Done()
		errLocations = fetchJSON("https://groupietrackers.herokuapp.com/api/locations", &AllLocations)
	}()

	// /api/dates
	go func() {
		defer wg.Done()
		errDates = fetchJSON("https://groupietrackers.herokuapp.com/api/dates", &AllDates)
	}()

	// /api/relation
	go func() {
		defer wg.Done()
		errRelations = fetchJSON("https://groupietrackers.herokuapp.com/api/relation", &AllRelations)
	}()

	wg.Wait()

	if errArtists != nil {
		return errArtists
	}
	if errLocations != nil {
		return errLocations
	}
	if errDates != nil {
		return errDates
	}
	if errRelations != nil {
		return errRelations
	}

	return nil
}
