package groupie

import (
	"encoding/json"
	"fmt"
	//"html/template"
	"io/ioutil"
	//"log"
	"net/http"
)
var (
	ArtistID              []int
	ArtistImage           []string
	ArtistName            []string
	ArtistMembers         [][]string
	ArtistCreationDate    []int
	ArtistFirstAlbum      []string
	ArtistLocations       [][]string
	ArtistConcertDates    [][]string
	ArtistsDatesLocations map[string][]string
)

type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Dates struct {
	Dates []dates `json:"index"`
}

type dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Locations struct {
	Locations []locations `json:"index"`
}

type locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Relation struct {
	Relation []relations `json:"index"`
}

type relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func main() {
	responseArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		panic("Couldn't get Artists info from API")
	}
	defer responseArtists.Body.Close()

	responseArtistsData, err := ioutil.ReadAll(responseArtists.Body)
	if err != nil {
		panic("Couldn't read data for Artists!")
	}

	var responseObjectArtists Artists
	json.Unmarshal(responseArtistsData, &responseObjectArtists)

	for i := 0; i < len(responseObjectArtists); i++ {
		ArtistID = append(ArtistID, responseObjectArtists[i].ID)
	}

	for i := 0; i < len(responseObjectArtists); i++ {
		ArtistImage = append(ArtistImage, responseObjectArtists[i].Image)
	}


	for i := 0; i < len(responseObjectArtists); i++ {
		ArtistMembers = append(ArtistMembers, responseObjectArtists[i].Members)
	}

	for i := 0; i < len(responseObjectArtists); i++ {
		ArtistCreationDate = append(ArtistCreationDate, responseObjectArtists[i].CreationDate)
	}

	fmt.Println(ArtistImage[0])
	fmt.Println(ArtistCreationDate[0])

}