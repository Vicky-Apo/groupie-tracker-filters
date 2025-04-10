package data

// Artist structure from /api/artists
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relation     string   `json:"relations"` // URL to the "relation" endpoint for this artist
}

// LocationsIndex structure from /api/locations
type LocationsIndex struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

// DatesIndex structure from /api/dates
type DatesIndex struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

// RelationIndex structure from /api/relation
type RelationIndex struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// CombinedData is a helper struct to show all data for a single artist
// (artist, locations, dates, and relation mapping).
type CombinedData struct {
	Artist      Artist
	Locations   []string
	Dates       []string
	DatesLocMap map[string][]string
}
