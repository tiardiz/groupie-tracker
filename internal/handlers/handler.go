package handlers

import (
	"errors"
	groupieapi "groupie-tracker/internal/groupieAPI"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		HandleError(w, 404, "Page not found")
		return
	}
	artists, err := groupieapi.IndexArtists()
	if err != nil {
		log.Println(err)
		HandleError(w, 500, "Internal server error")
		return
	}
	mainPageTemplate.Execute(w, artists)
}

func HandleArtist(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	if !(len(parts) == 2 && parts[0] == "artist") {
		HandleError(w, 404, "Not found")
		return
	}
	idString := parts[1]
	id, err := strconv.Atoi(idString)
	if errors.Is(err, strconv.ErrRange) {
		HandleError(w, 404, "Artist not found")
		return
	} else if errors.Is(err, strconv.ErrSyntax) {
		HandleError(w, 400, "Artist ID must be a number")
		return
	}

	data, errors := groupieapi.BundleArtistData(id)
	for len(errors) > 0 {
		err = <-errors
		if err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		HandleError(w, 404, "Artist not found")
		return
	}
	if err = artistTemplate.Execute(w, data); err != nil {
		log.Println(err)
	}
}

func HandleError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	err := errorPageTemplate.Execute(w, struct {
		ErrorCode    int
		ErrorMessage string
	}{
		code, msg,
	})
	if err != nil {
		log.Println(err)
	}
}

const (
	errorTemplatePath  = "frontend/templates/error.html"
	indexTemplatePath  = "frontend/templates/index.html"
	artistTemplatePath = "frontend/templates/artist.html"
)

var (
	errorPageTemplate = template.Must(template.ParseFiles(errorTemplatePath))
	mainPageTemplate  = template.Must(template.ParseFiles(indexTemplatePath))
	artistTemplate    = template.Must(template.ParseFiles(artistTemplatePath))
)
