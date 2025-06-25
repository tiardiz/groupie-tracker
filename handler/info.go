package handler

import (
	"html/template"
	"net/http"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		MethodNotAllowedHandler(w, r)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/info.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		InternalServerErrorHandler(w, r)
	}
}
