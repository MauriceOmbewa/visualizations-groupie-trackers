package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock API response
var mockAPIResponse = `{
	"lat": -1.2921,
	"lng": 36.8219
}`

// TestGeocodeLocation tests the GeocodeLocation function
func TestGeocodeLocation(t *testing.T) {
	// Create a mock server that returns the mock API response
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockAPIResponse))
	}))
	defer mockServer.Close()

	// Test the function with the mock server URL
	location := "Nairobi"
	expectedResult := GeocodeResult{
		Latitude:  -1.2921,
		Longitude: 36.8219,
	}

	// Call the GeocodeLocation function using the mock server URL
	result, err := GeocodeLocation(mockServer.URL, location)

	// Check for unexpected errors
	if err != nil {
		t.Fatalf("expected no error, but got %v", err)
	}

	// Compare the result with the expected result
	if result != expectedResult {
		t.Errorf("expected %v, but got %v", expectedResult, result)
	}
}
