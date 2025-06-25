package main

import (
	"fmt"
	groupieapi "groupie-tracker/internal/groupieAPI"
)

func main() {
	index, err := groupieapi.IndexArtists()
	fmt.Println(index, err)
}
