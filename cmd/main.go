package main

import (
	"fmt"
	groupieapi "groupie-tracker/internal/groupieAPI"
)

func main() {
	artist, _ := groupieapi.GetArtist(5)
	location, _ := groupieapi.GetLocations(5)
	fmt.Printf("%#v\n%#v\n", artist, location)
}
