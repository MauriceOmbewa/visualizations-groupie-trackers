package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GeocodeResult represents the structure for the geocoding response
type GeocodeResult struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// GeocodeLocation takes a URL, location, and fetches the geocode data
func GeocodeLocation(apiURL, location string) (GeocodeResult, error) {
	apiKey := "pk.eyJ1IjoiYXRob29oIiwiYSI6ImNtMWY2N3prZjJsN3MybHNjMWd3bThzOXcifQ.HNgAHQBkzGdrnuS1MtwYlQ" // Replace with your actual API key
	url := fmt.Sprintf("%s?location=%s&key=%s", apiURL, location, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return GeocodeResult{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GeocodeResult{}, err
	}

	var result GeocodeResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return GeocodeResult{}, err
	}

	return result, nil
}
