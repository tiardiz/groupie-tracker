package groupieapi

type artistDataBundle struct {
	Artist
	Locations ArtistLocations
	Dates     ArtistDates
	Relations Relation
}

func BundleArtistData(id int) (artistDataBundle, chan error) {
	artistChan := make(chan Artist)
	locationsChan := make(chan ArtistLocations)
	datesChan := make(chan ArtistDates)
	relationChan := make(chan Relation)
	errChan := make(chan error, 4)
	go func() {
		artist, err := GetArtist(id)
		if err != nil {
			errChan <- err
		}
		artistChan <- artist
	}()
	go func() {
		locations, err := GetLocations(id)
		if err != nil {
			errChan <- err
		}
		locationsChan <- locations
	}()
	go func() {
		dates, err := GetDates(id)
		if err != nil {
			errChan <- err
		}
		datesChan <- dates
	}()
	go func() {
		relations, err := GetRelation(id)
		if err != nil {
			errChan <- err
		}
		relationChan <- relations
	}()
	return artistDataBundle{<-artistChan, <-locationsChan, <-datesChan, <-relationChan}, errChan
}
