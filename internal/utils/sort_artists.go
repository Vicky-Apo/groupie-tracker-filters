package utils

import (
	"groupie-tracker/internal/data"
	"sort"
	"strings"
)

func SortingArtists(artists []data.Artist) []data.Artist {

	sortedArtists := make([]data.Artist, len(artists))
	copy(sortedArtists, artists)
	sort.Slice(sortedArtists, func(i, j int) bool {
		return strings.ToLower(sortedArtists[i].Name) < strings.ToLower(sortedArtists[j].Name)
	})
	return sortedArtists

}
