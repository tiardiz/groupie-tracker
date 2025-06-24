package main

import (
	"fmt"
	groupieapi "groupie-tracker/internal/groupieAPI"
)

func main() {
	fmt.Println(groupieapi.GetArtistsJson())
}
