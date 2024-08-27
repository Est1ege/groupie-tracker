package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/models"
	"net/http"
	"strconv"
	"text/template"
)

var (
	artistTmpl = template.Must(template.ParseFiles("web/templates/artist.html"))
)

func HandleArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed, "")
		return
	}

	// strID := strings.TrimPrefix(r.URL.Path, "/artist/")
	// if r.Method != http.MethodGet {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	return
	// }

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		ErrorHandler(w, http.StatusBadRequest, "missing id parametrs")
		return
	}

	id, _ := strconv.Atoi(idStr)
	if id < 1 || id > 52 {
		ErrorHandler(w, http.StatusBadRequest, "invalid id param")
		return
	}

	// art := len(id)-1

	artist, err := api.GetArtistById(int64(id), w)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound, "artist not found")
		return
	}

	locations, err := api.GetLocationsApi(int64(id), w)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound, "location not found")
		return
	}

	concertDates, err := api.GetConcertDates(int64(id), w)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound, "date not found")
		return
	}

	relations, err := api.GetRelationsApi(int64(id), w)
	if err != nil {
		ErrorHandler(w, http.StatusNotFound, "relations not found")
		return
	}

	data := models.ArtistTemplate{
		Artist:       artist,
		Locations:    locations,
		ConcertDates: concertDates,
		Relations:    relations,
	}
	if err := artistTmpl.Execute(w, data); err != nil {
		ErrorHandler(w, http.StatusNotFound, "execute error")
		return
	}
}
