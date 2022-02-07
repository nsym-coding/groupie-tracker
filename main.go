package main

import (
	"fmt"

	"git.learn.01founders.co/nsym_coding/groupie-tracker.git/groupie"
)

func main() {

	groupie.UnmarshalArtistData()

	for i := 0; i < len(groupie.TotalInfo.ArtistID); i++ {
		fmt.Println(groupie.TotalInfo.ArtistID[i])
		fmt.Println(groupie.TotalInfo.ArtistName[i])
		fmt.Println(groupie.TotalInfo.ArtistMembers[i])
		fmt.Println(groupie.TotalInfo.ArtistCreationDate[i])
		fmt.Println(groupie.TotalInfo.ArtistFirstAlbum[i])
		fmt.Println(groupie.TotalInfo.ArtistConcertDates[i])
		fmt.Println(groupie.TotalInfo.ArtistLocations[i])
		fmt.Println()
	}
	//groupie.Requests()
}
