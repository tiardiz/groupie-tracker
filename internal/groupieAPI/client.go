package groupieapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	ArtistsLink   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsLink = "https://groupietrackers.herokuapp.com/api/locations"
	DatesLink     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationsLink = "https://groupietrackers.herokuapp.com/api/relation"
)

func GetFromAPI[T availableAPI](id int) (result T, err error) {
	var link string // link goes here
	switch any(result).(type) {
	case Artist:
		link = ArtistsLink
	case ArtistLocations:
		link = LocationsLink
	case ArtistDates:
		link = DatesLink
	case Relation:
		link = RelationsLink
	default:
		return result, fmt.Errorf("unsupported type for API call")
	}
	link = link + "/" + strconv.Itoa(id)
	data, err := getData(link)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal([]byte(data), &result)
	return
}

func GetArtist(id int) (Artist, error) {
	return GetFromAPI[Artist](id)
}

func IndexArtists() (string, error) {
	return getData(ArtistsLink)
}

func IndexDates() (string, error) {
	return getData(DatesLink)
}

func IndexRelations() (string, error) {
	return getData(RelationsLink)
}

func IndexLocations() (string, error) {
	return getData(LocationsLink)
}

func getData(link string) (string, error) {
	response, err := http.Get(link)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	result, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type ArtistLocations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type ArtistDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type availableAPI interface {
	Artist | ArtistLocations | ArtistDates | Relation
}
