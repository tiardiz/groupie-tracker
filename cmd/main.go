package main

import (
	"fmt"
	groupieapi "groupie-tracker/internal/groupieAPI"
)

func main() {
	artist, _ := groupieapi.GetFromAPI[groupieapi.ArtistDates](5)
	location, _ := groupieapi.GetFromAPI[groupieapi.ArtistLocations](5)
	fmt.Printf("%#v\n%#v\n", artist, location)
}
