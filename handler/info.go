package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func InfoHandler(w http.ResponseWriter, r *http.Request) {

	var tmpl = template.Must(template.ParseFiles("templates/info.html"))

	path := r.URL.Path // "/info/1"

	parts := strings.Split(path, "/") // ["", "info", "1"]
	if len(parts) < 3 || parts[2] == "" {
		NotFoundHandler(w, r)
		return
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		NotFoundHandler(w, r)
		return
	}

	pageData, err := GetPageData()
	if err != nil {
		fmt.Println("error here 3", err)
		InternalServerErrorHandler(w, r)
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
		NotFoundHandler(w, r)
		return
	}

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

	fmt.Println(infoData.Locations)

	err = tmpl.Execute(w, infoData)
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}
