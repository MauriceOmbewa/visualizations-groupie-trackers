package fetch

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"groupie-tracker-search-bar/internal/models"
)

func TestFetchData(t *testing.T) {
	mockResponse := `[
		{"id": 1, "name": "Artist 1"},
		{"id": 2, "name": "Artist 2"}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	var artists []models.Artist
	err := FetchData(server.URL, &artists)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(artists) != 2 {
		t.Fatalf("expected 2 artists, got %v", len(artists))
	}
}

func TestFetchAllData(t *testing.T) {
	_, _, _, _, err := FetchAllData()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
