package handlers

import (
	"groupie-tracker/api"
	"groupie-tracker/models"
	"html/template"
	"net/http"
)

var (
	tmpl = template.Must(template.ParseFiles("web/templates/index.html"))
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusMethodNotAllowed, "")
		return
	}

	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		ErrorHandler(w, http.StatusNotFound, "")
		return
	}

	artists, err := api.GetArtistApi(w)
	if err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "")
		return
	}

	data := struct {
		Artist []models.Artist
	}{
		Artist: artists,
	}

	if err := tmpl.Execute(w, data); err != nil {
		ErrorHandler(w, http.StatusInternalServerError, "error rendering template")
		return
	}
}
