package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"groupie-tracker-search-bar/internal/fetch"
	"groupie-tracker-search-bar/internal/models"
)

var (
	artists       []models.Artist
	locationsData models.LocationsData
	datesData     models.DatesData
	relationsData models.RelationsData
)

var templateFuncs = template.FuncMap{
	"Join": strings.Join, // Register the Join function
}

var templates = template.Must(template.New("").Funcs(templateFuncs).ParseGlob("templates/*.html"))

// InitData loads data when the application starts
func InitData() error {
	var err error
	artists, locationsData, datesData, relationsData, err = fetch.FetchAllData()
	return err
}

// RenderError displays a custom error page with status code and message
func RenderError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	errDetail := models.ErrorDetail{
		Title:   http.StatusText(status),
		Message: message,
	}
	err := templates.ExecuteTemplate(w, strconv.Itoa(status)+".html", errDetail)
	if err != nil {
		http.Error(w, "An error occurred while rendering the error page", http.StatusInternalServerError)
	}
}

// IndexHandler handles requests to the home page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Error loading the homepage")
	}
}

// ArtistsHandler handles requests to the artists listing page
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the data was successfully loaded (e.g., network issues)
	if len(artists) == 0 {
		RenderError(w, http.StatusInternalServerError, "Failed to load artist data. Please check your internet connection.")
		return
	}

	err := templates.ExecuteTemplate(w, "artists.html", artists)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Error loading the artists page")
	}
}

// ArtistDetailHandler handles requests to individual artist detail pages
func ArtistDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		RenderError(w, http.StatusBadRequest, "Invalid artist ID")
		return
	}

	index := id - 1
	if index < 0 || index >= len(artists) {
		RenderError(w, http.StatusNotFound, "Artist not found")
		return
	}

	// Fetch the artist details
	artistDetail := models.ArtistDetail{
		Artist:    artists[index],
		Locations: locationsData.Index[index],
		Dates:     datesData.Index[index],
		Relations: relationsData.Index[index],
	}

	// Handle geocoding failures gracefully by stopping if there's no internet
	type ConcertLocation struct {
		LocationName string    `json:"locationName"`
		Coordinates  []float64 `json:"coordinates"`
	}
	var concertLocations []ConcertLocation
	geocodeChan := make(chan ConcertLocation)
	errorChan := make(chan error)

	// Launch concurrent geocoding tasks
	for _, location := range artistDetail.Locations.Locations {
		go func(loc string) {
			coords, err := fetch.GeocodeLocation(loc)
			if err != nil {
				errorChan <- err
				return
			}
			geocodeChan <- ConcertLocation{LocationName: loc, Coordinates: coords}
		}(location)
	}

	// Collect geocoded results or errors
	for range artistDetail.Locations.Locations {
		select {
		case geocode := <-geocodeChan:
			concertLocations = append(concertLocations, geocode)
		case err := <-errorChan:
			log.Printf("Error geocoding location: %v", err)
			RenderError(w, http.StatusInternalServerError, "Error geocoding concert locations. Please check your internet connection.")
			return
		}
	}

	// Convert concert locations to JSON
	concertLocationsJSON, err := json.Marshal(concertLocations)
	if err != nil {
		log.Printf("Error marshaling concert locations: %v", err)
		RenderError(w, http.StatusInternalServerError, "Error processing concert locations")
		return
	}

	// Pass the concert locations to the template
	data := struct {
		ArtistDetail         models.ArtistDetail
		ConcertLocationsJSON string
	}{
		ArtistDetail:         artistDetail,
		ConcertLocationsJSON: string(concertLocationsJSON),
	}

	// Render the artist detail page
	err = templates.ExecuteTemplate(w, "artist_detail.html", data)
	if err != nil {
		RenderError(w, http.StatusInternalServerError, "Error loading the artist detail page")
	}
}

// SearchHandler handles search requests for artists
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))
	var results []map[string]string

	if query != "" {
		for _, artist := range artists {
			types := artist.SearchResultType(query)
			for _, resultType := range types {
				results = append(results, map[string]string{
					"name": resultType,
					"id":   strconv.Itoa(artist.ID),
				})
			}
		}

		// Search by location
		for _, relation := range relationsData.Index {
			locationResults := relation.SearchArtistsByLocation(query, artists)
			for _, result := range locationResults {
				results = append(results, map[string]string{
					"name": result,
					"id":   strconv.Itoa(relation.ID),
				})
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
