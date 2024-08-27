package api

import (
	"encoding/json"
	"fmt"
	"groupie-tracker/models"
	"io"
	"net/http"
)

func GetArtistApi(w http.ResponseWriter) ([]models.Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artists []models.Artist
	err := GetData(url, &artists, w)
	return artists, err
}

func GetArtistById(id int64, w http.ResponseWriter) (*models.Artist, error) {
	artists, err := GetArtistApi(w)
	for _, artist := range artists {
		if artist.ID == id {
			return &artist, err
		}
	}
	return nil, err
}

func GetLocationsApi(id int64, w http.ResponseWriter) (models.Locations, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)
	var locations models.Locations
	err := GetData(url, &locations, w)
	return locations, err
}

func GetRelationsApi(id int64, w http.ResponseWriter) (models.Relations, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%d", id)
	var relations models.Relations
	err := GetData(url, &relations, w)
	return relations, err
}

func GetConcertDates(id int64, w http.ResponseWriter) (models.ConcertDates, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%d", id)
	var concertdates models.ConcertDates
	err := GetData(url, &concertdates, w)
	return concertdates, err
}

func GetData(url string, value interface{}, w http.ResponseWriter) error {
	responce, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка чтения урл:", err)
		return err
	}
	defer responce.Body.Close()

	body, err := io.ReadAll(responce.Body)
	if err != nil {
		fmt.Println("Ошибка джейсон:", err)
		return err
	}

	err = json.Unmarshal(body, &value)
	if err != nil {
		return err
	}
	return nil
}
