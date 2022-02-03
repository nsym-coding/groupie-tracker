package groupie

import (
	"encoding/json"
	"io/ioutil"
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

func UnmarshalArtistData() {

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
		ArtistFirstAlbum = append(ArtistFirstAlbum, responseObjectArtists[i].FirstAlbum)
	}

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

	for i := 0; i < len(responseObjectArtists); i++ {
		ArtistName = append(ArtistName, responseObjectArtists[i].Name)
	}

	responseRelations, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		panic("Couldn't get the relations data!")
	}

	responseData, err := ioutil.ReadAll(responseRelations.Body)
	if err != nil {
		panic("Couldn't read data for the Artists")
	}

	var responseObjectRelation Relation

	json.Unmarshal(responseData, &responseObjectRelation)

	//ArtistsDatesLocations := responseObjectRelation.Relation

	// ArtistLocations := make([]string, len(ArtistsDatesLocations.DatesLocations))

	// for i := 0; i < len(ArtistLocations); i++ {

	// 	fmt.Println(ArtistLocations[i])
	// }
	// //ArtistsDatesLofmt.Println(cations)

	// for k, v := range ArtistsDatesLocations{

	//}

	responseDates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		panic("Couldn't get Dates info from the API!")
	}
	defer responseDates.Body.Close()

	responseDatesData, err := ioutil.ReadAll(responseDates.Body)
	if err != nil {
		panic("Couldn't read data for Dates")
	}

	var responseObjectDates Dates
	json.Unmarshal(responseDatesData, &responseObjectDates)

	for i := 0; i < len(responseObjectDates.Dates); i++ {
		ArtistConcertDates = append(ArtistConcertDates, responseObjectDates.Dates[i].Dates)
	}

	responseLocations, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		panic("Couldn't get Location info from API")
	}
	defer responseLocations.Body.Close()

	responseLocationsData, err := ioutil.ReadAll(responseLocations.Body)
	if err != nil {
		panic("Couldn't read data for Locations!")
	}

	var responseObjectLocations Locations
	json.Unmarshal(responseLocationsData, &responseObjectLocations)

	//fmt.Println(responseObjectLocations.Locations[0].Locations)

	for i := 0; i < len(responseObjectLocations.Locations); i++ {
		ArtistLocations = append(ArtistLocations, responseObjectLocations.Locations[i].Locations)
	}

}
