package handler

import (
	// "html/template"

	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("templates/info.html"))

	path := r.URL.Path // "/info/1"

	parts := strings.Split(path, "/") // ["", "info", "1"]
	if len(parts) < 3 {
		NotFoundHandler(w, r)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		NotFoundHandler(w, r)
		return
	}

	var selected *Artist
	for _, a := range artists {
		if a.ID == id {
			selected = &a
			break
		}
	}
	if selected == nil {
		NotFoundHandler(w, r)
		return
	}
	err = tmpl.Execute(w, selected)
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}
