package handlers

import (
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
	if err != nil {
		HandleError(w, 404, "Artist not found")
		return
	}
	data, errorChannel := groupieapi.BundleArtistData(id)
	if groupieapi.ArtistNotFound(data.Artist) {
		HandleError(w, 404, "Artist not found")
		return
	}
	for len(errorChannel) > 0 {
		err = <-errorChannel
		if err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		HandleError(w, 500, "Internal server error")
		return
	}
	if err = artistTemplate.Execute(w, data); err != nil {
		log.Println(err)
		HandleError(w, 500, "Internal server error")
		return
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
