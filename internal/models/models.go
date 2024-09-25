package models

import (
	"strconv"
	"strings"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
}

type LocationsData struct {
	Index []Location `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DatesData struct {
	Index []Date `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationsData struct {
	Index []Relation `json:"index"`
}

type ArtistDetail struct {
	Artist    Artist
	Locations Location
	Dates     Date
	Relations Relation
}

type ErrorDetail struct {
	Title   string
	Message string
}

func (a Artist) SearchResultType(query string) []string {
	var resultTypes []string
	lowerQuery := strings.ToLower(query)

	// Check if query matches the artist/band name
	if strings.Contains(strings.ToLower(a.Name), lowerQuery) {
		resultTypes = append(resultTypes, a.Name+" - artist/band")
	}

	// Check if query matches any band member's name
	for _, member := range a.Members {
		if strings.Contains(strings.ToLower(member), lowerQuery) {
			resultTypes = append(resultTypes, member+" - member of "+a.Name)
		}
	}

	// Check if query matches the first album name
	if strings.Contains(strings.ToLower(a.FirstAlbum), lowerQuery) {
		resultTypes = append(resultTypes, a.FirstAlbum+" - first album date of "+a.Name)
	}

	// Check if query matches the creation date
	if strings.Contains(strconv.Itoa(a.CreationDate), lowerQuery) {
		resultTypes = append(resultTypes, strconv.Itoa(a.CreationDate)+" - creation date of "+a.Name)
	}

	return resultTypes
}

func SearchArtists(query string, artists []Artist) []string {
	var suggestions []string

	for _, artist := range artists {
		resultTypes := artist.SearchResultType(query)
		suggestions = append(suggestions, resultTypes...)
	}

	return suggestions
}

func (r Relation) SearchArtistsByLocation(query string, artists []Artist) []string {
	var results []string
	query = strings.ToLower(query)

	// Search for the location in the DatesLocations map
	for location, _ := range r.DatesLocations {
		if strings.Contains(strings.ToLower(location), query) {
			// Get the artist related to this location
			for _, artist := range artists {
				if artist.ID == r.ID {
					results = append(results, artist.Name + " - " + location)
					break
				}
			}
		}
	}
	return results
}
