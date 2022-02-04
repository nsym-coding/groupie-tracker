package main

import (
	"fmt"

	"git.learn.01founders.co/nsym_coding/groupie-tracker.git/groupie"
)




func main() {

	groupie.UnmarshalArtistData()
	for k, v := range groupie.ArtistsDatesLocations[""]{
		fmt.Println(k, v)
	}

	groupie.Requests()
}
