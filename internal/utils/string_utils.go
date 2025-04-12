package utils

import "strings"

var knownAbbreviations = map[string]bool{
	"USA": true,
	"UK":  true,
	"UAE": true,
	"EU":  true,
}

func FormatLocation(s string) string {
	// Replace dashes and underscores with spaces
	s = strings.ReplaceAll(s, "-", ", ")
	s = strings.ReplaceAll(s, "_", " ")

	words := strings.Fields(s)
	for i, word := range words {
		upper := strings.ToUpper(word)
		if knownAbbreviations[upper] {
			words[i] = upper
		} else {
			if len(word) > 1 {
				words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
			} else {
				words[i] = strings.ToUpper(word)
			}
		}
	}
	return strings.Join(words, " ")
}

// ReplaceSpaces replaces spaces with dashes in the given string.
func ReplaceSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "-")
}

// CleanDate removes all asterisk (*) characters from a date string.
func CleanDate(date string) string {
	cleaned := strings.ReplaceAll(date, "*", "")
	parts := strings.Fields(cleaned)

	if len(parts) == 3 {
		return parts[0] + "-" + parts[1] + "-" + parts[2]
	}
	return cleaned
}

// NormalizeQuery converts user input to the same format as stored data (e.g. "New York" -> "New_York")
func NormalizeQuery(s string) string {
	// Only replace spaces with underscores if there are no digits (i.e. it's not a date)
	if strings.ContainsAny(s, "0123456789") {
		return s
	}
	return strings.ReplaceAll(s, " ", "_")
}
func NormalizeString(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.Join(strings.Fields(s), " ") // trims & normalizes spacing
	return s
}
