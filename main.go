package main

import (
	"git.learn.01founders.co/nsym_coding/groupie-tracker.git/groupie"
)

func main() {

	groupie.UnmarshalArtistData()
	groupie.UnmarshalDatesData()
	groupie.UnmarshalRelationsData()
	groupie.UnmarshallLocationsData()
	groupie.Requests()

}
