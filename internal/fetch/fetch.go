package fetch

import (
	"encoding/json"
	"net/http"

	"github.com/Athooh/groupie-tracker/internal/models"
)

func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
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
