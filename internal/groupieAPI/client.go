package groupieapi

import (
	"io"
	"net/http"
)

const (
	Artists   = "https://groupietrackers.herokuapp.com/api/artists"
	Locations = "https://groupietrackers.herokuapp.com/api/locations"
	Dates     = "https://groupietrackers.herokuapp.com/api/dates"
	Relation  = "https://groupietrackers.herokuapp.com/api/relation"
)

func GetArtistsJson() (string, error) {
	return GetAPI(Artists)
}

func GetDatesJson() (string, error) {
	return GetAPI(Dates)
}

func GetRelationJson() (string, error) {
	return GetAPI(Relation)
}

func GetLocationsJson() (string, error) {
	return GetAPI(Locations)
}

func GetAPI(link string) (string, error) {
	response, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return string(result), nil
}
