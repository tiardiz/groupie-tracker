package functions

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

func GetApi() {

}
