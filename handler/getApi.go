package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ArtistsLink   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsLink = "https://groupietrackers.herokuapp.com/api/locations"
	DatesLink     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationsLink = "https://groupietrackers.herokuapp.com/api/relation"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}
type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type PageData struct {
	Artist    []Artist
	Relations []Relation
	Dates     []Dates
	Locations []Locations
}

func FetchArtists() ([]Artist, error) {
	resp, err := http.Get(ArtistsLink)
	if err != nil {
		return nil, fmt.Errorf("cannot get artists: %w", err)
	}
	defer resp.Body.Close()

	var artists []Artist
	if err := json.NewDecoder(resp.Body).Decode(&artists); err != nil {
		return nil, fmt.Errorf("cannot fetch artists: %w", err)
	}
	return artists, nil
}

type LocationsResponse struct {
	Index []Locations `json:"index"`
}

func FetchLocations() ([]Locations, error) {
	resp, err := http.Get(LocationsLink)
	if err != nil {
		return nil, fmt.Errorf("cannot get locations: %w", err)
	}
	defer resp.Body.Close()

	var locResp LocationsResponse
	if err := json.NewDecoder(resp.Body).Decode(&locResp); err != nil {
		return nil, fmt.Errorf("cannot fetch locations: %w", err)
	}
	return locResp.Index, nil
}

type DatesResponse struct {
	Index []Dates `json:"index"`
}

func FetchDates() ([]Dates, error) {
	resp, err := http.Get(DatesLink)
	if err != nil {
		return nil, fmt.Errorf("cannot get dates: %w", err)
	}
	defer resp.Body.Close()
	var datesResp DatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&datesResp); err != nil {
		return nil, fmt.Errorf("cannot fetch dates: %w", err)
	}
	return datesResp.Index, nil
}

type RelationResponse struct {
	Index []Relation `json:"index"`
}

func FetchRelations() ([]Relation, error) {
	resp, err := http.Get(RelationsLink)
	if err != nil {
		return nil, fmt.Errorf("cannot get relations: %w", err)
	}
	defer resp.Body.Close()

	var relResp RelationResponse
	if err := json.NewDecoder(resp.Body).Decode(&relResp); err != nil {
		return nil, fmt.Errorf("cannot fetch relations: %w", err)
	}
	return relResp.Index, nil
}

func GetPageData() (PageData, error) {
	var data PageData
	var err error

	data.Artist, err = FetchArtists()
	if err != nil {
		return data, err
	}

	data.Locations, err = FetchLocations()
	if err != nil {
		return data, err
	}

	data.Dates, err = FetchDates()
	if err != nil {
		return data, err
	}

	data.Relations, err = FetchRelations()
	if err != nil {
		return data, err
	}

	return data, nil
}
