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
	go asyncGetData(id, artistChan, errChan)
	go asyncGetData(id, locationsChan, errChan)
	go func() {
		dates, err := GetDates(id)
		if err != nil {
			errChan <- err
		}
		datesChan <- dates
	}()
	go asyncGetData(id, relationChan, errChan)
	return artistDataBundle{<-artistChan, <-locationsChan, <-datesChan, <-relationChan}, errChan
}

func asyncGetData[T HasAPI](id int, channel chan T, errChan chan error) {
	data, err := getFromAPI[T](id)
	if err != nil {
		errChan <- err
	}
	channel <- data
}
