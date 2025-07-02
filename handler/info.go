package handler

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		MethodNotAllowedHandler(w, r, errors.New(r.Method+"method not allowed"))
		return
	}
	tmpl := template.Must(template.ParseFiles("templates/info.html"))

	path := r.URL.Path // "/info/1"

	parts := strings.Split(path, "/") // ["", "info", "1"]
	if len(parts) < 3 || parts[2] == "" {
		err := errors.New("page not found")
		NotFoundHandler(w, r, err)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		NotFoundHandler(w, r, err)
		return
	}

	pageData, err := GetPageData()
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}

	var selected *Artist
	for i, a := range pageData.Artist {
		if a.ID == id {
			selected = &pageData.Artist[i]
			break
		}
	}

	if selected == nil {
		err := errors.New("page not found")
		NotFoundHandler(w, r, err)
		return
	}

	// relation
	var relForArtist []Relation
	for _, rel := range pageData.Relations {
		if rel.ID == selected.ID {
			relForArtist = append(relForArtist, rel)
		}
	}

	// locations
	locForArtistSlice := []Locations{}
	for _, loc := range pageData.Locations {
		if loc.ID == selected.ID {
			locForArtistSlice = append(locForArtistSlice, loc)
		}
	}
	// Dates
	dateForArtistSlice := []Dates{}
	for _, date := range pageData.Dates {
		if date.ID == selected.ID {
			dateForArtistSlice = append(dateForArtistSlice, date)
		}
	}

	infoData := struct {
		Artist    *Artist
		Relations []Relation
		Locations []Locations
		Dates     []Dates
	}{
		Artist:    selected,
		Relations: relForArtist,
		Locations: locForArtistSlice,
		Dates:     dateForArtistSlice,
	}

	err = tmpl.Execute(w, infoData)
	if err != nil {
		InternalServerErrorHandler(w, r, err)
	}
}
