package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"groupie-tracker-search-bar/internal/models"
)

const mapboxToken = "pk.eyJ1IjoiYXRob29oIiwiYSI6ImNtMWY2N3prZjJsN3MybHNjMWd3bThzOXcifQ.HNgAHQBkzGdrnuS1MtwYlQ"

type GeocodeResponse struct {
	Features []struct {
		Center []float64 `json:"center"`
	} `json:"features"`
}

func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

// In-memory cache for geocoding results
var geocodeCache = make(map[string][]float64)
var cacheMutex sync.RWMutex

// formatLocation takes a location string like 'auckland-new_zealand'
// and returns a properly formatted string 'auckland, new zealand'.
func formatLocation(location string) string {
	formattedLocation := strings.ReplaceAll(location, "-", ", ")        // Replace hyphens with ', '
	formattedLocation = strings.ReplaceAll(formattedLocation, "_", " ") // Replace underscores with spaces
	return formattedLocation
}

func FetchAllData() ([]models.Artist, models.LocationsData, models.DatesData, models.RelationsData, error) {
	var artists []models.Artist
	var locationsData models.LocationsData
	var datesData models.DatesData
	var relationsData models.RelationsData

	err := FetchData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	if err != nil {
		return nil, models.LocationsData{}, models.DatesData{}, models.RelationsData{}, err
	}
	err = FetchData("https://groupietrackers.herokuapp.com/api/locations", &locationsData)
	if err != nil {
		return nil, models.LocationsData{}, models.DatesData{}, models.RelationsData{}, err
	}
	err = FetchData("https://groupietrackers.herokuapp.com/api/dates", &datesData)
	if err != nil {
		return nil, models.LocationsData{}, models.DatesData{}, models.RelationsData{}, err
	}
	err = FetchData("https://groupietrackers.herokuapp.com/api/relation", &relationsData)
	if err != nil {
		return nil, models.LocationsData{}, models.DatesData{}, models.RelationsData{}, err
	}

	return artists, locationsData, datesData, relationsData, nil
}

// GeocodeLocation takes a raw location, formats it, and returns geographic coordinates.
// Implements caching to avoid redundant requests.
func GeocodeLocation(rawLocation string) ([]float64, error) {
	formattedLocation := formatLocation(rawLocation)

	// Check if the location is cached
	cacheMutex.RLock()
	if cachedCoords, found := geocodeCache[formattedLocation]; found {
		cacheMutex.RUnlock()
		return cachedCoords, nil
	}
	cacheMutex.RUnlock()

	var geocodeResponse GeocodeResponse
	var err error

	// Retry logic in case of network issues
	for attempts := 0; attempts < 3; attempts++ {
		geocodeURL := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s", url.QueryEscape(formattedLocation), mapboxToken)
		resp, err := http.Get(geocodeURL)
		if err == nil && resp.StatusCode == 200 {
			defer resp.Body.Close()
			err = json.NewDecoder(resp.Body).Decode(&geocodeResponse)
			if err == nil && len(geocodeResponse.Features) > 0 {
				coords := geocodeResponse.Features[0].Center

				// Cache the coordinates for future use
				cacheMutex.Lock()
				geocodeCache[formattedLocation] = coords
				cacheMutex.Unlock()

				return coords, nil
			}
		}

		// Wait before retrying
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to geocode location %s after 3 attempts: %v", formattedLocation, err)
}
