package router

import (
	"groupie-tracker/internal/handlers"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/css/", http.StripPrefix("/static/css/", http.FileServer(http.Dir("web/templates/static/css"))))
	//регистрируем два обработчика
	mux.HandleFunc("/", handlers.HandleIndex)
	mux.HandleFunc("/artist", handlers.HandleArtist)
	return mux
}
