package handlers

import (
	"html/template"
	"net/http"
)


var tmplErr = template.Must(template.ParseFiles("web/templates/error.html"))

type ErrorData struct {
	Code int
	Message string
}

func ErrorHandler(w http.ResponseWriter, code int, message string ) {
	w.WriteHeader(code)

	data := ErrorData{
		Code: code,
		Message: message,
	}

	tmplErr.Execute(w, data)


}